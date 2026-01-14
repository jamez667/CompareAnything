import React, { useState, useEffect, useMemo } from 'react';
import { getColumnVisibility, setColumnVisibility, getDefaultColumnVisibility } from '../utils/settings';
import { t } from '../utils/i18n';
import './ComparisonTable.css';

/**
 * ColumnToggle component - can be used separately if needed
 */
export function ColumnToggle({ fields, columnVisibility, onToggle, locale, showMenu, onMenuToggle }) {
  return (
    <div className="column-toggle-container">
      <button 
        className="column-toggle-button"
        onClick={() => onMenuToggle(!showMenu)}
      >
        {t('table.columns', locale) || 'Columns'} ▾
      </button>
      {showMenu && (
        <div className="column-menu">
          <div className="column-menu-header">
            {t('table.showColumns', locale) || 'Show Columns'}
          </div>
          {fields.map(field => (
            <label key={field.key} className="column-menu-item">
              <input
                type="checkbox"
                checked={columnVisibility[field.key] !== false}
                onChange={() => onToggle(field.key)}
              />
              <span>{field.label}</span>
            </label>
          ))}
        </div>
      )}
    </div>
  );
}

/**
 * Generic comparison table component
 * 
 * @param {Object} props
 * @param {Array} props.data - Array of data objects to display
 * @param {Array} props.fields - Array of field definitions: [{ key, label, filterType?, ... }]
 * @param {Function} props.formatValue - Function to format cell values: (key, value, item) => string
 * @param {Function} props.onSelectionChange - Callback when selection changes: (selectedIds: Set) => void
 * @param {Set} props.selectedIds - Currently selected item IDs
 * @param {string} props.locale - Current locale
 * @param {Function} props.getSortValue - Optional custom sort value function: (item, key) => any
 * @param {Object} props.filterConfig - Optional filter configuration
 * @param {string} props.noResultsKey - Optional translation key for "no results" message (defaults to 'table.noResults')
 */
function ComparisonTable({
  data,
  fields: allFieldsDef,
  formatValue,
  onSelectionChange,
  selectedIds = new Set(),
  locale = 'en-US',
  getSortValue: customGetSortValue,
  filterConfig = {},
  noResultsKey = 'table.noResults',
}) {
  const [filters, setFilters] = useState({});
  const [sortConfig, setSortConfig] = useState({ key: null, direction: 'asc' });
  const [columnVisibility, setColumnVisibilityState] = useState(() => {
    const saved = getColumnVisibility();
    return saved || getDefaultColumnVisibility();
  });
  const [showColumnMenu, setShowColumnMenu] = useState(false);
  const [openMultiSelect, setOpenMultiSelect] = useState(null);
  const [multiSelectSearch, setMultiSelectSearch] = useState({});

  // Close column menu when clicking outside
  useEffect(() => {
    const handleClickOutside = (event) => {
      if (showColumnMenu && !event.target.closest('.column-toggle-container')) {
        setShowColumnMenu(false);
      }
      if (openMultiSelect && !event.target.closest(`.multiselect-filter-container-${openMultiSelect}`)) {
        setOpenMultiSelect(null);
      }
    };
    document.addEventListener('mousedown', handleClickOutside);
    return () => document.removeEventListener('mousedown', handleClickOutside);
  }, [showColumnMenu, openMultiSelect]);

  // Build fields array with visibility filtering
  const allFields = useMemo(() => {
    return allFieldsDef.map(field => ({
      ...field,
      filterType: field.filterType || 'text', // text, range, select, boolean, multiSelect, customRange
      selectOptions: field.selectOptions || [],
    }));
  }, [allFieldsDef]);

  const fields = useMemo(() => {
    return allFields.filter(field => columnVisibility[field.key] !== false);
  }, [allFields, columnVisibility]);

  const handleFilterChange = (key, value) => {
    setFilters(prev => ({
      ...prev,
      [key]: value
    }));
  };

  const handleRangeFilterChange = (key, type, value) => {
    setFilters(prev => ({
      ...prev,
      [key]: {
        ...prev[key],
        [type]: value === '' ? undefined : parseFloat(value)
      }
    }));
  };

  const handleMultiSelectToggle = (key, value) => {
    setFilters(prev => {
      const currentValues = prev[key] || [];
      const newValues = Array.isArray(currentValues) ? [...currentValues] : [];
      const index = newValues.indexOf(value);
      
      if (index > -1) {
        newValues.splice(index, 1);
      } else {
        newValues.push(value);
      }
      
      return {
        ...prev,
        [key]: newValues.length > 0 ? newValues : undefined
      };
    });
  };

  const handleColumnVisibilityToggle = (key) => {
    const newVisibility = {
      ...columnVisibility,
      [key]: !columnVisibility[key]
    };
    setColumnVisibilityState(newVisibility);
    setColumnVisibility(newVisibility);
  };

  // Default sort value function
  const defaultGetSortValue = (item, key) => {
    const value = item[key];
    if (typeof value === 'number') {
      return value;
    }
    if (typeof value === 'boolean') {
      return value ? 1 : 0;
    }
    if (typeof value === 'string') {
      return value.toLowerCase();
    }
    return value;
  };

  const getSortValue = customGetSortValue || defaultGetSortValue;

  const handleSort = (key) => {
    let direction = 'asc';
    if (sortConfig.key === key && sortConfig.direction === 'asc') {
      direction = 'desc';
    }
    setSortConfig({ key, direction });
  };

  // Apply filtering
  const filteredData = useMemo(() => {
    let filtered = data;

    Object.keys(filters).forEach(key => {
      const filterValue = filters[key];
      const field = allFields.find(f => f.key === key);
      
      if (!field || !filterValue) return;

      // Range filters
      if (field.filterType === 'range') {
        if (filterValue && (filterValue.min !== undefined || filterValue.max !== undefined)) {
          filtered = filtered.filter(item => {
            const value = item[key];
            if (value === null || value === undefined) return false;
            const numValue = typeof value === 'number' ? value : parseFloat(value);
            if (isNaN(numValue)) return false;
            
            if (filterValue.min !== undefined && numValue < filterValue.min) return false;
            if (filterValue.max !== undefined && numValue > filterValue.max) return false;
            return true;
          });
        }
      }
      // Custom range filter (e.g., for focal length parsing)
      else if (field.filterType === 'customRange') {
        if (filterValue && (filterValue.min !== undefined || filterValue.max !== undefined)) {
          filtered = filtered.filter(item => {
            if (field.customRangeFilter) {
              return field.customRangeFilter(item, filterValue);
            }
            return true;
          });
        }
      }
      // Multi-select filters (array of selected values)
      else if (field.filterType === 'multiSelect') {
        if (filterValue && Array.isArray(filterValue) && filterValue.length > 0) {
          filtered = filtered.filter(item => filterValue.includes(item[key]));
        }
      }
      // Select/Boolean filters
      else if (field.filterType === 'select' || field.filterType === 'boolean') {
        if (filterValue && filterValue !== 'all') {
          if (field.filterType === 'boolean') {
            const filterBool = filterValue === 'yes';
            filtered = filtered.filter(item => item[key] === filterBool);
          } else {
            filtered = filtered.filter(item => item[key] === filterValue);
          }
        }
      }
      // Text filters
      else if (typeof filterValue === 'string' && filterValue.trim() !== '') {
        const term = filterValue.toLowerCase().trim();
        filtered = filtered.filter(item => {
          const value = item[key];
          if (value === null || value === undefined) return false;
          if (typeof value === 'string') {
            return value.toLowerCase().includes(term);
          }
          return String(value).toLowerCase().includes(term);
        });
      }
    });

    // Apply sorting
    if (sortConfig.key) {
      filtered = [...filtered].sort((a, b) => {
        const aValue = getSortValue(a, sortConfig.key);
        const bValue = getSortValue(b, sortConfig.key);

        if (aValue < bValue) {
          return sortConfig.direction === 'asc' ? -1 : 1;
        }
        if (aValue > bValue) {
          return sortConfig.direction === 'asc' ? 1 : -1;
        }
        return 0;
      });
    }

    return filtered;
  }, [data, filters, sortConfig, allFields, getSortValue]);

  const toggleSelection = (itemId) => {
    const newSet = new Set(selectedIds);
    if (newSet.has(itemId)) {
      newSet.delete(itemId);
    } else {
      newSet.add(itemId);
    }
    onSelectionChange(newSet);
  };

  const hasActiveFilters = Object.keys(filters).some(key => {
    const filterValue = filters[key];
    if (!filterValue) return false;
    if (Array.isArray(filterValue) && filterValue.length > 0) {
      return true;
    }
    if (typeof filterValue === 'object' && (filterValue.min !== undefined || filterValue.max !== undefined)) {
      return true;
    }
    if (typeof filterValue === 'string' && filterValue.trim() !== '' && filterValue !== 'all') {
      return true;
    }
    return false;
  });

  // Get unique values for multi-select fields
  const getUniqueValues = (key) => {
    const values = new Set();
    data.forEach(item => {
      const value = item[key];
      if (value !== null && value !== undefined && value !== '') {
        values.add(value);
      }
    });
    return Array.from(values).sort();
  };

  return (
    <div className="comparison-table-wrapper">
      <div className="table-controls">
        <div className="filter-info">
          {hasActiveFilters && (
            <>
              {filteredData.length} {t('table.of', locale) || 'of'} {data.length} {t('table.items', locale) || 'items'}
              <button 
                className="clear-filters-btn"
                onClick={() => setFilters({})}
              >
                {t('table.clearFilters', locale) || 'Clear Filters'}
              </button>
            </>
          )}
        </div>
        <ColumnToggle
          fields={allFields}
          columnVisibility={columnVisibility}
          onToggle={handleColumnVisibilityToggle}
          locale={locale}
          showMenu={showColumnMenu}
          onMenuToggle={setShowColumnMenu}
        />
      </div>

      <div className="comparison-table-container">
        <table className="comparison-table">
          <thead>
            <tr>
              <th className="checkbox-column">{t('table.compare', locale)}</th>
              {fields.map(field => {
                // Conditionally show/hide column headers based on data
                // We'll show all columns, but cells will be conditionally rendered
                return (
                  <th 
                    key={field.key} 
                    className={sortConfig.key === field.key ? `sortable sorted-${sortConfig.direction}` : 'sortable'}
                    onClick={() => handleSort(field.key)}
                  >
                    {field.label}
                    {sortConfig.key === field.key && (
                      <span className="sort-arrow">
                        {sortConfig.direction === 'asc' ? ' ↑' : ' ↓'}
                      </span>
                    )}
                  </th>
                );
              })}
            </tr>
            <tr className="filter-row">
              <th className="checkbox-column filter-header"></th>
              {fields.map(field => {
                // Range filters
                if (field.filterType === 'range') {
                  return (
                    <th key={field.key} className="filter-header">
                      <div className="range-filter">
                        <input
                          type="number"
                          placeholder={t('table.min', locale) || "Min"}
                          value={filters[field.key]?.min || ''}
                          onChange={(e) => handleRangeFilterChange(field.key, 'min', e.target.value)}
                          className="range-input"
                          onClick={(e) => e.stopPropagation()}
                          step={field.step || 1}
                        />
                        <span className="range-separator">-</span>
                        <input
                          type="number"
                          placeholder={t('table.max', locale) || "Max"}
                          value={filters[field.key]?.max || ''}
                          onChange={(e) => handleRangeFilterChange(field.key, 'max', e.target.value)}
                          className="range-input"
                          onClick={(e) => e.stopPropagation()}
                          step={field.step || 1}
                        />
                      </div>
                    </th>
                  );
                }
                // Custom range filter
                else if (field.filterType === 'customRange') {
                  return (
                    <th key={field.key} className="filter-header">
                      <div className="range-filter">
                        <input
                          type="number"
                          placeholder={t('table.min', locale) || "Min"}
                          value={filters[field.key]?.min || ''}
                          onChange={(e) => handleRangeFilterChange(field.key, 'min', e.target.value)}
                          className="range-input"
                          onClick={(e) => e.stopPropagation()}
                          step={field.step || 1}
                        />
                        <span className="range-separator">-</span>
                        <input
                          type="number"
                          placeholder={t('table.max', locale) || "Max"}
                          value={filters[field.key]?.max || ''}
                          onChange={(e) => handleRangeFilterChange(field.key, 'max', e.target.value)}
                          className="range-input"
                          onClick={(e) => e.stopPropagation()}
                          step={field.step || 1}
                        />
                      </div>
                    </th>
                  );
                }
                // Multi-select filters
                else if (field.filterType === 'multiSelect') {
                  const uniqueValues = getUniqueValues(field.key);
                  const selectedValues = filters[field.key] || [];
                  const searchTerm = (multiSelectSearch[field.key] || '').toLowerCase();
                  const filteredOptions = uniqueValues.filter(val => 
                    String(val).toLowerCase().includes(searchTerm)
                  );
                  const isOpen = openMultiSelect === field.key;

                  return (
                    <th key={field.key} className="filter-header">
                      <div className={`multiselect-filter-container multiselect-filter-container-${field.key}`}>
                        <button
                          type="button"
                          className="multiselect-filter-button"
                          onClick={(e) => {
                            e.stopPropagation();
                            setOpenMultiSelect(isOpen ? null : field.key);
                          }}
                        >
                          {selectedValues.length > 0 
                            ? `${selectedValues.length} ${t('table.selected', locale) || 'selected'}` 
                            : t('table.filter', locale) || 'Filter...'}
                          <span className="multiselect-arrow">{isOpen ? '▲' : '▼'}</span>
                        </button>
                        {isOpen && (
                          <div className="multiselect-dropdown">
                            <input
                              type="text"
                              className="multiselect-search"
                              placeholder={t('table.search', locale) || 'Search...'}
                              value={multiSelectSearch[field.key] || ''}
                              onChange={(e) => {
                                e.stopPropagation();
                                setMultiSelectSearch(prev => ({
                                  ...prev,
                                  [field.key]: e.target.value
                                }));
                              }}
                              onClick={(e) => e.stopPropagation()}
                            />
                            <div className="multiselect-options">
                              {filteredOptions.length > 0 ? (
                                filteredOptions.map(option => {
                                  const isSelected = selectedValues.includes(option);
                                  return (
                                    <label
                                      key={option}
                                      className="multiselect-option"
                                      onClick={(e) => e.stopPropagation()}
                                    >
                                      <input
                                        type="checkbox"
                                        checked={isSelected}
                                        onChange={() => handleMultiSelectToggle(field.key, option)}
                                      />
                                      <span>{option}</span>
                                    </label>
                                  );
                                })
                              ) : (
                                <div className="multiselect-no-results">
                                  {t(noResultsKey, locale) || 'No results'}
                                </div>
                              )}
                            </div>
                          </div>
                        )}
                      </div>
                    </th>
                  );
                }
                // Select/Boolean filters
                else if (field.filterType === 'select' || field.filterType === 'boolean') {
                  const options = field.filterType === 'boolean' 
                    ? [
                        { value: 'all', label: t('table.all', locale) || 'All' },
                        { value: 'yes', label: t('table.yes', locale) },
                        { value: 'no', label: t('table.no', locale) },
                      ]
                    : [
                        { value: 'all', label: t('table.all', locale) || 'All' },
                        ...field.selectOptions
                      ];

                  return (
                    <th key={field.key} className="filter-header">
                      <select
                        value={filters[field.key] || 'all'}
                        onChange={(e) => handleFilterChange(field.key, e.target.value)}
                        className="column-filter-select"
                        onClick={(e) => e.stopPropagation()}
                      >
                        {options.map(option => (
                          <option key={option.value} value={option.value}>
                            {typeof option.label === 'function' ? option.label(locale) : option.label}
                          </option>
                        ))}
                      </select>
                    </th>
                  );
                }
                // Text filter
                else {
                  return (
                    <th key={field.key} className="filter-header">
                      <input
                        type="text"
                        placeholder={t('table.filter', locale) || "Filter..."}
                        value={filters[field.key] || ''}
                        onChange={(e) => handleFilterChange(field.key, e.target.value)}
                        className="column-filter-input"
                        onClick={(e) => e.stopPropagation()}
                      />
                    </th>
                  );
                }
              })}
            </tr>
          </thead>
          <tbody>
            {filteredData.map(item => (
              <tr key={item.id}>
                <td className="checkbox-column">
                  <input
                    type="checkbox"
                    checked={selectedIds.has(item.id)}
                    onChange={() => toggleSelection(item.id)}
                  />
                </td>
                {fields.map(field => {
                  // Conditionally show/hide fields based on fuel type in table view
                  const isElectric = item.fuelType === 'Electric';
                  const isHybrid = item.fuelType === 'Hybrid' || item.fuelType === 'Plug-in Hybrid';
                  
                  // For electric: show battery capacity, hide fuel capacity (but still render cell with "-")
                  if (field.key === 'fuelCapacity' && isElectric && !isHybrid) {
                    return (
                      <td key={field.key}>-</td>
                    );
                  }
                  
                  // For gas: show fuel capacity, hide battery capacity (but still render cell with "-")
                  if (field.key === 'batteryCapacity' && !isElectric && !isHybrid) {
                    return (
                      <td key={field.key}>-</td>
                    );
                  }
                  
                  // For computed fields, pass a placeholder value
                  const value = field.isComputed ? 'computed' : item[field.key];
                  
                  // Skip if value is null/undefined/empty (unless it's a computed field or trim field)
                  // Trim field should show "Base" when empty, so don't skip it
                  if (!field.isComputed && field.key !== 'trim' && (value === null || value === undefined || value === '' || value === 0)) {
                    return (
                      <td key={field.key}>-</td>
                    );
                  }
                  
                  return (
                    <td key={field.key}>
                      {formatValue(field.key, value, item)}
                    </td>
                  );
                })}
              </tr>
            ))}
          </tbody>
        </table>
        {filteredData.length === 0 && (
          <div className="no-results">
            {t(noResultsKey, locale) || 'No results found'}
          </div>
        )}
      </div>
    </div>
  );
}

export default ComparisonTable;

