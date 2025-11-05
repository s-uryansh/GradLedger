'use client';

import { useState } from 'react';
import { motion, AnimatePresence } from 'framer-motion';
import toast from 'react-hot-toast';

interface OtpModalProps {
  email: string;
  onClose: () => void;
  onVerified: () => void;
}

export default function OtpModal({ email, onClose, onVerified }: OtpModalProps) {
  const [otp, setOtp] = useState('');
  const [loading, setLoading] = useState(false);

  const handleVerify = async () => {
    if (otp.length !== 6) {
      toast.error('Enter 6-digit OTP');
      return;
    }
    setLoading(true);
    try {
      const res = await fetch('/api/verify/check-otp', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ email, otp }),
      });
      const data = await res.json();
      if (!res.ok) throw new Error(data.error || 'Verification failed');
      toast.success('Email verified successfully!');
      onVerified();
      onClose();
    } catch (err: any) {
      toast.error(err.message);
    } finally {
      setLoading(false);
    }
  };

  return (
    <AnimatePresence>
      <motion.div
        initial={{ opacity: 0 }}
        animate={{ opacity: 1 }}
        exit={{ opacity: 0 }}
        className="fixed inset-0 flex items-center justify-center bg-black/60 backdrop-blur-md z-50"
      >
        <motion.div
          initial={{ scale: 0.9 }}
          animate={{ scale: 1 }}
          exit={{ scale: 0.9 }}
          transition={{ duration: 0.3 }}
          className="bg-white/10 backdrop-blur-lg border border-white/10 rounded-2xl p-8 w-[400px] text-center text-white shadow-lg"
        >
          <h2 className="text-2xl font-bold mb-4">Enter OTP</h2>
          <p className="text-gray-300 mb-4 text-sm">
            We’ve sent a 6-digit code to <span className="text-indigo-400">{email}</span>
          </p>
          <input
            type="text"
            value={otp}
            onChange={(e) => setOtp(e.target.value)}
            maxLength={6}
            className="w-full text-center py-2 bg-transparent border-b-2 border-gray-500 text-white focus:border-indigo-400 outline-none mb-6 tracking-widest text-lg"
          />
          <button
            onClick={handleVerify}
            disabled={loading}
            className="w-full bg-gradient-to-r from-indigo-600 to-purple-600 py-2 rounded-md font-semibold hover:from-indigo-700 hover:to-purple-700 transition-all"
          >
            {loading ? 'Verifying...' : 'Verify'}
          </button>
          <button onClick={onClose} className="mt-4 text-gray-400 text-sm hover:text-gray-200">
            Cancel
          </button>
        </motion.div>
      </motion.div>
    </AnimatePresence>
  );
}
