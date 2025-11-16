import { NextResponse } from 'next/server';
import { connectDB } from '@/lib/mongodb';
import Conversation from '@/models/Conversation';
import Message from '@/models/Message';

export async function GET(
  req: Request,
  context: { params: { id: string } | Promise<{ id: string }> }
) {
  try {
    await connectDB();

    const { id: convoId } = await context.params;

    const userId = req.headers.get('x-user-id');

    if (!convoId || !userId) {
      return NextResponse.json(
        { error: "Missing conversationId or userId" },
        { status: 400 }
      );
    }

    const convo = await Conversation.findById(convoId).populate('participants');

    if (!convo) {
      return NextResponse.json(
        { error: "Conversation not found", other: null, messages: [] },
        { status: 404 }
      );
    }

    const other = convo.participants.find(
      (p: any) => p._id.toString() !== userId
    );

    const messages = await Message.find({ conversationId: convoId })
      .sort({ createdAt: 1 });

    return NextResponse.json({ other, messages });

  } catch (err) {
    return NextResponse.json(
      { error: "Server error", details: String(err) },
      { status: 500 }
    );
  }
}
