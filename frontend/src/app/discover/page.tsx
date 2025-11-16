'use client';

import { useEffect, useState } from 'react';
import Navbar from '@/components/UI/Navbar';
import ProfileCard from '@/components/ProfileCard/ProfileCard';
import { useRouter } from 'next/navigation';
import ColorBends from '@/components/BackgroundAnimations/ColorBends';

export default function DiscoverPage() {
  const router = useRouter();
  const [viewer, setViewer] = useState<any>(null);
  const [users, setUsers] = useState<any[]>([]);
  const [query, setQuery] = useState("");
  const [page, setPage] = useState(1);
  const [totalPages, setTotalPages] = useState(1);

  useEffect(() => {
    const loadViewer = async () => {
      const res = await fetch("/api/auth/me", { credentials: "include" });
      const data = await res.json();
      if (!data.user) return router.push("/");
      setViewer(data.user);
    };
    loadViewer();
  }, []);

  useEffect(() => {
    if (!viewer) return;

    const load = async () => {
      const url = `/api/users/discover?viewerId=${viewer._id}&page=${page}&q=${query}`;
      const res = await fetch(url);
      const data = await res.json();

      if (page === 1) {
        setUsers(data.users);
      } else {
        setUsers(prev =>
          [...prev, ...data.users.filter((u: any) => !prev.some(p => p._id === u._id))]
        );
      }

      setTotalPages(data.totalPages);
    };

    load();
  }, [viewer, page, query]);

  if (!viewer) return <div className="text-white p-10">Loading...</div>;

  const openChat = async (userId: string) => {
    const res = await fetch("/api/messages/get-or-create", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ userA: viewer._id, userB: userId })
    });

    const convo = await res.json();
    router.push(`/messages/${convo._id}`);
  };

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

      <Navbar user={viewer} onLoginClick={() => router.push('/')} />

      <div className="mt-24 max-w-4xl mx-auto">

        <input
          value={query}
          onChange={(e) => {
            setQuery(e.target.value);
            setPage(1);
          }}
          placeholder="Search full name or email"
          className="w-full p-3 rounded-lg bg-white/10 border border-white/20 text-white mb-8"
        />

        <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
          {users
            .filter((u) => u._id !== viewer._id)
            .map((u) => (
              <div key={u._id} className="relative group">

                <ProfileCard
                  name={u.fullName}
                  title={`${u.program} — ${u.major}`}
                  handle={u.email.split("@")[0]}
                  status="Verified"
                  avatarUrl={u.selfieImage || u.profileImage}
                  showUserInfo={false}
                  enableTilt={true}
                  contactText="View Profile"
                  onContactClick={() => router.push(`/profile/${u._id}`)}
                />

                <button
                  onClick={() => openChat(u._id)}
                  className="
                    absolute bottom-3 right-3 bg-blue-600
                    text-white px-2 py-1 rounded-md opacity-0
                    group-hover:opacity-100 transition
                  "
                >
                  Chat
                </button>
              </div>
            ))}
        </div>

        {(page < totalPages) && users.length > 0 && (
          <button
            onClick={() => setPage((p) => p + 1)}
            className="w-full mt-8 py-3 bg-white/10 rounded-lg text-white hover:bg-white/20"
          >
            Load More
          </button>
        )}
      </div>
    </div>
  );
}
