import React from "react";

// Help component: shows a (?) icon with a tooltip on hover or focus
function Help({ text }) {
  const lines = (text || "").split('\n');
  return (
    <span className="cg-help" tabIndex={0} aria-label={text}>
      <span className="cg-help-dot">?</span>
      <span className="cg-help-tip">
        {lines.map((line, i) => (
          <span key={i}>
            {line}
            {i < lines.length - 1 ? <br /> : null}
          </span>
        ))}
      </span>
    </span>
  );
}

export default Help;
