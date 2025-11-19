"use client";

import { useEffect, useState } from "react";
import Navbar from "@/components/UI/Navbar";
import ColorBends from "@/components/BackgroundAnimations/ColorBends";
import { useRouter } from "next/navigation";

export default function ExplorePage() {
  const router = useRouter();
  const [viewer, setViewer] = useState<any>(null);
  const [items, setItems] = useState<any[]>([]);
  const [q, setQ] = useState("");
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    (async () => {
      const me = await fetch("/api/auth/me", { credentials: "include" }).then(r => r.json()).catch(() => ({}));
      setViewer(me.user || null);

      const res = await fetch(`/api/resources/list?page=1&q=&publicOnly=true`);
      const data = await res.json();
      setItems(data.users || []);
      setLoading(false);
    })();
  }, []);

  const search = async () => {
    const res = await fetch(`/api/resources/list?page=1&q=${encodeURIComponent(q)}&publicOnly=true`);
    const data = await res.json();
    setItems(data.users || []);
  };

  return (
    <div className="relative min-h-screen text-white">
      
      <div className="fixed inset-0 -z-30">
        <ColorBends
          colors={["#3e47f4", "#06b31a", "#b46d04"]}
          rotation={0}
          speed={0.3}
          scale={1}
          frequency={1}
          warpStrength={1}
          mouseInfluence={1}
          parallax={0.5}
          noise={0.1}
        />
        <div className="absolute inset-0 pointer-events-none" style={{
          background: "linear-gradient(to bottom, rgba(0,0,0,0.25), rgba(0,0,0,0.1))"
        }}/>
      </div>

      <Navbar user={viewer} />

      <div className="pt-24 max-w-4xl mx-auto px-4">
        <h1 className="text-3xl font-bold mb-6">Explore Resources</h1>

        <div className="flex gap-3 mb-6">
          <input
            value={q}
            onChange={e => setQ(e.target.value)}
            placeholder="Search public resources"
            className="flex-1 p-2 bg-white/10 rounded"
          />
          <button
            onClick={search}
            className="px-4 bg-indigo-600 rounded"
          >
            Search
          </button>
        </div>

        <div className="space-y-4">
          {loading ? (
            <div className="text-gray-300">Loading...</div>
          ) : items.length === 0 ? (
            <div className="text-gray-400">No public resources found.</div>
          ) : (
            items.map(r => (
              <div key={r._id} className="p-4 bg-white/5 rounded cursor-pointer"
                onClick={() => router.push(`/resources/${r._id}`)}
              >
                <div className="font-semibold">{r.title}</div>
                <div className="text-xs text-gray-300">{r.subject} • {r.tags?.join(", ")}</div>
                <div className="text-xs text-gray-400 mt-1">by {r.owner.fullName}</div>
              </div>
            ))
          )}
        </div>
      </div>
    </div>
  );
}
