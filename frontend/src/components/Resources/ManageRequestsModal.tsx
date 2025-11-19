// src/components/Resources/ManageRequestsModal.tsx
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
                <img src={r.user.profileImage || "/avatar.png"} className="w-10 h-10 rounded-full object-cover" />
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
