import Message from "@/models/Message";
import { connectDB } from "@/lib/mongodb";
import { NextResponse } from "next/server";

export async function POST(req: Request) {
  await connectDB();
  const { conversationId, userId } = await req.json();

  await Message.updateMany(
    { conversationId, receiver: userId, seen: false },
    { $set: { seen: true } }
  );

  return NextResponse.json({ ok: true });
}
