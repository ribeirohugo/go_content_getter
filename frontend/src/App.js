import React, { useEffect, useState } from "react";
import './App.css';

const API_URL = process.env.REACT_APP_API_URL || "/api";

function App() {
  const [urls, setUrls] = useState("");
  const [patterns, setPatterns] = useState([]);

  const [contentPatternSelect, setContentPatternSelect] = useState("");
  const [contentPatternCustom, setContentPatternCustom] = useState("");

  const [titlePatternSelect, setTitlePatternSelect] = useState("");
  const [titlePatternCustom, setTitlePatternCustom] = useState("");

  const [loading, setLoading] = useState(false);
  const [result, setResult] = useState(null);
  const [error, setError] = useState("");

  useEffect(() => {
    fetch(`${API_URL}/patterns`)
      .then((res) => res.json())
      .then((data) => {
        // data is expected to be an array of { Description, Regex }
        setPatterns(Array.isArray(data) ? data : []);
        if (Array.isArray(data) && data.length > 0) {
          const firstRegex = data[0].Regex || data[0].regex || "";
          setContentPatternSelect(firstRegex);
          setTitlePatternSelect(firstRegex);
        }
      })
      .catch(() => {});
  }, []);

  const handleSubmit = async (e) => {
    e.preventDefault();
    setLoading(true);
    setError("");
    setResult(null);
    const urlList = urls
      .split("\n")
      .map((u) => u.trim())
      .filter((u) => u);
    if (urlList.length === 0) {
      setError("Please enter at least one URL.");
      setLoading(false);
      return;
    }

    // Decide final patterns: if select is '__custom__' use custom value, otherwise use selected regex
    const finalContentPattern = contentPatternSelect === "__custom__" ? contentPatternCustom : contentPatternSelect;
    const finalTitlePattern = titlePatternSelect === "__custom__" ? titlePatternCustom : titlePatternSelect;

    try {
      const payload = {
        urls: urlList,
        contentPattern: finalContentPattern,
        titlePattern: finalTitlePattern,
      };
      console.log('Request payload:', payload);
      const res = await fetch(`${API_URL}/download-and-store`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(payload),
      });
      const data = await res.json();
      console.log('Response:', data);
      if (!res.ok) {
        setError(data.error || "Unknown error");
      } else {
        setResult(data.files || []);
      }
    } catch (err) {
      setError("Network error");
    }
    setLoading(false);
  };

  return (
    <>
      <nav className="cg-navbar">
        <div className="cg-navbar-content">
          <img src="/logo.svg" alt="CG Logo" className="cg-logo" />
          <a href="/" className="cg-navbar-link">Home</a>
        </div>
      </nav>
      <div className="cg-content">
        <div className="cg-card">
          <div className="cg-title">Download Content</div>
          <form onSubmit={handleSubmit} autoComplete="off">
            <label className="cg-label">URLs (one per line):</label>
            <textarea
              className="cg-textarea"
              rows={5}
              value={urls}
              onChange={(e) => setUrls(e.target.value)}
              placeholder={`https://example.com/page1\nhttps://example.com/page2`}
            />

            <label className="cg-label">Content Pattern:</label>
            <select
              className="cg-input"
              value={contentPatternSelect}
              onChange={(e) => setContentPatternSelect(e.target.value)}
            >
              {patterns.map((p, i) => (
                <option key={i} value={p.Regex || p.regex}>
                  {`${p.Description || p.description}: ${p.Regex || p.regex}`}
                </option>
              ))}
              <option value="__custom__">Custom...</option>
            </select>
            {contentPatternSelect === "__custom__" && (
              <input
                className="cg-input"
                type="text"
                value={contentPatternCustom}
                onChange={(e) => setContentPatternCustom(e.target.value)}
                placeholder="Enter custom regex"
              />
            )}

            <label className="cg-label">Title Pattern:</label>
            <select
              className="cg-input"
              value={titlePatternSelect}
              onChange={(e) => setTitlePatternSelect(e.target.value)}
            >
              {patterns.map((p, i) => (
                <option key={i} value={p.Regex || p.regex}>
                  {`${p.Description || p.description}: ${p.Regex || p.regex}`}
                </option>
              ))}
              <option value="__custom__">Custom...</option>
            </select>
            {titlePatternSelect === "__custom__" && (
              <input
                className="cg-input"
                type="text"
                value={titlePatternCustom}
                onChange={(e) => setTitlePatternCustom(e.target.value)}
                placeholder="Enter custom regex"
              />
            )}

            <button className="cg-btn" type="submit" disabled={loading}>
              {loading ? "Downloading..." : "Download"}
            </button>
          </form>
          {error && <div className="cg-error">{error}</div>}
          {result && (
            <div className="cg-results">
              <h3>Results:</h3>
              <ul className="cg-list">
                {result.map((file, index) => (
                  <li key={index}>
                    <a href={file.url} target="_blank" rel="noopener noreferrer">
                      {file.Filename}
                    </a>
                    <strong>{file.size}</strong>
                  </li>
                ))}
              </ul>
            </div>
          )}
        </div>
      </div>
    </>
  );
}

export default App;
