"use client";

import { useState, useEffect } from "react";
import Navbar from "@/components/UI/Navbar";
import ColorBends from "@/components/BackgroundAnimations/ColorBends";
import { useRouter } from "next/navigation";

export default function UploadResourcePage() {
  const router = useRouter();

  const [viewer, setViewer] = useState<any>(null);
  const [loadingViewer, setLoadingViewer] = useState(true);

  const [title, setTitle] = useState("");
  const [description, setDescription] = useState("");
  const [category, setCategory] = useState("PYQ");
  const [customCategory, setCustomCategory] = useState("");
  const [subject, setSubject] = useState("");
  const [tags, setTags] = useState<string[]>([]);
  const [tagInput, setTagInput] = useState("");
  const [isPublic, setIsPublic] = useState(true);
  const [file, setFile] = useState<File | null>(null);
  const [uploading, setUploading] = useState(false);

useEffect(() => {
  (async () => {
    const me = await fetch("/api/auth/me", { credentials: "include" }).then(r => r.json());
    if (!me.user) return router.push("/");
    setViewer(me.user);
    setLoadingViewer(false);
  })();
}, []);  


  const handleTagAdd = (e: any) => {
    if (e.key === "Enter" || e.key === ",") {
      e.preventDefault();
      const t = tagInput.trim().toLowerCase();
      if (t && !tags.includes(t)) {
        setTags([...tags, t]);
      }
      setTagInput("");
    }
  };

  const removeTag = (t: string) => {
    setTags(tags.filter((x) => x !== t));
  };

  const convertFileToBase64 = (file: File) =>
    new Promise<string>((resolve, reject) => {
      const reader = new FileReader();
      reader.readAsDataURL(file);
      reader.onload = () => {
        const base64 = (reader.result as string).split(",")[1]; // remove "data:.."
        resolve(base64);
      };
      reader.onerror = (error) => reject(error);
    });

  const submit = async () => {
    if (!title || !subject || !file) {
      alert("Title, subject and file are required.");
      return;
    }

    const finalCategory = category === "Other" ? customCategory.trim() : category;
    if (!finalCategory) {
      alert("Enter a category.");
      return;
    }

    setUploading(true);
    const fileBase64 = await convertFileToBase64(file);
    const res = await fetch("/api/resources/create", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({
        ownerId: viewer._id,
        title,
        description,
        category: finalCategory,
        subject,
        tags,
        isPublic,
        fileName: file.name,
        fileDataBase64: fileBase64,
      }),
    });


    const data = await res.json();
    // console.log("data", data)
    setUploading(false);

    if (data.success) {
      router.push(`/resources/${data.resource._id}`);
    } else {
      alert(data.error || "Error");
    }
  };

  if (loadingViewer) return <div className="text-white p-10">Loading...</div>;

  return (
    <div className="relative min-h-screen text-gray-200">
      
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

      <Navbar user={viewer} />

      <div className="pt-24 max-w-2xl mx-auto px-4">

        <h1 className="text-2xl font-bold mb-6">Upload Resource</h1>

        <div className="space-y-4 bg-white/5 p-6 rounded-xl border border-white/10">

          {/* Title */}
          <input
            value={title}
            onChange={(e) => setTitle(e.target.value)}
            placeholder="Title *"
            className="w-full p-2 rounded bg-white/10 outline-none"
          />

          {/* Description */}
          <textarea
            value={description}
            onChange={(e) => setDescription(e.target.value)}
            placeholder="Description (optional)"
            className="w-full p-2 rounded bg-white/10 outline-none"
          />

          {/* Category */}
          <div>
            <label className="text-sm">Category *</label>
            <select
              value={category}
              onChange={(e) => setCategory(e.target.value)}
              className="w-full mt-1 p-2 rounded bg-white/10"
            >
              <option>PYQ</option>
              <option>Assignments</option>
              <option>Lab</option>
              <option>Notes</option>
              <option>Other</option>
            </select>

            {category === "Other" && (
              <input
                value={customCategory}
                onChange={(e) => setCustomCategory(e.target.value)}
                placeholder="Enter custom category"
                className="w-full mt-2 p-2 rounded bg-white/10 outline-none"
              />
            )}
          </div>

          {/* Subject */}
          <input
            value={subject}
            onChange={(e) => setSubject(e.target.value)}
            placeholder="Subject *"
            className="w-full p-2 rounded bg-white/10 outline-none"
          />

          {/* Tags */}
          <div>
            <label className="text-sm">Tags (press Enter)</label>
            <input
              value={tagInput}
              onChange={(e) => setTagInput(e.target.value)}
              onKeyDown={handleTagAdd}
              placeholder="Add tag"
              className="w-full mt-1 p-2 rounded bg-white/10 outline-none"
            />
            <div className="flex flex-wrap gap-2 mt-2">
              {tags.map((t) => (
                <div key={t} className="px-2 py-1 bg-white/10 rounded text-sm flex items-center gap-2">
                  #{t}
                  <button onClick={() => removeTag(t)} className="text-red-300">x</button>
                </div>
              ))}
            </div>
          </div>

          {/* Visibility */}
          <div className="flex items-center gap-4">
            <label className="flex items-center gap-2">
              <input
                type="radio"
                checked={isPublic}
                onChange={() => setIsPublic(true)}
              />
              Public
            </label>
            <label className="flex items-center gap-2">
              <input
                type="radio"
                checked={!isPublic}
                onChange={() => setIsPublic(false)}
              />
              Private
            </label>
          </div>

          {/* File Upload */}
          <input
            type="file"
            onChange={(e) => setFile(e.target.files?.[0] || null)}
            className="w-full p-2 rounded bg-white/10"
            accept=".pdf,.doc,.docx,.png,.jpg,.jpeg"
          />

          {/* Submit */}
          <button
            onClick={submit}
            disabled={uploading}
            className="w-full py-2 bg-indigo-600 rounded"
          >
            {uploading ? "Uploading..." : "Upload"}
          </button>
        </div>
      </div>
    </div>
  );
}
