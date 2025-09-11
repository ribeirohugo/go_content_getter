// src/api.js
const API_BASE_URL = process.env.REACT_APP_API_URL || "/api";

export const ENDPOINTS = {
  PATTERNS: `${API_BASE_URL}/patterns`,
  DOWNLOAD_AND_STORE: `${API_BASE_URL}/download-and-store`,
  DOWNLOAD_URLS: `${API_BASE_URL}/download-urls`,
  VIDEO_INFO: `${API_BASE_URL}/video/info`,
  VIDEO_DOWNLOAD: `${API_BASE_URL}/video/download`,
  DOWNLOAD_VIDEO: `${API_BASE_URL}/download-video`,
  HELP: `${API_BASE_URL}/help`,
};

export async function fetchPatterns() {
  const res = await fetch(ENDPOINTS.PATTERNS);
  return res.json();
}

export async function downloadAndStore(payload) {
  return fetch(ENDPOINTS.DOWNLOAD_AND_STORE, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(payload),
  });
}

export async function downloadURLsPost(payload) {
  return fetch(ENDPOINTS.DOWNLOAD_URLS, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(payload),
  });
}

export async function fetchVideoInfo(payload) {
  return fetch(ENDPOINTS.VIDEO_INFO, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(payload),
  });
}

// Keep both names to match different views
export async function downloadVideoPost(payload) {
  return fetch(ENDPOINTS.DOWNLOAD_VIDEO, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(payload),
  });
}

export async function downloadVideoAlt(payload) {
  // alternative endpoint some servers use
  return fetch(ENDPOINTS.VIDEO_DOWNLOAD, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(payload),
  });
}

export async function getHelp() {
  const res = await fetch(ENDPOINTS.HELP);
  return res.json();
}
