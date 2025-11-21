import { NextResponse } from "next/server";
import { connectDB } from "@/lib/mongodb";
import Resource from "@/models/Resource";

export async function POST(req: Request) {
  await connectDB();

  const body = await req.json();
  const resourceId = body.resourceId;
  const requestId = body.requestId;
  const action = body.action; 
  const ownerId = body.ownerId;
  const approveFlag = typeof body.approve === "boolean" ? body.approve : undefined;

  if (!resourceId || !requestId || !ownerId) {
    return NextResponse.json({ error: "Missing resourceId / requestId / ownerId" }, { status: 400 });
  }

  const r = await Resource.findById(resourceId);
  if (!r) return NextResponse.json({ error: "Not found" }, { status: 404 });

  if (r.owner.toString() !== ownerId) {
    return NextResponse.json({ error: "Unauthorized" }, { status: 403 });
  }

  const reqDoc: any = r.requests.id(requestId);
  if (!reqDoc) return NextResponse.json({ error: "Request not found" }, { status: 404 });
  if (reqDoc.status !== "pending") return NextResponse.json({ error: "Already responded" }, { status: 400 });

  const approve = approveFlag !== undefined ? approveFlag : (action === "approve");

  reqDoc.status = approve ? "approved" : "rejected";
  reqDoc.respondedAt = new Date();

  if (approve) {
    const uid = reqDoc.user.toString();

    if (!Array.isArray(r.approvedUsers)) r.approvedUsers = [];
    if (!r.approvedUsers.map(String).includes(uid)) r.approvedUsers.push(reqDoc.user);

    if (!Array.isArray((r as any).allowedUsers)) (r as any).allowedUsers = [];
    if (!(r as any).allowedUsers.map(String).includes(uid)) (r as any).allowedUsers.push(reqDoc.user);

    if (Array.isArray((r as any).pendingRequests)) {
      (r as any).pendingRequests = (r as any).pendingRequests.filter((x: any) => x.toString() !== uid);
    }
  } else {
    const uid = reqDoc.user.toString();
    if (Array.isArray((r as any).pendingRequests)) {
      (r as any).pendingRequests = (r as any).pendingRequests.filter((x: any) => x.toString() !== uid);
    }
  }

  await r.save();

  return NextResponse.json({ success: true, request: reqDoc.toObject ? reqDoc.toObject() : reqDoc });
}
