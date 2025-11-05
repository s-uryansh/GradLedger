export const PYTHON_API_URL = process.env.NEXT_PUBLIC_PYTHON_API_URL || 'http://127.0.0.1:8000';

export async function verifyFaces(image1: File, image2: File) {
  const formData = new FormData();
  formData.append('image1', image1);
  formData.append('image2', image2);

  const res = await fetch(`${PYTHON_API_URL}/verify`, {
    method: 'POST',
    body: formData,
  });

  if (!res.ok) {
    const err = await res.json();
    throw new Error(err.detail || 'Face verification failed');
  }

  return res.json();
}
