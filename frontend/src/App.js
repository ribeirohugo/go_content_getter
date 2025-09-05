import React, { useEffect, useState } from "react";
import './App.css';

const API_URL = process.env.REACT_APP_API_URL || "/api";

function App() {
  const [urls, setUrls] = useState("");
  const [contentPattern, setContentPattern] = useState("");
  const [titlePattern, setTitlePattern] = useState("");
  const [loading, setLoading] = useState(false);
  const [result, setResult] = useState(null);
  const [error, setError] = useState("");

  useEffect(() => {
    fetch(`${API_URL}/default-patterns`)
      .then((res) => res.json())
      .then((data) => {
        setContentPattern(data.contentPattern || "");
        setTitlePattern(data.titlePattern || "");
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
    try {
      const payload = {
        urls: urlList,
        contentPattern,
        titlePattern,
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
            <input
              className="cg-input"
              type="text"
              value={contentPattern}
              onChange={(e) => setContentPattern(e.target.value)}
            />
            <label className="cg-label">Title Pattern:</label>
            <input
              className="cg-input"
              type="text"
              value={titlePattern}
              onChange={(e) => setTitlePattern(e.target.value)}
            />
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
