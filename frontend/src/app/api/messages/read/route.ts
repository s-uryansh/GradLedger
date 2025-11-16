import Message from "@/models/Message";
import User from "@/models/User";
import { connectDB } from "@/lib/mongodb";
import { NextResponse } from "next/server";

export async function POST(req: Request) {
  await connectDB();
  const { conversationId, userId } = await req.json();

  await Message.updateMany(
    { conversationId, receiver: userId, isRead: false },
    { $set: { isRead: true } }
  );

  await User.findByIdAndUpdate(userId, { unreadMessages: 0 });

  return NextResponse.json({ ok: true });
}
