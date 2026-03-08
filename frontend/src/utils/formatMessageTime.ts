/**
 * Форматирует строку времени в вид "минуты:секунды" (MM:SS).
 */
export function formatMessageTime(time: string): string {
  const date = new Date(time);
  if (Number.isNaN(date.getTime())) {
    return time;
  }
  const minutes = String(date.getMinutes()).padStart(2, "0");
  const seconds = String(date.getSeconds()).padStart(2, "0");
  return `${minutes}:${seconds}`;
}
