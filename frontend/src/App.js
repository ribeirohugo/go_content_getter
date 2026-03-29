import React, { useState } from "react";
import './App.css';
import DownloadURLsView from './DownloadURLsView';
import DownloadContentView from './DownloadContentView';
import DownloadYoutubeView from './DownloadYoutubeView';
import DownloadVideoView from './DownloadVideoView';

const API_URL = process.env.REACT_APP_API_URL || "/api";

function App() {
  const [view, setView] = useState('content');
  const [navOpen, setNavOpen] = useState(false);

  return (
    <>
      <nav className="cg-navbar">
        <div className="cg-navbar-content">
          <img src="/logo.svg" alt="CG Logo" className="cg-logo" />
          <div className={`cg-navbar-links ${navOpen ? 'open' : ''}`}>
            <button type="button" className="cg-navbar-link" onClick={() => { setView('content'); setNavOpen(false); }}>Content</button>
            <button type="button" className="cg-navbar-link" onClick={() => { setView('download-urls'); setNavOpen(false); }}>URLs</button>
            <button type="button" className="cg-navbar-link" onClick={() => { setView('video'); setNavOpen(false); }}>Video</button>
            <button type="button" className="cg-navbar-link" onClick={() => { setView('youtube'); setNavOpen(false); }}>YouTube</button>
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
