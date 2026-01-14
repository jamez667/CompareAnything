import React from 'react';
import { useNavigate } from 'react-router-dom';
import ComparePageHeader from './components/ComparePageHeader';
import ImageView from './components/ImageView';
import SpecRow from './components/SpecRow';
import { useComparePageData } from './hooks/useComparePageData';
import { formatCarValue } from './utils/carFormatters';
import { getCarFields } from './config/carFields';
import { t } from './utils/i18n';
import './CarComparePage.css';

function CarComparePage() {
  const navigate = useNavigate();
  const {
    data: cars,
    selectedItems: selectedCars,
    loading,
    currency,
    units,
    locale,
  } = useComparePageData('/api/cars');

  // Generate fields based on current settings
  const fields = getCarFields(locale, units, currency);

  if (loading) {
    return <div className="loading">{t('table.allCars', locale)}...</div>;
  }

  if (selectedCars.length === 0) {
    return (
      <div className="App">
        <div className="container">
          <div className="no-selection">
            <h2>{t('carComparePage.noSelection', locale)}</h2>
            <button className="back-button" onClick={() => navigate('/cars')}>
              {t('carComparePage.goBack', locale)}
            </button>
          </div>
        </div>
      </div>
    );
  }

  return (
    <>
      <ComparePageHeader
        titleKey="carComparePage.title"
        backToListKey="carComparePage.backToList"
        backPath="/cars"
        locale={locale}
      />

      <div className="container">
        <div className="comparison-grid">
          {selectedCars.map(car => (
            <div key={car.id} className="car-detail-card">
              <div className="car-detail-header">
                <h2 className="car-detail-name">
                  {car.model || car.name || 'Unknown Model'} {car.trim && car.trim.trim() !== '' ? car.trim : 'Base'}{car.variant ? ` ${car.variant}` : ''}
                </h2>
                <p className="car-detail-manufacturer">{car.manufacturer} {car.year}</p>
              </div>
              
              <div className="car-images-section">
                <ImageView
                  imageUrl={car.imageFront}
                  altText={`${car.model}${car.trim ? ` ${car.trim}` : ''} front`}
                  titleKey="carComparePage.frontView"
                  locale={locale}
                  imageClassName="car-image"
                />
                <ImageView
                  imageUrl={car.imageSide}
                  altText={`${car.model}${car.trim ? ` ${car.trim}` : ''} side`}
                  titleKey="carComparePage.sideView"
                  locale={locale}
                  imageClassName="car-image"
                />
                <ImageView
                  imageUrl={car.imageInterior}
                  altText={`${car.model}${car.trim ? ` ${car.trim}` : ''} interior`}
                  titleKey="carComparePage.interiorView"
                  locale={locale}
                  imageClassName="car-image"
                />
                <ImageView
                  imageUrl={car.imageExterior}
                  altText={`${car.model}${car.trim ? ` ${car.trim}` : ''} exterior`}
                  titleKey="carComparePage.exteriorView"
                  locale={locale}
                  imageClassName="car-image"
                />
              </div>

              <div className="car-detail-specs">
                {fields.map(field => {
                  // For computed fields like range, calculate the value
                  let value = car[field.key];
                  if (field.isComputed && field.key === 'range') {
                    value = 'computed'; // Placeholder, will be calculated in formatter
                  }
                  
                  // Conditionally show/hide fields based on fuel type
                  const isElectric = car.fuelType === 'Electric';
                  const isHybrid = car.fuelType === 'Hybrid' || car.fuelType === 'Plug-in Hybrid';
                  
                  // For electric: show battery capacity, hide fuel capacity
                  if (field.key === 'fuelCapacity' && isElectric && !isHybrid) {
                    return null;
                  }
                  
                  // For gas: show fuel capacity, hide battery capacity
                  if (field.key === 'batteryCapacity' && !isElectric && !isHybrid) {
                    return null;
                  }
                  
                  // Skip if value is null/undefined/empty (unless it's a computed field or trim field)
                  // Trim field should show "Base" when empty, so don't skip it
                  if (!field.isComputed && field.key !== 'trim' && (value === null || value === undefined || value === '' || value === 0)) {
                    return null;
                  }
                  
                  // Adjust label for electric cars - show energy consumption labels
                  let displayLabel = field.label;
                  if (isElectric && (field.key === 'mpgCity' || field.key === 'mpgHighway' || field.key === 'mpgCombined')) {
                    const baseLabel = field.key === 'mpgCity' ? t('table.energyCity', locale) || 'City Energy Consumption' :
                                     field.key === 'mpgHighway' ? t('table.energyHighway', locale) || 'Highway Energy Consumption' :
                                     t('table.energyCombined', locale) || 'Combined Energy Consumption';
                    const unit = units === 'metric' ? 'kWh/100km' : 'MPGe';
                    displayLabel = `${baseLabel} (${unit})`;
                  }
                  
                  return (
                    <SpecRow
                      key={field.key}
                      label={displayLabel}
                      value={formatCarValue(field.key, value, locale, currency, units, car)}
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

export default CarComparePage;
