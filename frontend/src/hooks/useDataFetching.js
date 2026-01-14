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
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [apiEndpoint, autoFetch]);

  return { data, loading, error, refetch: fetchData };
}
