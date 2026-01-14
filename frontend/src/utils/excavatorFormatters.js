import { formatPriceValue, formatWeightValue, formatLengthValue, formatEmptyValue } from './sharedFormatters';
import { t } from './i18n';

/**
 * Format an excavator field value for display
 */
export const formatExcavatorValue = (key, value, locale, currency, units) => {
  // Price formatting
  if (key === 'price') {
    return formatPriceValue(value, currency);
  }

  // Weight conversion (data stored in kg, convert to lbs for imperial)
  if (key === 'weight' || key === 'liftingCapacity') {
    return formatWeightValue(value, units, 'kg', 2);
  }

  // Dimension conversions (data stored in cm, convert to inches for imperial)
  if (key === 'length' || key === 'width' || key === 'height' || key === 'diggingDepth') {
    return formatLengthValue(value, units, 'cm', 1);
  }

  // Battery capacity
  if (key === 'batteryCapacity') {
    if (value === 0) return t('table.none', locale) || 'None';
    return `${value} mAh`;
  }

  // Operating time
  if (key === 'operatingTime') {
    if (value === 0) return t('table.none', locale) || 'None';
    return `${value} min`;
  }

  // Control channels
  if (key === 'controlChannels') {
    if (value === 0) return t('table.none', locale) || 'None';
    return `${value}`;
  }

  // Rotation angle
  if (key === 'rotationAngle') {
    return `${value}°`;
  }

  return formatEmptyValue(value);
};

/**
 * Get sort value for an excavator field
 */
export const getExcavatorSortValue = (excavator, key) => {
  // Handle numeric fields
  if (['price', 'weight', 'length', 'width', 'height', 'batteryCapacity', 
       'operatingTime', 'controlChannels', 'diggingDepth', 'liftingCapacity', 
       'rotationAngle', 'year'].includes(key)) {
    const value = excavator[key];
    return typeof value === 'number' ? value : parseFloat(value) || 0;
  }
  
  // String fields
  if (typeof excavator[key] === 'string') {
    return excavator[key].toLowerCase();
  }
  
  return excavator[key];
};
