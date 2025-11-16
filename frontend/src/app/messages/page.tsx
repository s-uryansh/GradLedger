'use client';

import { useEffect, useState } from 'react';
import Navbar from '@/components/UI/Navbar';
import { useRouter } from 'next/navigation';
import { FiSearch } from 'react-icons/fi';

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
      else setConversations((prev) => [...prev, ...data]);

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
      conversations.filter((c) =>
        c.participants.some(
          (p: any) =>
            p._id !== user._id &&
            p.fullName.toLowerCase().includes(s)
        )
      )
    );
  }, [search, conversations, user]);

  if (loadingUser)
    return (
      <div className="text-gray-300 min-h-screen flex items-center justify-center">
        Loading...
      </div>
    );

  return (
    <div className="relative min-h-screen text-white">
      <Navbar user={user} onLoginClick={() => {}} />

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
            const other = convo.participants.find(
              (p: any) => p._id !== user._id
            );

            return (
              <div
                key={convo._id}
                onClick={() => router.push(`/messages/${convo._id}`)}
                className="p-3 bg-white/5 rounded-lg flex items-center gap-3 cursor-pointer hover:bg-white/10 transition"
              >
                <img
                  src={other.profileImage}
                  className="w-12 h-12 rounded-full object-cover"
                />
                <div>
                  <p className="text-sm font-semibold">{other.fullName}</p>
                  <p className="text-xs text-gray-400">
                    {convo.lastMessage?.text
                      ? convo.lastMessage.text.slice(0, 30)
                      : 'No messages yet'}
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
