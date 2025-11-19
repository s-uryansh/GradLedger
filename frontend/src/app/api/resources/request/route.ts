import { NextResponse } from "next/server";
import { connectDB } from "@/lib/mongodb";
import Resource from "@/models/Resource";

export async function GET(req: Request) {
  await connectDB();
  const url = new URL(req.url);
  const ownerId = url.searchParams.get("ownerId");
  if (!ownerId) return NextResponse.json({ error: "Missing ownerId" }, { status: 400 });

  const resources = await Resource.find({ owner: ownerId }).select("title requests");
  const out = resources.map((r: any) => ({
    resourceId: r._id,
    title: r.title,
    requests: r.requests.filter((req: any) => req.status === "pending")
  }));
  return NextResponse.json({ success: true, items: out });
}
