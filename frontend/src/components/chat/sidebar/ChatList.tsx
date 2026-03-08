import { ChatListItem } from "./ChatListItem.tsx";
import type { Chat } from "../../../types/chat.ts";
import { useChat } from "../../../contexts/ChatContext";
import "./chatList.css";

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
  const { chats, selectedChatIndex, setSelectedChatIndex } = useChat();

  return (
    <ul className="chat-list">
      {chats.map((chat, index) => {
        const status = MOCK_CHAT_STATUS[chat.id];
        const lastMsg =
          chat.messages?.length > 0 ? chat.messages[chat.messages.length - 1] : null;
        return (
          <ChatListItem
            key={chat.id}
            isSelected={selectedChatIndex === index}
            name={getChatDisplayName(chat, CURRENT_USER_ID)}
            time={lastMsg?.message.time ?? ""}
            hasCheckmark={status?.isRead ?? true}
            lastMessage={lastMsg?.message.text ?? "Нет сообщений"}
            isOnline={status?.isOnline ?? false}
            onClick={() => setSelectedChatIndex(index)}
          />
        );
      })}
    </ul>
  );
}
