import { t } from '../utils/i18n';
import { CURRENCIES } from '../utils/settings';

/**
 * Generate car field definitions for the comparison table
 */
export const getCarFields = (locale, units, currency) => {
  const currencyInfo = CURRENCIES[currency] || CURRENCIES.USD;
  
  // Determine unit labels based on unit system
  const lengthUnit = units === 'metric' ? 'cm' : 'in';
  const weightUnit = units === 'metric' ? 'kg' : 'lbs';
  const mpgUnit = units === 'metric' ? 'L/100km' : 'mpg';
  const speedUnit = units === 'metric' ? 'km/h' : 'mph';
  const cargoUnit = units === 'metric' ? 'L' : 'ft³';
  const fuelUnit = units === 'metric' ? 'L' : 'gal';

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
      key: 'trim', 
      label: t('table.trim', locale),
      filterType: 'text'
    },
    { 
      key: 'variant', 
      label: t('table.variant', locale),
      filterType: 'text'
    },
    { 
      key: 'year', 
      label: t('table.year', locale),
      filterType: 'range',
      step: 1
    },
    { 
      key: 'type', 
      label: t('table.type', locale),
      filterType: 'select',
      selectOptions: [
        { value: 'Sedan', label: t('table.typeSedan', locale) || 'Sedan' },
        { value: 'SUV', label: t('table.typeSUV', locale) || 'SUV' },
        { value: 'Truck', label: t('table.typeTruck', locale) || 'Truck' },
        { value: 'Coupe', label: t('table.typeCoupe', locale) || 'Coupe' },
        { value: 'Hatchback', label: t('table.typeHatchback', locale) || 'Hatchback' },
        { value: 'Convertible', label: t('table.typeConvertible', locale) || 'Convertible' },
      ]
    },
    { 
      key: 'bodyStyle', 
      label: t('table.bodyStyle', locale),
      filterType: 'select',
      selectOptions: [
        { value: 'Sedan', label: t('table.bodyStyleSedan', locale) || 'Sedan' },
        { value: 'Coupe', label: t('table.bodyStyleCoupe', locale) || 'Coupe' },
        { value: 'Hatchback', label: t('table.bodyStyleHatchback', locale) || 'Hatchback' },
        { value: 'Convertible', label: t('table.bodyStyleConvertible', locale) || 'Convertible' },
        { value: 'SUV', label: t('table.bodyStyleSUV', locale) || 'SUV' },
        { value: 'Pickup', label: t('table.bodyStylePickup', locale) || 'Pickup' },
      ]
    },
    { 
      key: 'engine', 
      label: t('table.engine', locale),
      filterType: 'text'
    },
    { 
      key: 'horsepower', 
      label: t('table.horsepower', locale),
      filterType: 'range',
      step: 1
    },
    { 
      key: 'torque', 
      label: t('table.torque', locale),
      filterType: 'range',
      step: 1
    },
    { 
      key: 'transmission', 
      label: t('table.transmission', locale),
      filterType: 'select',
      selectOptions: [
        { value: 'Automatic', label: t('table.transmissionAutomatic', locale) || 'Automatic' },
        { value: 'Manual', label: t('table.transmissionManual', locale) || 'Manual' },
        { value: 'CVT', label: t('table.transmissionCVT', locale) || 'CVT' },
        { value: 'Single-Speed', label: t('table.transmissionSingleSpeed', locale) || 'Single-Speed' },
      ]
    },
    { 
      key: 'drivetrain', 
      label: t('table.drivetrain', locale),
      filterType: 'select',
      selectOptions: [
        { value: 'FWD', label: t('table.drivetrainFWD', locale) || 'FWD' },
        { value: 'RWD', label: t('table.drivetrainRWD', locale) || 'RWD' },
        { value: 'AWD', label: t('table.drivetrainAWD', locale) || 'AWD' },
        { value: '4WD', label: t('table.drivetrain4WD', locale) || '4WD' },
      ]
    },
    { 
      key: 'mpgCity', 
      label: `${t('table.mpgCity', locale)} (${mpgUnit})`,
      filterType: 'range',
      step: 0.1
    },
    { 
      key: 'mpgHighway', 
      label: `${t('table.mpgHighway', locale)} (${mpgUnit})`,
      filterType: 'range',
      step: 0.1
    },
    { 
      key: 'mpgCombined', 
      label: `${t('table.mpgCombined', locale)} (${mpgUnit})`,
      filterType: 'range',
      step: 0.1
    },
    { 
      key: 'fuelCapacity', 
      label: `${t('table.fuelCapacity', locale)} (${fuelUnit})`,
      filterType: 'range',
      step: 0.1
    },
    { 
      key: 'batteryCapacity', 
      label: `${t('table.batteryCapacity', locale)} (kWh)`,
      filterType: 'range',
      step: 0.1
    },
    { 
      key: 'range', 
      label: t('table.range', locale),
      filterType: 'range',
      step: 1,
      isComputed: true
    },
    { 
      key: 'seating', 
      label: t('table.seating', locale),
      filterType: 'range',
      step: 1
    },
    { 
      key: 'cargoCapacity', 
      label: `${t('table.cargoCapacity', locale)} (${cargoUnit})`,
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
      key: 'wheelbase', 
      label: `${t('table.wheelbase', locale)} (${lengthUnit})`,
      filterType: 'range',
      step: 0.1
    },
    { 
      key: 'weight', 
      label: `${t('table.weight', locale)} (${weightUnit})`,
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
      key: 'topSpeed', 
      label: `${t('table.topSpeed', locale)} (${speedUnit})`,
      filterType: 'range',
      step: 1
    },
    { 
      key: 'zeroToSixty', 
      label: t('table.zeroToSixty', locale),
      filterType: 'range',
      step: 0.1
    },
    { 
      key: 'safetyRating', 
      label: t('table.safetyRating', locale),
      filterType: 'text'
    },
    { 
      key: 'fuelType', 
      label: t('table.fuelType', locale),
      filterType: 'select',
      selectOptions: [
        { value: 'Gasoline', label: t('table.fuelTypeGasoline', locale) || 'Gasoline' },
        { value: 'Electric', label: t('table.fuelTypeElectric', locale) || 'Electric' },
        { value: 'Hybrid', label: t('table.fuelTypeHybrid', locale) || 'Hybrid' },
        { value: 'Diesel', label: t('table.fuelTypeDiesel', locale) || 'Diesel' },
        { value: 'Plug-in Hybrid', label: t('table.fuelTypePlugInHybrid', locale) || 'Plug-in Hybrid' },
      ]
    },
  ];
};
