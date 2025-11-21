import { NextResponse } from "next/server";
import { connectDB } from "@/lib/mongodb";
import Resource from "@/models/Resource";

export async function GET(req: Request) {
  await connectDB();
  const url = new URL(req.url);
  const ownerId = url.searchParams.get("ownerId");
  if (!ownerId) return NextResponse.json({ error: "Missing ownerId" }, { status: 400 });

  const resources = await Resource.find({ owner: ownerId }).select("title requests");
  const out = resources.map((r: any) => ({
    resourceId: r._id,
    title: r.title,
    requests: r.requests.filter((req: any) => req.status === "pending")
  }));
  return NextResponse.json({ success: true, items: out });
}

export async function POST(req: Request) {
  try {
    await connectDB();
    const { resourceId, userId, message = "" } = await req.json();

    if (!resourceId || !userId) {
      return NextResponse.json({ error: "Missing resourceId or userId" }, { status: 400 });
    }

    const r = await Resource.findById(resourceId);
    if (!r) return NextResponse.json({ error: "Not found" }, { status: 404 });

    // owner already has access
    if (r.owner.toString() === userId) {
      return NextResponse.json({ error: "Owner already has access" }, { status: 400 });
    }

    // already approved
    if (Array.isArray(r.approvedUsers) && r.approvedUsers.map(String).includes(userId)) {
      return NextResponse.json({ error: "Already approved" }, { status: 400 });
    }

    const existing = (r.requests || []).find((req: any) => req.user.toString() === userId);
    if (existing) {
      return NextResponse.json({ success: true, request: existing });
    }

    const newReq = {
      user: userId,
      message,
      status: "pending",
      requestedAt: new Date(),
    };

    r.requests.push(newReq);

    if (!Array.isArray((r as any).pendingRequests)) (r as any).pendingRequests = [];
    if (!(r as any).pendingRequests.map(String).includes(userId)) {
      (r as any).pendingRequests.push(userId);
    }

    await r.save();

    return NextResponse.json({ success: true, request: newReq });
  } catch (err: any) {
    return NextResponse.json({ error: err.message || "Server error" }, { status: 500 });
  }
}
