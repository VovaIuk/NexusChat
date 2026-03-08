import { useChat } from "../../contexts/ChatContext";
import MessageList from "./MessageList";
import MessageInput from "./MessageInput";

export default function ChatView() {
  const { chats, selectedChatIndex } = useChat();
  const selectedChat =
    selectedChatIndex != null && chats[selectedChatIndex] != null
      ? chats[selectedChatIndex]
      : null;

  return (
    <div className="chat-view">
      <header className="chat-view__header">
        {selectedChat?.name ?? "Выберите чат"}
      </header>
      <MessageList messages={selectedChat?.messages ?? []} />
      <MessageInput />
    </div>
  );
}
