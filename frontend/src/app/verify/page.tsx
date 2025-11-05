'use client';

import { useState, useEffect } from 'react';
import { motion } from 'framer-motion';
import ColorBends from '@/components/BackgroundAnimations/ColorBends';
import Navbar from '@/components/UI/Navbar';
import { verifyFaces } from '@/lib/pythonAPI';
import { useRouter } from 'next/navigation';

export default function VerifyPage() {
  const [image1, setImage1] = useState<File | null>(null);
  const [image2, setImage2] = useState<File | null>(null);
  const [result, setResult] = useState<any>(null);
  const [loading, setLoading] = useState(false);
  const [user, setUser] = useState<any>(null);
  const router = useRouter();

  useEffect(() => {
    const fetchUser = async () => {
      try {
        const res = await fetch('/api/auth/me', { credentials: 'include' });
        const data = await res.json();
        if (data?.user) setUser(data.user);
      } catch (err) {
        console.error('Failed to fetch user:', err);
      }
    };
    fetchUser();
  }, []);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!image1 || !image2 || result?.match === false) return; // block mismatch retries

    setLoading(true);
    setResult(null);

    try {
      const res = await verifyFaces(image1, image2);
      const confidence = Math.round(res.confidence * 100) / 100;
      const isMatch = res.match && confidence >= 0.1;

      const normalized = {
        ...res,
        confidence,
        match: isMatch,
        message: !isMatch
          ? 'No match detected. Confidence too low.'
          : confidence >= 0.7
          ? 'High confidence match — verified successfully.'
          : confidence >= 0.4
          ? 'Moderate confidence match — likely same person.'
          : 'Low confidence match — proceed with caution.',
      };

      setResult(normalized);

      if (isMatch && user?.email) {
        const res2 = await fetch('/api/verify/face', {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          credentials: 'include',
          body: JSON.stringify({
            email: user.email,
            result: { match: isMatch, confidence },
          }),
        });

        const data = await res2.json();
        if (data.success) {
          setTimeout(() => router.push('/'), 1500);
        }
      }
    } catch (err) {
      console.error('Verification error:', err);
      setResult({
        match: false,
        confidence: 0,
        message: 'Error verifying faces. Please try again.',
      });
    } finally {
      setLoading(false);
    }
  };

  const resetImages = () => {
    setImage1(null);
    setImage2(null);
    setResult(null);
  };

  const verifyDisabled =
    loading ||
    !image1 ||
    !image2 ||
    (result && result.match === false); // block if mismatch

  return (
    <div className="relative min-h-screen w-full text-white flex flex-col items-center justify-center">
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

      <Navbar onLoginClick={() => {}} />

      <motion.div
        initial={{ opacity: 0, y: 20 }}
        animate={{ opacity: 1, y: 0 }}
        transition={{ duration: 0.6 }}
        className="bg-white/5 backdrop-blur-lg border border-white/10 rounded-3xl p-10 shadow-xl max-w-lg w-full text-center mt-20"
      >
        <h2 className="text-3xl font-bold mb-6 bg-gradient-to-r from-indigo-400 to-purple-400 text-transparent bg-clip-text">
          Face Verification
        </h2>

        <form onSubmit={handleSubmit} className="space-y-4">
          {/* Capture Selfie */}
          <div className="flex flex-col items-center space-y-3">
            <p className="text-gray-300">Capture Live Selfie</p>
            {!image1 ? (
              <>
                <video
                  id="video"
                  autoPlay
                  playsInline
                  className="rounded-lg border border-white/10 shadow-md w-64 h-48 object-cover"
                />
                <div className="flex gap-3 mt-2">
                  <button
                    type="button"
                    onClick={async () => {
                      try {
                        const stream = await navigator.mediaDevices.getUserMedia({ video: true });
                        const video = document.getElementById('video') as HTMLVideoElement;
                        if (video) video.srcObject = stream;
                      } catch (err) {
                        console.error('Camera access error:', err);
                      }
                    }}
                    className="px-4 py-2 bg-gradient-to-r from-indigo-700 to-purple-600 rounded-md font-medium hover:from-indigo-800 hover:to-purple-700 hover:shadow-[0_0_10px_rgba(99,102,241,0.6)] transition-all"
                  >
                    Start Camera
                  </button>

                  <button
                    type="button"
                    onClick={() => {
                      const video = document.getElementById('video') as HTMLVideoElement;
                      if (!video || !video.srcObject) return;
                      const canvas = document.createElement('canvas');
                      canvas.width = video.videoWidth;
                      canvas.height = video.videoHeight;
                      const ctx = canvas.getContext('2d');
                      if (!ctx) return;
                      ctx.drawImage(video, 0, 0, canvas.width, canvas.height);
                      canvas.toBlob((blob) => {
                        if (blob) {
                          const file = new File([blob], 'selfie.jpg', { type: 'image/jpeg' });
                          setImage1(file);
                          const stream = video.srcObject as MediaStream;
                          stream.getTracks().forEach((t) => t.stop());
                          video.srcObject = null;
                        }
                      }, 'image/jpeg');
                    }}
                    className="px-4 py-2 bg-gradient-to-r from-green-600 to-emerald-500 rounded-md font-medium hover:from-green-700 hover:to-emerald-600 hover:shadow-[0_0_10px_rgba(34,197,94,0.6)] transition-all"
                  >
                    Capture
                  </button>
                </div>
              </>
            ) : (
              <>
                <img
                  src={URL.createObjectURL(image1)}
                  alt="Captured selfie"
                  className="rounded-lg w-64 h-48 object-cover border border-white/10 shadow-md"
                />
                <button
                  type="button"
                  onClick={resetImages}
                  className="text-sm text-red-400 hover:text-red-300 mt-2"
                >
                  Retake
                </button>
              </>
            )}
          </div>

          {/* Upload ID Card */}
          <div className="flex flex-col items-center space-y-3 mt-6">
            <p className="text-gray-300">Upload ID Card</p>
            <label
              htmlFor="id-card"
              className="cursor-pointer px-5 py-2 rounded-md bg-gradient-to-r from-indigo-700 to-purple-600 text-white font-medium shadow-md hover:from-indigo-800 hover:to-purple-700 hover:shadow-[0_0_15px_rgba(99,102,241,0.6)] transition-all"
            >
              Choose File
            </label>
            <input
              id="id-card"
              type="file"
              accept="image/*"
              onChange={(e) => setImage2(e.target.files?.[0] || null)}
              className="hidden"
            />
            {image2 && <p className="text-sm text-gray-400">{image2.name}</p>}
          </div>

          <button
            type="submit"
            disabled={verifyDisabled}
            className={`w-full py-2 rounded-md font-semibold shadow-md mt-4 transition-all ${
              verifyDisabled
                ? 'bg-gray-600 cursor-not-allowed opacity-70'
                : 'bg-gradient-to-r from-indigo-600 to-purple-600 hover:from-indigo-700 hover:to-purple-700 hover:shadow-[0_0_15px_rgba(99,102,241,0.6)]'
            }`}
          >
            {loading ? 'Verifying...' : 'Verify'}
          </button>
        </form>

        {result && (
          <motion.div
            initial={{ opacity: 0, scale: 0.9 }}
            animate={{ opacity: 1, scale: 1 }}
            transition={{ duration: 0.4, ease: 'easeOut' }}
            className={`mt-8 mx-auto text-center rounded-xl p-5 w-full shadow-lg backdrop-blur-md ${
              result.match
                ? 'bg-green-500/20 border border-green-400/40'
                : 'bg-red-500/20 border border-red-400/40'
            }`}
          >
            <h3
              className={`text-2xl font-bold mb-2 ${
                result.match ? 'text-green-400' : 'text-red-400'
              }`}
            >
              {result.match ? 'Face Match Successful' : 'Face Mismatch'}
            </h3>
          </motion.div>
        )}
      </motion.div>

      <footer className="absolute bottom-6 text-sm text-gray-400">
        GradLedger © {new Date().getFullYear()}
      </footer>
    </div>
  );
}
