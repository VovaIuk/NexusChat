/**
 * Форматирует строку времени в вид «часы:минуты» (HH:MM), 24-часовой формат.
 */
export function formatMessageTime(time: string): string {
  const date = new Date(time);
  if (Number.isNaN(date.getTime())) {
    return time;
  }
  const hours = String(date.getHours()).padStart(2, "0");
  const minutes = String(date.getMinutes()).padStart(2, "0");
  return `${hours}:${minutes}`;
}
