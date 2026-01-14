import React from 'react';
import { t } from '../utils/i18n';
import './ImageView.css';

/**
 * Shared image view component for compare pages
 * @param {string} imageUrl - URL of the image
 * @param {string} altText - Alt text for the image
 * @param {string} titleKey - Translation key for the view title
 * @param {string} locale - Locale for translations
 * @param {string} imageClassName - Optional CSS class for the image
 */
function ImageView({ imageUrl, altText, titleKey, locale, imageClassName = 'comparison-image' }) {
  if (!imageUrl) {
    return null;
  }

  return (
    <div className="image-view">
      <h3>{t(titleKey, locale)}</h3>
      <div className="image-container">
        <img src={imageUrl} alt={altText} className={imageClassName} />
      </div>
    </div>
  );
}

export default ImageView;
