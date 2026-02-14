import MessageList from "./MessageList";
import MessageInput from "./MessageInput";

export default function ChatView() {
  return (
    <div className="chat-view">
      <header className="chat-view__header">
        Иван Иванов
      </header>
      <MessageList/>
      <MessageInput/>
    </div>
  );
}
