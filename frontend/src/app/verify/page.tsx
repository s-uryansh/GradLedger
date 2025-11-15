'use client';

import { useState, useEffect } from 'react';
import { motion } from 'framer-motion';
import ColorBends from '@/components/BackgroundAnimations/ColorBends';
import Navbar from '@/components/UI/Navbar';
import { verifyFaces } from '@/lib/pythonAPI';
import { useRouter } from 'next/navigation';
import toast from 'react-hot-toast';

export default function VerifyPage() {
  const [image1, setImage1] = useState<File | null>(null);
  const [image2, setImage2] = useState<File | null>(null);
  const [result, setResult] = useState<any>(null);
  const [loading, setLoading] = useState(false);
  const [user, setUser] = useState<any>(null);
  const [loadingUser, setLoadingUser] = useState(true);
  const router = useRouter();

  useEffect(() => {
    const fetchUser = async () => {
      try {
        const res = await fetch('/api/auth/me', { credentials: 'include' });
        const data = await res.json();

        if (!data.user) {
          router.push('/');
          return;
        }

        if (!data.user.mailVerified) {
          toast.error('Please verify your email first.');
          router.push('/profile');
          return;
        }

        setUser(data.user);
      } catch (err) {
        console.error('Failed to fetch user:', err);
        router.push('/');
      } finally {
        setLoadingUser(false);
      }
    };
    fetchUser();
  }, [router]);

  if (loadingUser)
    return (
      <div className="flex items-center justify-center min-h-screen text-gray-300 text-lg">
        Checking verification status...
      </div>
    );


  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!image1 || !image2) return;

    setLoading(true);
    setResult(null);

    try {
      const selfieBase64 = await new Promise<string>((resolve) => {
        const reader = new FileReader();
        reader.onloadend = () => resolve(reader.result as string);
        reader.readAsDataURL(image1);
      });
      // Verify face match
      const res = await verifyFaces(image1, image2);
      const confidence = Math.round(res.confidence * 100) / 100;
      const isMatch = res.match && confidence >= 0.1;

      setResult({
        match: isMatch,
        confidence,
        message: !isMatch
          ? 'No match detected. Confidence too low.'
          : confidence >= 0.7
          ? 'High confidence match — verified successfully.'
          : 'Moderate confidence match.',
      });

      if (!isMatch) {
        toast.error('Face mismatch. Please try again.');
        return;
      }

      // Upload ID and extract data (crop + OCR)
      const formData = new FormData();
      formData.append('file', image2);
      formData.append('email', user.email);

      const uploadRes = await fetch('/api/upload/id-face', {
        method: 'POST',
        body: formData,
      });
      const uploadData = await uploadRes.json();

      const res2 = await fetch('/api/verify/face', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        credentials: 'include',
        body: JSON.stringify({
          email: user.email,
          result: { match: isMatch, confidence },
          profileImage: uploadData.url,
          selfieImage: selfieBase64,
          rollNumber: uploadData.rollNumber,
          program: uploadData.program,
          major: uploadData.major,
        }),
      });

      const data = await res2.json();
      if (data.success) {
        toast.success('Verification completed successfully!');
        setTimeout(() => router.push('/'), 1500);
      } else {
        toast.error('Failed to update profile.');
      }
    } catch (err) {
      console.error('Verification error:', err);
      toast.error('Error verifying faces. Please try again.');
    } finally {
      setLoading(false);
    }
  };

  const resetImages = () => {
    setImage1(null);
    setImage2(null);
    setResult(null);
  };

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
          {/* Selfie */}
          <div className="flex flex-col items-center space-y-3">
            <p className="text-gray-300">Capture Live Selfie</p>
            {!image1 ? (
              <>
                <video id="video" autoPlay playsInline className="rounded-lg border border-white/10 w-64 h-48" />
                <div className="flex gap-3 mt-2">
                  <button
                    type="button"
                    onClick={async () => {
                      const stream = await navigator.mediaDevices.getUserMedia({ video: true });
                      const video = document.getElementById('video') as HTMLVideoElement;
                      if (video) video.srcObject = stream;
                    }}
                    className="px-4 py-2 bg-gradient-to-r from-indigo-700 to-purple-600 rounded-md font-medium"
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
                      });
                    }}
                    className="px-4 py-2 bg-gradient-to-r from-green-600 to-emerald-500 rounded-md font-medium"
                  >
                    Capture
                  </button>
                </div>
              </>
            ) : (
              <>
                <img src={URL.createObjectURL(image1)} className="w-64 h-48 rounded-lg object-cover" />
                <button onClick={resetImages} className="text-sm text-red-400 mt-2">Retake</button>
              </>
            )}
          </div>

          {/* ID Card */}
          <div className="flex flex-col items-center space-y-3 mt-6">
            <p className="text-gray-300">Upload ID Card</p>
            <label htmlFor="id-card" className="cursor-pointer px-5 py-2 rounded-md bg-gradient-to-r from-indigo-700 to-purple-600 text-white font-medium shadow-md">
              Choose File
            </label>
            <input id="id-card" type="file" accept="image/*" onChange={(e) => setImage2(e.target.files?.[0] || null)} className="hidden" />
            {image2 && <p className="text-sm text-gray-400">{image2.name}</p>}
          </div>

          <button
            type="submit"
            disabled={loading || !image1 || !image2}
            className={`w-full py-2 rounded-md font-semibold mt-4 ${
              loading ? 'bg-gray-600 cursor-not-allowed' : 'bg-gradient-to-r from-indigo-600 to-purple-600 hover:from-indigo-700 hover:to-purple-700'
            }`}
          >
            {loading ? 'Verifying...' : 'Verify'}
          </button>
        </form>

        {result && (
          <div className={`mt-8 p-4 rounded-xl ${result.match ? 'bg-green-500/20' : 'bg-red-500/20'}`}>
            <h3 className={`text-xl font-bold ${result.match ? 'text-green-400' : 'text-red-400'}`}>
              {result.match ? 'Verification Successful' : 'Verification Failed'}
            </h3>
            <p className="text-gray-300 text-sm mt-2">Confidence: {result.confidence}</p>
          </div>
        )}
      </motion.div>
    </div>
  );
}
