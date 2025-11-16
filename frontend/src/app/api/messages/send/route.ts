import { NextResponse } from "next/server";
import { connectDB } from "@/lib/mongodb";
import Message from "@/models/Message";

export async function POST(req: Request) {
  try {
    await connectDB();
    const body = await req.json();

    const saved = await Message.create({
      conversationId: body.conversationId,
      sender: body.sender,
      receiver: body.receiver,
      text: body.text,
      seen: false,
    });

    return NextResponse.json({ success: true, message: saved });

  } catch (err: any) {
    return NextResponse.json({ error: err.message }, { status: 500 });
  }
}
