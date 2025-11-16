"use client";

import ColorBends from "@/components/BackgroundAnimations/ColorBends";

export default function BackgroundWrapper({ children }: { children: React.ReactNode }) {
  return (
    <div className="relative min-h-screen w-full overflow-x-hidden text-white">

      <div className="fixed inset-0 -z-30">
        <ColorBends
          colors={["#3e47f4", "#06b31a", "#b46d04"]}
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
              "linear-gradient(to bottom, rgba(0,0,0,0.25) 0%, rgba(0,0,0,0.05) 25%, rgba(0,0,0,0.10) 70%, rgba(0,0,0,0.15) 100%)",
          }}
        />
      </div>

      {children}
    </div>
  );
}
