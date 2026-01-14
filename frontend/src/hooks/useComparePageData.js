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
