'use client';

import { useState, useEffect } from 'react';
import { useRouter } from 'next/navigation';
import { motion } from 'framer-motion';
import { toast } from 'react-hot-toast';
import ColorBends from '@/components/BackgroundAnimations/ColorBends';
import Navbar from '@/components/UI/Navbar';
import OtpModal from '@/components/Auth/OtpModal';

export default function ProfilePage() {
  const [user, setUser] = useState<any>(null);
  const [isOtpOpen, setIsOtpOpen] = useState(false);
  const [sendingOtp, setSendingOtp] = useState(false); // <— new state
  const router = useRouter();

  useEffect(() => {
    const fetchUser = async () => {
      try {
        const res = await fetch('/api/auth/me', { credentials: 'include' });
        const data = await res.json();
        if (!data.user) router.push('/');
        else setUser(data.user);
      } catch (err) {
        console.error('Error fetching user:', err);
        router.push('/');
      }
    };
    fetchUser();
  }, [router]);

  const handleLogout = async () => {
    try {
      await fetch('/api/auth/logout', { method: 'POST', credentials: 'include' });
      toast.success('Logged out successfully!');
      setTimeout(() => router.push('/'), 800);
    } catch {
      toast.error('Logout failed. Try again.');
    }
  };

  const handleSendOtp = async () => {
    if (sendingOtp) return;
    setSendingOtp(true);
    try {
      const res = await fetch('/api/verify/send-otp', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ email: user.email }),
      });
      const data = await res.json();
      if (!res.ok) throw new Error(data.error || 'Failed to send OTP');
      toast.success('OTP sent to your email!');
      setIsOtpOpen(true);
    } catch (err: any) {
      toast.error(err.message);
    } finally {
      setSendingOtp(false);
    }
  };

  if (!user) {
    return (
      <div className="flex items-center justify-center min-h-screen text-gray-300 text-lg">
        Loading profile...
      </div>
    );
  }

  return (
    <div className="relative min-h-screen w-full overflow-hidden text-white">
      {/* Background */}
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

      {/* Navbar */}
      <div className="relative z-40">
        <Navbar user={user} onLoginClick={() => router.push('/')} />
      </div>

      {/* Profile Content */}
      <main className="relative z-20 flex flex-col items-center justify-center min-h-[80vh] px-6">
        <motion.div
          initial={{ opacity: 0, y: 20 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.6 }}
          className="bg-white/5 backdrop-blur-lg border border-white/10 rounded-3xl p-10 max-w-md w-full shadow-[0_0_40px_rgba(0,0,0,0.4)] text-center"
        >
          <h2 className="text-3xl font-bold mb-6 bg-gradient-to-r from-indigo-400 to-purple-400 text-transparent bg-clip-text">
            Your Profile
          </h2>

          <div className="space-y-3 text-gray-300 text-sm text-left">
            <p><span className="font-semibold text-white">Full Name:</span> {user.fullName}</p>
            <p><span className="font-semibold text-white">Email:</span> {user.email}</p>
            <p><span className="font-semibold text-white">Role:</span> {user.role}</p>
            <p>
              <span className="font-semibold text-white">Wallet Address:</span>{' '}
              {user.walletAddress?.slice(0, 8)}...{user.walletAddress?.slice(-6)}
            </p>
            <p><span className="font-semibold text-white">Mail Verified:</span> {user.mailVerified ? 'True' : 'False'}</p>
            <p><span className="font-semibold text-white">Face Verified:</span> {user.faceVerified ? 'True' : 'False'}</p>
          </div>

          {/* User Tags */}
          {user.tags && user.tags.length > 0 && (
            <div className="flex flex-wrap justify-center gap-2 mt-6">
              {user.tags.map((tag: string, idx: number) => (
                <span
                  key={idx}
                  className="px-3 py-1 bg-gray-700/50 border border-gray-500/50 text-sm rounded-full text-indigo-300"
                >
                  {tag}
                </span>
              ))}
            </div>
          )}

          <div className="flex flex-col items-center gap-3 mt-8">
            {!user.mailVerified && (
              <>
                <button
                  onClick={handleSendOtp}
                  disabled={sendingOtp}
                  className={`px-6 py-2 rounded-lg font-semibold text-white transition-all ${
                    sendingOtp
                      ? 'bg-gray-600 cursor-not-allowed'
                      : 'bg-gradient-to-r from-indigo-600 to-purple-600 hover:from-indigo-700 hover:to-purple-700'
                  }`}
                >
                  {sendingOtp ? 'Sending OTP...' : 'Verify Email'}
                </button>

                {isOtpOpen && (
                  <OtpModal
                    email={user.email}
                    onClose={() => setIsOtpOpen(false)}
                    onVerified={() => {
                      setUser({ ...user, mailVerified: true });
                      setIsOtpOpen(false);
                      router.push('/verify');
                    }}
                  />
                )}
              </>
            )}

            {user.mailVerified && !user.faceVerified && (
              <button
                onClick={() => router.push('/verify')}
                className="px-6 py-2 rounded-lg bg-gradient-to-r from-green-600 to-emerald-500 font-semibold text-white hover:from-green-700 hover:to-emerald-600 transition-all"
              >
                Verify Face
              </button>
            )}

            <button
              onClick={handleLogout}
              className="text-red-400 text-sm hover:text-red-300 mt-3 transition-all"
            >
              Log out
            </button>
          </div>
        </motion.div>
      </main>

      <footer className="absolute bottom-6 text-sm text-gray-400 text-center w-full">
        GradLedger © {new Date().getFullYear()}
      </footer>
    </div>
  );
}
