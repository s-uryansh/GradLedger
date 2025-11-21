const GO_BASE = "http://localhost:8080";

async function api(path: string, options: RequestInit = {}) {
  const res = await fetch(`${GO_BASE}${path}`, {
    ...options,
    headers: {
      "Content-Type": "application/json",
      ...(options.headers || {})
    }
  });
  return res.json();
}

export function upvote(walletAddress: string) {
  return api("/reputation/add", {
    method: "POST",
    body: JSON.stringify({
      mentor: walletAddress,
      amount: 1
    })
  });
}

export function downvote(walletAddress: string) {
  return api("/reputation/sub", {
    method: "POST",
    body: JSON.stringify({
      mentor: walletAddress,
      amount: 1
    })
  });
}
