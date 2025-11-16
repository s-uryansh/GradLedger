import { NextResponse } from "next/server";
import Message from "@/models/Message";
import { connectDB } from "@/lib/mongodb";

export async function GET(req: Request, { params }: { params: any }) {
  await connectDB();
  const { conversationId } = params;

  const messages = await Message.find({ conversationId })
    .sort({ createdAt: 1 });

  return NextResponse.json(messages);
}
