import React from 'react';
import SettingsHeader from './SettingsHeader';
import HomeCard from './components/HomeCard';
import { getLocale } from './utils/settings';
import { t } from './utils/i18n';
import './HomePage.css';

function HomePage() {
  const locale = getLocale();

  return (
    <div className="App">
      <SettingsHeader />
      <header className="App-header">
        <h1>{t('homePage.title', locale)}</h1>
        <p>{t('homePage.subtitle', locale)}</p>
      </header>

      <div className="container">
        <div className="home-grid">
          <HomeCard
            icon="📷"
            titleKey="homePage.lenses"
            descriptionKey="homePage.lensesDescription"
            path="/lenses"
            locale={locale}
          />
          
          <HomeCard
            icon="🚗"
            titleKey="homePage.cars"
            descriptionKey="homePage.carsDescription"
            path="/cars"
            locale={locale}
          />
          
          <HomeCard
            icon="🚜"
            titleKey="homePage.excavators"
            descriptionKey="homePage.excavatorsDescription"
            path="/excavators"
            locale={locale}
          />
        </div>
      </div>
    </div>
  );
}

export default HomePage;
