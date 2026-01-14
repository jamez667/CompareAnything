import React from 'react';
import { useNavigate } from 'react-router-dom';
import ComparePageHeader from './components/ComparePageHeader';
import ImageView from './components/ImageView';
import SpecRow from './components/SpecRow';
import { useComparePageData } from './hooks/useComparePageData';
import { formatExcavatorValue } from './utils/excavatorFormatters';
import { getExcavatorFields } from './config/excavatorFields';
import { t } from './utils/i18n';
import './ExcavatorComparePage.css';

function ExcavatorComparePage() {
  const navigate = useNavigate();
  const {
    data: excavators,
    selectedItems: selectedExcavators,
    loading,
    currency,
    units,
    locale,
  } = useComparePageData('/api/excavators');

  // Generate fields based on current settings
  const fields = getExcavatorFields(locale, units, currency);

  if (loading) {
    return <div className="loading">{t('table.allExcavators', locale) || 'Loading excavators'}...</div>;
  }

  if (selectedExcavators.length === 0) {
    return (
      <div className="App">
        <div className="container">
          <div className="no-selection">
            <h2>{t('excavatorComparePage.noSelection', locale) || 'No excavators selected'}</h2>
            <button className="back-button" onClick={() => navigate('/excavators')}>
              {t('excavatorComparePage.goBack', locale) || 'Go Back'}
            </button>
          </div>
        </div>
      </div>
    );
  }

  return (
    <>
      <ComparePageHeader
        titleKey="excavatorComparePage.title"
        backToListKey="excavatorComparePage.backToList"
        backPath="/excavators"
        locale={locale}
      />

      <div className="container">
        <div className="comparison-grid">
          {selectedExcavators.map(excavator => (
            <div key={excavator.id} className="excavator-detail-card">
              <div className="excavator-detail-header">
                <h2 className="excavator-detail-name">
                  {excavator.manufacturer} {excavator.model} {excavator.scale}
                </h2>
                <p className="excavator-detail-manufacturer">{excavator.name}</p>
              </div>
              
              <div className="excavator-images-section">
                <ImageView
                  imageUrl={excavator.imageFront}
                  altText={`${excavator.manufacturer} ${excavator.model} front`}
                  titleKey="excavatorComparePage.frontView"
                  locale={locale}
                  imageClassName="excavator-image"
                />
                <ImageView
                  imageUrl={excavator.imageSide}
                  altText={`${excavator.manufacturer} ${excavator.model} side`}
                  titleKey="excavatorComparePage.sideView"
                  locale={locale}
                  imageClassName="excavator-image"
                />
                <ImageView
                  imageUrl={excavator.imageTop}
                  altText={`${excavator.manufacturer} ${excavator.model} top`}
                  titleKey="excavatorComparePage.topView"
                  locale={locale}
                  imageClassName="excavator-image"
                />
                <ImageView
                  imageUrl={excavator.imageDetail}
                  altText={`${excavator.manufacturer} ${excavator.model} detail`}
                  titleKey="excavatorComparePage.detailView"
                  locale={locale}
                  imageClassName="excavator-image"
                />
              </div>

              <div className="excavator-detail-specs">
                {fields.map(field => {
                  const value = excavator[field.key];
                  
                  // Always show the field, even if empty (to maintain alignment)
                  return (
                    <SpecRow
                      key={field.key}
                      label={field.label}
                      value={formatExcavatorValue(field.key, value, locale, currency, units)}
                    />
                  );
                })}
              </div>
            </div>
          ))}
        </div>
      </div>
    </>
  );
}

export default ExcavatorComparePage;
