export const runtime = 'nodejs';

import { NextResponse } from 'next/server';
import fs from 'fs/promises';
import path from 'path';
import sharp from 'sharp';
import tesseract from 'node-tesseract-ocr';
import { connectDB } from '@/lib/mongodb';
import User from '@/models/User';

export async function POST(req: Request) {
  try {
    const formData = await req.formData();
    const file = formData.get('file') as File;
    const email = formData.get('email') as string;
    if (!file || !email)
      return NextResponse.json({ error: 'Missing file or email' }, { status: 400 });

    const bytes = Buffer.from(await file.arrayBuffer());
    const uploadDir = path.join(process.cwd(), 'tmp');
    await fs.mkdir(uploadDir, { recursive: true });
    const filePath = path.join(uploadDir, file.name);
    await fs.writeFile(filePath, bytes);

    const meta = await sharp(filePath).metadata();
    const width = meta.width || 800;
    const height = meta.height || 1000;

    const faceRegion = {
      left: Math.round(width * 0.25),
      top: Math.round(height * 0.18),
      width: Math.round(width * 0.5),
      height: Math.round(height * 0.28),
    };

    const croppedPath = path.join(uploadDir, `cropped_${file.name}`);
    await sharp(filePath)
      .extract(faceRegion)
      .resize(256, 256)
      .toFile(croppedPath);

    const text = await tesseract.recognize(filePath, { lang: 'eng' });
    console.log('OCR output:', text);

    const rollMatch = text.match(/(\d{10})/);
    const programMatch = text.match(/Program\s*[:\-]?\s*([A-Z]+)/i);
    const majorMatch = text.match(/Major\s*[:\-]?\s*([A-Za-z\s]+)/i);

    const rollNumber = rollMatch ? rollMatch[1] : '';
    const program = programMatch ? programMatch[1].trim() : '';
    const major = majorMatch ? majorMatch[1].trim() : '';

    const imageBase64 = (await fs.readFile(croppedPath)).toString('base64');
    const profileImage = `data:image/jpeg;base64,${imageBase64}`;

    await connectDB();
    const user = await User.findOne({ email });
    if (user) {
      user.rollNumber = rollNumber;
      user.program = program;
      user.major = major;
      user.profileImage = profileImage;
      await user.save();
    }

    return NextResponse.json({
      success: true,
      rollNumber,
      program,
      major,
      url: profileImage,
    });
  } catch (err: any) {
    console.error('Upload error:', err);
    return NextResponse.json({ error: err.message }, { status: 500 });
  }
}
