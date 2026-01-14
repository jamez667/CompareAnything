import { getCookie, setCookie } from './cookies';
import { LOCALES, DEFAULT_LOCALE } from './localization';
import { getLanguageFromLocale } from './i18n';

const CURRENCY_KEY = 'app_currency';
const UNITS_KEY = 'app_units';
const LOCALE_KEY = 'app_locale';
const COLUMN_VISIBILITY_KEY = 'app_column_visibility';

export const DEFAULT_CURRENCY = LOCALES[DEFAULT_LOCALE].currency;
export const DEFAULT_UNITS = LOCALES[DEFAULT_LOCALE].units;

export const CURRENCIES = {
  USD: { symbol: '$', name: 'US Dollar', code: 'USD' },
  EUR: { symbol: '€', name: 'Euro', code: 'EUR' },
  GBP: { symbol: '£', name: 'British Pound', code: 'GBP' },
  JPY: { symbol: '¥', name: 'Japanese Yen', code: 'JPY' },
  CAD: { symbol: 'C$', name: 'Canadian Dollar', code: 'CAD' },
};

// Exchange rates (simplified - in production, fetch from API)
export const EXCHANGE_RATES = {
  USD: 1.0,
  EUR: 0.92,
  GBP: 0.79,
  JPY: 149.5,
  CAD: 1.36,
};

export const UNIT_SYSTEMS = {
  metric: {
    weight: { unit: 'g', name: 'Metric', label: 'Weight (g)' },
    length: { unit: 'mm', name: 'Millimeters', label: 'Dimensions (mm)' },
    convertWeight: (g) => Math.round(g),
    convertLength: (mm) => Math.round(mm),
    displayName: 'Metric',
  },
  imperial: {
    weight: { unit: 'oz', name: 'Imperial', label: 'Weight (oz)' },
    length: { unit: 'in', name: 'Inches', label: 'Dimensions (in)' },
    convertWeight: (g) => parseFloat((g / 28.35).toFixed(1)),
    convertLength: (mm) => parseFloat((mm / 25.4).toFixed(2)),
    displayName: 'Imperial',
  },
};

export const getLocale = () => {
  const saved = getCookie(LOCALE_KEY);
  return saved && LOCALES[saved] ? saved : DEFAULT_LOCALE;
};

export const setLocale = (localeCode) => {
  if (LOCALES[localeCode]) {
    const locale = LOCALES[localeCode];
    setCookie(LOCALE_KEY, localeCode);
    setCookie(CURRENCY_KEY, locale.currency);
    setCookie(UNITS_KEY, locale.units);
  }
};

export const getCurrency = () => {
  const localeCode = getCookie(LOCALE_KEY);
  if (localeCode && LOCALES[localeCode]) {
    return LOCALES[localeCode].currency;
  }
  const saved = getCookie(CURRENCY_KEY);
  return saved && CURRENCIES[saved] ? saved : DEFAULT_CURRENCY;
};

export const setCurrency = (currency) => {
  if (CURRENCIES[currency]) {
    setCookie(CURRENCY_KEY, currency);
  }
};

export const getUnits = () => {
  const localeCode = getCookie(LOCALE_KEY);
  if (localeCode && LOCALES[localeCode]) {
    return LOCALES[localeCode].units;
  }
  const saved = getCookie(UNITS_KEY);
  return saved && UNIT_SYSTEMS[saved] ? saved : DEFAULT_UNITS;
};

export const setUnits = (units) => {
  if (UNIT_SYSTEMS[units]) {
    setCookie(UNITS_KEY, units);
  }
};

export const formatPrice = (priceUSD, currency) => {
  const rate = EXCHANGE_RATES[currency] || 1;
  const convertedPrice = priceUSD * rate;
  const currencyInfo = CURRENCIES[currency] || CURRENCIES[DEFAULT_CURRENCY];
  
  if (currency === 'JPY') {
    return `${currencyInfo.symbol}${Math.round(convertedPrice).toLocaleString()}`;
  }
  
  return `${currencyInfo.symbol}${convertedPrice.toFixed(2)}`;
};

export const formatWeight = (weightGrams, units) => {
  const unitSystem = UNIT_SYSTEMS[units] || UNIT_SYSTEMS[DEFAULT_UNITS];
  const converted = unitSystem.convertWeight(weightGrams);
  return `${converted}${unitSystem.weight.unit}`;
};

export const formatLength = (lengthMM, units) => {
  const unitSystem = UNIT_SYSTEMS[units] || UNIT_SYSTEMS[DEFAULT_UNITS];
  const converted = unitSystem.convertLength(lengthMM);
  return `${converted}${unitSystem.length.unit}`;
};

// Column visibility management
export const getColumnVisibility = () => {
  const saved = getCookie(COLUMN_VISIBILITY_KEY);
  if (saved) {
    try {
      return JSON.parse(saved);
    } catch (e) {
      return null;
    }
  }
  return null;
};

export const setColumnVisibility = (visibility) => {
  setCookie(COLUMN_VISIBILITY_KEY, JSON.stringify(visibility));
};

// Default: all columns visible
export const getDefaultColumnVisibility = () => {
  return {
    name: true,
    focalLength: true,
    maxAperture: true,
    minAperture: true,
    imageStab: true,
    weatherSealing: true,
    length: true,
    diameter: true,
    weight: true,
    price: true,
    type: true,
  };
};

