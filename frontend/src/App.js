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
        @import url('https://fonts.googleapis.com/css2?family=Inter:wght@400;500;600;700&display=swap');
        html, body {
          font-family: 'Inter', 'Segoe UI', 'Roboto', Arial, sans-serif;
          background: #f6f8fa;
          color: #23243a;
          letter-spacing: 0.01em;
          font-size: 17px;
        }
        .cg-navbar {
          position: fixed;
          top: 0;
          left: 0;
          width: 100%;
          height: 64px;
          background: #23243a;
          display: flex;
          align-items: center;
          box-shadow: 0 4px 16px 0 rgba(34,34,59,0.07);
          z-index: 1000;
        }
        .cg-navbar-content {
          width: 100%;
          max-width: 1200px;
          margin: 0 auto;
          display: flex;
          align-items: center;
          padding: 0 32px;
        }
        .cg-navbar-logo {
          font-size: 1.35rem;
          font-weight: 700;
          color: #fff;
          letter-spacing: 0.04em;
          margin-right: 36px;
          text-shadow: 0 2px 8px rgba(34,34,59,0.10);
          user-select: none;
        }
        .cg-navbar-link {
          color: #e0e0f0;
          text-decoration: none;
          font-size: 1.08rem;
          font-weight: 500;
          margin-right: 24px;
          padding: 8px 14px;
          border-radius: 6px;
          transition: background 0.18s, color 0.18s;
        }
        .cg-navbar-link:hover {
          background: #35355a;
          color: #5f6fff;
        }
        .cg-card {
          background: #fff;
          max-width: 800px;
          margin: 56px auto 0 auto;
          border-radius: 18px;
          box-shadow: 0 4px 24px 0 rgba(0,0,0,0.08);
          padding: 40px 32px 32px 32px;
        }
        .cg-title {
          font-size: 2.2rem;
          font-weight: 700;
          color: #23243a;
          margin-bottom: 28px;
          text-align: center;
          letter-spacing: 0.02em;
        }
        .cg-label {
          font-weight: 600;
          color: #3d3e5a;
          margin-bottom: 7px;
          display: block;
          font-size: 1.04rem;
          letter-spacing: 0.01em;
        }
        .cg-input, .cg-textarea {
          box-sizing: border-box;
          width: 100%;
          border: 1.5px solid #c9c9c9;
          border-radius: 8px;
          padding: 12px 14px;
          font-size: 1.04rem;
          margin-bottom: 20px;
          background: #f8f9fa;
          transition: border 0.2s, box-shadow 0.2s;
          display: block;
          font-family: inherit;
        }
        .cg-input:focus, .cg-textarea:focus {
          border: 1.5px solid #5f6fff;
          outline: none;
          background: #fff;
          box-shadow: 0 2px 8px 0 rgba(95,111,255,0.08);
        }
        .cg-btn {
          width: 100%;
          background: linear-gradient(90deg, #5f6fff 0%, #6c63ff 100%);
          color: #fff;
          border: none;
          border-radius: 8px;
          padding: 14px 0;
          font-size: 1.13rem;
          font-weight: 600;
          cursor: pointer;
          box-shadow: 0 2px 8px 0 rgba(95,111,255,0.08);
          transition: background 0.2s, box-shadow 0.2s;
          display: block;
          letter-spacing: 0.01em;
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
          font-size: 1.04rem;
          text-align: center;
          letter-spacing: 0.01em;
        }
        .cg-results {
          margin-top: 36px;
        }
        .cg-results h3 {
          color: #23243a;
          font-size: 1.18rem;
          margin-bottom: 14px;
          font-weight: 600;
          letter-spacing: 0.01em;
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
          font-size: 1.01rem;
        }
        .cg-list strong {
          color: #23243a;
          font-weight: 600;
        }
        .cg-list a {
          color: #5f6fff;
          text-decoration: none;
          font-weight: 500;
          letter-spacing: 0.01em;
        }
        @media (max-width: 900px) {
          .cg-card {
            max-width: 98vw;
            padding: 18px 2vw 18px 2vw;
          }
          .cg-navbar-content {
            padding: 0 8px;
          }
        }
        .cg-content {
          margin-top: 80px;
        }
      `}</style>
      <nav className="cg-navbar">
        <div className="cg-navbar-content">
          <span className="cg-navbar-logo">Content Getter</span>
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
      </div>
    </>
  );
}

export default App;
