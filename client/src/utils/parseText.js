export function parseText(text) {
  return text[0]?.trim().replace(/^\[\s*"\n\t|\n\t"\s*\]$/g, "");
}
