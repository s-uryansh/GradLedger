'use client';

import { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';
import Navbar from '@/components/UI/Navbar';
import ColorBends from '@/components/BackgroundAnimations/ColorBends';

export default function PublicProfile({ params }: any) {
  const router = useRouter();
  const { id } = params;

  const [viewer, setViewer] = useState<any>(null);
  const [profile, setProfile] = useState<any>(null);

  useEffect(() => {
    const load = async () => {
      const me = await fetch('/api/auth/me', { credentials: 'include' }).then(r => r.json());
      if (!me.user) return router.push('/');
      setViewer(me.user);

      const p = await fetch(`/api/users/${id}`).then(r => r.json());
      setProfile(p.user);
    };
    load();
  }, [id]);

  const startChat = async () => {
    const res = await fetch("/api/messages/get-or-create", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ userA: viewer._id, userB: profile._id })
    });

    const convo = await res.json();
    router.push(`/messages/${convo._id}`);
  };

  if (!profile || !viewer) return <div className="text-white p-10">Loading...</div>;

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

      <Navbar user={viewer} />

      <div className="mt-24 flex flex-col items-center">
        <img
          src={profile.selfieImage || profile.profileImage}
          className="w-32 h-32 rounded-full border border-white/20 object-cover"
        />

        <h2 className="mt-4 text-2xl font-bold">{profile.fullName}</h2>
        <p className="text-gray-300 text-sm">{profile.program} — {profile.major}</p>

        <button
          onClick={startChat}
          className="mt-6 px-6 py-2 bg-blue-600 rounded-lg"
        >
          Chat Now
        </button>
      </div>
    </div>
  );
}
