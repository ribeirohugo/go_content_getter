import React, { useEffect, useState } from "react";

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
      const res = await fetch(`${API_URL}/download`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          urls: urlList,
          contentPattern,
          titlePattern,
        }),
      });
      const data = await res.json();
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
    <div style={{ maxWidth: 600, margin: "40px auto", fontFamily: "sans-serif" }}>
      <h2>Download Many Targets</h2>
      <form onSubmit={handleSubmit}>
        <div style={{ marginBottom: 16 }}>
          <label>URLs (one per line):</label>
          <br />
          <textarea
            rows={5}
            style={{ width: "100%" }}
            value={urls}
            onChange={(e) => setUrls(e.target.value)}
            placeholder="https://example.com/page1\nhttps://example.com/page2"
          />
        </div>
        <div style={{ marginBottom: 16 }}>
          <label>Content Pattern:</label>
          <br />
          <input
            type="text"
            style={{ width: "100%" }}
            value={contentPattern}
            onChange={(e) => setContentPattern(e.target.value)}
          />
        </div>
        <div style={{ marginBottom: 16 }}>
          <label>Title Pattern:</label>
          <br />
          <input
            type="text"
            style={{ width: "100%" }}
            value={titlePattern}
            onChange={(e) => setTitlePattern(e.target.value)}
          />
        </div>
        <button type="submit" disabled={loading}>
          {loading ? "Downloading..." : "Download"}
        </button>
      </form>
      {error && <div style={{ color: "red", marginTop: 16 }}>{error}</div>}
      {result && (
        <div style={{ marginTop: 24 }}>
          <h3>Results</h3>
          {result.length === 0 ? (
            <div>No files found.</div>
          ) : (
            <ul>
              {result.map((file, idx) => (
                <li key={idx}>
                  <strong>{file.Filename || "Untitled"}</strong>
                  {file.url && (
                    <span> - <a href={file.url} target="_blank" rel="noopener noreferrer">Download</a></span>
                  )}
                </li>
              ))}
            </ul>
          )}
        </div>
      )}
    </div>
  );
}

export default App;
