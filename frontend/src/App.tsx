import { useState, useCallback } from "react";

export default function App() {
  const [url, setUrl] = useState<string | null>(null);
  const [error, setError] = useState<string | null>(null);

  const upload = useCallback(async (file: File) => {
    const form = new FormData();
    form.append("image", file);

    const res = await fetch("/upload", { method: "POST", body: form });
    if (!res.ok) {
      setError("upload failed");
      return;
    }

    const data = await res.json();
    setUrl(data.url);
    setError(null);
  }, []);

  // ctrl+v
  const onPaste = useCallback((e: React.ClipboardEvent) => {
    const file = e.clipboardData.files[0];
    if (file) upload(file);
  }, [upload]);

  // drag and drop
  const onDrop = useCallback((e: React.DragEvent) => {
    e.preventDefault();
    const file = e.dataTransfer.files[0];
    if (file) upload(file);
  }, [upload]);

  const onDragOver = (e: React.DragEvent) => e.preventDefault();

  // file input
  const onChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0];
    if (file) upload(file);
  };

  return (
    <div onPaste={onPaste} onDrop={onDrop} onDragOver={onDragOver}>
      <p>paste, drop or select an image</p>
      <input type="file" accept="image/*" onChange={onChange} />
      {url && <p>http://{url}</p>}
      {error && <p>{error}</p>}
    </div>
  );
}