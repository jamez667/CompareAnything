import React from 'react';
import { useNavigate } from 'react-router-dom';
import SettingsHeader from '../SettingsHeader';
import { t } from '../utils/i18n';
import { getLocale } from '../utils/settings';
import './ComparePageHeader.css';

/**
 * Shared header component for compare pages
 * @param {string} titleKey - Translation key for the title
 * @param {string} backToListKey - Translation key for back button
 * @param {string} backPath - Path to navigate back to
 * @param {string} locale - Optional locale override
 */
function ComparePageHeader({ titleKey, backToListKey, backPath, locale = null }) {
  const navigate = useNavigate();
  const currentLocale = locale || getLocale();

  return (
    <div className="App">
      <SettingsHeader />
      <header className="App-header">
        <h1>{t(titleKey, currentLocale)}</h1>
        <button className="back-button-header" onClick={() => navigate(backPath)}>
          {t(backToListKey, currentLocale)}
        </button>
      </header>
    </div>
  );
}

export default ComparePageHeader;
