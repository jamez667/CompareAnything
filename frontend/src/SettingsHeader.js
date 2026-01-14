import React, { useState, useEffect } from 'react';
import './SettingsHeader.css';
import { getLocale, setLocale } from './utils/settings';
import { LOCALES } from './utils/localization';
import { t } from './utils/i18n';

function SettingsHeader() {
  const [locale, setLocaleState] = useState(getLocale());

  useEffect(() => {
    setLocaleState(getLocale());
  }, []);

  const handleLocaleChange = (newLocaleCode) => {
    setLocale(newLocaleCode);
    setLocaleState(newLocaleCode);
    window.dispatchEvent(new CustomEvent('settingsChanged'));
  };

  return (
    <div className="settings-header">
      <div className="settings-header-content">
        <div className="settings-controls">
          <div className="settings-group">
            <label htmlFor="locale-select">{t('settings.region', locale)}:</label>
            <select
              id="locale-select"
              value={locale}
              onChange={(e) => handleLocaleChange(e.target.value)}
              className="settings-select"
            >
              {Object.entries(LOCALES).map(([code, localeInfo]) => (
                <option key={code} value={code}>
                  {localeInfo.name}
                </option>
              ))}
            </select>
          </div>
        </div>
      </div>
    </div>
  );
}

export default SettingsHeader;

