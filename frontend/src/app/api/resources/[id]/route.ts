import { NextResponse } from "next/server";
import { connectDB } from "@/lib/mongodb";
import Resource from "@/models/Resource";
import { getReputation } from "@/lib/go";

interface Params {
  id: string;
}

export async function GET(
  req: Request,
  { params }: { params: Params | Promise<Params> }
) {
  await connectDB();

  const { id } = await params;
  const url = new URL(req.url);
  const viewerId = url.searchParams.get("viewerId");

  const r = await Resource.findById(id)
    .populate("owner", "_id fullName profileImage selfieImage walletAddress")
    .populate("approvedUsers", "_id")
    .populate("requests.user", "_id fullName");

  if (!r) {
    return NextResponse.json({ error: "Not found" }, { status: 404 });
  }

  const owner = r.owner as any;

  let ownerReputation = 0;
  try {
    if (owner?.walletAddress) {
      const repData = await getReputation(owner.walletAddress);
      ownerReputation = Number(repData?.score ?? 0);
    }
  } catch (err) {
    console.warn("Reputation fetch failed:", err);
  }

  const isOwner = viewerId && owner?._id?.toString() === viewerId;
  const isApproved =
    viewerId &&
    Array.isArray(r.approvedUsers) &&
    r.approvedUsers.some((u: any) => u._id.toString() === viewerId);

  const allowed = r.isPublic || isOwner || isApproved;

  const viewerRequest =
    viewerId
      ? (r.requests || []).find((req: any) => {
          const uid = req.user?._id?.toString() ?? req.user?.toString();
          return uid === viewerId;
        })
      : null;

  return NextResponse.json({
    resource: {
      _id: r._id,
      title: r.title,
      description: r.description,
      category: r.category,
      subject: r.subject,
      tags: r.tags,
      owner: {
        _id: owner._id,
        fullName: owner.fullName,
        profileImage: owner.profileImage,
        selfieImage: owner.selfieImage,
        walletAddress: owner.walletAddress,
        reputation: ownerReputation,
      },
      isPublic: r.isPublic,
      createdAt: r.createdAt,
      allowed,
      requestStatus: viewerRequest?.status || "none",
      fileUrl: allowed ? r.fileUrl : null,
      requests: isOwner ? r.requests : [],
    },
  });
}
