'use client';

import { useEffect, useRef, useState } from 'react';
import { useParams, useRouter } from 'next/navigation';
import io from 'socket.io-client';
import Navbar from '@/components/UI/Navbar';
import { FiSend } from 'react-icons/fi';

let socket: any = null;

export default function ChatPage() {
  const params = useParams();
  const router = useRouter();
  const convoId = params.conversationId as string;

  const [user, setUser] = useState<any>(null);
  const [other, setOther] = useState<any>(null);
  const [messages, setMessages] = useState<any[]>([]);
  const [input, setInput] = useState('');

  const bottomRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    if (!user || !other) return;

    fetch('/api/messages/mark-seen', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ conversationId: convoId, userId: user._id })
    });
  }, [user, other, convoId]);

  useEffect(() => {
    const fetchUser = async () => {
      const res = await fetch('/api/auth/me', { credentials: 'include' });
      const data = await res.json();
      if (!data.user) return router.push('/');
      if (!data.user.mailVerified || !data.user.faceVerified) return router.push('/profile');
      setUser(data.user);
    };
    fetchUser();
  }, [router]);

  useEffect(() => {
    if (!user) return;

    const load = async () => {
      const res = await fetch(`/api/messages/conversation/${convoId}`, {
        headers: {
          'x-user-id': user._id,
        },
      });

      const data = await res.json();
      setOther(data.other);
      setMessages(data.messages);
    };

    load();
  }, [user, convoId]);

  useEffect(() => {
    if (!user) return;

    socket = io('/', {
      path: '/socket',
      transports: ['websocket'],
    });

    socket.emit('online', user._id);

    socket.on('receive-message', (msg: any) => {
      if (msg.conversationId === convoId) {
        setMessages((prev) => [...prev, msg]);
      }
    });

    return () => socket.disconnect();
  }, [user, other, convoId]);

  useEffect(() => {
    bottomRef.current?.scrollIntoView({ behavior: 'smooth' });
  }, [messages]);

  const send = async () => {
    if (!input.trim() || !user || !other) return;

    const payload = {
      conversationId: convoId,
      sender: user._id,
      receiver: other._id,       // FIXED
      text: input.trim(),
    };

    setMessages((prev) => [...prev, payload]);
    setInput('');

    socket.emit('send-message', payload);

    await fetch('/api/messages/send', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(payload),
    });
  };

  if (!user || !other) {
    return (
      <div className="text-gray-300 min-h-screen flex items-center justify-center">
        Loading...
      </div>
    );
  }

  return (
    <div className="relative min-h-screen text-white">
      <Navbar user={user} onLoginClick={() => {}} />

      <div className="pt-24 max-w-2xl mx-auto px-4">
        <div className="flex items-center gap-3 p-3 bg-white/10 rounded-lg mb-4">
          <img
            src={other.profileImage}
            className="w-12 h-12 rounded-full object-cover"
          />
          <div>
            <p className="font-semibold">{other.fullName}</p>
          </div>
        </div>

        <div className="h-[60vh] overflow-y-auto bg-white/5 rounded-lg p-4 flex flex-col gap-2 scrollbar-none">
          {messages.map((m, i) => (
            <div
              key={i}
              className={`max-w-[70%] px-3 py-2 rounded-lg text-sm ${
                m.sender === user._id
                  ? 'self-end bg-blue-600'
                  : 'self-start bg-gray-700'
              }`}
            >
              {m.text}
            </div>
          ))}
          <div ref={bottomRef}></div>
        </div>

        <div className="flex items-center gap-3 mt-4 bg-white/10 px-4 py-2 rounded-lg">
          <input
            className="flex-1 bg-transparent outline-none text-gray-200"
            placeholder="Type a message..."
            value={input}
            onChange={(e) => setInput(e.target.value)}
            onKeyDown={(e) => e.key === 'Enter' && send()}
          />

          <button onClick={send} className="text-blue-400 hover:text-blue-300 transition">
            <FiSend size={22} />
          </button>
        </div>
      </div>
    </div>
  );
}
