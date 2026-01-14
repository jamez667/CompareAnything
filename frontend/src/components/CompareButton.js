import React from 'react';
import { useNavigate } from 'react-router-dom';
import { t } from '../utils/i18n';
import { getLocale } from '../utils/settings';
import './CompareButton.css';

/**
 * Shared compare button component for list pages
 * @param {number} selectedCount - Number of selected items
 * @param {string} comparePath - Path to navigate to (with ?ids= will be appended)
 * @param {Set} selectedIds - Set of selected item IDs
 * @param {string} locale - Optional locale override
 */
function CompareButton({ selectedCount, comparePath, selectedIds, locale = null }) {
  const navigate = useNavigate();
  const currentLocale = locale || getLocale();

  const handleCompare = () => {
    if (selectedCount > 0) {
      const ids = Array.from(selectedIds).join(',');
      navigate(`${comparePath}?ids=${ids}`);
    }
  };

  if (selectedCount === 0) {
    return null;
  }

  return (
    <button className="compare-button" onClick={handleCompare}>
      {t('table.compareSelected', currentLocale, { count: selectedCount })}
    </button>
  );
}

export default CompareButton;
