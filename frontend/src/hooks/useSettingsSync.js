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
