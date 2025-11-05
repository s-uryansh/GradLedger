import { NextResponse } from 'next/server';
import jwt from 'jsonwebtoken';
import { connectDB } from '@/lib/mongodb';
import User from '@/models/User';

const JWT_SECRET = process.env.JWT_SECRET!;

export async function GET(req: Request) {
  try {
    const cookieHeader = req.headers.get('cookie');
    const token = cookieHeader?.split('auth_token=')[1]?.split(';')[0];
    if (!token) return NextResponse.json({ user: null });

    const decoded: any = jwt.verify(token, JWT_SECRET);
    await connectDB();
    const user = await User.findById(decoded.id).select('-password');

    if (!user) return NextResponse.json({ user: null });
    return NextResponse.json({ user });
  } catch (err) {
    return NextResponse.json({ user: null });
  }
}
