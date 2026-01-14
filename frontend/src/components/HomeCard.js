import React from 'react';
import { useNavigate } from 'react-router-dom';
import { t } from '../utils/i18n';
import { getLocale } from '../utils/settings';
import './HomeCard.css';

/**
 * Shared home card component
 * @param {string} icon - Emoji or icon to display
 * @param {string} titleKey - Translation key for the title
 * @param {string} descriptionKey - Translation key for the description
 * @param {string} path - Path to navigate to
 * @param {string} locale - Optional locale override
 */
function HomeCard({ icon, titleKey, descriptionKey, path, locale = null }) {
  const navigate = useNavigate();
  const currentLocale = locale || getLocale();

  return (
    <div className="home-card" onClick={() => navigate(path)}>
      <div className="home-card-icon">{icon}</div>
      <h2>{t(titleKey, currentLocale)}</h2>
      <p>{t(descriptionKey, currentLocale)}</p>
    </div>
  );
}

export default HomeCard;
