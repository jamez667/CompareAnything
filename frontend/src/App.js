import React from 'react';
import SettingsHeader from './SettingsHeader';
import ComparisonTable from './components/ComparisonTable';
import CompareButton from './components/CompareButton';
import { useListPageData } from './hooks/useListPageData';
import { getLensFields } from './config/lensFields';
import { formatLensValue, getLensSortValue } from './utils/lensFormatters';
import { t } from './utils/i18n';
import './App.css';

function App() {
  const {
    data: lenses,
    loading,
    selectedIds,
    handleSelectionChange,
    fields,
    formatValue,
    locale,
    getSortValue,
    noResultsKey,
  } = useListPageData(
    '/api/lenses',
    getLensFields,
    (key, value, item, locale, currency, units) => formatLensValue(key, value, locale, currency, units),
    getLensSortValue,
    'table.noLensesFound'
  );

  if (loading) {
    return <div className="loading">{t('table.allLenses', locale)}...</div>;
  }

  return (
    <div className="App">
      <SettingsHeader />
      <header className="App-header">
        <h1>{t('header.title', locale)}</h1>
        <p>{t('header.subtitle', locale)}</p>
      </header>

      <div className="container">
        <div className="lens-table-section">
          <div className="table-header">
            <h2>{t('table.allLenses', locale)}</h2>
            <CompareButton
              selectedCount={selectedIds.size}
              comparePath="/lenses/compare"
              selectedIds={selectedIds}
              locale={locale}
            />
          </div>
          
          <ComparisonTable
            data={lenses}
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

export default App;
