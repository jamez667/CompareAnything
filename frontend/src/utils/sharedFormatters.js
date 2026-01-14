import { formatPrice, formatWeight, formatLength } from './settings';
import { t } from './i18n';

/**
 * Common unit conversion utilities
 */

// Weight conversions
export const kgToLbs = (kg) => {
  return kg * 2.20462;
};

export const lbsToKg = (lbs) => {
  return lbs * 0.453592;
};

// Length conversions
export const cmToInches = (cm) => {
  return cm / 2.54;
};

export const inchesToCm = (inches) => {
  return inches * 2.54;
};

// MPG conversions
export const mpgToL100km = (mpg) => {
  if (!mpg || mpg === 0) return 0;
  return 235.214583 / mpg;
};

export const l100kmToMpg = (l100km) => {
  if (!l100km || l100km === 0) return 0;
  return 235.214583 / l100km;
};

// MPGe to kWh/100km for electric cars
export const mpgeToKwh100km = (mpge) => {
  if (!mpge || mpge === 0) return 0;
  return 2093 / mpge;
};

// Volume conversions
export const litersToGallons = (liters) => {
  return liters / 3.78541;
};

export const gallonsToLiters = (gallons) => {
  return gallons * 3.78541;
};

export const litersToCubicFeet = (liters) => {
  return liters / 28.3168;
};

export const cubicFeetToLiters = (cubicFeet) => {
  return cubicFeet * 28.3168;
};

// Speed conversions
export const mphToKmh = (mph) => {
  return mph * 1.60934;
};

export const kmhToMph = (kmh) => {
  return kmh / 1.60934;
};

/**
 * Format price (shared utility)
 */
export const formatPriceValue = (value, currency) => {
  return formatPrice(value, currency);
};

/**
 * Format weight with unit conversion
 * @param {number} value - Weight value
 * @param {string} units - 'metric' or 'imperial'
 * @param {string} sourceUnit - 'kg' or 'lbs' (the unit the data is stored in)
 * @param {number} decimals - Number of decimal places (default: 0 for kg, 2 for lbs)
 */
export const formatWeightValue = (value, units, sourceUnit = 'lbs', decimals = null) => {
  if (value === null || value === undefined || value === '') {
    return '-';
  }

  if (units === 'metric') {
    const kg = sourceUnit === 'lbs' ? lbsToKg(value) : value;
    const dec = decimals !== null ? decimals : 0;
    return `${kg.toFixed(dec)} kg`;
  } else {
    const lbs = sourceUnit === 'kg' ? kgToLbs(value) : value;
    const dec = decimals !== null ? decimals : 2;
    return `${lbs.toFixed(dec)} lbs`;
  }
};

/**
 * Format length/distance with unit conversion
 * @param {number} value - Length value
 * @param {string} units - 'metric' or 'imperial'
 * @param {string} sourceUnit - 'cm' or 'in' (the unit the data is stored in)
 * @param {number} decimals - Number of decimal places (default: 1)
 */
export const formatLengthValue = (value, units, sourceUnit = 'in', decimals = 1) => {
  if (value === null || value === undefined || value === '') {
    return '-';
  }

  if (units === 'metric') {
    const cm = sourceUnit === 'in' ? inchesToCm(value) : value;
    return `${cm.toFixed(decimals)} cm`;
  } else {
    const inches = sourceUnit === 'cm' ? cmToInches(value) : value;
    return `${inches.toFixed(decimals)} in`;
  }
};

/**
 * Format empty/null values
 */
export const formatEmptyValue = (value, emptyText = '-') => {
  if (value === null || value === undefined || value === '') {
    return emptyText;
  }
  return value;
};

/**
 * Format boolean as Yes/No
 */
export const formatBoolean = (value, locale) => {
  return value ? t('table.yes', locale) : t('table.no', locale);
};

/**
 * Format number with unit
 */
export const formatNumberWithUnit = (value, unit, decimals = 0) => {
  if (value === null || value === undefined || value === '') {
    return '-';
  }
  return `${value.toFixed(decimals)} ${unit}`;
};

/**
 * Lens-specific formatters
 * These work with grams (weight) and millimeters (length)
 * Re-exported for consistency with other formatters
 */
export const formatLensPrice = formatPrice;
export const formatLensWeight = formatWeight;
export const formatLensLength = formatLength;
