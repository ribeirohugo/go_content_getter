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
    <>
      <style>{`
        body {
          background: #f6f8fa;
        }
        .cg-card {
          background: #fff;
          max-width: 480px;
          margin: 48px auto;
          border-radius: 18px;
          box-shadow: 0 4px 24px 0 rgba(0,0,0,0.08);
          padding: 32px 28px 28px 28px;
        }
        .cg-title {
          font-size: 2rem;
          font-weight: 700;
          color: #22223b;
          margin-bottom: 24px;
          text-align: center;
        }
        .cg-label {
          font-weight: 500;
          color: #4a4e69;
          margin-bottom: 6px;
          display: block;
        }
        .cg-input, .cg-textarea {
          width: 100%;
          border: 1.5px solid #c9c9c9;
          border-radius: 8px;
          padding: 10px 12px;
          font-size: 1rem;
          margin-bottom: 18px;
          background: #f8f9fa;
          transition: border 0.2s;
        }
        .cg-input:focus, .cg-textarea:focus {
          border: 1.5px solid #5f6fff;
          outline: none;
          background: #fff;
        }
        .cg-btn {
          width: 100%;
          background: linear-gradient(90deg, #5f6fff 0%, #6c63ff 100%);
          color: #fff;
          border: none;
          border-radius: 8px;
          padding: 12px 0;
          font-size: 1.1rem;
          font-weight: 600;
          cursor: pointer;
          box-shadow: 0 2px 8px 0 rgba(95,111,255,0.08);
          transition: background 0.2s, box-shadow 0.2s;
        }
        .cg-btn:disabled {
          background: #bfc5ff;
          cursor: not-allowed;
        }
        .cg-error {
          color: #e63946;
          background: #fff0f3;
          border-radius: 6px;
          padding: 10px 14px;
          margin-top: 18px;
          font-size: 1rem;
          text-align: center;
        }
        .cg-results {
          margin-top: 32px;
        }
        .cg-results h3 {
          color: #22223b;
          font-size: 1.2rem;
          margin-bottom: 12px;
        }
        .cg-list {
          list-style: none;
          padding: 0;
        }
        .cg-list li {
          background: #f4f6fb;
          border-radius: 7px;
          margin-bottom: 10px;
          padding: 10px 14px;
          display: flex;
          align-items: center;
          justify-content: space-between;
        }
        .cg-list strong {
          color: #22223b;
        }
        .cg-list a {
          color: #5f6fff;
          text-decoration: none;
          font-weight: 500;
        }
        @media (max-width: 600px) {
          .cg-card {
            padding: 18px 6vw 18px 6vw;
          }
        }
      `}</style>
      <div className="cg-card">
        <div className="cg-title">Download Many Targets</div>
        <form onSubmit={handleSubmit} autoComplete="off">
          <label className="cg-label">URLs (one per line):</label>
          <textarea
            className="cg-textarea"
            rows={5}
            value={urls}
            onChange={(e) => setUrls(e.target.value)}
            placeholder="https://example.com/page1\nhttps://example.com/page2"
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
                    {file.name}
                  </a>
                  <strong>{file.size}</strong>
                </li>
              ))}
            </ul>
          </div>
        )}
      </div>
    </>
  );
}

export default App;
