import { formatPriceValue, mpgToL100km, mpgeToKwh100km, formatWeightValue, formatLengthValue, formatEmptyValue, gallonsToLiters, litersToCubicFeet, cubicFeetToLiters, mphToKmh } from './sharedFormatters';
import { t } from './i18n';

/**
 * Format a car field value for display
 */
export const formatCarValue = (key, value, locale, currency, units, car = null) => {
  // Trim - show "Base" if empty (check this before the general null check)
  if (key === 'trim') {
    return value && value.toString().trim() !== '' ? value.toString().trim() : 'Base';
  }

  if (value === null || value === undefined || value === '') {
    return '-';
  }

  // Price formatting
  if (key === 'price') {
    return formatPriceValue(value, currency);
  }

  // MPG conversions (data is stored in MPG/MPGe)
  // For electric cars: convert to kWh/100km in metric
  // For other cars: convert to L/100km in metric
  if (key === 'mpgCity' || key === 'mpgHighway' || key === 'mpgCombined') {
    const isElectric = car && car.fuelType === 'Electric';
    
    if (units === 'metric') {
      if (isElectric) {
        const kwh100km = mpgeToKwh100km(value);
        return `${kwh100km.toFixed(1)} kWh/100km`;
      } else {
        const l100km = mpgToL100km(value);
        return `${l100km.toFixed(1)} L/100km`;
      }
    }
    // For imperial units, show MPGe for electric, MPG for others
    return `${value.toFixed(1)} ${isElectric ? 'MPGe' : 'mpg'}`;
  }

  // Weight conversion (data stored in lbs, convert to kg for metric)
  if (key === 'weight') {
    return formatWeightValue(value, units, 'lbs', 0);
  }

  // Dimension conversions (data stored in inches, convert to cm for metric)
  if (key === 'length' || key === 'width' || key === 'height' || key === 'wheelbase') {
    return formatLengthValue(value, units, 'in', 1);
  }

  // Fuel capacity conversion (data stored in gallons, convert to liters for metric)
  if (key === 'fuelCapacity') {
    if (units === 'metric') {
      const liters = gallonsToLiters(value);
      return `${liters.toFixed(1)} L`;
    }
    return `${value.toFixed(1)} gal`;
  }

  // Battery capacity (always in kWh)
  if (key === 'batteryCapacity') {
    return `${value.toFixed(1)} kWh`;
  }

  // Range calculation (computed field)
  if (key === 'range') {
    if (!car) return '-';
    const isElectric = car.fuelType === 'Electric';
    
    let range;
    if (isElectric) {
      // For electric: BatteryCapacity (kWh) ÷ (kWh/100km) × 100 = range in km
      const batteryCapacity = car.batteryCapacity || 0;
      const mpgCombined = car.mpgCombined || 0;
      if (batteryCapacity === 0 || mpgCombined === 0) return '-';
      
      const kwh100km = mpgeToKwh100km(mpgCombined);
      range = (batteryCapacity / kwh100km) * 100; // range in km
      
      if (units === 'metric') {
        return `${Math.round(range)} km`;
      } else {
        // Convert km to miles
        const miles = range / 1.60934;
        return `${Math.round(miles)} mi`;
      }
    } else {
      // For gas/hybrid: FuelCapacity (gallons) × MPGCombined = range in miles
      const fuelCapacity = car.fuelCapacity || 0;
      const mpgCombined = car.mpgCombined || 0;
      if (fuelCapacity === 0 || mpgCombined === 0) return '-';
      
      range = fuelCapacity * mpgCombined; // range in miles
      
      if (units === 'metric') {
        // Convert miles to km
        const km = range * 1.60934;
        return `${Math.round(km)} km`;
      } else {
        return `${Math.round(range)} mi`;
      }
    }
  }

  // Cargo capacity conversion (data stored in cubic feet, convert to liters for metric)
  if (key === 'cargoCapacity') {
    if (units === 'metric') {
      const liters = cubicFeetToLiters(value);
      return `${liters.toFixed(1)} L`;
    }
    return `${value.toFixed(1)} ft³`;
  }

  // Speed conversion (data stored in mph, convert to km/h for metric)
  if (key === 'topSpeed') {
    if (units === 'metric') {
      const kmh = mphToKmh(value);
      return `${Math.round(kmh)} km/h`;
    }
    return `${value} mph`;
  }

  // 0-60 time
  if (key === 'zeroToSixty') {
    return `${value.toFixed(1)}s`;
  }

  // Torque - might need unit conversion (ft-lb to Nm)
  if (key === 'torque') {
    // For now, display as-is. Could add Nm conversion if needed
    return `${value} ft-lb`;
  }

  return value;
};

/**
 * Get sort value for a car field
 */
export const getCarSortValue = (car, key) => {
  const value = car[key];
  
  // Numeric fields
  if (['price', 'weight', 'length', 'width', 'height', 'wheelbase', 
       'horsepower', 'torque', 'mpgCity', 'mpgHighway', 'mpgCombined',
       'fuelCapacity', 'batteryCapacity', 'cargoCapacity', 'topSpeed', 'zeroToSixty',
       'seating', 'year', 'range'].includes(key)) {
    return typeof value === 'number' ? value : parseFloat(value) || 0;
  }
  
  // String fields
  if (typeof value === 'string') {
    return value.toLowerCase();
  }
  
  return value;
};
