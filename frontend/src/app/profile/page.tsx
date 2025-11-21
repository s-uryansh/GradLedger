'use client';

import { useEffect, useState } from 'react';
import { motion } from 'framer-motion';
import ColorBends from '@/components/BackgroundAnimations/ColorBends';
import Navbar from '@/components/UI/Navbar';
import { toast } from 'react-hot-toast';
import { useRouter } from 'next/navigation';
import OtpModal from '@/components/Auth/OtpModal';
import Link from 'next/link';

export default function ProfilePage() {
  const [user, setUser] = useState<any>(null);
  const [reputation, setReputation] = useState<number | null>(null);
  const [sendingOtp, setSendingOtp] = useState(false);
  const [isOtpOpen, setIsOtpOpen] = useState(false);
  const [flipped, setFlipped] = useState(false);
  const [uploads, setUploads] = useState<any[]>([]);
  const router = useRouter();

  useEffect(() => {
    const fetchUser = async () => {
      try {
        const res = await fetch('/api/auth/me', { credentials: 'include' });
        const data = await res.json();
        if (!data.user) return router.push('/');
        setUser(data.user);
        if (data.user.walletAddress) {
          const repRes = await fetch(`http://localhost:8080/reputation/${data.user.walletAddress}`);
          const repData = await repRes.json();
          setReputation(repData.score ? Number(repData.score) : 0);
        }

        try {
          const rres = await fetch(`/api/resources/owner?ownerId=${data.user._id}`);
          const rdata = await rres.json();
          setUploads((rdata.resources || rdata.items || rdata.resources) as any[]); // tolerate shapes
        } catch (e) {
          console.warn('Failed to load uploads', e);
        }
      } catch {
        router.push('/');
      }
    };

    fetchUser();
  }, [router]);


  const openMessages = () => router.push("/messages");

  const handleLogout = async () => {
    try {
      await fetch('/api/auth/logout', { method: 'POST', credentials: 'include' });
      toast.success('Logged out!');
      setTimeout(() => router.push('/'), 600);
    } catch {
      toast.error('Logout failed.');
    }
  };

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
      if (!res.ok) throw new Error(data.error || 'Error sending OTP');
      toast.success('OTP sent!');
      setIsOtpOpen(true);
    } catch (err: any) {
      toast.error(err.message);
    } finally {
      setSendingOtp(false);
    }
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

  if (!user) {
    return <div className="flex justify-center items-center min-h-screen text-gray-300">Loading…</div>;
  }

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

      <Navbar user={user} onLoginClick={() => router.push('/')} />

      <main className="relative z-20 flex flex-col items-center justify-center min-h-[80vh] px-6">
        <motion.div
          initial={{ opacity: 0, y: 20 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.6 }}
          className="bg-white/5 backdrop-blur-lg border border-white/10 rounded-3xl p-10 max-w-md w-full text-center shadow-lg"
        >

          <div className="relative w-full flex flex-col items-center">
            <button
              onClick={openMessages}
              className="absolute top-0 right-0 text-blue-400 hover:text-blue-300"
            >
              <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" 
                strokeWidth={1.5} stroke="currentColor" fill="none" 
                className="w-9 h-9">
                <path strokeLinecap="round" strokeLinejoin="round"
                  d="M8.25 9.75h7.5m-7.5 3h4.5m-2.25-9a9 9 0 110 18 
                     8.96 8.96 0 01-4.772-1.383L3 21l1.633-3.042A8.97 
                     8.97 0 013.75 12c0-4.97 4.03-9 9-9z" />
              </svg>
            </button>

            <motion.div
              className="w-28 h-28 relative"
              whileTap={{ scale: 0.95 }}
              onClick={() => setFlipped(!flipped)}
              style={{ perspective: 800 }}
            >
              <motion.div
                animate={{ rotateY: flipped ? 180 : 0 }}
                transition={{ duration: 0.6 }}
                className="relative w-full h-full rounded-full"
                style={{ transformStyle: 'preserve-3d' }}
              >
                {user.profileImage ? (
                  <img src={user.profileImage} className="absolute inset-0 w-full h-full rounded-full object-cover border-2 border-indigo-400" />
                ) : (
                  <div className="absolute inset-0 w-full h-full rounded-full flex items-center justify-center border-2 border-indigo-400" style={{ background: "#374151" }}>
                    <span className="text-white font-semibold text-xl">{makePlaceholder(user._id, user.fullName)}</span>
                  </div>
                )}

                {user.selfieImage ? (
                  <img src={user.selfieImage} className="absolute inset-0 w-full h-full rounded-full object-cover border-2 border-purple-400" style={{ transform: 'rotateY(180deg)' }} />
                ) : (
                  <div className="absolute inset-0 w-full h-full rounded-full flex items-center justify-center border-2 border-purple-400" style={{ transform: 'rotateY(180deg)', background: "#4b5563" }}>
                    <span className="text-white font-semibold text-xl">{makePlaceholder(user._id, user.fullName)}</span>
                  </div>
                )}
                </motion.div>
              </motion.div>

            <h3 className="mt-4 text-xl font-bold">{user.fullName}</h3>
            <p className="text-gray-400">{user.rollNumber}</p>
            <p className="text-gray-400">{user.program} — {user.major}</p>
          </div>

          <div className="mt-6 space-y-2 text-sm text-gray-300 text-left">
            <p><span className="text-white font-semibold">Email:</span> {user.email}</p>
            <p><span className="text-white font-semibold">Role:</span> {user.role}</p>
            <div className="mt-4 text-sm text-gray-300">
              <p>
                <span className="text-white font-semibold">Reputation:</span>
                {reputation === null ? " Loading..." : ` ${reputation}`}
              </p>
            </div>
          </div>

          {/* Uploaded resources */}
          <div className="mt-6 text-left w-full">
            <div className="flex items-center justify-between mb-2">
              <div className="text-white font-semibold">My Uploads</div>
              <div>
                <button
                  onClick={() => router.push('/resources/manage')}
                  className="text-xs px-2 py-1 bg-amber-600 rounded"
                >
                  Manage Requests
                </button>
              </div>
            </div>
            {uploads.length === 0 ? (
              <div className="text-gray-400 text-sm">No uploads yet.</div>
            ) : (
              <div className="space-y-2">
                {uploads.map((u:any) => (
                  <div key={u._id} className="flex items-center justify-between bg-white/6 p-2 rounded">
                    <div className="text-sm">{u.title}</div>
                    <div className="flex items-center gap-2">
                      <Link href={`/resources/${u._id}`} className="text-xs px-2 py-1 bg-indigo-600 rounded">
                        View
                      </Link>
                      <div className="text-xs text-gray-300">{u.isPublic ? 'Public' : 'Private'}</div>
                    </div>
                  </div>
                ))}
              </div>
            )}
          </div>

          <div className="flex flex-col mt-8 gap-3">
            {!user.mailVerified && (
              <>
                <button
                  onClick={handleSendOtp}
                  disabled={sendingOtp}
                  className="px-6 py-2 rounded-lg bg-indigo-600"
                >
                  {sendingOtp ? 'Sending...' : 'Verify Email'}
                </button>

                {isOtpOpen && (
                  <OtpModal
                    email={user.email}
                    onClose={() => setIsOtpOpen(false)}
                    onVerified={() => {
                      setUser({ ...user, mailVerified: true });
                      setIsOtpOpen(false);
                    }}
                  />
                )}
              </>
            )}

            {!user.faceVerified && user.mailVerified && (
              <button
                onClick={() => router.push('/verify')}
                className="px-6 py-2 rounded-lg bg-green-600"
              >
                Verify Face
              </button>
            )}

            <button
              onClick={handleLogout}
              className="text-red-300 text-sm mt-4"
            >
              Log out
            </button>
          </div>

        </motion.div>
      </main>

      <footer className="absolute bottom-6 text-center w-full text-gray-400 text-sm">
        GradLedger © {new Date().getFullYear()}
      </footer>
    </div>
  );
}
