'use client';

import { motion } from 'framer-motion';
import { useRouter } from 'next/navigation';
import { Frown } from 'lucide-react';
import ColorBends from '@/components/BackgroundAnimations/ColorBends';

export default function ErrorPage({
  title = 'Verification Failed',
  message = 'Something went wrong during verification. Please try again later.',
  retryPath = '/',
}: {
  title?: string;
  message?: string;
  retryPath?: string;
}) {
  const router = useRouter();

  return (
    <div className="relative min-h-screen w-full overflow-hidden text-white flex flex-col items-center justify-center px-6">
      {/* === Shared Animated Background === */}
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

      {/* === Error Card === */}
      <motion.div
        initial={{ opacity: 0, y: 30 }}
        animate={{ opacity: 1, y: 0 }}
        transition={{ duration: 0.6 }}
        className="relative bg-white/5 backdrop-blur-lg border border-white/10 rounded-3xl shadow-[0_0_30px_rgba(0,0,0,0.5)] p-10 max-w-lg text-center z-10"
      >
        <motion.div
          initial={{ scale: 0 }}
          animate={{ scale: 1 }}
          transition={{ type: 'spring', stiffness: 180, damping: 12, delay: 0.2 }}
          className="w-20 h-20 mx-auto flex items-center justify-center rounded-full bg-red-500/10 border border-red-500/30 mb-6"
        >
          <Frown className="text-red-400 w-10 h-10" />
        </motion.div>

        <h1 className="text-3xl font-bold mb-3 bg-gradient-to-r from-red-400 to-pink-500 text-transparent bg-clip-text">
          {title}
        </h1>
        <p className="text-gray-300 mb-8">{message}</p>

        <motion.button
          whileHover={{ scale: 1.05 }}
          whileTap={{ scale: 0.95 }}
          onClick={() => router.push(retryPath)}
          className="px-6 py-3 rounded-lg bg-gradient-to-r from-indigo-600 to-purple-600 font-semibold text-white shadow-lg hover:from-indigo-700 hover:to-purple-700 transition-all"
        >
          Try Again
        </motion.button>
      </motion.div>

      {/* === Footer === */}
      <div className="absolute bottom-6 text-sm text-gray-400">
        GradLedger © {new Date().getFullYear()}
      </div>
    </div>
  );
}
