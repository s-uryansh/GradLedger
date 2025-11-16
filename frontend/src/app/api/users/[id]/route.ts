import { NextResponse } from "next/server";
import { connectDB } from "@/lib/mongodb";
import User from "@/models/User";

export async function GET(req: Request, { params }: any) {
  await connectDB();
  const { id } = params;

  const user = await User.findById(id)
    .select("fullName profileImage selfieImage tags program major rollNumber _id");

  if (!user) return NextResponse.json({ error: "Not found" }, { status: 404 });

  return NextResponse.json({ user });
}
