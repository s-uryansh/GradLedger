"use client";

interface Props {
  score: number;
}

export default function ReputationBadge({ score }: Props) {
  let label = "New";
  let color = "bg-gray-600";

  if (score >= 5)  { label = "Contributor"; color = "bg-blue-600"; }
  if (score >= 20) { label = "Trusted";     color = "bg-green-600"; }
  if (score >= 50) { label = "Senior";      color = "bg-orange-600"; }
  if (score >= 200){ label = "Legend";      color = "bg-red-700"; }

  return (
    <span className={`px-2 py-1 text-xs rounded-full text-white ${color}`}>
      {label} • {score}
    </span>
  );
}
