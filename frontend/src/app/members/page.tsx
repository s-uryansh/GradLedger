"use client";

import { useEffect, useState } from "react";
import ProfileCard from "@/components/ProfileCard/ProfileCard";
import Navbar from "@/components/UI/Navbar";
import { useRouter } from "next/navigation";

export default function MembersPage() {
  const router = useRouter();

  const [viewer, setViewer] = useState<any>(null);
  const [users, setUsers] = useState<any[]>([]);
  const [query, setQuery] = useState("");
  const [page, setPage] = useState(1);
  const [totalPages, setTotalPages] = useState(1);

  useEffect(() => {
    const fetchViewer = async () => {
      const res = await fetch("/api/auth/me", { credentials: "include" });
      const data = await res.json();
      if (!data.user) return router.push("/");
      if (!data.user.mailVerified || !data.user.faceVerified)
        return router.push("/profile");

      setViewer(data.user);
    };
    fetchViewer();
  }, []);

  useEffect(() => {
    if (!viewer) return;

    const loadUsers = async () => {
      const res = await fetch(`/api/users/list?page=${page}&query=${query}`);
      const data = await res.json();

      setUsers(data.users);
      setTotalPages(data.totalPages);
    };
    loadUsers();
  }, [viewer, page, query]);

  if (!viewer) return <div className="text-white p-10">Loading...</div>;

  return (
    <div className="text-white min-h-screen px-6">
      <Navbar user={viewer} />

      <div className="mt-24 max-w-4xl mx-auto">

        <input
          value={query}
          onChange={(e) => {
            setQuery(e.target.value);
            setPage(1);
          }}
          placeholder="Search full name"
          className="w-full p-3 rounded-lg bg-white/10 border border-white/20 text-white mb-8"
        />

        <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
          {users.map((u) => (
            <ProfileCard
              key={u._id}
              name={u.fullName}
              title={`${u.program || ""} ${u.major ? "— " + u.major : ""}`}
              handle={u.fullName.toLowerCase().replace(/\s+/g, "")}
              status="Verified"
              avatarUrl={u.selfieImage || u.profileImage}
              showUserInfo={false}
              enableTilt={true}
              contactText={u._id === viewer._id ? "Messages" : "Chat"}
              onContactClick={() =>
                router.push(`/profile/${u._id}`)
              }
            />
          ))}
        </div>

        <div className="flex justify-between mt-10">
          <button
            disabled={page <= 1}
            className="px-4 py-2 bg-white/10 rounded disabled:opacity-40"
            onClick={() => setPage(page - 1)}
          >
            Prev
          </button>

          <button
            disabled={page >= totalPages}
            className="px-4 py-2 bg-white/10 rounded disabled:opacity-40"
            onClick={() => setPage(page + 1)}
          >
            Next
          </button>
        </div>
      </div>
    </div>
  );
}
