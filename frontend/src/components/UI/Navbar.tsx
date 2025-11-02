'use client';

interface NavbarProps {
  onLoginClick: () => void;
}

export default function Navbar({ onLoginClick }: NavbarProps) {
  return (
    <nav className="absolute top-0 left-0 w-full z-30 flex justify-between items-center px-10 py-5 bg-transparent">
      <h1 className="text-2xl font-bold text-white tracking-wide drop-shadow-[0_2px_4px_rgba(0,0,0,0.4)]">
        GradLedger
      </h1>
      <button
        onClick={onLoginClick}
        className="px-5 py-2 bg-gradient-to-r from-indigo-700 to-purple-600 text-white font-semibold rounded-lg shadow-md hover:from-indigo-800 hover:to-purple-700 transition-all"
      >
        Login / Register
      </button>
    </nav>
  );
}
