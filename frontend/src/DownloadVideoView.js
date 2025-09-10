import React, { useState } from "react";

// Help component: shows a small (?) icon with hover/focus tooltip
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

function humanFileSize(bytes) {
  if (!bytes || bytes === 0) return "-";
  const thresh = 1024;
  if (Math.abs(bytes) < thresh) return bytes + " B";
  const units = ["KB", "MB", "GB", "TB"];
  let u = -1;
  do {
    bytes /= thresh;
    ++u;
  } while (Math.abs(bytes) >= thresh && u < units.length - 1);
  return bytes.toFixed(1) + " " + units[u];
}

export default function DownloadVideoView({ apiUrl }) {
  const API_URL = apiUrl || process.env.REACT_APP_API_URL || "/api";
  const [url, setUrl] = useState("");
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");
  const [video, setVideo] = useState(null);
  const [qualities, setQualities] = useState([]);
  const [selectedQuality, setSelectedQuality] = useState("");
  const [formats, setFormats] = useState([]);
  const [selectedFormat, setSelectedFormat] = useState("");
  const [audioFormats, setAudioFormats] = useState([]);
  const [selectedAudioFormat, setSelectedAudioFormat] = useState("");
  const [result, setResult] = useState(null);
  const [store, setStore] = useState(false);

  const fetchInfo = async (e) => {
    e && e.preventDefault();
    setError("");
    setResult(null);
    setVideo(null);
    setFormats([]);
    setQualities([]);
    setSelectedFormat("");
    setSelectedQuality("");

    if (!url) {
      setError("Please enter a URL.");
      return;
    }

    setLoading(true);
    try {
      const res = await fetch(`${API_URL}/youtube/info`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ url }),
      });
      if (!res.ok) {
        try {
          const d = await res.json();
          setError(d.error || "Error fetching video info");
        } catch (e) {
          setError("Server error");
        }
        setLoading(false);
        return;
      }
      const data = await res.json();
      const v = data.video;
      setVideo(v || null);
      const fmts = Array.isArray(v?.formats) ? v.formats : [];
      setFormats(fmts);

      // derive qualities (unique heights) and include audio option
      const qset = new Set();
      fmts.forEach((f) => {
        if (!f || (f.height === 0 || !f.height) && f.acodec && !f.vcodec) {
          qset.add("audio");
        } else if (f.height) {
          qset.add(String(f.height));
        }
      });
      const qarr = Array.from(qset).sort((a, b) => {
        if (a === "audio") return 1;
        if (b === "audio") return -1;
        return Number(b) - Number(a);
      });
      setQualities(qarr);
      setSelectedQuality(qarr[0] || "");

      // derive audio-only formats (have acodec and no video codec)
      const afmts = fmts.filter((f) => f && f.acodec && f.acodec !== 'none' && (!f.vcodec || f.vcodec === 'none' || f.vcodec === ''));
      setAudioFormats(afmts);
      setSelectedAudioFormat(afmts[0] ? (afmts[0].format_id || afmts[0].FormatID || '') : '');
    } catch (err) {
      setError("Network error");
    }
    setLoading(false);
  };

  // when quality changes, pick formats
  React.useEffect(() => {
    if (!selectedQuality) {
      setSelectedFormat("");
      setFormats((f) => f);
      return;
    }
    // formats already loaded in state; filter them
    setSelectedFormat("");
    // no need to change formats list; we'll filter when rendering options
  }, [selectedQuality]);

  const handleFormatChange = (e) => {
    const val = e.target.value;
    setSelectedFormat(val);
    // if chosen format is one of audioFormats, set selectedAudioFormat, otherwise clear it
    const isAudio = audioFormats.find((af) => (af.format_id === val || af.FormatID === val));
    setSelectedAudioFormat(isAudio ? val : '');
  };

  const handleDownload = async (e) => {
    e && e.preventDefault();
    setError("");
    setResult(null);

    if (!selectedFormat) {
      setError("Please select a format to download.");
      return;
    }

    const fmt = formats.find((f) => f.format_id === selectedFormat || f.FormatID === selectedFormat);
    if (!fmt || !fmt.url) {
      setError("Selected format has no direct URL available.");
      return;
    }

    setLoading(true);
    try {
      const payload = {
        // use the URL entered in the form (video page), not the direct format URL
        url: url,
        videoFormat: selectedFormatselectedFormat || (fmt.format_id || fmt.FormatID || ''),
        audioFormat: selectedAudioFormat || '',
        title: video?.title || '',
        store: store,
      };

      const res = await fetch(`${API_URL}/youtube/download`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(payload),
      });

      if (!res.ok) {
        try {
          const d = await res.json();
          setError(d.error || "Error downloading the format");
        } catch (e) {
          setError("Server error");
        }
      } else {
        const contentType = (res.headers.get("content-type") || "").toLowerCase();
        if (contentType.includes("application/zip") || contentType.includes("application/octet-stream")) {
          const arr = await res.arrayBuffer();
          const blob = new Blob([arr], { type: "application/zip" });
          const urlBlob = window.URL.createObjectURL(blob);
          const a = document.createElement("a");
          a.href = urlBlob;
          a.download = `${video?.title || 'video'}.zip`;
          document.body.appendChild(a);
          a.click();
          a.remove();
          window.URL.revokeObjectURL(urlBlob);
          setResult([]);
        } else {
          const d = await res.json();
          setResult(d.files || []);
        }
      }
    } catch (err) {
      setError("Network error");
    }

    setLoading(false);
  };

  const filteredFormats = formats.filter((f) => {
    if (!selectedQuality) return true;
    if (selectedQuality === "audio") {
      return (!f.height || f.height === 0) && f.acodec && (!f.vcodec || f.vcodec === "none");
    }
    return String(f.height) === String(selectedQuality);
  });

  return (
    <div className="cg-card">
      <div className="cg-title">Download YouTube / Video</div>
      <form onSubmit={(e) => { e.preventDefault(); }} autoComplete="off">
        <label className="cg-label">Video URL: <Help text={"Paste the video URL (YouTube or similar) and press 'Get info' to list available formats."} /></label>
        <input
          className="cg-input"
          type="text"
          value={url}
          onChange={(e) => setUrl(e.target.value)}
          placeholder="https://www.youtube.com/watch?v=..."
        />
        <div style={{ marginTop: 8 }}>
          <button className="cg-btn" onClick={fetchInfo} disabled={loading}>{loading ? 'Loading...' : 'Get info'}</button>
        </div>

        {video && (
          <div style={{ marginTop: 12 }}>
            <div><strong>{video.title}</strong> — {video.uploader}</div>

            <label className="cg-label" style={{ marginTop: 8 }}>Quality</label>
            <select className="cg-input" value={selectedQuality} onChange={(e) => setSelectedQuality(e.target.value)}>
              {qualities.map((q, i) => (
                <option key={i} value={q}>{q === 'audio' ? 'Audio only' : q + 'p'}</option>
              ))}
            </select>

            <label className="cg-label" style={{ marginTop: 8 }}>Video Format</label>
            <select className="cg-input" value={selectedFormat} onChange={handleFormatChange}>
              <option value="">-- select --</option>
              {filteredFormats.map((f, i) => (
                <option key={i} value={f.format_id || f.FormatID || i}>
                  {`${f.ext || f.Ext || ''} ${f.height ? f.height + 'p' : ''} ${f.vcodec ? f.vcodec : ''} ${f.acodec ? '(' + f.acodec + ')' : ''} ${humanFileSize(f.filesize)}`}
                </option>
              ))}
            </select>

            {audioFormats && audioFormats.length > 0 && (
              <>
                <label className="cg-label" style={{ marginTop: 8 }}>Audio format</label>
                <select className="cg-input" value={selectedAudioFormat} onChange={(e) => setSelectedAudioFormat(e.target.value)}>
                  <option value="">-- select audio --</option>
                  {audioFormats.map((f, i) => (
                    <option key={i} value={f.format_id || f.FormatID || i}>
                      {`${f.ext || f.Ext || ''} ${f.acodec ? '(' + f.acodec + ')' : ''} ${humanFileSize(f.filesize)}`}
                    </option>
                  ))}
                </select>
              </>
            )}

            <div style={{ marginTop: 12 }}>
              <button className="cg-btn" onClick={handleDownload} disabled={loading || !selectedFormat}>{loading ? 'Downloading...' : 'Download'}</button>
            </div>

            <label style={{ display: 'block', marginTop: '8px' }}>
              <input
                type="checkbox"
                checked={store}
                onChange={(e) => setStore(e.target.checked)}
                style={{ marginRight: '6px' }}
              />
              Store locally? <Help text={"If checked, the server will store the file and return links. Otherwise a direct download will be performed."} />
            </label>

          </div>
        )}

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
