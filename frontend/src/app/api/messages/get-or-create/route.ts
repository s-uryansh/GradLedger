import { NextResponse } from "next/server";
import { connectDB } from "@/lib/mongodb";
import Conversation from "@/models/Conversation";
import User from "@/models/User";

export async function POST(req: Request) {
  try {
    await connectDB();
    const { userA, userB } = await req.json();

    if (!userA || !userB) {
      return NextResponse.json({ error: "Missing users" }, { status: 400 });
    }

    const existsA = await User.findById(userA);
    const existsB = await User.findById(userB);
    if (!existsA || !existsB) {
      return NextResponse.json({ error: "Invalid user(s)" }, { status: 400 });
    }

    let convo = await Conversation.findOne({
      participants: { $all: [userA, userB] }
    });

    if (!convo) {
      convo = await Conversation.create({
        participants: [userA, userB]
      });
    }

    return NextResponse.json({
      _id: convo._id.toString()
    });

  } catch (err: any) {
    console.error("get-or-create error:", err);
    return NextResponse.json({ error: err.message }, { status: 500 });
  }
}
