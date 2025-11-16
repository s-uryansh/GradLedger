import React, { useRef, useEffect, useCallback, useMemo } from "react";
import "./ProfileCard.css";

interface Props {
  avatarUrl: string;
  miniAvatarUrl?: string;
  name: string;
  title: string;
  handle: string;
  status?: string;
  contactText?: string;
  enableTilt?: boolean;
  enableMobileTilt?: boolean;
  showUserInfo?: boolean;
  onContactClick?: () => void;
}

const ProfileCard: React.FC<Props> = ({
  avatarUrl,
  miniAvatarUrl,
  name,
  title,
  handle,
  status = "Online",
  contactText = "Contact",
  enableTilt = true,
  enableMobileTilt = false,
  showUserInfo = true,
  onContactClick
}) => {
  const wrapRef = useRef<HTMLDivElement>(null);
  const shellRef = useRef<HTMLDivElement>(null);
  const enterTimer = useRef<number | null>(null);
  const leaveTimer = useRef<number | null>(null);

  const tilt = useMemo(() => {
    if (!enableTilt) return null;

    let raf: number | null = null;
    let currentX = 0;
    let currentY = 0;
    let targetX = 0;
    let targetY = 0;
    let running = false;

    const updateVars = (x: number, y: number) => {
      const wrap = wrapRef.current;
      const shell = shellRef.current;
      if (!wrap || !shell) return;

      const w = shell.clientWidth || 1;
      const h = shell.clientHeight || 1;

      const px = Math.min(Math.max((x / w) * 100, 0), 100);
      const py = Math.min(Math.max((y / h) * 100, 0), 100);

      wrap.style.setProperty("--pointer-x", `${px}%`);
      wrap.style.setProperty("--pointer-y", `${py}%`);
      wrap.style.setProperty("--rotate-x", `${-(px - 50) / 6}deg`);
      wrap.style.setProperty("--rotate-y", `${(py - 50) / 6}deg`);
    };

    const animate = (ts: number) => {
      currentX += (targetX - currentX) * 0.12;
      currentY += (targetY - currentY) * 0.12;
      updateVars(currentX, currentY);

      if (Math.abs(targetX - currentX) > 0.5 || Math.abs(targetY - currentY) > 0.5) {
        raf = requestAnimationFrame(animate);
      } else {
        running = false;
        raf = null;
      }
    };

    return {
      setTarget(x: number, y: number) {
        targetX = x;
        targetY = y;
        if (!running) {
          running = true;
          raf = requestAnimationFrame(animate);
        }
      },
      center() {
        const shell = shellRef.current;
        if (!shell) return;
        this.setTarget(shell.clientWidth / 2, shell.clientHeight / 2);
      },
      cancel() {
        if (raf) cancelAnimationFrame(raf);
      }
    };
  }, [enableTilt]);

  const move = useCallback(
    (e: PointerEvent) => {
      if (!tilt || !shellRef.current) return;
      const rect = shellRef.current.getBoundingClientRect();
      tilt.setTarget(e.clientX - rect.left, e.clientY - rect.top);
    },
    [tilt]
  );

  const enter = useCallback(() => {
    if (!tilt || !shellRef.current) return;
    shellRef.current.classList.add("active");
    if (enterTimer.current) clearTimeout(enterTimer.current);
  }, [tilt]);

  const leave = useCallback(() => {
    if (!tilt || !shellRef.current) return;
    tilt.center();
    leaveTimer.current = requestAnimationFrame(() => {
      shellRef.current?.classList.remove("active");
    });
  }, [tilt]);

  useEffect(() => {
    const shell = shellRef.current;
    if (!shell || !tilt) return;

    shell.addEventListener("pointerenter", enter);
    shell.addEventListener("pointermove", move);
    shell.addEventListener("pointerleave", leave);

    tilt.center();

    return () => {
      shell.removeEventListener("pointerenter", enter);
      shell.removeEventListener("pointermove", move);
      shell.removeEventListener("pointerleave", leave);
      tilt.cancel();
    };
  }, [tilt, enter, move, leave]);

  return (
    <div ref={wrapRef} className="pc-wrap">
      <div ref={shellRef} className="pc-shell">
        <div className="pc-card">
          <div className="pc-inner">
            <div className="pc-avatar-box">
              <img src={avatarUrl} className="pc-avatar" alt="" />
            </div>

            {showUserInfo && (
              <div className="pc-info">
                <div className="pc-mini">
                  <img
                    src={miniAvatarUrl || avatarUrl}
                    className="pc-mini-avatar"
                    alt=""
                  />
                  <div>
                    <div className="pc-handle">@{handle}</div>
                    <div className="pc-status">{status}</div>
                  </div>
                </div>

                <button onClick={onContactClick} className="pc-btn">
                  {contactText}
                </button>
              </div>
            )}

            <div className="pc-details">
              <h3>{name}</h3>
              <p>{title}</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default ProfileCard;
