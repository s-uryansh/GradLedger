import { NextResponse } from "next/server";
import { connectDB } from "@/lib/mongodb";
import Resource from "@/models/Resource";
import fs from "fs";
import path from "path";
import crypto from "crypto";

export async function POST(req: Request) {
  try {
    await connectDB();
    const body = await req.json();

    const {
      ownerId,
      title,
      description = "",
      category = "Other",
      subject = "",
      tags = [],
      isPublic = true,
      fileName,
      fileDataBase64 
    } = body;

    if (!ownerId || !title || !fileName || !fileDataBase64)
      return NextResponse.json({ error: "Missing fields" }, { status: 400 });

    const uploadsDir = path.join(process.cwd(), "public", "uploads");
    if (!fs.existsSync(uploadsDir)) fs.mkdirSync(uploadsDir, { recursive: true });

    const ext = path.extname(fileName) || "";
    const unique = crypto.randomBytes(10).toString("hex");
    const saveName = `${Date.now()}-${unique}${ext}`;
    const filePath = path.join(uploadsDir, saveName);
    const fileBuffer = Buffer.from(fileDataBase64, "base64");

    fs.writeFileSync(filePath, fileBuffer);

    const fileUrl = `/uploads/${saveName}`;

    const resource = await Resource.create({
      owner: ownerId,
      title,
      description,
      category,
      subject,
      tags,
      fileName,
      fileUrl,
      isPublic,
      allowedUsers: [],
    });

    return NextResponse.json({ success: true, resource });
  } catch (err: any) {
    return NextResponse.json({ error: err.message || "Server error" }, { status: 500 });
  }
}
