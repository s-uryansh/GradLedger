"use client";

import { useState } from "react";
import { ChevronUpIcon, ChevronDownIcon } from "@heroicons/react/24/solid";
import { addReputation, subReputation } from "@/lib/go";

export default function VoteBox({
  walletAddress,
  initialScore,
  viewerId,
  profileId,
}: {
  walletAddress: string;
  initialScore: number;
  viewerId: string;
  profileId: string;
}) {
  const voteKey = `vote-${viewerId}-${profileId}`;
  const saved = typeof window !== "undefined" ? localStorage.getItem(voteKey) : null;

  const [score, setScore] = useState(initialScore);
  const [vote, setVote] = useState(saved || "none"); 
  const [loading, setLoading] = useState(false);

  const handleUpvote = async () => {
    if (vote === "up" || loading) return;

    const delta = vote === "down" ? 2 : 1;
    setScore(score + delta);
    setVote("up");
    localStorage.setItem(voteKey, "up");

    setLoading(true);
    try {
      await addReputation({ mentor: walletAddress, amount: 1 });
    } catch {
      setScore(initialScore);
      setVote(saved || "none");
    }
    setLoading(false);
  };

  const handleDownvote = async () => {
    if (vote === "down" || loading) return;

    const delta = vote === "up" ? -2 : -1;
    setScore(score + delta);
    setVote("down");
    localStorage.setItem(voteKey, "down");

    setLoading(true);
    try {
      await subReputation({ mentor: walletAddress, amount: 1 });
    } catch {
      // rollback
      setScore(initialScore);
      setVote(saved || "none");
    }
    setLoading(false);
  };

  return (
    <div className="flex flex-col items-center w-10 select-none">
      <button
        onClick={handleUpvote}
        className={`p-1 rounded transition ${
          vote === "up" ? "text-orange-400" : "text-gray-400 hover:text-gray-300"
        }`}
      >
        <ChevronUpIcon className="h-6 w-6" />
      </button>

      <div className="text-lg font-bold">{score}</div>

      <button
        onClick={handleDownvote}
        className={`p-1 rounded transition ${
          vote === "down" ? "text-blue-400" : "text-gray-400 hover:text-gray-300"
        }`}
      >
        <ChevronDownIcon className="h-6 w-6" />
      </button>
    </div>
  );
}
