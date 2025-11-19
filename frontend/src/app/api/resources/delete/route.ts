// src/app/api/resources/delete/route.ts
import { NextResponse } from "next/server";
import { connectDB } from "@/lib/mongodb";
import Resource from "@/models/Resource";
import fs from "fs";
import path from "path";

export async function POST(req: Request) {
  try {
    await connectDB();
    const body = await req.json();
    const { resourceId, ownerId } = body;
    if (!resourceId || !ownerId) return NextResponse.json({ error: "Missing fields" }, { status: 400 });

    const r = await Resource.findById(resourceId);
    if (!r) return NextResponse.json({ error: "Not found" }, { status: 404 });
    if (r.owner.toString() !== ownerId) return NextResponse.json({ error: "Unauthorized" }, { status: 403 });

    if (r.fileUrl && r.fileUrl.startsWith("/uploads/")) {
      try {
        const filePath = path.join(process.cwd(), "public", r.fileUrl.replace("/uploads/", "uploads/"));
        if (fs.existsSync(filePath)) fs.unlinkSync(filePath);
      } catch (e) {
      }
    }

    await Resource.findByIdAndDelete(resourceId);

    return NextResponse.json({ success: true });
  } catch (err: any) {
    return NextResponse.json({ error: err.message || "Server error" }, { status: 500 });
  }
}
