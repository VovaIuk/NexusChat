import { useState } from "react";
import { ChatListItem } from "./ChatListItem.tsx";
import type { Chat } from "../../../types/chat.ts";
import "./chatList.css";

const MOCK_CHATS: Chat[] = [
  {
    id: 1,
    name: "Test2",
    users: [
      { id: 0, tag: "me", name: "Вы" },
      { id: 1, tag: "test2", name: "Test2" },
    ],
    messages: [
      {
        user_author: { id: 0, tag: "me", name: "Вы" },
        message: {
          id: 1,
          text: "https://github.com/NoroSaro...",
          time: "21:25",
        },
      },
    ],
  },
  {
    id: 2,
    name: "Alexandr Shevelev",
    users: [
      { id: 0, tag: "me", name: "Вы" },
      { id: 2, tag: "alex", name: "Alexandr Shevelev" },
    ],
    messages: [
      {
        user_author: { id: 2, tag: "alex", name: "Alexandr Shevelev" },
        message: {
          id: 1,
          text: "Можешь сам любые подобрать))",
          time: "15:49",
        },
      },
    ],
  },
  {
    id: 3,
    name: "Igor",
    users: [
      { id: 0, tag: "me", name: "Вы" },
      { id: 3, tag: "igor", name: "Igor" },
    ],
    messages: [
      {
        user_author: { id: 3, tag: "igor", name: "Igor" },
        message: {
          id: 1,
          text: "Тогда пушу)",
          time: "14:50",
        },
      },
    ],
  },
];

const CURRENT_USER_ID = 0;

const MOCK_CHAT_STATUS: Record<
  number,
  { isRead: boolean; isOnline: boolean }
> = {
  1: { isRead: true, isOnline: false },
  2: { isRead: false, isOnline: true },
  3: { isRead: true, isOnline: true },
};

function getChatDisplayName(chat: Chat, currentUserId: number): string {
  const others = chat.users.filter((u) => u.id !== currentUserId);
  return others.length
    ? others.map((u) => u.name).join(", ")
    : chat.users.map((u) => u.name).join(", ");
}

export default function ChatList() {
  const [selectedId, setSelectedId] = useState<number>(MOCK_CHATS[1].id);

  return (
    <ul className="chat-list">
      {MOCK_CHATS.map((chat) => {
        const status = MOCK_CHAT_STATUS[chat.id];
        const lastMsg =
          chat.messages?.length > 0 ? chat.messages[chat.messages.length - 1] : null;
        return (
          <ChatListItem
            key={chat.id}
            isSelected={selectedId === chat.id}
            name={getChatDisplayName(chat, CURRENT_USER_ID)}
            time={lastMsg?.message.time ?? ""}
            hasCheckmark={status?.isRead ?? true}
            lastMessage={lastMsg?.message.text ?? "Нет сообщений"}
            isOnline={status?.isOnline ?? false}
            onClick={() => setSelectedId(chat.id)}
          />
        );
      })}
    </ul>
  );
}
