import { NextResponse } from "next/server";
import { connectDB } from "@/lib/mongodb";
import Resource from "@/models/Resource";

interface Params {
  id: string;
}

export async function GET(
  req: Request,
  { params }: { params: Params | Promise<Params> }
) {
  await connectDB();

  const { id } = (await params) as Params;
  const url = new URL(req.url);
  const viewerId = url.searchParams.get("viewerId");

  const r = await Resource.findById(id)
    .populate("owner", "fullName profileImage")
    .populate("approvedUsers", "_id");

  if (!r) {
    return NextResponse.json({ error: "Not found" }, { status: 404 });
  }

  const isOwner =
    viewerId !== null && r.owner._id.toString() === viewerId;

  const isApproved =
    viewerId !== null &&
    r.approvedUsers.some((u: { _id: any }) => u._id.toString() === viewerId);

  const allowed = r.isPublic || isOwner || isApproved;

  const viewerRequest =
    viewerId !== null
      ? r.requests.find((req: { user: any }) => req.user.toString() === viewerId)
      : null;

  return NextResponse.json({
    resource: {
      _id: r._id,
      title: r.title,
      description: r.description,
      category: r.category,
      subject: r.subject,
      tags: r.tags,
      owner: r.owner,
      isPublic: r.isPublic,
      createdAt: r.createdAt,
      allowed,
      requestStatus: viewerRequest?.status || "none",
      fileUrl: allowed ? r.fileUrl : null,
    },
  });
}
