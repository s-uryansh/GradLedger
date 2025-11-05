'use client';

import { motion, AnimatePresence } from 'framer-motion';
import { toast } from 'react-hot-toast';

interface ProfileDialogProps {
  user: {
    fullName: string;
    email: string;
    role: string;
    walletAddress: string;
  };
  onClose: () => void;
}

export default function ProfileDialog({ user, onClose }: ProfileDialogProps) {
  const handleLogout = async () => {
    try {
      await fetch('/api/auth/logout', { method: 'POST', credentials: 'include' });
      toast.success('Logged out successfully!');
      setTimeout(() => window.location.reload(), 800);
    } catch {
      toast.error('Logout failed.');
    }
  };

  return (
    <div className="fixed inset-0 flex items-center justify-center bg-black/60 backdrop-blur-md z-50">
      <AnimatePresence>
        <motion.div
          key="profile"
          initial={{ opacity: 0, scale: 0.9, y: 20 }}
          animate={{ opacity: 1, scale: 1, y: 0 }}
          exit={{ opacity: 0, scale: 0.9, y: -20 }}
          transition={{ duration: 0.3 }}
          className="relative bg-gradient-to-br from-[#0a0b14]/95 via-[#0e1c1c]/90 to-[#131336]/95 rounded-3xl shadow-xl overflow-hidden w-[460px] p-8 border border-white/10"
        >
          <button
            onClick={onClose}
            className="absolute top-3 right-4 text-gray-300 hover:text-white text-xl font-bold"
          >
            ✕
          </button>

          <h2 className="text-2xl font-bold text-white mb-6 text-center">Profile</h2>

          <div className="space-y-3 text-gray-300 text-sm">
            <p><span className="font-semibold text-white">Name:</span> {user.fullName}</p>
            <p><span className="font-semibold text-white">Email:</span> {user.email}</p>
            <p><span className="font-semibold text-white">Role:</span> {user.role}</p>
            <p>
              <span className="font-semibold text-white">Wallet:</span>{' '}
              {user.walletAddress?.slice(0, 10)}...
            </p>
          </div>

          <div className="flex flex-col items-center mt-8 gap-3">
            <button
              onClick={() => toast.success('Verification started!')}
              className="px-6 py-2 rounded-lg bg-gradient-to-r from-indigo-600 to-purple-600 font-semibold text-white hover:from-indigo-700 hover:to-purple-700 transition-all"
            >
              Verify
            </button>

            <button
              onClick={handleLogout}
              className="text-red-400 text-sm hover:text-red-300 transition-all mt-2"
            >
              Log out
            </button>
          </div>
        </motion.div>
      </AnimatePresence>
    </div>
  );
}
