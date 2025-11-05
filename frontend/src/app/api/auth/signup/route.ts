import { NextResponse } from 'next/server';
import { connectDB } from '@/lib/mongodb';
import User from '@/models/User';

export async function POST(req: Request) {
  try {
    const { fullName, email, role, walletAddress } = await req.json();

    if (!fullName || !email || !walletAddress || !role) {
      return NextResponse.json({ error: 'Missing required fields' }, { status: 400 });
    }

    if (!email.endsWith('@snu.edu.in')) {
      return NextResponse.json({ error: 'Only snu.edu.in emails are allowed' }, { status: 400 });
    }

    await connectDB();

    const existing = await User.findOne({ email });
    if (existing) {
      return NextResponse.json({ error: 'User already exists' }, { status: 400 });
    }

    await User.create({
      fullName,
      email,
      walletAddress,
      faceVerified: false,
      mailVerified: false,
      role,
      reputation: 10,
    });

    return NextResponse.json({ message: 'Signup successful', user: { fullName, email, role, walletAddress } });
  } catch (error) {
    console.error('Signup Error:', error);
    return NextResponse.json({ error: 'Internal server error' }, { status: 500 });
  }
}
