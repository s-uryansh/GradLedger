'use client';

import { useEffect, useState } from 'react';
import { useRouter, useParams } from 'next/navigation';
import Navbar from '@/components/UI/Navbar';
import ColorBends from '@/components/BackgroundAnimations/ColorBends';
import VoteBox from "@/components/Resources/VoteBox";

export default function PublicProfile() {
  const { id } = useParams() as { id: string };
  const router = useRouter();

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
        { (profile.selfieImage || profile.profileImage) ? (
          <img src={profile.selfieImage || profile.profileImage} className="w-32 h-32 rounded-full border border-white/20 object-cover" />
        ) : (
          <div className="w-32 h-32 rounded-full border border-white/20 flex items-center justify-center" style={{ background: "#374151" }}>
            <span className="text-white font-bold text-2xl">{makePlaceholder(profile._id, profile.fullName)}</span>
          </div>
        )}

        <h2 className="mt-4 text-2xl font-bold">{profile.fullName}</h2>
        <p className="text-gray-300 text-sm">{profile.program} — {profile.major}</p>

        {/* tags */}
        {Array.isArray(profile.tags) && profile.tags.length > 0 && (
          <div className="mt-3 flex flex-wrap gap-2">
            {profile.tags.map((t: string) => (
              <span key={t} className="text-xs bg-white/6 px-2 py-1 rounded text-gray-100">
                {t}
              </span>
            ))}
          </div>
        )}

          <div className="mt-4 flex items-center gap-4">
            <div className="flex items-center gap-3 mt-1">
              <VoteBox
                walletAddress={profile.walletAddress}
                initialScore={Number(profile.reputation)}
                viewerId={viewer?._id}
                profileId={profile._id}
              />
            </div>
          </div>

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
