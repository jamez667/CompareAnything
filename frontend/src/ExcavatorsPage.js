import React from 'react';
import SettingsHeader from './SettingsHeader';
import ComparisonTable from './components/ComparisonTable';
import CompareButton from './components/CompareButton';
import { useListPageData } from './hooks/useListPageData';
import { getExcavatorFields } from './config/excavatorFields';
import { formatExcavatorValue, getExcavatorSortValue } from './utils/excavatorFormatters';
import { t } from './utils/i18n';
import './App.css';

function ExcavatorsPage() {
  const {
    data: excavators,
    loading,
    selectedIds,
    handleSelectionChange,
    fields,
    formatValue,
    locale,
    getSortValue,
    noResultsKey,
  } = useListPageData(
    '/api/excavators',
    getExcavatorFields,
    (key, value, item, locale, currency, units) => formatExcavatorValue(key, value, locale, currency, units),
    getExcavatorSortValue,
    'table.noExcavatorsFound'
  );

  if (loading) {
    return <div className="loading">{t('table.allExcavators', locale) || 'Loading excavators'}...</div>;
  }

  return (
    <div className="App">
      <SettingsHeader />
      <header className="App-header">
        <h1>{t('excavatorsPage.title', locale) || 'RC Excavators'}</h1>
        <p>{t('excavatorsPage.subtitle', locale) || 'Compare RC excavators'}</p>
      </header>

      <div className="container">
        <div className="lens-table-section">
          <div className="table-header">
            <h2>{t('table.allExcavators', locale) || 'All Excavators'}</h2>
            <CompareButton
              selectedCount={selectedIds.size}
              comparePath="/excavators/compare"
              selectedIds={selectedIds}
              locale={locale}
            />
          </div>
          
          <ComparisonTable
            data={excavators}
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

export default ExcavatorsPage;
