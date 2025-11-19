"use client";

import React from "react";

export default function ResourceCard({
  resource,
  onOpen,
  onManageRequests,
  onToggle,
  onDelete,
}: {
  resource: any;
  onOpen: (id: string) => void;
  onManageRequests: (r: any) => void;
  onToggle: (id: string, makePublic: boolean) => void;
  onDelete: (id: string) => void;
}) {
  return (
    <div className="bg-white/5 p-4 rounded-lg flex flex-col gap-3">
      <div className="flex justify-between items-start">
        <div>
          <div className="font-semibold text-white">{resource.title}</div>
          <div className="text-xs text-gray-300">{resource.subject} • {resource.category}</div>
          <div className="text-xs text-gray-400 mt-1">{resource.tags?.map((t:string)=>`#${t}`).join(" ")}</div>
        </div>

        <div className="text-xs text-gray-300">{new Date(resource.createdAt).toLocaleDateString()}</div>
      </div>

      <div className="flex gap-2">
        <button onClick={() => onOpen(resource._id)} className="px-3 py-1 bg-indigo-600 rounded text-white">View / Open</button>
        <button onClick={() => onManageRequests(resource)} className="px-3 py-1 bg-white/8 rounded text-white">Manage Requests ({resource.pendingRequests?.length || 0})</button>
        <button onClick={() => onToggle(resource._id, !resource.isPublic)} className="px-3 py-1 bg-amber-600 rounded text-white">{resource.isPublic ? "Make Private" : "Make Public"}</button>
        <button onClick={() => onDelete(resource._id)} className="ml-auto px-3 py-1 bg-red-600 rounded text-white">Delete</button>
      </div>
    </div>
  );
}
