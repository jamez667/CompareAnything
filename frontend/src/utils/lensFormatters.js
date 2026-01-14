import { formatLensPrice, formatLensWeight, formatLensLength } from './sharedFormatters';
import { t } from './i18n';

/**
 * Format a lens field value for display
 */
export const formatLensValue = (key, value, locale, currency, units) => {
  if (key === 'imageStab') {
    return value ? t('table.yes', locale) : t('table.no', locale);
  }
  if (key === 'price') {
    return formatLensPrice(value, currency);
  }
  if (key === 'weight') {
    return formatLensWeight(value, units);
  }
  if (key === 'length' || key === 'diameter') {
    return formatLensLength(value, units);
  }
  if (key === 'maxAperture' || key === 'minAperture') {
    return `f/${value}`;
  }
  return value;
};

/**
 * Get sort value for a lens field
 */
export const getLensSortValue = (lens, key) => {
  const value = lens[key];
  if (key === 'price' || key === 'weight' || key === 'length' || key === 'diameter' || key === 'maxAperture' || key === 'minAperture') {
    return typeof value === 'number' ? value : parseFloat(value) || 0;
  }
  if (key === 'imageStab') {
    return value ? 1 : 0;
  }
  if (typeof value === 'string') {
    return value.toLowerCase();
  }
  return value;
};

