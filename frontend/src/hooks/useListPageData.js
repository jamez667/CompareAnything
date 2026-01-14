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
