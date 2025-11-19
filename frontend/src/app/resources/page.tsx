'use client';
import React, { useEffect, useState } from "react";
import ProfileCard from "@/components/ProfileCard/ProfileCard"; 
import Navbar from "@/components/UI/Navbar";
import ColorBends from "@/components/BackgroundAnimations/ColorBends";
import { useRouter } from "next/navigation";

export default function ResourcesPage() {
  const [viewer, setViewer] = useState<any>(null);
  const [resources, setResources] = useState<any[]>([]);
  const [page, setPage] = useState(1);
  const [q, setQ] = useState("");
  const [totalPages, setTotalPages] = useState(1);
  const router = useRouter();

  useEffect(() => {
    (async () => {
      const me = await fetch("/api/auth/me", { credentials: "include" }).then(r => r.json()).catch(() => ({}));
      setViewer(me.user || null);

    })();
  }, [router]);

  useEffect(() => {
    if (!viewer) return;
    (async () => {
      const res = await fetch(`/api/resources/list?page=${page}&q=${encodeURIComponent(q)}&viewerId=${viewer._id}`);
      const data = await res.json();
      setResources(prev => page === 1 ? data.users : [...prev, ...data.users]);
      setTotalPages(data.totalPages || 1);
    })();
  }, [viewer, page, q]);

  const openResource = async (_id: string) => {
    if (!viewer) return;
    const res = await fetch(`/api/resources/${_id}?viewerId=${viewer._id}`);
    const data = await res.json();
    if (data.resource.allowed) {
      window.open(data.resource.fileUrl, "_blank");
      return;
    }
    // request access
    const message = prompt("Request message (optional)");
    if (!message && !confirm("Send empty request?")) return;
    await fetch(`/api/resources/${_id}/request`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ userId: viewer._id, message })
    });
    alert("Request sent");
  };

  if (!viewer) return <div className="text-white p-8">Loading...</div>;
  
  return (
    <div className="relative min-h-screen w-full overflow-x-hidden text-white">
      <div className="fixed inset-0 -z-30">
        <ColorBends colors={['#3e47f4','#06b31a','#b46d04']} rotation={0} speed={0.3} scale={1} frequency={1} warpStrength={1} mouseInfluence={1} parallax={0.5} noise={0.1}/>
        <div className="absolute inset-0 pointer-events-none" style={{background: 'linear-gradient(to bottom, rgba(0,0,0,0.25) 0%, rgba(0,0,0,0.05) 25%, rgba(0,0,0,0.10) 70%, rgba(0,0,0,0.15) 100%)'}}/>
      </div>

      <Navbar user={viewer} onLoginClick={() => router.push("/")} />
      <div className="mt-24 max-w-4xl mx-auto p-4">
        <input value={q} onChange={e => { setQ(e.target.value); setPage(1); }} placeholder="Search title / subject / tags" className="w-full p-2 rounded bg-white/5 text-white mb-4"/>
        <div className="space-y-4">
          {resources.map(r => (
            <div key={r._id} className="p-4 bg-white/5 rounded">
              <div className="flex justify-between">
                <div>
                  <div className="font-semibold">{r.title}</div>
                  <div className="text-xs text-gray-300">{r.subject} • {r.tags?.join(", ")}</div>
                </div>
                <div>
                  <button onClick={() => openResource(r._id)} className="px-3 py-1 bg-indigo-600 rounded">Open / Request</button>
                </div>
              </div>
            </div>
          ))}
        </div>

        {page < totalPages && (
          <button onClick={() => setPage(p => p + 1)} className="mt-6 w-full py-2 bg-white/10 rounded">Load More</button>
        )}
      </div>
    </div>
  );
}
