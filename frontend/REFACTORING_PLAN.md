# Critical Issues Refactoring Plan

**Date:** Created  
**Status:** Ready for Implementation  
**Estimated Time:** 4-6 hours

---

## Overview

This plan addresses the critical modularity issues identified in the audit:
1. Settings change handler duplication (6 files)
2. List page component duplication (3 files)
3. Compare page duplication (3 files)
4. Inconsistent formatter imports

---

## Phase 1: Create Custom Hooks (Foundation)

### Step 1.1: Create `hooks` Directory
**File:** `frontend/src/hooks/` (new directory)

**Action:** Create the directory structure for custom hooks.

---

### Step 1.2: Create `useSettingsSync` Hook
**File:** `frontend/src/hooks/useSettingsSync.js` (new)

**Purpose:** Eliminate settings change handler duplication across 6 files.

**Implementation:**
```javascript
import { useState, useEffect } from 'react';
import { getCurrency, getUnits, getLocale } from '../utils/settings';

/**
 * Custom hook to sync component state with global settings
 * Listens to 'settingsChanged' events and updates local state
 * 
 * @returns {Object} { currency, units, locale }
 */
export function useSettingsSync() {
  const [currency, setCurrencyState] = useState(getCurrency());
  const [units, setUnitsState] = useState(getUnits());
  const [locale, setLocaleState] = useState(getLocale());

  useEffect(() => {
    const handleSettingsChange = () => {
      setCurrencyState(getCurrency());
      setUnitsState(getUnits());
      setLocaleState(getLocale());
    };
    
    window.addEventListener('settingsChanged', handleSettingsChange);
    return () => window.removeEventListener('settingsChanged', handleSettingsChange);
  }, []);

  return { currency, units, locale };
}
```

**Files to Update:**
- `App.js` - Replace settings state management
- `CarsPage.js` - Replace settings state management
- `ExcavatorsPage.js` - Replace settings state management
- `CarComparePage.js` - Replace settings state management
- `ExcavatorComparePage.js` - Replace settings state management
- `ComparePage.js` - Replace settings state management

**Impact:** Removes ~30 lines of duplicated code per file = ~180 lines total

---

### Step 1.3: Create `useDataFetching` Hook
**File:** `frontend/src/hooks/useDataFetching.js` (new)

**Purpose:** Standardize data fetching pattern with loading and error states.

**Implementation:**
```javascript
import { useState, useEffect } from 'react';

/**
 * Custom hook for fetching data from API endpoints
 * 
 * @param {string} apiEndpoint - API endpoint URL
 * @param {Object} options - Optional configuration
 * @param {boolean} options.autoFetch - Whether to fetch on mount (default: true)
 * @returns {Object} { data, loading, error, refetch }
 */
export function useDataFetching(apiEndpoint, options = {}) {
  const { autoFetch = true } = options;
  const [data, setData] = useState([]);
  const [loading, setLoading] = useState(autoFetch);
  const [error, setError] = useState(null);

  const fetchData = async () => {
    try {
      setLoading(true);
      setError(null);
      const response = await fetch(apiEndpoint);
      if (!response.ok) {
        throw new Error(`Failed to fetch: ${response.statusText}`);
      }
      const result = await response.json();
      setData(result);
      setLoading(false);
    } catch (err) {
      console.error(`Error fetching from ${apiEndpoint}:`, err);
      setError(err.message);
      setLoading(false);
    }
  };

  useEffect(() => {
    if (autoFetch) {
      fetchData();
    }
  }, [apiEndpoint, autoFetch]);

  return { data, loading, error, refetch: fetchData };
}
```

**Files to Update:**
- `App.js` - Replace fetchLenses
- `CarsPage.js` - Replace fetchCars
- `ExcavatorsPage.js` - Replace fetchExcavators
- `CarComparePage.js` - Replace fetchCars
- `ExcavatorComparePage.js` - Replace fetchExcavators
- `ComparePage.js` - Replace fetchLenses

**Impact:** Removes ~15 lines of duplicated code per file = ~90 lines total

---

### Step 1.4: Create `useListPageData` Hook
**File:** `frontend/src/hooks/useListPageData.js` (new)

**Purpose:** Consolidate all list page logic (data, selection, settings, fields).

**Implementation:**
```javascript
import { useState, useMemo } from 'react';
import { useSettingsSync } from './useSettingsSync';
import { useDataFetching } from './useDataFetching';

/**
 * Custom hook for list page components (App, CarsPage, ExcavatorsPage)
 * Combines data fetching, settings sync, selection, and field generation
 * 
 * @param {string} apiEndpoint - API endpoint URL
 * @param {Function} getFields - Function to generate fields: (locale, units, currency) => Array
 * @param {Function} formatValue - Function to format values: (key, value, item, locale, currency, units) => string
 * @param {Function} getSortValue - Function to get sort value: (item, key) => any
 * @param {string} noResultsKey - Translation key for "no results" message
 * @returns {Object} All data and handlers needed for list page
 */
export function useListPageData(apiEndpoint, getFields, formatValue, getSortValue, noResultsKey) {
  const { currency, units, locale } = useSettingsSync();
  const { data, loading, error } = useDataFetching(apiEndpoint);
  const [selectedIds, setSelectedIds] = useState(new Set());

  const handleSelectionChange = (newSelectedIds) => {
    setSelectedIds(newSelectedIds);
  };

  const fields = useMemo(() => {
    return getFields(locale, units, currency);
  }, [locale, units, currency, getFields]);

  const formatValueWrapper = (key, value, item) => {
    return formatValue(key, value, item, locale, currency, units);
  };

  return {
    data,
    loading,
    error,
    selectedIds,
    setSelectedIds,
    handleSelectionChange,
    fields,
    formatValue: formatValueWrapper,
    getSortValue,
    currency,
    units,
    locale,
    noResultsKey,
  };
}
```

**Files to Update:**
- `App.js` - Replace entire component logic
- `CarsPage.js` - Replace entire component logic
- `ExcavatorsPage.js` - Replace entire component logic

**Impact:** Reduces each file from ~110 lines to ~40 lines = ~210 lines saved

---

### Step 1.5: Create `useComparePageData` Hook
**File:** `frontend/src/hooks/useComparePageData.js` (new)

**Purpose:** Consolidate compare page logic (data, URL params, settings).

**Implementation:**
```javascript
import { useState, useEffect } from 'react';
import { useSearchParams } from 'react-router-dom';
import { useSettingsSync } from './useSettingsSync';
import { useDataFetching } from './useDataFetching';

/**
 * Custom hook for compare page components
 * Handles data fetching, URL parameter parsing, and selected items
 * 
 * @param {string} apiEndpoint - API endpoint URL
 * @returns {Object} { data, selectedItems, loading, error, currency, units, locale }
 */
export function useComparePageData(apiEndpoint) {
  const { currency, units, locale } = useSettingsSync();
  const { data, loading, error } = useDataFetching(apiEndpoint);
  const [selectedItems, setSelectedItems] = useState([]);
  const [searchParams] = useSearchParams();

  useEffect(() => {
    if (data.length > 0) {
      const ids = searchParams.get('ids');
      if (ids) {
        const idArray = ids.split(',');
        const selected = data.filter(item => idArray.includes(item.id));
        setSelectedItems(selected);
      } else {
        setSelectedItems([]);
      }
    }
  }, [data, searchParams]);

  return {
    data,
    selectedItems,
    loading,
    error,
    currency,
    units,
    locale,
  };
}
```

**Files to Update:**
- `CarComparePage.js` - Replace data fetching and URL parsing
- `ExcavatorComparePage.js` - Replace data fetching and URL parsing
- `ComparePage.js` - Replace data fetching and URL parsing (after standardization)

**Impact:** Removes ~30 lines of duplicated code per file = ~90 lines total

---

## Phase 2: Standardize Formatter Imports

### Step 2.1: Add Lens Formatters to `sharedFormatters.js`
**File:** `frontend/src/utils/sharedFormatters.js` (modify)

**Purpose:** Add lens-specific formatter wrappers for consistency.

**Action:** Add these functions to the end of `sharedFormatters.js`:

```javascript
/**
 * Lens-specific formatters (wrappers for settings.js functions)
 * These work with grams (weight) and millimeters (length)
 */

/**
 * Format lens weight (data in grams)
 */
export const formatLensWeight = (weightGrams, units) => {
  const { formatWeight } = require('./settings');
  return formatWeight(weightGrams, units);
};

/**
 * Format lens length/diameter (data in millimeters)
 */
export const formatLensLength = (lengthMM, units) => {
  const { formatLength } = require('./settings');
  return formatLength(lengthMM, units);
};
```

**Note:** Actually, better approach - just re-export from settings or create proper wrappers. Let me revise:

**Better Implementation:**
```javascript
import { formatPrice, formatWeight, formatLength } from './settings';

// ... existing code ...

/**
 * Lens-specific formatters
 * These work with grams (weight) and millimeters (length)
 * Re-exported for consistency with other formatters
 */
export const formatLensWeight = formatWeight;
export const formatLensLength = formatLength;
export const formatLensPrice = formatPrice;
```

---

### Step 2.2: Update `lensFormatters.js`
**File:** `frontend/src/utils/lensFormatters.js` (modify)

**Purpose:** Use sharedFormatters for consistency.

**Current:**
```javascript
import { formatPrice, formatWeight, formatLength } from './settings';
```

**Change to:**
```javascript
import { formatLensPrice, formatLensWeight, formatLensLength } from './sharedFormatters';
import { t } from './i18n';
```

**Update function calls:**
- `formatPrice(value, currency)` → `formatLensPrice(value, currency)`
- `formatWeight(value, units)` → `formatLensWeight(value, units)`
- `formatLength(value, units)` → `formatLensLength(value, units)`

**Impact:** Consistent import pattern across all formatters

---

## Phase 3: Refactor List Pages

### Step 3.1: Refactor `App.js`
**File:** `frontend/src/App.js` (modify)

**Current Structure:** ~108 lines
**Target Structure:** ~40 lines

**Changes:**
1. Import `useListPageData` hook
2. Remove all state management (use hook instead)
3. Remove `fetchLenses` function
4. Remove settings change handler
5. Simplify component to just render

**Before:**
```javascript
const [lenses, setLenses] = useState([]);
const [selectedIds, setSelectedIds] = useState(new Set());
const [loading, setLoading] = useState(true);
const [currency, setCurrency] = useState(getCurrency());
// ... settings change handler ...
const fetchLenses = async () => { ... };
```

**After:**
```javascript
const {
  data: lenses,
  loading,
  selectedIds,
  handleSelectionChange,
  fields,
  formatValue,
  locale,
  getSortValue,
  noResultsKey,
} = useListPageData(
  '/api/lenses',
  getLensFields,
  formatLensValue,
  getLensSortValue,
  'table.noLensesFound'
);
```

---

### Step 3.2: Refactor `CarsPage.js`
**File:** `frontend/src/CarsPage.js` (modify)

**Same pattern as App.js:**
- Replace with `useListPageData` hook
- Use `/api/cars` endpoint
- Use `getCarFields`, `formatCarValue`, `getCarSortValue`
- Use `'table.noCarsFound'` for noResultsKey

**Note:** Keep the special handling for `range` computed field in `formatValue` wrapper if needed.

---

### Step 3.3: Refactor `ExcavatorsPage.js`
**File:** `frontend/src/ExcavatorsPage.js` (modify)

**Same pattern:**
- Replace with `useListPageData` hook
- Use `/api/excavators` endpoint
- Use `getExcavatorFields`, `formatExcavatorValue`, `getExcavatorSortValue`
- Use `'table.noExcavatorsFound'` for noResultsKey

---

## Phase 4: Refactor Compare Pages

### Step 4.1: Refactor `CarComparePage.js`
**File:** `frontend/src/CarComparePage.js` (modify)

**Changes:**
1. Import `useComparePageData` hook
2. Remove data fetching logic
3. Remove settings change handler
4. Remove URL parameter parsing
5. Simplify component

**Before:**
```javascript
const [cars, setCars] = useState([]);
const [selectedCars, setSelectedCars] = useState([]);
// ... settings change handler ...
// ... fetchCars function ...
// ... URL parsing useEffect ...
```

**After:**
```javascript
const {
  data: cars,
  selectedItems: selectedCars,
  loading,
  currency,
  units,
  locale,
} = useComparePageData('/api/cars');
```

---

### Step 4.2: Refactor `ExcavatorComparePage.js`
**File:** `frontend/src/ExcavatorComparePage.js` (modify)

**Same pattern as CarComparePage:**
- Use `useComparePageData` hook
- Use `/api/excavators` endpoint

---

### Step 4.3: Standardize `ComparePage.js` (Lens)
**File:** `frontend/src/ComparePage.js` (modify)

**Changes:**
1. Use `useComparePageData` hook
2. Use `ComparePageHeader` component (like other compare pages)
3. Use `formatLensValue` from `lensFormatters.js` instead of inline function
4. Use `getLensFields` from config instead of inline fields
5. Extract `DimensionOverlay` to separate component or enhance `ImageView`

**Note:** This is a larger refactor to make it consistent with other compare pages.

---

## Phase 5: Testing & Validation

### Step 5.1: Test List Pages
- [ ] Test `/lenses` page loads and displays data
- [ ] Test `/cars` page loads and displays data
- [ ] Test `/excavators` page loads and displays data
- [ ] Test selection works on all pages
- [ ] Test compare button navigation
- [ ] Test settings changes update all pages
- [ ] Test filtering and sorting

### Step 5.2: Test Compare Pages
- [ ] Test `/lenses/compare` with selected items
- [ ] Test `/cars/compare` with selected items
- [ ] Test `/excavators/compare` with selected items
- [ ] Test back navigation
- [ ] Test empty selection state
- [ ] Test settings changes update compare pages

### Step 5.3: Test Formatters
- [ ] Verify lens formatting works correctly
- [ ] Verify car formatting works correctly
- [ ] Verify excavator formatting works correctly
- [ ] Test unit conversions
- [ ] Test currency conversions

---

## Implementation Order

### Recommended Sequence:

1. **Phase 1** (Steps 1.1-1.5): Create all hooks first
   - This provides the foundation
   - Can test hooks independently
   - Time: ~1.5 hours

2. **Phase 2** (Steps 2.1-2.2): Standardize formatters
   - Quick win, low risk
   - Time: ~30 minutes

3. **Phase 3** (Steps 3.1-3.3): Refactor list pages
   - Apply hooks to list pages
   - Test after each page
   - Time: ~1 hour

4. **Phase 4** (Steps 4.1-4.3): Refactor compare pages
   - Apply hooks to compare pages
   - Standardize ComparePage.js
   - Time: ~1.5 hours

5. **Phase 5**: Testing
   - Comprehensive testing
   - Fix any issues
   - Time: ~1 hour

**Total Estimated Time: 4-6 hours**

---

## Risk Mitigation

### Potential Issues:

1. **Breaking Changes:**
   - **Mitigation:** Test each phase before moving to next
   - Keep old code commented initially

2. **Format Function Signatures:**
   - **Issue:** `formatLensValue` might have different signature
   - **Mitigation:** Check all usages, update hook if needed

3. **Computed Fields:**
   - **Issue:** Car `range` field needs special handling
   - **Mitigation:** Keep wrapper function in CarsPage if needed

4. **URL Parameter Parsing:**
   - **Issue:** Different compare pages might need different logic
   - **Mitigation:** Make hook flexible, add options if needed

---

## Success Criteria

✅ All 6 files using `useSettingsSync` instead of duplicated code  
✅ All 3 list pages using `useListPageData`  
✅ All 3 compare pages using `useComparePageData`  
✅ `lensFormatters.js` using `sharedFormatters.js`  
✅ No functionality broken  
✅ Code reduction: ~570 lines of duplicated code removed  
✅ Consistent patterns across all pages  

---

## Rollback Plan

If issues arise:
1. Keep original files as `.backup` versions
2. Use git branches for each phase
3. Can revert individual phases if needed

---

## Notes

- All hooks should be pure and testable
- Consider adding JSDoc comments to all hooks
- May want to add error boundaries for better error handling
- Consider adding loading skeletons instead of just "Loading..." text

---

**Ready to begin implementation? Start with Phase 1, Step 1.1**
