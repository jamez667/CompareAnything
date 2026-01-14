// BCP 47 locale codes (language-country format)
export const LOCALES = {
  'en-US': {
    code: 'en-US',
    name: 'United States - English',
    currency: 'USD',
    units: 'imperial',
    language: 'en',
    country: 'US',
  },
  'es-US': {
    code: 'es-US',
    name: 'United States - Spanish',
    currency: 'USD',
    units: 'imperial',
    language: 'es',
    country: 'US',
  },
  'en-CA': {
    code: 'en-CA',
    name: 'Canada - English',
    currency: 'CAD',
    units: 'metric',
    language: 'en',
    country: 'CA',
  },
  'fr-CA': {
    code: 'fr-CA',
    name: 'Canada - French',
    currency: 'CAD',
    units: 'metric',
    language: 'fr',
    country: 'CA',
  },
  'en-GB': {
    code: 'en-GB',
    name: 'United Kingdom - English',
    currency: 'GBP',
    units: 'metric',
    language: 'en',
    country: 'GB',
  },
  'en-EU': {
    code: 'en-EU',
    name: 'Europe - English',
    currency: 'EUR',
    units: 'metric',
    language: 'en',
    country: 'EU',
  },
  'de-DE': {
    code: 'de-DE',
    name: 'Germany - German',
    currency: 'EUR',
    units: 'metric',
    language: 'de',
    country: 'DE',
  },
  'fr-FR': {
    code: 'fr-FR',
    name: 'France - French',
    currency: 'EUR',
    units: 'metric',
    language: 'fr',
    country: 'FR',
  },
  'ja-JP': {
    code: 'ja-JP',
    name: 'Japan - Japanese',
    currency: 'JPY',
    units: 'metric',
    language: 'ja',
    country: 'JP',
  },
};

export const DEFAULT_LOCALE = 'en-US';

export const getLocaleByCurrencyAndUnits = (currency, units) => {
  return Object.values(LOCALES).find(
    locale => locale.currency === currency && locale.units === units
  ) || LOCALES[DEFAULT_LOCALE];
};

