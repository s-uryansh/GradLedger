import { NextResponse } from "next/server";
import { connectDB } from "@/lib/mongodb";
import Conversation from "@/models/Conversation";
import Message from "@/models/Message";
import User from "@/models/User";

export async function GET(req: Request) {
  await connectDB();

  const userId = req.headers.get("x-user-id");
  const page = parseInt(req.headers.get("x-page") || "1", 10);
  const limit = 10;
  const skip = (page - 1) * limit;

  const convos = await Conversation.find({ participants: userId })
    .populate("participants")
    .sort({ updatedAt: -1 })
    .skip(skip)
    .limit(limit);

  const results = [];

  for (const convo of convos) {
    const last = await Message.findOne({ conversationId: convo._id })
      .sort({ createdAt: -1 });

    if (!last) continue; 

    results.push({
      _id: convo._id,
      participants: convo.participants,
      lastMessage: last,
    });
  }

  return NextResponse.json(results);
}
