'use client';

import { useState } from 'react';
import { motion, AnimatePresence } from 'framer-motion';

interface AuthDialogProps {
  isOpen: boolean;
  onClose: () => void;
}

export default function AuthDialog({ isOpen, onClose }: AuthDialogProps) {
  const [isRegister, setIsRegister] = useState(false);
  const [animating, setAnimating] = useState(false);

  if (!isOpen) return null;

  const fadeUp = (delay: number) => ({
    initial: { opacity: 0, y: 20 },
    animate: { opacity: 1, y: 0, transition: { delay, duration: 0.4 } },
    exit: { opacity: 0, y: -20, transition: { duration: 0.3 } },
  });

  const handleToggle = () => {
    if (animating) return;
    setAnimating(true);
    setTimeout(() => {
      setIsRegister(!isRegister);
      setTimeout(() => setAnimating(false), 800);
    }, 200);
  };

  return (
    <div className="fixed inset-0 flex items-center justify-center bg-black/60 backdrop-blur-md z-50 m-0 p-0">
      <AnimatePresence>
        <motion.div
          key="auth"
          initial={{ opacity: 0, scale: 0.9, y: 20 }}
          animate={{ opacity: 1, scale: 1, y: 0 }}
          exit={{ opacity: 0, scale: 0.9, y: -20 }}
          transition={{ duration: 0.3 }}
          className="relative bg-gradient-to-br from-[#0a0b14]/90 via-[#0e1c1c]/85 to-[#131336]/90 rounded-3xl shadow-[0_0_35px_rgba(0,0,0,0.6)] overflow-hidden w-[780px] h-[460px] flex border border-white/10 backdrop-blur-xl"
        >
          <button
            onClick={onClose}
            className="absolute top-3 right-4 text-gray-300 hover:text-white text-xl font-bold z-50"
          >
            ✕
          </button>

          {/* Left Side */}
          <motion.div
            animate={{ x: isRegister ? '100%' : '0%' }}
            transition={{ duration: 0.8, ease: 'easeInOut' }}
            className="w-1/2 h-full flex flex-col justify-center px-10 absolute left-0 bg-white/5 backdrop-blur-md rounded-l-3xl border-r border-white/10"
          >
            <AnimatePresence mode="wait">
              <motion.div
                key={isRegister ? 'register' : 'login'}
                {...fadeUp(0)}
                className="w-full"
              >
                <h2 className="text-3xl font-bold text-center text-white mb-6">
                  {isRegister ? 'Sign Up' : 'Sign In'}
                </h2>

                <form className="space-y-5">
                  {isRegister && (
                    <motion.div {...fadeUp(0.2)} className="relative">
                      <input
                        type="text"
                        required
                        className="peer w-full border-b-2 border-gray-500 text-white outline-none focus:border-indigo-400 bg-transparent"
                      />
                      <label className="absolute left-0 top-1/2 -translate-y-1/2 text-gray-400 peer-focus:text-xs peer-focus:-translate-y-6 transition-all">
                        Username
                      </label>
                    </motion.div>
                  )}

                  <motion.div {...fadeUp(0.3)} className="relative">
                    <input
                      type="email"
                      required
                      className="peer w-full border-b-2 border-gray-500 text-white outline-none focus:border-indigo-400 bg-transparent"
                    />
                    <label className="absolute left-0 top-1/2 -translate-y-1/2 text-gray-400 peer-focus:text-xs peer-focus:-translate-y-6 transition-all">
                      Email
                    </label>
                  </motion.div>

                  <motion.div {...fadeUp(0.4)} className="relative">
                    <input
                      type="password"
                      required
                      className="peer w-full border-b-2 border-gray-500 text-white outline-none focus:border-indigo-400 bg-transparent"
                    />
                    <label className="absolute left-0 top-1/2 -translate-y-1/2 text-gray-400 peer-focus:text-xs peer-focus:-translate-y-6 transition-all">
                      Password
                    </label>
                  </motion.div>

                  <motion.button
                    {...fadeUp(0.5)}
                    type="submit"
                    className="w-full bg-gradient-to-r from-indigo-600 to-purple-600 text-white py-2 rounded-md shadow-md hover:from-indigo-700 hover:to-purple-700 transition-all"
                  >
                    {isRegister ? 'Sign Up' : 'Sign In'}
                  </motion.button>

                  <motion.p
                    {...fadeUp(0.6)}
                    className="text-center text-gray-400 mt-5 text-sm"
                  >
                    {isRegister
                      ? 'Already have an account?'
                      : "Don’t have an account?"}{' '}
                    <button
                      type="button"
                      onClick={handleToggle}
                      className="text-indigo-400 font-semibold hover:underline"
                    >
                      {isRegister ? 'Login' : 'Sign Up'}
                    </button>
                  </motion.p>
                </form>
              </motion.div>
            </AnimatePresence>
          </motion.div>

          {/* Right Side */}
          <motion.div
            animate={{
              x: isRegister ? '-100%' : '0%',
              background: isRegister
                ? 'linear-gradient(135deg, #3b1d75, #5b1bbd)'
                : 'linear-gradient(135deg, #1e3a8a, #6d28d9)',
            }}
            transition={{ duration: 0.8, ease: 'easeInOut' }}
            className="w-1/2 h-full absolute right-0 flex flex-col justify-center items-center text-center text-white px-10 rounded-r-3xl"
          >
            <AnimatePresence mode="wait">
              <motion.div
                key={isRegister ? 'welcome' : 'hello'}
                initial={{ opacity: 0 }}
                animate={{ opacity: 1 }}
                exit={{ opacity: 0 }}
                transition={{ duration: 0.6, delay: 0.5 }}
              >
                <motion.h2 className="text-3xl font-bold mb-3 drop-shadow-lg">
                  {isRegister ? 'Welcome Back!' : 'Hello, Friend!'}
                </motion.h2>
                <motion.p className="text-base text-white/90">
                  {isRegister
                    ? 'Enter your details to sign in.'
                    : 'Enter your details to get started.'}
                </motion.p>
              </motion.div>
            </AnimatePresence>
          </motion.div>
        </motion.div>
      </AnimatePresence>
    </div>
  );
}
