import { NextResponse } from "next/server";
import { connectDB } from "@/lib/mongodb";
import Message from "@/models/Message";

export async function GET(req: Request) {
  await connectDB();
  const url = new URL(req.url);
  const userId = url.searchParams.get("userId");

  if (!userId) return NextResponse.json({ count: 0 });

  const count = await Message.countDocuments({
    receiver: userId,
    seen: false,
  });

  return NextResponse.json({ count });
}
