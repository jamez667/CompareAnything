import enTranslations from '../locales/en.json';
import esTranslations from '../locales/es.json';
import frTranslations from '../locales/fr.json';
import deTranslations from '../locales/de.json';
import jaTranslations from '../locales/ja.json';

const translations = {
  en: enTranslations,
  es: esTranslations,
  fr: frTranslations,
  de: deTranslations,
  ja: jaTranslations,
};

export const getLanguageFromLocale = (locale) => {
  // Extract language code from locale (e.g., 'en-US' -> 'en')
  return locale.split('-')[0] || 'en';
};

export const t = (key, locale = 'en', params = {}) => {
  const lang = getLanguageFromLocale(locale);
  const translation = translations[lang] || translations.en;
  
  // Navigate through nested keys (e.g., 'header.title')
  const keys = key.split('.');
  let value = translation;
  
  for (const k of keys) {
    if (value && typeof value === 'object') {
      value = value[k];
    } else {
      // Fallback to English if key not found
      const enValue = translations.en;
      let fallbackValue = enValue;
      for (const fallbackKey of keys) {
        if (fallbackValue && typeof fallbackValue === 'object') {
          fallbackValue = fallbackValue[fallbackKey];
        } else {
          return key; // Return key if not found even in English
        }
      }
      value = fallbackValue;
      break;
    }
  }
  
  // If value is a string, replace params
  if (typeof value === 'string') {
    return value.replace(/\{(\w+)\}/g, (match, paramKey) => {
      return params[paramKey] !== undefined ? params[paramKey] : match;
    });
  }
  
  return value || key;
};

