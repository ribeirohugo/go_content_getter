import React, { useState } from "react";

export default function DownloadURLsView({ apiUrl }) {
  const [urls, setUrls] = useState("");
  const [loading, setLoading] = useState(false);
  const [result, setResult] = useState(null);
  const [error, setError] = useState("");

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
      const payload = { urls: urlList };
      const res = await fetch(`${apiUrl}/download-urls`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(payload),
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
    <div className="cg-card">
      <div className="cg-title">Download URLs</div>
      <form onSubmit={handleSubmit} autoComplete="off">
        <label className="cg-label">URLs (one per line):</label>
        <textarea
          className="cg-textarea"
          rows={6}
          value={urls}
          onChange={(e) => setUrls(e.target.value)}
          placeholder={`https://example.com/file1.jpg\nhttps://example.com/file2.png`}
        />

        <button className="cg-btn" type="submit" disabled={loading}>
          {loading ? "Downloading..." : "Download URLs"}
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
  );
}

