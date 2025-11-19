const GO_BASE = "http://localhost:8080";

async function api(path: string, options: RequestInit = {}) {
  const res = await fetch(`${GO_BASE}${path}`, {
    ...options,
    headers: {
      "Content-Type": "application/json",
      ...(options.headers || {})
    }
  });

  if (!res.ok) {
    let err = {};
    try {
      err = await res.json();
    } catch {}
    throw new Error(`API error ${res.status}`);
  }

  return res.json();
}

// ---------- USERS ----------
export function getUserInfo(address: string) {
  return api(`/user/${address}`);
}

export function verifyUserTx(payload: any) {
  return api(`/user/verify`, {
    method: "POST",
    body: JSON.stringify(payload),
  });
}

// ---------- CONTENT ----------
export function getContent(id: number) {
  return api(`/content/${id}`);
}

export function listPublicContent() {
  return api(`/content/public`);
}

export function getUserUploads(address: string) {
  return api(`/content/user/${address}`);
}

export function uploadContent(payload: any) {
  return api(`/content/upload`, {
    method: "POST",
    body: JSON.stringify(payload),
  });
}

export function checkAccess(contentId: number, viewer: string) {
  return api(`/content/access?contentId=${contentId}&viewer=${viewer}`);
}

export function grantContent(payload: any) {
  return api(`/content/grant`, {
    method: "POST",
    body: JSON.stringify(payload),
  });
}

export function revokeContent(payload: any) {
  return api(`/content/revoke`, {
    method: "POST",
    body: JSON.stringify(payload),
  });
}

// ---------- MENTORSHIP ----------
export function requestSession(payload: any) {
  return api(`/mentorship/request`, {
    method: "POST",
    body: JSON.stringify(payload),
  });
}

export function acceptSession(payload: any) {
  return api(`/mentorship/accept`, {
    method: "POST",
    body: JSON.stringify(payload),
  });
}

export function completeSession(payload: any) {
  return api(`/mentorship/complete`, {
    method: "POST",
    body: JSON.stringify(payload),
  });
}

export function giveFeedback(payload: any) {
  return api(`/mentorship/feedback`, {
    method: "POST",
    body: JSON.stringify(payload),
  });
}

// ---------- REPUTATION ----------
export function getReputation(address: string) {
  return api(`/reputation/${address}`);
}

export function addReputation(payload: any) {
  return api(`/reputation/add`, {
    method: "POST",
    body: JSON.stringify(payload),
  });
}

export function subReputation(payload: any) {
  return api(`/reputation/sub`, {
    method: "POST",
    body: JSON.stringify(payload),
  });
}
