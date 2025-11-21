import { NextResponse } from "next/server";
import User from "@/models/User";
import { getReputation } from "@/lib/go";

export async function GET(_req: Request, context: { params: { id: string } | Promise<{ id: string }> }) {
  const { id } = await context.params;

  if (!id || id === "undefined") {
    return NextResponse.json({ error: "Invalid user id" }, { status: 400 });
  }

  let user;
  try {
    user = await User.findById(id).lean().exec();
  } catch {
    return NextResponse.json({ error: "Invalid ObjectId" }, { status: 400 });
  }

  if (!user) {
    return NextResponse.json({ error: "User not found" }, { status: 404 });
  }

  const walletAddress = (user as any)?.walletAddress;
  let rep = "0";
const res = await fetch("http://localhost:8080/reputation/0xef5b8c51ae60f3f7af6417e225817cb0363e22f6");
const data = await res.json();
console.log("REPUTATION RESPONSE:", data);

  if (walletAddress) {
    try {
      // const data = await getReputation(walletAddress);
      const data = await getReputation(walletAddress);
      rep = String(data.score ?? "0");
    } catch {}
  }
console.log(walletAddress, rep)
  return NextResponse.json({ user: { ...user, reputation: rep } });
}
