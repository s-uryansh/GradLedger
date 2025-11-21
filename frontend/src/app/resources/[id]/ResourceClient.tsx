"use client";

import { useEffect, useState } from "react";
import Navbar from "@/components/UI/Navbar";
import ColorBends from "@/components/BackgroundAnimations/ColorBends";
import { useRouter } from "next/navigation";
import OtpModal from "@/components/Auth/OtpModal";
import VoteBox from "@/components/Resources/VoteBox";
import { addReputation, subReputation } from "@/lib/go"; 

export default function ResourceClient({ id }: { id: string }) {
  const router = useRouter();

  const [viewer, setViewer] = useState<any>(null);
  const [resource, setResource] = useState<any>(null);
  const [loading, setLoading] = useState(true);
  const [requesting, setRequesting] = useState(false);
  const [respondLoading, setRespondLoading] = useState(false);
  const [showOtp, setShowOtp] = useState(false);
  const [voting, setVoting] = useState(false);

  useEffect(() => {
    const load = async () => {
      const me = await fetch("/api/auth/me", { credentials: "include" }).then(r => r.json()).catch(() => ({}));
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

  const doUpvote = async () => {
    if (!resource?.owner?.walletAddress) return alert("Owner wallet unknown");
    setVoting(true);
    try {
      await addReputation({ mentor: resource.owner.walletAddress, amount: 1 });
      const refreshed = await fetch(`/api/resources/${id}${viewer ? `?viewerId=${viewer._id}` : ""}`).then(r => r.json());
      setResource(refreshed.resource);
    } catch (e) {
      console.error("Upvote error:", e);
      alert("Upvote failed");
    } finally {
      setVoting(false);
    }
  };

  const doDownvote = async () => {
    if (!resource?.owner?.walletAddress) return alert("Owner wallet unknown");
    setVoting(true);
    try {
      await subReputation({ mentor: resource.owner.walletAddress, amount: 1 });
      const refreshed = await fetch(`/api/resources/${id}${viewer ? `?viewerId=${viewer._id}` : ""}`).then(r => r.json());
      setResource(refreshed.resource);
    } catch (e) {
      console.error("Downvote error:", e);
      alert("Downvote failed");
    } finally {
      setVoting(false);
    }
  };

  if (loading) {
    return <div className="text-white min-h-screen flex items-center justify-center">Loading...</div>;
  }

  return (
    <div className="relative min-h-screen w-full overflow-x-hidden text-white">
      <div className="fixed inset-0 -z-30">
        <ColorBends colors={['#3e47f4', '#06b31a', '#b46d04']} rotation={0} speed={0.3} scale={1} frequency={1} warpStrength={1} mouseInfluence={1} parallax={0.5} noise={0.1}/>
      </div>

      <Navbar user={viewer} onLoginClick={() => router.push("/")} />

      <main className="pt-24 max-w-3xl mx-auto px-6">
        <div className="bg-white/5 backdrop-blur-lg border border-white/10 rounded-2xl p-6">
          <h1 className="text-2xl font-bold">{resource.title}</h1>
          <p className="text-sm text-gray-300 mt-2">{resource.description}</p>

          <div className="mt-6 flex items-center gap-3">
            <img src={resource.owner.profileImage} className="w-12 h-12 rounded-full object-cover" />
            <div>
              <div className="font-semibold">{resource.owner.fullName}</div>
              <div className="flex items-center gap-3 mt-1">
                <div className="flex items-center gap-3 mt-1">
                  <VoteBox
                    walletAddress={resource.owner.walletAddress}
                    initialScore={Number(resource.owner.reputation)}
                    viewerId={viewer?._id}
                    profileId={resource.owner._id}
                  />
                </div>
              </div>
            </div>
          </div>

          <div className="mt-6">
            {resource.allowed ? (
              <a href={resource.fileUrl} target="_blank" rel="noreferrer" className="inline-block px-4 py-2 bg-indigo-600 rounded text-white">View / Download</a>
            ) : viewer && viewer._id === resource.owner._id ? (
              <>
                <div className="text-sm text-gray-300">Requests</div>
                <div className="mt-2 space-y-2">
                  {(resource.requests || []).filter((r:any)=>r.status === "pending").map((r:any)=>(
                    <div key={r._id} className="flex items-center justify-between bg-white/6 p-2 rounded">
                      <div>
                        <div className="font-medium">{r.user?.fullName || r.user}</div>
                        <div className="text-xs text-gray-400">{r.message}</div>
                      </div>
                      <div className="flex gap-2">
                        <button onClick={() => handleRespond(r._id, "approve")} disabled={respondLoading} className="px-3 py-1 bg-green-600 rounded">Approve</button>
                        <button onClick={() => handleRespond(r._id, "reject")} disabled={respondLoading} className="px-3 py-1 bg-red-600 rounded">Reject</button>
                      </div>
                    </div>
                  ))}
                </div>
              </>
            ) : (
              <>
                {resource.requestStatus === "pending" ? (
                  <div className="text-sm text-yellow-300">Request pending</div>
                ) : resource.requestStatus === "approved" ? (
                  <div className="text-sm text-green-300">Access approved — refresh to view</div>
                ) : (
                  <button onClick={askAccess} disabled={requesting} className="px-4 py-2 bg-indigo-600 rounded">{requesting ? "Requesting..." : "Request access"}</button>
                )}
              </>
            )}
          </div>
        </div>
      </main>
    </div>
  );
}
