"use client";

import React, { useEffect, useState } from "react";
import Navbar from "@/components/UI/Navbar";
import ColorBends from "@/components/BackgroundAnimations/ColorBends";
import ResourceCard from "@/components/Resources/ResouceCard";
import ManageRequestsModal from "@/components/Resources/ManageRequestsModal";
import toast from "react-hot-toast";
import { useRouter } from "next/navigation";

export default function ResourcesManagePage() {
  const [viewer, setViewer] = useState<any>(null);
  const [resources, setResources] = useState<any[]>([]);
  const [loading, setLoading] = useState(true);
  const [modalOpen, setModalOpen] = useState(false);
  const [activeResource, setActiveResource] = useState<any | null>(null);
  const router = useRouter();

  useEffect(() => {
    (async () => {
      const me = await fetch("/api/auth/me", { credentials: "include" }).then(r => r.json()).catch(()=>({}));
      if (!me.user) return router.push("/");
      setViewer(me.user);
    })();
  }, [router]);

  useEffect(() => {
    if (!viewer) return;
    loadResources();
  }, [viewer]);

  const loadResources = async () => {
    setLoading(true);
    try {
      const res = await fetch(`/api/resources/owner?ownerId=${viewer._id}`);
      const data = await res.json();
      if (data.success) setResources(data.resources || []);
    } catch (e) {
      console.error(e);
    } finally {
      setLoading(false);
    }
  };

  const openResource = async (id: string) => {
    // reuse existing resource view endpoint
    const res = await fetch(`/api/resources/${id}?viewerId=${viewer._id}`);
    const data = await res.json();
    if (data.resource?.allowed) {
      window.open(data.resource.fileUrl, "_blank");
      return;
    }
    toast("Resource is private. Use Manage Requests to approve.");
  };

  const openManageRequests = (r: any) => {
    setActiveResource(r);
    setModalOpen(true);
  };

  const toggleVisibility = async (resourceId: string, makePublic: boolean) => {
    try {
      const res = await fetch("/api/resources/toggle", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ resourceId, ownerId: viewer._id, makePublic }),
      });
      const data = await res.json();
      if (!data?.success) throw new Error(data?.error || "Failed");
      toast.success("Updated");
      setResources((prev) => prev.map(p => p._id === resourceId ? { ...p, isPublic: data.isPublic } : p));
    } catch (err: any) {
      toast.error(err.message || "Error");
    }
  };

  const deleteResource = async (resourceId: string) => {
    if (!confirm("Delete this resource? This cannot be undone.")) return;
    try {
      const res = await fetch("/api/resources/delete", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ resourceId, ownerId: viewer._id }),
      });
      const data = await res.json();
      if (!data?.success) throw new Error(data?.error || "Failed");
      toast.success("Deleted");
      setResources((prev) => prev.filter(r => r._id !== resourceId));
    } catch (err: any) {
      toast.error(err.message || "Error");
    }
  };

  const onModalUpdate = async () => {
    setModalOpen(false);
    setActiveResource(null);
    await loadResources();
  };

  if (!viewer) return <div className="text-white p-8">Loading...</div>;

  return (
    <div className="relative min-h-screen w-full overflow-x-hidden text-white">
      <div className="fixed inset-0 -z-30">
        <ColorBends colors={['#3e47f4','#06b31a','#b46d04']} rotation={0} speed={0.3} scale={1} frequency={1} warpStrength={1} mouseInfluence={1} parallax={0.5} noise={0.1}/>
        <div className="absolute inset-0 pointer-events-none" style={{background: 'linear-gradient(to bottom, rgba(0,0,0,0.25) 0%, rgba(0,0,0,0.05) 25%, rgba(0,0,0,0.10) 70%, rgba(0,0,0,0.15) 100%)'}}/>
      </div>

      <Navbar user={viewer} onLoginClick={() => router.push("/")} />

      <main className="pt-24 max-w-4xl mx-auto p-4">
        <div className="mb-6 flex justify-between items-center">
          <div>
            <h1 className="text-2xl font-bold">My Resources</h1>
            <div className="text-sm text-gray-300">{resources.length} resources • {resources.reduce((acc,r)=>acc + (r.pendingRequests?.length||0),0)} pending requests</div>
          </div>
        </div>

        {loading ? (
          <div className="text-gray-300">Loading resources…</div>
        ) : resources.length === 0 ? (
          <div className="text-gray-300">You have no resources yet.</div>
        ) : (
          <div className="space-y-4">
            {resources.map(r => (
              <ResourceCard
                key={r._id}
                resource={r}
                onOpen={openResource}
                onManageRequests={() => openManageRequests(r)}
                onToggle={toggleVisibility}
                onDelete={deleteResource}
              />
            ))}
          </div>
        )}
      </main>

      {activeResource && (
        <ManageRequestsModal
          open={modalOpen}
          onClose={() => { setModalOpen(false); setActiveResource(null); }}
          resourceId={activeResource._id}
          ownerId={viewer._id}
          requests={activeResource.pendingRequests || []}
          onUpdate={onModalUpdate}
        />
      )}
    </div>
  );
}
