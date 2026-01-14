# Requirements for Adding Car Comparisons

This document outlines all the pieces needed to add car comparison functionality to the CompareAnything application.

## 1. Backend Changes (Go)

### 1.1 Data Structure
- **File**: `backend/main.go`
- **Action**: Create `Car` struct with fields:
  - `ID` (string) - Unique identifier
  - `Name` (string) - Full car name/model
  - `Manufacturer` (string) - Car manufacturer/brand
  - `Year` (int) - Model year
  - `Type` (string) - Vehicle type (Sedan, SUV, Truck, etc.)
  - `BodyStyle` (string) - Body style (Coupe, Hatchback, etc.)
  - `Engine` (string) - Engine description (e.g., "3.5L V6", "2.0L Turbo I4")
  - `Horsepower` (int) - Horsepower
  - `Torque` (int) - Torque (ft-lb or Nm)
  - `Transmission` (string) - Transmission type (Automatic, Manual, CVT)
  - `Drivetrain` (string) - FWD, RWD, AWD, 4WD
  - `MPGCity` (float64) - City fuel economy (mpg or L/100km)
  - `MPGHighway` (float64) - Highway fuel economy
  - `MPGCombined` (float64) - Combined fuel economy
  - `FuelCapacity` (float64) - Fuel tank capacity (gallons or liters)
  - `Seating` (int) - Number of seats
  - `CargoCapacity` (float64) - Cargo space (cubic feet or liters)
  - `Length` (float64) - Vehicle length (inches or cm)
  - `Width` (float64) - Vehicle width
  - `Height` (float64) - Vehicle height
  - `Wheelbase` (float64) - Wheelbase
  - `Weight` (int) - Curb weight (lbs or kg)
  - `Price` (float64) - Base MSRP
  - `ImageFront` (string) - URL to front view image
  - `ImageSide` (string) - URL to side view image
  - `ImageInterior` (string) - URL to interior image
  - `ImageExterior` (string) - URL to exterior/detail image
  - `TopSpeed` (int) - Top speed (mph or km/h)
  - `ZeroToSixty` (float64) - 0-60 mph time (seconds)
  - `SafetyRating` (string) - NHTSA/IIHS safety rating
  - `FuelType` (string) - Gasoline, Electric, Hybrid, Diesel, Plug-in Hybrid

### 1.2 API Endpoints
- **File**: `backend/main.go`
- **Action**: Add new API routes:
  - `GET /api/cars` - Get all cars
  - `GET /api/cars/{id}` - Get a specific car by ID

### 1.3 Data Storage
- **File**: `backend/main.go`
- **Action**: Create sample car data array with various car models (similar to the lenses array)

### 1.4 Handler Functions
- **File**: `backend/main.go`
- **Action**: Create handler functions:
  - `getCars(w http.ResponseWriter, r *http.Request)`
  - `getCar(w http.ResponseWriter, r *http.Request)`

---

## 2. Frontend Configuration Files

### 2.1 Car Field Definitions
- **File**: `frontend/src/config/carFields.js` (NEW FILE)
- **Action**: Create field definitions similar to `lensFields.js`:
  - Define all car fields with labels, filter types (range, select, text, boolean)
  - Handle unit system conversions (metric/imperial)
  - Handle currency formatting
  - Define select options for fields like:
    - `Type` (Sedan, SUV, Truck, etc.)
    - `BodyStyle` (Coupe, Hatchback, Convertible, etc.)
    - `Transmission` (Automatic, Manual, CVT, etc.)
    - `Drivetrain` (FWD, RWD, AWD, 4WD)
    - `FuelType` (Gasoline, Electric, Hybrid, Diesel, Plug-in Hybrid)
  - Custom range filters for fields like MPG, 0-60 time, etc.

### 2.2 Car Value Formatters
- **File**: `frontend/src/utils/carFormatters.js` (NEW FILE)
- **Action**: Create formatter functions similar to `lensFormatters.js`:
  - `formatCarValue(key, value, locale, currency, units)` - Format car field values
  - `getCarSortValue(car, key)` - Get sortable value for car fields
  - Handle special formatting for:
    - MPG (convert to L/100km for metric)
    - Dimensions (convert inches/cm, feet/meters)
    - Weight (convert lbs/kg)
    - Speed (convert mph/kmh)
    - Power (horsepower/torque)
    - Time (0-60 acceleration)
    - Price (currency formatting)
    - Numeric fields with units

---

## 3. Frontend Page Components

### 3.1 Car List Page
- **File**: `frontend/src/CarsPage.js` (NEW FILE)
- **Action**: Create component similar to `App.js` but for cars:
  - Fetch cars from `/api/cars`
  - Use `ComparisonTable` component (reusable)
  - Use `getCarFields()` for field definitions
  - Use `formatCarValue()` for value formatting
  - Handle navigation to compare page with car IDs
  - Handle settings changes (currency, units, locale)

### 3.2 Car Compare Detail Page
- **File**: `frontend/src/CarComparePage.js` (NEW FILE)
- **Action**: Create component similar to `ComparePage.js` but for cars:
  - Display selected cars side-by-side
  - Show car images (front, side, interior, exterior views)
  - Display all car specifications in a grid layout
  - Handle dimension overlays if needed (similar to lens length/diameter)
  - Format values based on settings (units, currency, locale)
  - Back navigation to car list

---

## 4. Routing Configuration

### 4.1 Update Router
- **File**: `frontend/src/index.js`
- **Action**: Add new routes:
  - `/cars` - Route to `CarsPage` component
  - `/cars/compare` - Route to `CarComparePage` component
  - Consider: Home page or navigation to choose between lenses/cars
  - Consider: Update `/compare` route to be more generic or separate by type

### 4.2 Navigation Updates
- **File**: `frontend/src/App.js` (might need home page instead)
- **Action**: 
  - Consider creating a home page component to choose between lenses and cars
  - OR add navigation between lens and car views
  - Update navigation links throughout the app

---

## 5. Localization Files

### 5.1 English Translations
- **File**: `frontend/src/locales/en.json`
- **Action**: Add car-related translation keys:
  - `carsPage.*` - Car list page translations
  - `carComparePage.*` - Car compare page translations
  - `table.*` - Additional table labels for car-specific fields:
    - `manufacturer`, `year`, `engine`, `horsepower`, `torque`
    - `transmission`, `drivetrain`, `mpgCity`, `mpgHighway`, `mpgCombined`
    - `fuelCapacity`, `seating`, `cargoCapacity`, `wheelbase`
    - `topSpeed`, `zeroToSixty`, `safetyRating`, `fuelType`
    - `bodyStyle`, `width`, `height`, `imageFront`, `imageSide`
    - `imageInterior`, `imageExterior`
  - Vehicle type options (Sedan, SUV, Truck, etc.)
  - Body style options (Coupe, Hatchback, Convertible, etc.)
  - Transmission options (Automatic, Manual, CVT, etc.)
  - Drivetrain options (FWD, RWD, AWD, 4WD)
  - Fuel type options (Gasoline, Electric, Hybrid, Diesel, Plug-in Hybrid)

### 5.2 Other Language Files
- **Files**: 
  - `frontend/src/locales/de.json`
  - `frontend/src/locales/es.json`
  - `frontend/src/locales/fr.json`
  - `frontend/src/locales/ja.json`
- **Action**: Add corresponding translations for all car-related keys in each language

---

## 6. Styling (CSS)

### 6.1 Car Page Styles
- **File**: `frontend/src/CarsPage.css` (NEW FILE - if needed)
- **Action**: Create styles for car list page (may reuse `App.css` if generic enough)

### 6.2 Car Compare Page Styles
- **File**: `frontend/src/CarComparePage.css` (NEW FILE)
- **Action**: Create styles similar to `ComparePage.css` but adapted for cars:
  - Grid layout for car comparison cards
  - Image display styles (front, side, interior, exterior)
  - Spec table/grid styling
  - Responsive design for car cards
  - Consider different image aspect ratios for cars vs lenses

---

## 7. Settings/Utilities (May need updates)

### 7.1 Settings Utilities
- **File**: `frontend/src/utils/settings.js`
- **Action**: Review if any car-specific settings are needed:
  - Check if unit conversions handle car-specific units correctly
  - Verify currency formatting works for car prices
  - Consider if separate settings needed for car preferences

### 7.2 Column Visibility Settings
- **File**: `frontend/src/utils/settings.js`
- **Action**: May need to add car-specific column visibility storage keys
  - Or make column visibility keys include type prefix (e.g., `car_horsepower` vs `lens_focalLength`)

---

## 8. Optional Enhancements

### 8.1 Home/Navigation Page
- **File**: `frontend/src/HomePage.js` (NEW FILE - OPTIONAL)
- **Action**: Create a landing page to choose between comparing lenses or cars
- **Alternative**: Add navigation header/menu to switch between lens and car views

### 8.2 Shared Type Constants
- **File**: `frontend/src/config/types.js` (NEW FILE - OPTIONAL)
- **Action**: Define constants for:
  - Comparison types (LENS, CAR)
  - Field types
  - Filter types
  - Could help with future extensibility

### 8.3 Generic Comparison Components (Refactoring Opportunity)
- **Consider**: Refactoring `ComparePage.js` to be more generic and reusable
- **Action**: Create a generic `ComparisonDetailPage` component that accepts:
  - Field definitions
  - Formatter functions
  - Image fields configuration
  - Could reduce code duplication between lens and car compare pages

---

## Summary Checklist

### Backend (Go)
- [ ] Create `Car` struct in `backend/main.go`
- [ ] Add `getCars()` handler function
- [ ] Add `getCar()` handler function  
- [ ] Add `/api/cars` route
- [ ] Add `/api/cars/{id}` route
- [ ] Create sample car data array

### Frontend Configuration
- [ ] Create `frontend/src/config/carFields.js`
- [ ] Create `frontend/src/utils/carFormatters.js`

### Frontend Components
- [ ] Create `frontend/src/CarsPage.js`
- [ ] Create `frontend/src/CarComparePage.js`
- [ ] Create `frontend/src/CarsPage.css` (if needed)
- [ ] Create `frontend/src/CarComparePage.css`

### Routing
- [ ] Update `frontend/src/index.js` with car routes
- [ ] Add navigation between lens/car views (or create home page)

### Localization
- [ ] Update `frontend/src/locales/en.json`
- [ ] Update `frontend/src/locales/de.json`
- [ ] Update `frontend/src/locales/es.json`
- [ ] Update `frontend/src/locales/fr.json`
- [ ] Update `frontend/src/locales/ja.json`

### Testing & Polish
- [ ] Test car data fetching
- [ ] Test car comparison table filtering/sorting
- [ ] Test car comparison detail page
- [ ] Test unit conversions (metric/imperial)
- [ ] Test currency formatting
- [ ] Test localization in all languages
- [ ] Verify responsive design
- [ ] Add sample car images/placeholders

---

## Notes

- The `ComparisonTable` component is already generic and reusable - no changes needed!
- Most utility functions (settings, i18n, cookies) should work as-is
- Consider the relationship between `/compare` (lenses) and `/cars/compare` routes - may want to make routing more consistent
- Car images will likely need different aspect ratios than lens images
- Some car fields may need custom formatters (e.g., MPG to L/100km conversion, 0-60 time formatting)
