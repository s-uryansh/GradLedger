import { NextResponse } from "next/server";
import { connectDB } from "@/lib/mongodb";
import Resource from "@/models/Resource";

export async function GET(req: Request) {
  await connectDB();

  const url = new URL(req.url);
  const page = Number(url.searchParams.get("page") || "1");
  const q = (url.searchParams.get("q") || "").trim();
  const viewerId = url.searchParams.get("viewerId") || null;

  const onlyPublic = url.searchParams.get("publicOnly") === "true";
  const limit = 5;
  const skip = (page - 1) * limit;

  const searchQuery: any = {};

  if (q) {
    const reg = new RegExp(q, "i");
    searchQuery.$or = [
      { title: reg },
      { subject: reg },
      { tags: reg },
      { description: reg },
    ];
  }

  if (!viewerId) {
    const filter: any = { isPublic: true };
    if (q) filter.$or = searchQuery.$or;

    const total = await Resource.countDocuments(filter);
    const items = await Resource.find(filter)
      .sort({ createdAt: -1 })
      .skip(skip)
      .limit(limit)
      .populate("owner", "fullName profileImage walletAddress")
      .select("_id title subject tags owner fileUrl isPublic createdAt")
      .lean();

    return NextResponse.json({ users: items, total, totalPages: Math.ceil(total / limit) });
  }

  if (onlyPublic) {
    const filter: any = { isPublic: true };
    if (q) filter.$or = searchQuery.$or;

    const total = await Resource.countDocuments(filter);
    const items = await Resource.find(filter)
      .sort({ createdAt: -1 })
      .skip(skip)
      .limit(limit)
      .populate("owner", "fullName profileImage walletAddress")
      .select("_id title subject tags owner fileUrl isPublic createdAt")
      .lean();

    return NextResponse.json({ users: items, total, totalPages: Math.ceil(total / limit) });
  }

  let baseFilter: any;
  if (q) {
    const reg = new RegExp(q, "i");
    baseFilter = {
      $and: [
        {
          $or: [
            { title: reg },
            { subject: reg },
            { tags: reg },
            { description: reg },
          ],
        },
      ],
    };
  } else {
    baseFilter = {};
  }

  const total = await Resource.countDocuments(baseFilter);
  const items = await Resource.find(baseFilter)
    .sort({ createdAt: -1 })
    .skip(skip)
    .limit(limit)
    .populate("owner", "fullName profileImage walletAddress")
    .select("_id title subject tags owner fileUrl isPublic createdAt approvedUsers")
    .lean();

  const mapped = items.map((it: any) => {
    const ownerId = it.owner?._id?.toString?.() ?? it.owner?.toString?.();
    const approvedArr = Array.isArray(it.approvedUsers) ? it.approvedUsers.map((x: any) => x?.toString?.() || x) : [];
    const allowed = it.isPublic || ownerId === viewerId || approvedArr.includes(viewerId);

    return {
      _id: it._id,
      title: it.title,
      subject: it.subject,
      tags: it.tags,
      owner: it.owner,
      fileUrl: allowed ? it.fileUrl : null, 
      isPublic: it.isPublic,
      isAllowedForViewer: allowed,
      createdAt: it.createdAt,
    };
  });

  return NextResponse.json({ users: mapped, total, totalPages: Math.ceil(total / limit) });
}
