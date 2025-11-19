'use client';

import { Link } from 'lucide-react';
import { useRouter } from 'next/navigation';
import { useMemo, useEffect, useState } from 'react';
import { FiMessageSquare } from 'react-icons/fi';

export default function Navbar({ user, onLoginClick }: any) {
  const router = useRouter();
  const [unread, setUnread] = useState(0);

  const randomColor = useMemo(() => {
    if (!user?.fullName) return '#6b7280';
    const colors = ['#60a5fa', '#34d399', '#f87171', '#a78bfa', '#f472b6', '#22d3ee'];
    return colors[user.fullName.charCodeAt(0) % colors.length];
  }, [user]);

  useEffect(() => {
    if (!user) return;

    const load = async () => {
      const res = await fetch(`/api/messages/unread-count?userId=${user._id}`);
      const data = await res.json();
      setUnread(data.count);
    };

    load();

    const id = setInterval(load, 3000);
    return () => clearInterval(id);
  }, [user]);

  return (
    <nav className="absolute top-0 left-0 w-full z-30 flex justify-between items-center px-10 py-5">
      <h1
        onClick={() => router.push('/')}
        className="text-2xl font-bold text-white cursor-pointer tracking-wide"
      >
        GradLedger
      </h1>

      {user ? (
        <div className="flex items-center gap-6">
          <button
            onClick={() => router.push('/explore')}
            className="px-4 py-2 bg-transparent-600 rounded-lg text-sm font-semibold hover:bg-indigo-500 transition"
          >
            Explore
          </button>
          <button
            onClick={() => router.push('/resources/upload')}
            className="px-4 py-2 bg-transparent-600 rounded-lg text-sm font-semibold hover:bg-indigo-500 transition"
          >
            Upload
          </button>
          
          <button
            onClick={() => router.push('/messages')}
            className="relative text-white hover:text-gray-200 transition"
          >
            <FiMessageSquare size={23} />

            {unread > 0 && (
              <span className="absolute -top-1.5 -right-1.5 bg-red-500 text-[10px] font-bold px-1.5 py-0.5 rounded-full">
                {unread}
              </span>
            )}
          </button>
          <button
            onClick={() => router.push('/discover')}
            className="relative text-white hover:text-gray-200 transition"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 24 24"
              fill="none"
              stroke="currentColor"
              strokeWidth={1.5}
              className="w-6 h-6"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                d="M12 20.25c4.556 0 8.25-3.694 8.25-8.25S16.556 3.75 12 3.75 3.75 7.444 3.75 12s3.694 8.25 8.25 8.25z"
              />
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                d="M12 15.75a3.75 3.75 0 100-7.5 3.75 3.75 0 000 7.5z"
              />
            </svg>
          </button>

          <button
            onClick={() => router.push('/profile')}
            className="flex items-center justify-center w-10 h-10 rounded-full bg-gray-800 hover:bg-gray-700 transition shadow-md"
            style={{ color: randomColor, fontWeight: 700, fontSize: '1rem' }}
          >
            {user.fullName[0].toUpperCase()}
          </button>
        </div>
      ) : (
        <button
          onClick={onLoginClick}
          className="px-5 py-2 bg-gradient-to-r from-indigo-700 to-purple-600 text-white font-semibold rounded-lg shadow-md"
        >
          Login / Register
        </button>
      )}
    </nav>
  );
}
