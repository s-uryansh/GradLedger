'use client';

import { useEffect, useState, useRef } from "react";
import { motion } from "framer-motion";
import { useSocket } from "@/lib/useSocket";

interface ChatModalProps {
  me: any;            
  partner: any;       
  conversationId: string;
  onClose: () => void;
}

interface ChatMessage {
  _id?: string;
  conversationId: string;
  sender: string;
  receiver: string;
  text: string;
  createdAt?: string;
}

export default function ChatModal({
  me,
  partner,
  conversationId,
  onClose
}: ChatModalProps) {

  const [messages, setMessages] = useState<ChatMessage[]>([]);
  const [message, setMessage] = useState("");
  const bottomRef = useRef<HTMLDivElement | null>(null);

  const socket = useSocket(me._id); 
  useEffect(() => {
    (async () => {
      const res = await fetch(`/api/messages/${conversationId}`);
      const data = await res.json();
      setMessages(data || []);
    })();
  }, [conversationId]);

  useEffect(() => {
    if (!socket.current) return;

    socket.current.on("receive-message", (msg: ChatMessage) => {
      if (msg.conversationId === conversationId) {
        setMessages(prev => [...prev, msg]);
      }
    });

    return () => {
      socket.current?.off("receive-message");
    };
  }, [socket, conversationId]);

  const sendMessage = async () => {
    if (!message.trim()) return;

    const msg: ChatMessage = {
      conversationId,
      sender: me._id,
      receiver: partner._id,
      text: message.trim(),
    };

    setMessages(prev => [...prev, msg]);
    setMessage("");

    socket.current?.emit("send-message", msg);

    await fetch("/api/messages/send", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(msg),
    });
  };

  useEffect(() => {
    bottomRef.current?.scrollIntoView({ behavior: "smooth" });
  }, [messages]);

  return (
    <motion.div
      className="fixed inset-0 flex items-center justify-center bg-black/60 z-50"
      initial={{ opacity: 0 }}
      animate={{ opacity: 1 }}
    >
      <div className="bg-white/10 backdrop-blur-lg w-[420px] h-[520px] rounded-xl p-4 flex flex-col border border-white/10">
        <h3 className="text-center text-white font-semibold mb-3">
          Chat with {partner.fullName}
        </h3>

        {/* Messages */}
        <div className="flex-1 overflow-y-auto space-y-2 p-2">
          {messages.map((m, i) => (
            <div
              key={i}
              className={`px-3 py-2 text-sm rounded-lg max-w-[70%] ${
                m.sender === me._id
                  ? "self-end bg-indigo-600 text-white"
                  : "self-start bg-gray-700 text-gray-200"
              }`}
            >
              {m.text}
            </div>
          ))}
          <div ref={bottomRef} />
        </div>

        {/* Input */}
        <div className="flex gap-2 mt-2">
          <input
            className="flex-1 px-3 py-2 bg-gray-900 text-white rounded-md border border-gray-700"
            value={message}
            onChange={(e) => setMessage(e.target.value)}
            placeholder="Type message..."
            onKeyDown={(e) => e.key === "Enter" && sendMessage()}
          />
          <button
            onClick={sendMessage}
            className="px-4 py-2 bg-blue-600 text-white rounded-md"
          >
            Send
          </button>
        </div>

        {/* Close */}
        <button
          onClick={onClose}
          className="text-gray-400 text-sm mt-2 hover:text-gray-200"
        >
          Close
        </button>
      </div>
    </motion.div>
  );
}
