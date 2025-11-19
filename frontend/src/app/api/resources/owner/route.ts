// src/app/api/resources/owner/route.ts
import { NextResponse } from "next/server";
import { connectDB } from "@/lib/mongodb";
import Resource from "@/models/Resource";
import User from "@/models/User";

export async function GET(req: Request) {
  await connectDB();
  const url = new URL(req.url);
  const ownerId = url.searchParams.get("ownerId");

  if (!ownerId) {
    return NextResponse.json({ error: "Missing ownerId" }, { status: 400 });
  }

  const resources = await Resource.find({ owner: ownerId })
    .sort({ createdAt: -1 })
    .populate("requests.user", "fullName profileImage")
    .populate("owner", "fullName profileImage");

  const formatted = resources.map((r: any) => ({
    _id: r._id,
    title: r.title,
    description: r.description,
    category: r.category,
    subject: r.subject,
    tags: r.tags,
    fileUrl: r.fileUrl,
    isPublic: r.isPublic,
    createdAt: r.createdAt,
    pendingRequests: (r.requests || []).filter((req: any) => req.status === "pending").map((req: any) => ({
      _id: req._id,
      user: req.user,
      message: req.message,
      requestedAt: req.requestedAt
    }))
  }));

  return NextResponse.json({ success: true, resources: formatted });
}
