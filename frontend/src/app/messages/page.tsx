'use client';

import { useEffect, useState } from 'react';
import Navbar from '@/components/UI/Navbar';
import { useRouter } from 'next/navigation';
import { FiSearch } from 'react-icons/fi';
import ColorBends from '@/components/BackgroundAnimations/ColorBends';

export default function MessagesPage() {
  const [user, setUser] = useState<any>(null);
  const [loadingUser, setLoadingUser] = useState(true);
  const [conversations, setConversations] = useState<any[]>([]);
  const [search, setSearch] = useState('');
  const [filtered, setFiltered] = useState<any[]>([]);
  const [page, setPage] = useState(1);
  const router = useRouter();

  useEffect(() => {
    const fetchUser = async () => {
      const res = await fetch('/api/auth/me', { credentials: 'include' });
      const data = await res.json();

      if (!data.user) return router.push('/');
      if (!data.user.mailVerified || !data.user.faceVerified)
        return router.push('/profile');

      setUser(data.user);
      setLoadingUser(false);
    };

    fetchUser();
  }, [router]);

  useEffect(() => {
    if (!user) return;

    const loadConversations = async () => {
      const res = await fetch('/api/messages/list', {
        headers: {
          'x-user-id': user._id,
          'x-page': String(page),
        },
      });

      const data = await res.json();

      if (page === 1) setConversations(data);
      else setConversations(prev => [...prev, ...data]);

      setFiltered(data);
    };

    loadConversations();
  }, [user, page]);

  useEffect(() => {
    if (search.trim() === '') {
      setFiltered(conversations);
      return;
    }

    const s = search.toLowerCase();
    setFiltered(
      conversations.filter(c =>
        c.participants.some(
          (p: any) =>
            p._id !== user._id && p.fullName.toLowerCase().includes(s)
        )
      )
    );
  }, [search, conversations, user]);

  if (loadingUser) {
    return (
      <div className="text-gray-300 min-h-screen flex items-center justify-center">
        Loading...
      </div>
    );
  }

  // helper: deterministic 3-letter placeholder
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

  const AvatarEl = ({ img, seed, name }: { img?: string; seed?: string; name?: string }) => {
    if (img) return <img src={img} className="w-12 h-12 rounded-full object-cover" />;
    const code = makePlaceholder(seed, name);
    const bg = "#4b5563";
    return <div style={{ background: bg }} className="w-12 h-12 rounded-full flex items-center justify-center text-white font-semibold">{code}</div>;
  };

  return (
    <div className="relative min-h-screen w-full overflow-x-hidden text-white">
      {/* === Background === */}
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

      {/* === Navbar === */}
      <Navbar user={user} onLoginClick={() => {}} />

      {/* === Content === */}
      <div className="pt-24 max-w-xl mx-auto px-4">
        <div className="flex items-center gap-3 bg-white/10 px-3 py-2 rounded-lg">
          <FiSearch size={18} className="text-gray-300" />
          <input
            type="text"
            placeholder="Search user..."
            value={search}
            onChange={(e) => setSearch(e.target.value)}
            className="bg-transparent w-full outline-none text-gray-200"
          />
        </div>

        <div className="mt-6 space-y-3">
          {filtered.map((convo) => {
            // defensive: participants or "other" might be missing from server data
            const other =
              (convo.participants && convo.participants.find((p: any) => p._id !== user._id)) ||
              { profileImage: null, _id: convo._id, fullName: 'Unknown' };

            return (
              <div
                key={convo._id}
                onClick={() => router.push(`/messages/${convo._id}`)}
                className="p-3 bg-white/5 rounded-lg flex items-center gap-3 cursor-pointer hover:bg-white/10 transition"
              >
                <AvatarEl img={other.profileImage} seed={other._id} name={other.fullName} />
                <div>
                  <p className="text-sm font-semibold">{other.fullName}</p>
                  <p className="text-xs text-gray-400">
                    {convo.lastMessage?.text ? convo.lastMessage.text.slice(0, 30) : 'No messages yet'}
                  </p>
                </div>
              </div>
            );
          })}

          <button
            onClick={() => setPage((p) => p + 1)}
            className="w-full mt-3 py-2 text-center bg-white/10 rounded-lg hover:bg-white/20"
          >
            Load More
          </button>
        </div>
      </div>
    </div>
  );
}
