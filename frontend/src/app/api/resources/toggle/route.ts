import { NextResponse } from "next/server";
import { connectDB } from "@/lib/mongodb";
import Resource from "@/models/Resource";

export async function POST(req: Request) {
  try {
    await connectDB();
    const body = await req.json();
    const { resourceId, ownerId, makePublic } = body;
    if (!resourceId || !ownerId) return NextResponse.json({ error: "Missing fields" }, { status: 400 });

    const r = await Resource.findById(resourceId);
    if (!r) return NextResponse.json({ error: "Not found" }, { status: 404 });
    if (r.owner.toString() !== ownerId) return NextResponse.json({ error: "Unauthorized" }, { status: 403 });

    r.isPublic = Boolean(makePublic);
    await r.save();

    return NextResponse.json({ success: true, isPublic: r.isPublic });
  } catch (err: any) {
    return NextResponse.json({ error: err.message || "Server error" }, { status: 500 });
  }
}
