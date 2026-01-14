import React from 'react';
import './SpecRow.css';

/**
 * Shared spec row component for compare pages
 * @param {string} label - Label for the spec
 * @param {string|ReactNode} value - Value to display
 */
function SpecRow({ label, value }) {
  return (
    <div className="spec-row">
      <span className="spec-label">{label}:</span>
      <span className="spec-value">{value}</span>
    </div>
  );
}

export default SpecRow;
