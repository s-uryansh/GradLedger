"use client";

import React from "react";
import toast from "react-hot-toast";

type Req = {
  _id: string;
  user: { _id: string; fullName: string; profileImage?: string };
  message?: string;
  requestedAt?: string;
};

export default function ManageRequestsModal({
  open,
  onClose,
  resourceId,
  ownerId,
  requests,
  onUpdate,
}: {
  open: boolean;
  onClose: () => void;
  resourceId: string;
  ownerId: string;
  requests: Req[];
  onUpdate: () => void;
}) {
  if (!open) return null;

  // deterministic short placeholder (3 letters) from id/name
  const makePlaceholder = (seed?: string, name?: string) => {
    const source = seed || name || Math.random().toString();
    let hash = 0;
    for (let i = 0; i < source.length; i++) hash = (hash * 31 + source.charCodeAt(i)) | 0;
    const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ";
    const a = letters[Math.abs(hash) % letters.length];
    const b = letters[Math.abs(hash >> 8) % letters.length];
    const c = letters[Math.abs(hash >> 16) % letters.length];
    return `${a}${b}${c}`;
  };

  const Avatar = ({ img, seed, name, size = 40 }: { img?: string | null; seed?: string; name?: string; size?: number }) => {
    if (img) return <img src={img} className="w-10 h-10 rounded-full object-cover" />;
    const code = makePlaceholder(seed, name);
    const bgColors = ["#6b7280", "#60a5fa", "#34d399", "#f87171", "#a78bfa", "#f472b6", "#22d3ee"];
    const color = bgColors[(seed ? Math.abs(seed.split('').reduce((s, c) => s + c.charCodeAt(0), 0)) : code.charCodeAt(0)) % bgColors.length];
    return (
      <div style={{ background: color }} className="w-10 h-10 rounded-full flex items-center justify-center text-white font-semibold">
        {code}
      </div>
    );
  };

  const respond = async (requestId: string, action: "approve" | "reject") => {
    try {
      const res = await fetch("/api/resources/respond", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ resourceId, requestId, action, ownerId }),
      });
      const data = await res.json();
      if (!data?.success) throw new Error(data?.error || "Failed");
      toast.success(`${action === "approve" ? "Approved" : "Rejected"}`);
      onUpdate();
    } catch (err: any) {
      toast.error(err.message || "Error");
    }
  };

  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center bg-black/50 p-4">
      <div className="bg-white/5 backdrop-blur rounded-lg w-full max-w-2xl p-4">
        <div className="flex justify-between items-center mb-3">
          <h3 className="text-white font-semibold">Pending Requests</h3>
          <button onClick={onClose} className="text-gray-300">Close</button>
        </div>

        <div className="space-y-3 max-h-80 overflow-y-auto">
          {requests.length === 0 && <div className="text-gray-300">No pending requests</div>}
          {requests.map((r) => (
            <div key={r._id} className="flex items-center justify-between bg-white/6 p-3 rounded">
              <div className="flex items-center gap-3">
                <Avatar img={(r.user as any).profileImage || null} seed={(r.user as any)._id} name={(r.user as any).fullName} />
                <div>
                  <div className="text-sm text-white font-medium">{r.user.fullName}</div>
                  <div className="text-xs text-gray-300">{r.message || "No message"}</div>
                  <div className="text-xs text-gray-400 mt-1">{new Date(r.requestedAt || Date.now()).toLocaleString()}</div>
                </div>
              </div>

              <div className="flex gap-2">
                <button onClick={() => respond(r._id, "approve")} className="px-3 py-1 bg-green-600 rounded text-white">Approve</button>
                <button onClick={() => respond(r._id, "reject")} className="px-3 py-1 bg-red-600 rounded text-white">Reject</button>
              </div>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}
