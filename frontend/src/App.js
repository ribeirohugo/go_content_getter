import React, { useEffect, useState } from "react";
import './App.css';
import DownloadURLsView from './DownloadURLsView';
import DownloadContentView from './DownloadContentView';
import DownloadYoutubeView from './DownloadYoutubeView';
import DownloadVideoView from './DownloadVideoView';

const API_URL = process.env.REACT_APP_API_URL || "/api";

function App() {
  const [view, setView] = useState('content');
  const [urls, setUrls] = useState("");
  const [patterns, setPatterns] = useState([]);
  const [navOpen, setNavOpen] = useState(false);

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

          // prefer "Image from src attribute" for content default when available
          const imgSrc = data.find(
            (p) => (p.Description || p.description || "") === "Image from src attribute"
          );
          const contentDefault = (imgSrc && (imgSrc.Regex || imgSrc.regex)) || firstRegex;

          // prefer "HTML title" for title default when available
          const htmlTitle = data.find(
            (p) => (p.Description || p.description || "") === "HTML title"
          );
          const titleDefault = (htmlTitle && (htmlTitle.Regex || htmlTitle.regex)) || firstRegex;

          setContentPatternSelect(contentDefault);
          setTitlePatternSelect(titleDefault);
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
          <div className={`cg-navbar-links ${navOpen ? 'open' : ''}`}>
            <a href="#" className="cg-navbar-link" onClick={(e)=>{e.preventDefault(); setView('content'); setNavOpen(false);}}>Content</a>
            <a href="#" className="cg-navbar-link" onClick={(e)=>{e.preventDefault(); setView('download-urls'); setNavOpen(false);}}>URLs</a>
            <a href="#" className="cg-navbar-link" onClick={(e)=>{e.preventDefault(); setView('video'); setNavOpen(false);}}>Video</a>
            <a href="#" className="cg-navbar-link" onClick={(e)=>{e.preventDefault(); setView('youtube'); setNavOpen(false);}}>YouTube</a>
          </div>

          <button className="cg-trigram" aria-label="Menu" onClick={() => setNavOpen(!navOpen)}>
            <svg width="28" height="18" viewBox="0 0 28 18" fill="none" xmlns="http://www.w3.org/2000/svg">
              <rect x="0" y="1" width="28" height="2" rx="1" fill="#FFFFFF" />
              <rect x="0" y="8" width="28" height="2" rx="1" fill="#FFFFFF" />
              <rect x="0" y="15" width="28" height="2" rx="1" fill="#FFFFFF" />
            </svg>
          </button>
        </div>
      </nav>
      <div className="cg-content">
        {view === 'download-urls' ? (
          <DownloadURLsView apiUrl={API_URL} />
        ) : view === 'youtube' ? (
          <DownloadYoutubeView apiUrl={API_URL} />
        ) : view === 'video' ? (
          <DownloadVideoView apiUrl={API_URL} />
        ) : (
          <DownloadContentView apiUrl={API_URL} />
        )}
      </div>
    </>
  );
}

export default App;
