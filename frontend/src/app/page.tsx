'use client';

import { useState } from 'react';
import dynamic from 'next/dynamic';
import { motion } from 'framer-motion';
import ColorBends from '@/components/BackgroundAnimations/ColorBends';
import Navbar from '@/components/UI/Navbar';
import AuthDialog from '@/components/Auth/AuthDialog';

const MagicBento = dynamic(() => import('@/components/BackgroundAnimations/MagicBento'), { ssr: false });

const Reveal = ({ children, delay = 0 }: { children: React.ReactNode; delay?: number }) => (
  <motion.div
    initial={{ opacity: 0, y: 18 }}
    whileInView={{ opacity: 1, y: 0 }}
    viewport={{ once: true, amount: 0.2 }}
    transition={{ duration: 0.6, delay }}
  >
    {children}
  </motion.div>
);

export default function LandingPage() {
  const [isAuthOpen, setIsAuthOpen] = useState(false);

  return (
    <div className="relative min-h-screen w-full overflow-x-hidden text-white">
      {/* === Background === */}
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
              'linear-gradient(to bottom, rgba(0,0,0,0.25) 0%, rgba(0,0,0,0.05) 25%, rgba(0,0,0,0.10) 70%, rgba(0,0,0,0.15) 100%)'
          }}
        />
      </div>

      {/* === Navbar === */}
      <div className="relative z-40">
        <Navbar onLoginClick={() => setIsAuthOpen(true)} />
      </div>

      {/* === Main Scroll Sections === */}
      <main className="relative z-20">
        {/* Hero Section */}
        <section className="min-h-screen flex items-center justify-center px-6 md:px-16">
          <div className="max-w-6xl w-full grid grid-cols-1 md:grid-cols-2 gap-8 items-center">
            <div className="relative">
              <Reveal>
                <h1 className="text-4xl md:text-6xl font-extrabold mb-4 leading-tight drop-shadow-[0_6px_20px_rgba(0,0,0,0.6)]">
                  Welcome to{' '}
                  <span className="text-transparent bg-clip-text bg-gradient-to-r from-indigo-400 via-purple-400 to-pink-500">
                    GradLedger
                  </span>
                </h1>
                <p className="text-gray-200 max-w-xl">
                  A decentralized mentorship & resource platform that pairs verified alumni with students —
                  trust, transparency, and fair rewards, built on-chain.
                </p>
                <div className="mt-6 flex gap-3">
                  <button
                    onClick={() => setIsAuthOpen(true)}
                    className="px-5 py-3 bg-gradient-to-r from-indigo-700 to-purple-600 rounded-lg font-semibold shadow"
                  >
                    Get Started
                  </button>
                  <a
                    href="#features"
                    className="px-5 py-3 rounded-lg border border-white/10 text-sm text-gray-200"
                  >
                    Learn more
                  </a>
                </div>
              </Reveal>
            </div>

            {/* <div className="w-full h-[420px] md:h-[520px] rounded-2xl overflow-hidden bg-white/5 backdrop-blur-md border border-white/10">
              <div className="w-full h-full bg-black/20 animate-pulse rounded-2xl" /> */}
            </div>
          {/* </div> */}
        </section>

        {/* Features */}
        <section id="features" className="py-20 px-6 md:px-16">
          <div className="max-w-6xl mx-auto grid grid-cols-1 lg:grid-cols-3 gap-6">
            <Reveal>
              <div className="rounded-2xl p-6 bg-white/5 backdrop-blur-md border border-white/10">
                <h3 className="text-2xl font-semibold mb-2">Why GradLedger</h3>
                <p className="text-gray-200">
                  Verified mentors, on-chain reputation - less noise, more trust. Ideal for
                  students and alumni who value provenance.
                </p>
              </div>
            </Reveal>

            <Reveal delay={0.08}>
              <div className="rounded-2xl p-6 bg-white/5 backdrop-blur-md border border-white/10">
                <h3 className="text-2xl font-semibold mb-2">Problems we fix</h3>
                <ul className="text-gray-200 space-y-2 text-sm">
                  <li>• Impersonation via fake credentials</li>
                  <li>• Centralized data control & censorship</li>
                  <li>• Unfair visibility mechanics</li>
                </ul>
              </div>
            </Reveal>

            <Reveal delay={0.16}>
              <div className="rounded-2xl p-6 bg-white/5 backdrop-blur-md border border-white/10">
                <h3 className="text-2xl font-semibold mb-2">On-chain / Off-chain</h3>
                <p className="text-gray-200 text-sm">
                  Contracts store verification, sessions, reputation, and content metadata. Heavy assets and
                  face verification live off-chain.
                </p>
              </div>
            </Reveal>
          </div>
        </section>

        {/* Interactive Grid */}
        <section className="py-20 px-6 md:px-16">
          <div className="max-w-6xl mx-auto">
            <Reveal>
              <h2 className="text-3xl font-bold mb-4">Experience the Platform</h2>
              <p className="text-gray-300 mb-6">
                Explore how verified mentors, students, and on-chain credentials interact in real time.
              </p>
            </Reveal>

            <div className="rounded-2xl p-6 bg-white/5 backdrop-blur-md border border-white/10">
              <MagicBento
                textAutoHide
                enableStars
                enableSpotlight
                enableBorderGlow
                enableTilt
                enableMagnetism
                clickEffect
                spotlightRadius={300}
                particleCount={12}
                glowColor="132, 0, 255"
              />
            </div>
          </div>
        </section>

        {/* CTA */}
        <section className="py-24 px-6 md:px-16">
          <div className="max-w-4xl mx-auto text-center">
            <Reveal>
              <h3 className="text-2xl font-bold mb-3">Ready to onboard verified mentors?</h3>
              <p className="text-gray-300 mb-6">
                Sign up and verify via our face-verification pipeline, or contact us to integrate your
                university.
              </p>
              <button
                onClick={() => setIsAuthOpen(true)}
                className="px-6 py-3 rounded-lg bg-gradient-to-r from-indigo-700 to-purple-600 font-semibold"
              >
                Start Verification
              </button>
            </Reveal>
          </div>
        </section>

        {/* Footer */}
        <footer className="py-8 px-6 md:px-16 text-center text-sm text-gray-400">
          <div className="max-w-5xl mx-auto">
            GradLedger built with transparency & privacy in mind.
          </div>
        </footer>
      </main>

      {/* Auth Dialog */}
      {isAuthOpen && <AuthDialog isOpen={isAuthOpen} onClose={() => setIsAuthOpen(false)} />}
    </div>
  );
}
