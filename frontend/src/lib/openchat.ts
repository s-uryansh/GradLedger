export async function openChat(myId: string, partnerId: string) {
  const res = await fetch("/api/messages/get-or-create", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ userA: myId, userB: partnerId }),
  });

  if (!res.ok) throw new Error("Failed to create or fetch conversation");

  return await res.json();
}
