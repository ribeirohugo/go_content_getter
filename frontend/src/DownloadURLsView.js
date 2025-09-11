import React, { useState } from "react";
import Help from "./Help";

export default function DownloadURLsView({ apiUrl }) {
  const [urls, setUrls] = useState("");
  const [loading, setLoading] = useState(false);
  const [result, setResult] = useState(null);
  const [error, setError] = useState("");
  const [store, setStore] = useState(true);

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
      const payload = { urls: urlList, Store: store };
      const res = await fetch(`${apiUrl}/download-urls`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(payload),
      });

      if (!res.ok) {
        // try parse json error
        try {
          const errData = await res.json();
          setError(errData.error || "Unknown error");
        } catch (e) {
          setError("Server error");
        }
      } else {
        const contentType = (res.headers.get("content-type") || "").toLowerCase();
        if (contentType.includes("application/zip") || contentType.includes("application/octet-stream")) {
          // binary zip - download
          const arr = await res.arrayBuffer();
          const blob = new Blob([arr], { type: "application/zip" });
          const url = window.URL.createObjectURL(blob);
          const a = document.createElement("a");
          a.href = url;
          a.download = "files.zip";
          document.body.appendChild(a);
          a.click();
          a.remove();
          window.URL.revokeObjectURL(url);
          setResult([]);
        } else {
          const data = await res.json();
          setResult(data.files || []);
        }
      }
    } catch (err) {
      setError("Network error");
    }

    setLoading(false);
  };

  return (
    <div className="cg-card">
      <div className="cg-title">Download from URLs</div>
      <form onSubmit={handleSubmit} autoComplete="off">
        <label className="cg-label">URLs (one per line): <Help text={"Enter one direct file URL per line.\nEach URL will be downloaded and returned as a file or zip."} /></label>
        <textarea
          className="cg-textarea"
          rows={6}
          value={urls}
          onChange={(e) => setUrls(e.target.value)}
          placeholder={`https://example.com/file1.jpg\nhttps://example.com/file2.png`}
        />

        <button className="cg-btn" type="submit" disabled={loading}>
          {loading ? "Downloading..." : "Download"}
        </button>
        <label style={{ display: 'block', marginTop: '8px' }}>
          <input
            type="checkbox"
            checked={store}
            onChange={(e) => setStore(e.target.checked)}
            style={{ marginRight: '6px' }}
          />
          Store locally? <Help text={"If checked, files will be stored on the server and returned as links.\nOtherwise they may be returned as a zip for direct download."} />
        </label>
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
