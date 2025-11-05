'use client';

import { useRouter } from 'next/navigation';
import { useMemo } from 'react';

interface NavbarProps {
  user?: {
    fullName: string;
  } | null;
  onLoginClick: () => void;
}

export default function Navbar({ user, onLoginClick }: NavbarProps) {
  const router = useRouter();

  const randomColor = useMemo(() => {
    if (!user?.fullName) return '#6b7280'; 
    const colors = [
      '#60a5fa', 
      '#34d399', 
      '#f87171', 
      '#a78bfa', 
      '#f472b6', 
      '#22d3ee', 
    ];
    const index = user.fullName.charCodeAt(0) % colors.length;
    return colors[index];
  }, [user]);

  return (
    <nav className="absolute top-0 left-0 w-full z-30 flex justify-between items-center px-10 py-5 bg-transparent">
      {/* Logo */}
      <h1
        onClick={() => router.push('/')}
        className="text-2xl font-bold text-white cursor-pointer tracking-wide drop-shadow-[0_2px_4px_rgba(0,0,0,0.4)]"
      >
        GradLedger
      </h1>

      {/* Right side */}
      {user ? (
        <button
          onClick={() => router.push('/profile')}
          className="flex items-center justify-center w-10 h-10 rounded-full bg-gray-800 hover:bg-gray-700 transition-all shadow-md"
          style={{
            color: randomColor,
            fontWeight: 700,
            fontSize: '1rem',
          }}
        >
          {user.fullName[0].toUpperCase()}
        </button>
      ) : (
        <button
          onClick={onLoginClick}
          className="px-5 py-2 bg-gradient-to-r from-indigo-700 to-purple-600 text-white font-semibold rounded-lg shadow-md hover:from-indigo-800 hover:to-purple-700 transition-all"
        >
          Login / Register
        </button>
      )}
    </nav>
  );
}
