import { useEffect, useRef } from "react";
import io from "socket.io-client";

export const useSocket = (email: string) => {
  const socketRef = useRef<any>(null);

  useEffect(() => {
    const socket = io("/", {
      path: "/socket",
      transports: ["websocket"]
    });

    socketRef.current = socket;

    socket.emit("online", email);

    return () => {
      socket.disconnect();
    };
  }, [email]);

  return socketRef;
};
