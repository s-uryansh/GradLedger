import { NextResponse } from 'next/server';
import { connectDB } from '@/lib/mongodb';
import User from '@/models/User';

export async function POST(req: Request) {
  try {
    const body = await req.json();
    const { email, result } = body || {};
    if (!email || !result) {
      return NextResponse.json({ error: 'Missing email or result' }, { status: 400 });
    }

    const match = Boolean(result.match);
    if(!match){
        return NextResponse.json({ success: false, error: 'Face mismatch' }, { status: 400 });
    }

    await connectDB();
    const user = await User.findOne({ email });
    if (!user) {
      console.error('verify/face: user not found for email', email);
      return NextResponse.json({ success: false, error: 'User not found' }, { status: 404 });
    }

    user.faceVerified = true;

    await user.save();
    console.log(`verify/face: faceVerified set true for ${email}`);

    return NextResponse.json({ success: true, email, faceVerified: user.faceVerified });
  } catch (err: any) {
    console.error('verify/face error:', err);
    return NextResponse.json({ success: false, error: err.message || 'Internal error' }, { status: 500 });
  }
}
