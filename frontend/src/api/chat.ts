import type {Chat, Message} from "../types/chat";

const apiBase = import.meta.env.VITE_API_BASE;


// getChats(token: string, limitMessages?: number)
export async function getChats(
    token: string,
    limitMessages: number = 5
  ): Promise<{ chats: Chat[] }> {
    const params = new URLSearchParams();
    if (limitMessages !== 5) {
      params.set("limit_messages", String(limitMessages));
    }
    const query = params.toString();
    const url = `${apiBase}/v1/chats${query ? `?${query}` : ""}`;
  
    const res = await fetch(url, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        "Authorization": `Bearer ${token}`,
      },
    });
  
    if (!res.ok) {
      const err = await res.json().catch(() => ({}));
      throw new Error(err?.message ?? "Ошибка загрузки чатов");
    }
  
    return res.json();
}

export async function getChatMessages(token: string, chatId: number, limit: number, beforeMessageId?: number | null): Promise<{messages: Message[]}> {
  const params = new URLSearchParams({limit: String(limit) });
  if (beforeMessageId != null && beforeMessageId > 0){
    params.set("before_message_id", String(beforeMessageId));
  }
  const url = `${apiBase}/v1/chats/${chatId}/messages${params.toString() ? `?${params.toString()}` : ""}`;
  const res = await fetch(url, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
      "Authorization": `Bearer ${token}`,
    },
  });
  if (!res.ok) {
    const err = await res.json().catch(() => ({}));
    throw new Error(err?.message ?? "Ошибка загрузки сообщений");
  }
  return res.json();
}
