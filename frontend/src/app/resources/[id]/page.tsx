import ResourceClient from "./ResourceClient";

export default async function Page({ params }: { params: { id: string } | Promise<{ id: string }> }) {
	const unwrapped = await params;
	return <ResourceClient id={unwrapped.id} />;
}
