"use client";

import React from "react";
import { useEffect, useState } from "react";
import ColorBends from "@/components/BackgroundAnimations/ColorBends";
import Navbar from "@/components/UI/Navbar";
import { useRouter } from "next/navigation";
import OtpModal from "@/components/Auth/OtpModal";

export default function ResourceClient({ id }: { id: string }) {
  const router = useRouter();

  const [viewer, setViewer] = useState<any>(null);
  const [resource, setResource] = useState<any>(null);
  const [loading, setLoading] = useState(true);
  const [requesting, setRequesting] = useState(false);
  const [respondLoading, setRespondLoading] = useState(false);
  const [showOtp, setShowOtp] = useState(false);

  useEffect(() => {
    const load = async () => {
      const me = await fetch("/api/auth/me", { credentials: "include" })
        .then(r => r.json())
        .catch(() => ({}));

      if (me.user) setViewer(me.user);

      const url = `/api/resources/${id}${me.user ? `?viewerId=${me.user._id}` : ""}`;
      const res = await fetch(url);
      const data = await res.json();

      if (data?.resource) setResource(data.resource);
      setLoading(false);
    };

    load();
  }, [id]);

  const askAccess = async () => {
    if (!viewer) return setShowOtp(true);
    if (resource.allowed) return;

    setRequesting(true);
    const res = await fetch("/api/resources/request", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ resourceId: id, userId: viewer._id, message: "" })
    });

    const data = await res.json();
    setRequesting(false);

    if (data?.success) {
      setResource((r: any) => ({ ...r, requestStatus: data.request?.status || "pending" }));
    }
  };

  const handleRespond = async (requestId: string, action: "approve" | "reject") => {
    setRespondLoading(true);

    const res = await fetch("/api/resources/respond", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ resourceId: id, requestId, action, ownerId: viewer._id })
    });

    const data = await res.json();
    setRespondLoading(false);

    if (data?.success) {
      const updated = await fetch(`/api/resources/${id}?viewerId=${viewer._id}`).then(r => r.json());
      setResource(updated.resource);
    }
  };

  if (loading) {
    return <div className="text-white min-h-screen flex items-center justify-center">Loading...</div>;
  }

  return (
    <div className="relative min-h-screen w-full overflow-x-hidden text-white">
      
      <div className="fixed inset-0 -z-30">
        <ColorBends
          colors={['#3e47f4', '#06b31a', '#b46d04']}
          rotation={0}
          speed={0.3}
          scale={1}
          frequency={1}
          warpStrength={1}
          mouseInfluence={1}
          parallax={0.5}
          noise={0.1}
        />
        <div
          className="absolute inset-0 pointer-events-none"
          style={{
            background:
              'linear-gradient(to bottom, rgba(0,0,0,0.25) 0%, rgba(0,0,0,0.05) 25%, rgba(0,0,0,0.10) 70%, rgba(0,0,0,0.15) 100%)',
          }}
        />
      </div>

      <Navbar user={viewer} onLoginClick={() => router.push("/")} />

      <main className="pt-24 max-w-3xl mx-auto px-6">
        <div className="bg-white/5 backdrop-blur-lg border border-white/10 rounded-2xl p-6">

          <h1 className="text-2xl font-bold">{resource.title}</h1>
          <p className="text-sm text-gray-300 mt-2">{resource.description}</p>

          <div className="mt-4 flex gap-2 flex-wrap">
            <div className="text-xs px-2 py-1 rounded bg-white/10">{resource.category}</div>
            {resource.tags?.map((t: string, i: number) => (
              <div key={i} className="text-xs px-2 py-1 rounded bg-white/10">#{t}</div>
            ))}
            <div className="text-xs px-2 py-1 rounded bg-white/10">Subject: {resource.subject}</div>
          </div>

          <div className="mt-6 flex items-center gap-3">
            <img src={resource.owner.profileImage} className="w-12 h-12 rounded-full object-cover" />
            <div>
              <div className="font-semibold">{resource.owner.fullName}</div>
              <div className="text-xs text-gray-400">Owner</div>
            </div>
          </div>

          <div className="mt-6">
            {resource.allowed ? (
              <a
                href={resource.fileUrl}
                target="_blank"
                rel="noreferrer"
                className="inline-block px-4 py-2 bg-indigo-600 rounded text-white"
              >
                View / Download
              </a>
            ) : viewer && viewer._id === resource.owner._id ? (
              <>
                <div className="text-sm text-gray-300">Requests</div>
                <div className="mt-2 space-y-2">
                  {(resource.requests || [])
                    .filter((r: any) => r.status === "pending")
                    .map((r: any) => (
                      <div key={r._id} className="flex items-center justify-between bg-white/6 p-2 rounded">
                        <div>
                          <div className="font-medium">{r.user}</div>
                          <div className="text-xs text-gray-400">{r.message}</div>
                        </div>
                        <div className="flex gap-2">
                          <button
                            onClick={() => handleRespond(r._id, "approve")}
                            disabled={respondLoading}
                            className="px-3 py-1 bg-green-600 rounded"
                          >
                            Approve
                          </button>
                          <button
                            onClick={() => handleRespond(r._id, "reject")}
                            disabled={respondLoading}
                            className="px-3 py-1 bg-red-600 rounded"
                          >
                            Reject
                          </button>
                        </div>
                      </div>
                    ))}

                  {resource.requests.filter((r: any) => r.status === "pending").length === 0 && (
                    <div className="text-xs text-gray-400">No pending requests</div>
                  )}
                </div>
              </>
            ) : (
              <>
                {resource.requestStatus === "pending" ? (
                  <div className="text-sm text-yellow-300">Request pending</div>
                ) : resource.requestStatus === "approved" ? (
                  <div className="text-sm text-green-300">Access approved — refresh to view</div>
                ) : (
                  <button
                    onClick={askAccess}
                    disabled={requesting}
                    className="px-4 py-2 bg-indigo-600 rounded"
                  >
                    {requesting ? "Requesting..." : "Request access"}
                  </button>
                )}
              </>
            )}
          </div>
        </div>
      </main>

      {showOtp && <OtpModal email="" onClose={() => setShowOtp(false)} onVerified={() => setShowOtp(false)} />}
    </div>
  );
}
