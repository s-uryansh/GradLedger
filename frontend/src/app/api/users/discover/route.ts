import { NextResponse } from "next/server";
import User from "@/models/User";
import { connectDB } from "@/lib/mongodb";

export async function GET(req: Request) {
  await connectDB();
  const url = new URL(req.url);

  const viewerId = url.searchParams.get("viewerId");
  const page = Number(url.searchParams.get("page") || 1);
  const limit = 5;
  const skip = (page - 1) * limit;

  const query = url.searchParams.get("q")?.trim()?.toLowerCase() || "";

  let filter: any = {
    mailVerified: true,
    faceVerified: true,
    _id: { $ne: viewerId }    
  };

  if (query) {
    filter.$or = [
      { fullName: { $regex: query, $options: "i" } },
      { email: { $regex: query, $options: "i" } }
    ];
  }

  let users;
  let total;

  if (!query) {
    total = await User.countDocuments(filter);

    users = await User.find(filter)
      .select("fullName profileImage selfieImage program major email _id")
      .skip(skip)
      .limit(limit)
      .lean();

    users.sort(() => 0.5 - Math.random());
  } else {
    total = await User.countDocuments(filter);

    users = await User.find(filter)
      .select("fullName profileImage selfieImage program major email _id")
      .skip(skip)
      .limit(limit)
      .lean();
  }

  const totalPages = Math.ceil(total / limit);

  return NextResponse.json({ users, totalPages });
}
