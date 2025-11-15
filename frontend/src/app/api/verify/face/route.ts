import { NextResponse } from 'next/server';
import { connectDB } from '@/lib/mongodb';
import User from '@/models/User';

export async function POST(req: Request) {
  try {
    const { email, result, profileImage, selfieImage, rollNumber, program, major } = await req.json();
    if (!email || !result?.match)
      return NextResponse.json({ success: false }, { status: 400 });

    await connectDB();
    const user = await User.findOne({ email });
    if (!user)
      return NextResponse.json({ success: false, error: 'User not found' }, { status: 404 });

    user.faceVerified = true;
    user.profileImage = profileImage || user.profileImage;
    user.selfieImage = selfieImage || user.selfieImage;
    if (rollNumber) user.rollNumber = rollNumber;
    if (program) user.program = program;
    if (major) user.major = major;
    await user.save();

    return NextResponse.json({ success: true });
  } catch (err) {
    console.error('verify/face error:', err);
    return NextResponse.json({ success: false, error: 'Server error' }, { status: 500 });
  }
}
