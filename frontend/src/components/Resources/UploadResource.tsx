'use client';
import React, { useState } from "react";
import { useRouter } from "next/navigation";

export default function UploadResource({ ownerId }: { ownerId: string }) {
  const [title, setTitle] = useState("");
  const [subject, setSubject] = useState("");
  const [tags, setTags] = useState("");
  const [isPublic, setIsPublic] = useState(true);
  const [file, setFile] = useState<File | null>(null);
  const router = useRouter();
  const [loading, setLoading] = useState(false);

  const handleFile = (f: File | null) => setFile(f);

  const submit = async () => {
    if (!ownerId || !title || !file) return alert("owner, title, file required");
    setLoading(true);
    const buf = await file.arrayBuffer();
    const b64 = Buffer.from(buf).toString("base64");
    const res = await fetch("/api/resources/create", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({
        ownerId,
        title,
        description: "",
        category: "Notes",
        subject,
        tags: tags.split(",").map(t => t.trim()).filter(Boolean),
        isPublic,
        fileName: file.name,
        fileDataBase64: b64
      })
    });
    const data = await res.json();
    setLoading(false);
    if (!res.ok) return alert(data.error || "Upload failed");
    router.push("/resources"); 
  };

  return (
    <div className="p-4 text-white">
      <input value={title} onChange={e => setTitle(e.target.value)} placeholder="Title" className="mb-2 w-full p-2 rounded bg-white/5"/>
      <input value={subject} onChange={e => setSubject(e.target.value)} placeholder="Subject" className="mb-2 w-full p-2 rounded bg-white/5"/>
      <input value={tags} onChange={e => setTags(e.target.value)} placeholder="tags comma separated" className="mb-2 w-full p-2 rounded bg-white/5"/>
      <div className="mb-2">
        <label className="mr-3"><input type="checkbox" checked={isPublic} onChange={e => setIsPublic(e.target.checked)}/> Public</label>
      </div>
      <input type="file" onChange={(e) => handleFile(e.target.files?.[0] || null)} className="mb-2"/>
      <button onClick={submit} disabled={loading} className="px-4 py-2 bg-blue-600 rounded">
        {loading ? "Uploading..." : "Upload"}
      </button>
    </div>
  );
}
