import { NextResponse } from 'next/server';
import { connectDB } from '@/lib/mongodb';
import User from '@/models/User';

export async function POST(req: Request) {
  try {
    const { email, otp } = await req.json();
    if (!email || !otp)
      return NextResponse.json({ error: 'Missing fields' }, { status: 400 });

    await connectDB();
    const user = await User.findOne({ email });
    if (!user) return NextResponse.json({ error: 'User not found' }, { status: 404 });

    if (!user.otp || user.otp !== otp)
      return NextResponse.json({ error: 'Invalid OTP' }, { status: 401 });

    if (user.otpExpiry < new Date())
      return NextResponse.json({ error: 'OTP expired' }, { status: 410 });

    user.mailVerified = true;
    user.otp = null;
    user.otpExpiry = null;
    await user.save();

    return NextResponse.json({ message: 'Email verified successfully' });
  } catch (err) {
    console.error(err);
    return NextResponse.json({ error: 'Verification failed' }, { status: 500 });
  }
}
