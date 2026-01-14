import React from 'react';
import { useNavigate } from 'react-router-dom';
import ComparePageHeader from './components/ComparePageHeader';
import SpecRow from './components/SpecRow';
import { useComparePageData } from './hooks/useComparePageData';
import { formatLensValue } from './utils/lensFormatters';
import { getLensFields } from './config/lensFields';
import { formatLength } from './utils/settings';
import { t } from './utils/i18n';
import './ComparePage.css';

function ComparePage() {
  const navigate = useNavigate();
  const {
    data: lenses,
    selectedItems: selectedLenses,
    loading,
    currency,
    units,
    locale,
  } = useComparePageData('/api/lenses');

  // Get fields from config (for specs display)
  const allFields = getLensFields(locale, units, currency);
  // Filter to only the fields we want to show in compare view
  const fields = allFields.filter(field => 
    ['focalLength', 'maxAperture', 'minAperture', 'imageStab', 'weatherSealing', 'weight', 'price', 'type'].includes(field.key)
  );

  const DimensionOverlay = ({ imageUrl, viewType, length, diameter, units, locale }) => {
    if (!imageUrl) {
      const viewTypeLabel = viewType === 'side' ? t('comparePage.sideView', locale) : t('comparePage.topView', locale);
      return (
        <div className="image-placeholder">
          <p>{t('comparePage.noImageAvailable', locale, { viewType: viewTypeLabel })}</p>
          <div className="dimension-info">
            {length > 0 && <span>{t('comparePage.length', locale)}: {formatLength(length, units)}</span>}
            {diameter > 0 && <span>{t('comparePage.diameter', locale)}: {formatLength(diameter, units)}</span>}
          </div>
        </div>
      );
    }

    return (
      <div className="image-container">
        <img src={imageUrl} alt="" className="lens-image" />
        <div className="dimension-overlay">
          {viewType === 'side' && length > 0 && (
            <div className="dimension-line horizontal">
              <div className="dimension-line-bar"></div>
              <span className="dimension-label">{formatLength(length, units)}</span>
            </div>
          )}
          {viewType === 'side' && diameter > 0 && (
            <div className="dimension-line vertical">
              <div className="dimension-line-bar"></div>
              <span className="dimension-label">{formatLength(diameter, units)}</span>
            </div>
          )}
          {viewType === 'top' && diameter > 0 && (
            <div className="dimension-line circular">
              <div className="dimension-line-bar"></div>
              <span className="dimension-label">{formatLength(diameter, units)}</span>
            </div>
          )}
        </div>
      </div>
    );
  };

  if (loading) {
    return <div className="loading">{t('table.allLenses', locale)}...</div>;
  }

  if (selectedLenses.length === 0) {
    return (
      <div className="App">
        <div className="container">
          <div className="no-selection">
            <h2>{t('comparePage.noSelection', locale)}</h2>
            <button className="back-button" onClick={() => navigate('/lenses')}>
              {t('comparePage.goBack', locale)}
            </button>
          </div>
        </div>
      </div>
    );
  }

  return (
    <>
      <ComparePageHeader
        titleKey="comparePage.title"
        backToListKey="comparePage.backToList"
        backPath="/lenses"
        locale={locale}
      />

      <div className="container">
        <div className="comparison-grid">
          {selectedLenses.map(lens => (
            <div key={lens.id} className="lens-detail-card">
              <div className="lens-detail-header">
                <h2 className="lens-detail-name">{lens.name}</h2>
              </div>
              
              <div className="lens-images-section">
                <div className="image-view">
                  <h3>{t('comparePage.sideView', locale)}</h3>
                  <DimensionOverlay
                    imageUrl={lens.imageSide}
                    viewType="side"
                    length={lens.length}
                    diameter={lens.diameter}
                    units={units}
                    locale={locale}
                  />
                </div>
                <div className="image-view">
                  <h3>{t('comparePage.topView', locale)}</h3>
                  <DimensionOverlay
                    imageUrl={lens.imageTop}
                    viewType="top"
                    length={lens.length}
                    diameter={lens.diameter}
                    units={units}
                    locale={locale}
                  />
                </div>
              </div>

              <div className="lens-detail-specs">
                {fields.map(field => (
                  <SpecRow
                    key={field.key}
                    label={field.label}
                    value={formatLensValue(field.key, lens[field.key], locale, currency, units)}
                  />
                ))}
                <SpecRow
                  label={t('comparePage.length', locale)}
                  value={lens.length > 0 ? formatLength(lens.length, units) : '-'}
                />
                <SpecRow
                  label={t('comparePage.diameter', locale)}
                  value={lens.diameter > 0 ? formatLength(lens.diameter, units) : '-'}
                />
              </div>
            </div>
          ))}
        </div>
      </div>
    </>
  );
}

export default ComparePage;
