import React from 'react';
import SettingsHeader from './SettingsHeader';
import ComparisonTable from './components/ComparisonTable';
import CompareButton from './components/CompareButton';
import { useListPageData } from './hooks/useListPageData';
import { getCarFields } from './config/carFields';
import { formatCarValue, getCarSortValue } from './utils/carFormatters';
import { t } from './utils/i18n';
import './App.css';

function CarsPage() {
  const {
    data: cars,
    loading,
    selectedIds,
    handleSelectionChange,
    fields,
    formatValue,
    locale,
    getSortValue,
    noResultsKey,
  } = useListPageData(
    '/api/cars',
    getCarFields,
    (key, value, item, locale, currency, units) => {
      // For computed fields like range, pass the car object
      if (key === 'range') {
        return formatCarValue(key, 'computed', locale, currency, units, item);
      }
      return formatCarValue(key, value, locale, currency, units, item);
    },
    getCarSortValue,
    'table.noCarsFound'
  );

  if (loading) {
    return <div className="loading">{t('table.allCars', locale)}...</div>;
  }

  return (
    <div className="App">
      <SettingsHeader />
      <header className="App-header">
        <h1>{t('carsPage.title', locale)}</h1>
        <p>{t('carsPage.subtitle', locale)}</p>
      </header>

      <div className="container">
        <div className="lens-table-section">
          <div className="table-header">
            <h2>{t('table.allCars', locale)}</h2>
            <CompareButton
              selectedCount={selectedIds.size}
              comparePath="/cars/compare"
              selectedIds={selectedIds}
              locale={locale}
            />
          </div>
          
          <ComparisonTable
            data={cars}
            fields={fields}
            formatValue={formatValue}
            onSelectionChange={handleSelectionChange}
            selectedIds={selectedIds}
            locale={locale}
            getSortValue={getSortValue}
            noResultsKey={noResultsKey}
          />
        </div>
      </div>
    </div>
  );
}

export default CarsPage;
