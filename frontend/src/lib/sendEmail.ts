import nodemailer from 'nodemailer';

export async function sendOtpEmail(to: string, otp: string) {
  try {
    const transporter = nodemailer.createTransport({
      service: 'gmail',
      auth: {
        user: process.env.EMAIL_USER,
        pass: process.env.EMAIL_PASS
      }
    });
console.log('EMAIL_USER:', process.env.EMAIL_USER);
console.log('EMAIL_PASS:', process.env.EMAIL_PASS);

    await transporter.sendMail({
      from: process.env.EMAIL_USER,
      to,
      subject: 'Verify your OTP',
      text: `Your verification code is ${otp}. It expires in 5 minutes.`
    });

    return { success: true };
  } catch (err) {
    console.error('Email sending failed:', err);
    return { success: false, error: 'Failed to send OTP' };
  }
}
