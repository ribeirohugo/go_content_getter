import React, { useState, useEffect } from 'react';
import Help from './Help';
import { downloadVideo } from './api';

// Video qualities list
const VIDEO_QUALITIES = ['144p','240p','360p','480p','720p','1080p','1440p','2160p'];
// Audio bitrates (kbps)
const AUDIO_BITRATES = ['64','96','128','160','192','256','320'];

export default function DownloadVideoView() {
  const [urls, setUrls] = useState('');
  const [format, setFormat] = useState('mp4');
  const [videoQuality, setVideoQuality] = useState(''); // empty = auto
  const [audioQuality, setAudioQuality] = useState(''); // required always now
  const [store, setStore] = useState(false);
  const [downloading, setDownloading] = useState(false);
  const [error, setError] = useState('');
  const [logs, setLogs] = useState([]);
  const [storedFiles, setStoredFiles] = useState([]);

  const appendLog = (m) => setLogs(l => [...l, m]);

  const handleDownload = async (e) => {
    e && e.preventDefault();
    setError('');
    setLogs([]);
    setStoredFiles([]);

    const list = urls.split('\n').map(x=>x.trim()).filter(Boolean);
    if (!list.length) { setError('Enter at least one URL.'); return; }
    if (format === 'mp4' && !videoQuality) { setError('Select a video quality.'); return; }
    if (!audioQuality) { setError('Select an audio quality.'); return; }

    setDownloading(true);
    for (const u of list) {
      appendLog(`Processing: ${u}`);
      try {
        const payload = {
          urls: [u],
          videoQuality: (format === 'mp4' && videoQuality) ? videoQuality.replace('p','') : '',
          audioQuality: audioQuality,
          format: format,
          store: store,
        };
        const res = await downloadVideo(payload);
        if (!res.ok) {
          let msg = 'Download error';
          try { const d = await res.json(); msg = d.error || msg; } catch(_){}
          appendLog(`Error: ${msg}`); continue;
        }
        if (store) {
          try {
            const data = await res.json();
            const files = data.files || data.Files || [];
            setStoredFiles(prev => [...prev, ...files]);
            appendLog(`Stored (${files.length || 1})`);
          } catch(err){ appendLog('Failed to parse JSON response'); }
        } else {
          const blob = await res.blob();
          const ext = format === 'mp3' ? 'mp3' : 'mp4';
          const fname = `video_${Date.now()}_${Math.random().toString(36).slice(2)}.${ext}`;
          const a = document.createElement('a');
          a.href = URL.createObjectURL(blob); a.download = fname; document.body.appendChild(a); a.click(); a.remove();
          appendLog('Download finished');
        }
      } catch(err) {
        appendLog(`Failed (${u}): ${err.message}`);
      }
    }
    setDownloading(false);
  };

  useEffect(() => {
    if (format === 'mp3' && videoQuality) {
      setVideoQuality('');
    }
  }, [format, videoQuality]);

  return (
    <div className="cg-card">
      <div className="cg-title">Video / Audio Download</div>
      <form onSubmit={handleDownload} autoComplete="off">
        <label className="cg-label">URLs (one per line): <Help text={'Paste video page URLs supported by yt-dlp. Each line is processed sequentially.'} /></label>
        <textarea className="cg-textarea" name="urls" rows={6} value={urls} onChange={(e)=>setUrls(e.target.value)} placeholder={`https://www.youtube.com/watch?v=...\nhttps://vimeo.com/...`} />

        <label className="cg-label" style={{marginTop:8}}>Format</label>
        {/* Format as toggle buttons */}
        <div style={{display:'flex', gap:'8px', marginBottom:'4px'}}>
          <button
            type="button"
            onClick={()=>setFormat('mp4')}
            className="cg-btn"
            style={{
              background: format==='mp4'? '#2563eb' : '#444',
              border: 'none',
              padding: '6px 14px'
            }}
          >mp4</button>
          <button
            type="button"
            onClick={()=>setFormat('mp3')}
            className="cg-btn"
            style={{
              background: format==='mp3'? '#2563eb' : '#444',
              border: 'none',
              padding: '6px 14px'
            }}
          >mp3</button>
        </div>
        {/* Hidden field for potential form compatibility */}
        <input type="hidden" value={format} readOnly />

        {format === 'mp4' && (
          <>
            <label className="cg-label" style={{marginTop:8}}>Video quality (required)</label>
            <select className="cg-input" value={videoQuality} required={format==='mp4'} onChange={(e)=>setVideoQuality(e.target.value)}>
              <option value="" disabled>{'Select quality'}</option>
              {VIDEO_QUALITIES.map(q=> <option key={q} value={q}>{q}</option>)}
            </select>
          </>
        )}

        <label className="cg-label" style={{marginTop:8}}>Audio Quality (required)</label>
        <select className="cg-input" value={audioQuality} required onChange={(e)=>setAudioQuality(e.target.value)}>
          <option value="" disabled>Select quality</option>
          {AUDIO_BITRATES.map(b=> <option key={b} value={b}>{b} kbps</option>)}
        </select>

        <button className="cg-btn" type="submit" disabled={downloading} style={{marginTop:12}}>
          {downloading ? 'Downloading...' : 'Start'}
        </button>

        <label style={{ display: 'block', marginTop: '8px' }}>
          <input type="checkbox" checked={store} onChange={(e)=>setStore(e.target.checked)} style={{marginRight:6}} />
          Store on server? <Help text={'If checked each file is stored server-side and returned as JSON instead of triggering a direct browser download.'} />
        </label>
      </form>

      {error && <div className="cg-error" style={{marginTop:12}}>{error}</div>}

      <div style={{marginTop:16}}>
        <h4>Logs</h4>
        <div style={{background:'#111',color:'#0f0',padding:10,fontFamily:'monospace',fontSize:12,maxHeight:200,overflowY:'auto',borderRadius:4}}>
          {logs.length===0 && <div>(empty)</div>}
          {logs.map((l,i)=><div key={i}>{l}</div>)}
        </div>
      </div>

      {store && storedFiles.length>0 && (
        <div className="cg-results" style={{marginTop:20}}>
          <h3>Stored Files:</h3>
          <ul className="cg-list">
            {storedFiles.map((f,i)=> (
              <li key={i}>
                <a href={f.url || f.URL} target="_blank" rel="noopener noreferrer">{f.Filename || f.filename}</a>
                <strong>{f.size || f.Size}</strong>
              </li>
            ))}
          </ul>
        </div>
      )}
    </div>
  );
}
