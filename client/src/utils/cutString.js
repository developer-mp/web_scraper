export function cutString(text) {
  if (text.length > 100) {
    return text.substring(0, 100) + "...";
  }
  return text;
}
