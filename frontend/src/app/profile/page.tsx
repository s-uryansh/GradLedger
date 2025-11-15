'use client';

import { useEffect, useState } from 'react';
import { motion } from 'framer-motion';
import ColorBends from '@/components/BackgroundAnimations/ColorBends';
import Navbar from '@/components/UI/Navbar';
import { toast } from 'react-hot-toast';
import { useRouter } from 'next/navigation';
import OtpModal from '@/components/Auth/OtpModal';

export default function ProfilePage() {
  const [user, setUser] = useState<any>(null);
  const [sendingOtp, setSendingOtp] = useState(false);
  const [isOtpOpen, setIsOtpOpen] = useState(false);
  const [flipped, setFlipped] = useState(false);
  const router = useRouter();

  // Fetch current user
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

  // Logout
  const handleLogout = async () => {
    try {
      await fetch('/api/auth/logout', { method: 'POST', credentials: 'include' });
      toast.success('Logged out successfully!');
      setTimeout(() => router.push('/'), 800);
    } catch {
      toast.error('Logout failed. Try again.');
    }
  };

  // Send OTP
  const handleSendOtp = async () => {
    if (!user?.email) return;
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
      {/* Animated Background */}
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
      </div>

      {/* Navbar */}
      <div className="relative z-40">
        <Navbar user={user} onLoginClick={() => router.push('/')} />
      </div>

      {/* Profile Card */}
      <main className="relative z-20 flex flex-col items-center justify-center min-h-[80vh] px-6">
        <motion.div
          initial={{ opacity: 0, y: 20 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.6 }}
          className="bg-white/5 backdrop-blur-lg border border-white/10 rounded-3xl p-10 max-w-md w-full shadow-[0_0_40px_rgba(0,0,0,0.4)] text-center"
        >
        {/* Profile Image */}
        <div className="flex flex-col items-center mb-6">
          <motion.div
            className="w-28 h-28 relative cursor-pointer"
            whileTap={{ scale: 0.95 }}
            onClick={() => setFlipped(!flipped)}
            style={{ perspective: 800 }}
          >
            <motion.div
              animate={{ rotateY: flipped ? 180 : 0 }}
              transition={{ duration: 0.6 }}
              className="relative w-full h-full rounded-full shadow-lg"
              style={{ transformStyle: 'preserve-3d' }}
            >
              {user.profileImage ? (
                <motion.img
                  src={user.profileImage}
                  alt="Profile"
                  className="absolute inset-0 w-full h-full rounded-full object-cover border-2 border-indigo-400 backface-hidden"
                />
              ) : (
                <div className="absolute inset-0 w-full h-full rounded-full bg-gray-700 border-2 border-gray-600 flex items-center justify-center text-gray-400 text-sm">
                  No ID Photo
                </div>
              )}
              {user.selfieImage ? (
                <motion.img
                  src={user.selfieImage}
                  alt="Selfie"
                  className="absolute inset-0 w-full h-full rounded-full object-cover border-2 border-purple-400 backface-hidden"
                  style={{ transform: 'rotateY(180deg)' }}
                />
              ) : (
                <div
                  className="absolute inset-0 w-full h-full rounded-full bg-gray-700 border-2 border-gray-600 flex items-center justify-center text-gray-400 text-sm"
                  style={{ transform: 'rotateY(180deg)' }}
                >
                  No Selfie
                </div>
              )}
            </motion.div>
          </motion.div>
            
          <h3 className="text-xl font-bold mt-4">{user.fullName}</h3>
          {user.rollNumber && <p className="text-gray-400">{user.rollNumber}</p>}
          {(user.program || user.major) && (
            <p className="text-gray-400">{user.program} — {user.major}</p>
          )}
        </div>

          {/* Basic Info */}
          <div className="space-y-3 text-gray-300 text-sm text-left">
            <p><span className="font-semibold text-white">Email:</span> {user.email}</p>
            <p><span className="font-semibold text-white">Role:</span> {user.role}</p>
            <p>
              <span className="font-semibold text-white">Wallet Address:</span>{' '}
              {user.walletAddress?.slice(0, 8)}...{user.walletAddress?.slice(-6)}
            </p>
            <p><span className="font-semibold text-white">Mail Verified:</span> {user.mailVerified ? 'True' : 'False'}</p>
            <p><span className="font-semibold text-white">Face Verified:</span> {user.faceVerified ? 'True' : 'False'}</p>
          </div>

          {/* Tags */}
          {user.tags && user.tags.length > 0 && (
            <div className="flex flex-wrap justify-center gap-2 mt-6">
              {user.tags.map((tag: string, i: number) => (
                <span
                  key={i}
                  className="px-3 py-1 bg-gray-700/50 border border-gray-500/50 text-sm rounded-full text-indigo-300"
                >
                  {tag}
                </span>
              ))}
            </div>
          )}

          {!(user.mailVerified && user.faceVerified) && (
            <div className="mt-6">
              <div className="flex justify-between text-xs text-gray-400 mb-1">
                <span>Email</span>
                <span>Face</span>
                <span>Complete</span>
              </div>
              <div className="w-full bg-gray-700 h-2 rounded-full overflow-hidden">
                <div
                  className={`h-full transition-all duration-500 ${
                    user.mailVerified
                      ? 'bg-indigo-500 w-2/3'
                      : 'bg-gray-500 w-1/3'
                  }`}
                />
              </div>
            </div>
          )}

          {/* Actions */}
          <div className="flex flex-col items-center gap-3 mt-8">
            {!user.mailVerified ? (
              <>
                <button
                  onClick={handleSendOtp}
                  disabled={sendingOtp}
                  className={`px-6 py-2 rounded-lg bg-gradient-to-r from-indigo-600 to-purple-600 font-semibold text-white transition-all ${
                    sendingOtp
                      ? 'opacity-60 cursor-not-allowed'
                      : 'hover:from-indigo-700 hover:to-purple-700'
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
                      toast.success('Email verified!');
                    }}
                  />
                )}
              </>
            ) : !user.faceVerified ? (
              <button
                onClick={() => router.push('/verify')}
                className="px-6 py-2 rounded-lg bg-gradient-to-r from-green-600 to-emerald-500 font-semibold text-white hover:from-green-700 hover:to-emerald-600 transition-all"
              >
                Verify Face
              </button>
            ) : null}

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
