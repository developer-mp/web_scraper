export function downloadFile(filename, text) {
  const file = `${filename}.txt`;
  const textToSave = text;
  const blob = new Blob([textToSave], { type: "text/plain;charset=utf-8" });
  const link = document.createElement("a");
  link.download = file;
  link.href = window.URL.createObjectURL(blob);
  link.click();
}
