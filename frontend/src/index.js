import React from 'react';
import ReactDOM from 'react-dom/client';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import './index.css';
import HomePage from './HomePage';
import App from './App';
import ComparePage from './ComparePage';
import CarsPage from './CarsPage';
import CarComparePage from './CarComparePage';
import ExcavatorsPage from './ExcavatorsPage';
import ExcavatorComparePage from './ExcavatorComparePage';

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<HomePage />} />
        <Route path="/lenses" element={<App />} />
        <Route path="/lenses/compare" element={<ComparePage />} />
        <Route path="/cars" element={<CarsPage />} />
        <Route path="/cars/compare" element={<CarComparePage />} />
        <Route path="/excavators" element={<ExcavatorsPage />} />
        <Route path="/excavators/compare" element={<ExcavatorComparePage />} />
      </Routes>
    </BrowserRouter>
  </React.StrictMode>
);


