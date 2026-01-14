import { t } from '../utils/i18n';
import { UNIT_SYSTEMS, CURRENCIES } from '../utils/settings';

/**
 * Generate lens field definitions for the comparison table
 */
export const getLensFields = (locale, units, currency) => {
  const unitSystem = UNIT_SYSTEMS[units] || UNIT_SYSTEMS.metric;
  const currencyInfo = CURRENCIES[currency] || CURRENCIES.USD;

  return [
    { 
      key: 'name', 
      label: t('table.name', locale),
      filterType: 'text'
    },
    { 
      key: 'focalLength', 
      label: t('table.focalLength', locale),
      filterType: 'customRange',
      step: 1,
      customRangeFilter: (item, filterValue) => {
        const focalLengthStr = item.focalLength || '';
        // Parse focal length range (e.g., "24-70mm" or "50mm")
        const match = focalLengthStr.match(/(\d+)(?:-(\d+))?mm?/);
        if (!match) return false;
        
        const min = parseFloat(match[1]);
        const max = match[2] ? parseFloat(match[2]) : min;
        
        if (filterValue.min !== undefined && max < filterValue.min) return false;
        if (filterValue.max !== undefined && min > filterValue.max) return false;
        return true;
      }
    },
    { 
      key: 'maxAperture', 
      label: t('table.maxAperture', locale),
      filterType: 'range',
      step: 0.1
    },
    { 
      key: 'minAperture', 
      label: t('table.minAperture', locale),
      filterType: 'range',
      step: 0.1
    },
    { 
      key: 'imageStab', 
      label: t('table.imageStabilization', locale),
      filterType: 'boolean'
    },
    { 
      key: 'weatherSealing', 
      label: t('table.weatherSealing', locale),
      filterType: 'select',
      selectOptions: [
        { value: 'dust', label: t('table.weatherSealingDust', locale) || 'Dust' },
        { value: 'drip', label: t('table.weatherSealingDrip', locale) || 'Drip' },
        { value: 'rain', label: t('table.weatherSealingRain', locale) || 'Rain' },
        { value: 'sealed', label: t('table.weatherSealingSealed', locale) || 'Sealed' },
        { value: 'submersible', label: t('table.weatherSealingSubmersible', locale) || 'Submersible' },
      ]
    },
    { 
      key: 'length', 
      label: `${t('table.length', locale)} (${unitSystem.length.unit})`,
      filterType: 'range',
      step: 1
    },
    { 
      key: 'diameter', 
      label: `${t('table.diameter', locale)} (${unitSystem.length.unit})`,
      filterType: 'range',
      step: 1
    },
    { 
      key: 'weight', 
      label: `${t('table.weight', locale)} (${unitSystem.weight.unit})`,
      filterType: 'range',
      step: 1
    },
    { 
      key: 'price', 
      label: `${t('table.price', locale)} (${currencyInfo.symbol})`,
      filterType: 'range',
      step: 0.01
    },
    { 
      key: 'type', 
      label: t('table.type', locale),
      filterType: 'text'
    },
  ];
};

