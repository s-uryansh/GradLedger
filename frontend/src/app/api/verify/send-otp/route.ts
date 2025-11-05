import { NextResponse } from 'next/server';
import nodemailer from 'nodemailer';
import { connectDB } from '@/lib/mongodb';
import User from '@/models/User';

export async function POST(req: Request) {
  try {
    const { email } = await req.json();
    if (!email)
      return NextResponse.json({ error: 'Email required' }, { status: 400 });

    await connectDB();
    const user = await User.findOne({ email });
    if (!user)
      return NextResponse.json({ error: 'User not found' }, { status: 404 });

    const otp = Math.floor(100000 + Math.random() * 900000).toString();
    user.otp = otp;
    user.otpExpiry = new Date(Date.now() + 10 * 60 * 1000); // 10 minutes
    await user.save();

    const GMAIL_USER = process.env.GMAIL_USER;
    const GMAIL_PASS = process.env.GMAIL_PASS;
    if (!GMAIL_USER || !GMAIL_PASS) {
      console.error('Missing GMAIL_USER or GMAIL_PASS');
      return NextResponse.json(
        { error: 'Email credentials missing on server' },
        { status: 500 }
      );
    }

    const transporter = nodemailer.createTransport({
      host: 'smtp.gmail.com',
      port: 587,
      secure: false,
      requireTLS: true,
      auth: { user: GMAIL_USER, pass: GMAIL_PASS },
      tls: { ciphers: 'TLSv1.2', rejectUnauthorized: true },
      connectionTimeout: 10000,
      greetingTimeout: 5000,
      socketTimeout: 10000,
    });

    const mailOptions = {
      from: `"GradLedger Verification" <${GMAIL_USER}>`,
      to: email,
      subject: 'Verify your email address - GradLedger',
      html: `
        <div style="font-family:Arial,sans-serif;max-width:600px;margin:auto;">
          <h2 style="color:#4F46E5;text-align:center;">GradLedger Email Verification</h2>
          <p>Dear ${user.fullName.split(' ')[0]},</p>
          <p>We received a request to verify your email address associated with GradLedger.</p>
          <p style="font-size:18px;font-weight:bold;text-align:center;">Your verification code is:</p>
          <div style="font-size:26px;font-weight:bold;color:#4F46E5;text-align:center;margin:12px 0;">${otp}</div>
          <p>This code is valid for <strong>10 minutes</strong>. Please do not share it with anyone.</p>
          <p>Thank you,<br/>The GradLedger Team</p>
          <hr/>
          <p style="font-size:12px;color:#777;">If you did not request this verification, please ignore this email.</p>
        </div>
      `,
    };

    await transporter.sendMail(mailOptions);
    console.log(`OTP sent to ${email}`);

    return NextResponse.json({ message: 'OTP sent successfully' });
  } catch (error: any) {
    console.error('Email sending error:', error);
    return NextResponse.json(
      { error: error.message || 'Failed to send OTP' },
      { status: 500 }
    );
  }
}
