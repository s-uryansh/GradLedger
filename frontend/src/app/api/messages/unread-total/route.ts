import { NextResponse } from 'next/server';
import Message from '@/models/Message';
import { connectDB } from '@/lib/mongodb';

export async function GET(req: Request) {
  await connectDB();

  const userId = req.headers.get('x-user-id');

  const count = await Message.countDocuments({
    receiver: userId,
    isRead: false
  });

  return NextResponse.json({ count });
}
