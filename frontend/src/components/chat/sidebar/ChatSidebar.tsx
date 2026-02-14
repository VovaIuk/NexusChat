import ChatSearch from "./ChatSearch";
import ChatList from "./ChatList";

export default function ChatSidebar() {
  return (
    <div className="chat-sidebar">
      <ChatSearch />
      <ChatList />
    </div>
  );
}