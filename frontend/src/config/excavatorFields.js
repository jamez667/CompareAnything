import { t } from '../utils/i18n';
import { CURRENCIES } from '../utils/settings';

/**
 * Generate excavator field definitions for the comparison table
 */
export const getExcavatorFields = (locale, units, currency) => {
  const currencyInfo = CURRENCIES[currency] || CURRENCIES.USD;
  
  // Determine unit labels based on unit system
  const lengthUnit = units === 'metric' ? 'cm' : 'in';
  const weightUnit = units === 'metric' ? 'kg' : 'lbs';

  return [
    { 
      key: 'manufacturer', 
      label: t('table.manufacturer', locale),
      filterType: 'multiSelect'
    },
    { 
      key: 'model', 
      label: t('table.model', locale),
      filterType: 'text'
    },
    { 
      key: 'scale', 
      label: t('table.scale', locale),
      filterType: 'select',
      selectOptions: [
        { value: '1:14', label: '1:14' },
        { value: '1:16', label: '1:16' },
        { value: '1:18', label: '1:18' },
        { value: '1:20', label: '1:20' },
        { value: '1:24', label: '1:24' },
      ]
    },
    { 
      key: 'year', 
      label: t('table.year', locale),
      filterType: 'range',
      step: 1
    },
    { 
      key: 'weight', 
      label: `${t('table.weight', locale)} (${weightUnit})`,
      filterType: 'range',
      step: 0.1
    },
    { 
      key: 'length', 
      label: `${t('table.length', locale)} (${lengthUnit})`,
      filterType: 'range',
      step: 0.1
    },
    { 
      key: 'width', 
      label: `${t('table.width', locale)} (${lengthUnit})`,
      filterType: 'range',
      step: 0.1
    },
    { 
      key: 'height', 
      label: `${t('table.height', locale)} (${lengthUnit})`,
      filterType: 'range',
      step: 0.1
    },
    { 
      key: 'batteryType', 
      label: t('table.batteryType', locale),
      filterType: 'select',
      selectOptions: [
        { value: 'LiPo', label: 'LiPo' },
        { value: 'NiMH', label: 'NiMH' },
        { value: 'AA', label: 'AA Batteries' },
      ]
    },
    { 
      key: 'batteryCapacity', 
      label: t('table.batteryCapacity', locale) + ' (mAh)',
      filterType: 'range',
      step: 100
    },
    { 
      key: 'operatingTime', 
      label: t('table.operatingTime', locale) + ' (min)',
      filterType: 'range',
      step: 1
    },
    { 
      key: 'controlChannels', 
      label: t('table.controlChannels', locale),
      filterType: 'range',
      step: 1
    },
    { 
      key: 'diggingDepth', 
      label: `${t('table.diggingDepth', locale)} (${lengthUnit})`,
      filterType: 'range',
      step: 0.1
    },
    { 
      key: 'liftingCapacity', 
      label: `${t('table.liftingCapacity', locale)} (${weightUnit})`,
      filterType: 'range',
      step: 0.01
    },
    { 
      key: 'rotationAngle', 
      label: t('table.rotationAngle', locale) + ' (°)',
      filterType: 'range',
      step: 90
    },
    { 
      key: 'material', 
      label: t('table.material', locale),
      filterType: 'select',
      selectOptions: [
        { value: 'Metal', label: t('table.materialMetal', locale) || 'Metal' },
        { value: 'Plastic', label: t('table.materialPlastic', locale) || 'Plastic' },
        { value: 'Metal/Plastic', label: t('table.materialMixed', locale) || 'Metal/Plastic' },
      ]
    },
    { 
      key: 'remoteControl', 
      label: t('table.remoteControl', locale),
      filterType: 'select',
      selectOptions: [
        { value: '2.4GHz', label: '2.4GHz' },
        { value: 'None', label: t('table.none', locale) || 'None' },
      ]
    },
    { 
      key: 'functions', 
      label: t('table.functions', locale),
      filterType: 'text'
    },
    { 
      key: 'price', 
      label: `${t('table.price', locale)} (${currencyInfo.symbol})`,
      filterType: 'range',
      step: 0.01
    },
  ];
};
