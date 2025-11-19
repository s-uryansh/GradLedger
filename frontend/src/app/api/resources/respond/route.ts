import { NextResponse } from "next/server";
import { connectDB } from "@/lib/mongodb";
import Resource from "@/models/Resource";

export async function POST(req: Request) {
  await connectDB();
  const { resourceId, requestId, approve } = await req.json();
  if (!resourceId || !requestId) return NextResponse.json({ error: "Missing" }, { status: 400 });

  const r = await Resource.findById(resourceId);
  if (!r) return NextResponse.json({ error: "Not found" }, { status: 404 });

  const reqDoc = r.requests.id(requestId);
  if (!reqDoc) return NextResponse.json({ error: "Request not found" }, { status: 404 });

  if (reqDoc.status !== "pending") return NextResponse.json({ error: "Already responded" }, { status: 400 });

  reqDoc.status = approve ? "approved" : "rejected";
  reqDoc.respondedAt = new Date();

  if (approve) {
    const uid = reqDoc.user.toString();
    if (!r.allowedUsers.map(String).includes(uid)) r.allowedUsers.push(reqDoc.user);
  }

  await r.save();

  return NextResponse.json({ success: true });
}
