import type { Message } from "../chat";

export const WS_MESSAGE_TYPE = {
    AUTH: "auth",
    CHAT_MESSAGE: "chat_message",
} as const;

export interface ClientWsEnvelope{
    type: typeof WS_MESSAGE_TYPE.AUTH | typeof WS_MESSAGE_TYPE.CHAT_MESSAGE;
    data: string; 
}

export interface AuthPayload{
    token: string;
}

export interface ChatMessagePayload {
    chat_id: number;
    text: string;
}

/** Сервер → клиент: broadcast (см. wsserver.BroadcastMessage) */
export interface ServerBroadcastMessage {
    message_id: number;
    user_id: number;
    username: string;
    usertag: string;
    chat_id: number;
    text: string;
    time: string;
}

export function broadcastToMessage(b: ServerBroadcastMessage): Message {
    return {
      user_author: {
        id: b.user_id,
        tag: b.usertag,
        name: b.username,
      },
      message: {
        id: b.message_id,
        text: b.text,
        time: b.time,
      },
    };
  }

  export function isServerBroadcast(raw: unknown): raw is ServerBroadcastMessage {
    if (raw == null || typeof raw !== "object") return false;
    const o = raw as Record<string, unknown>;
    return (
      typeof o.message_id === "number" &&
      typeof o.chat_id === "number" &&
      typeof o.text === "string"
    );
  }