import '../components/chat/chat.css';
import '../components/chat/theme.css';
import ChatView from '../components/chat/ChatView';
import ChatSidebar from '../components/chat/sidebar/ChatSidebar';
import { ChatProvider } from '../contexts/ChatContext';
import { SocketProvider } from '../contexts/SocketContext';

export default function Chat() {
  return (
    <ChatProvider>
      <SocketProvider>
        <div className="theme-provider" data-theme="light">
          <div className="chat-page">
            <ChatSidebar />
            <ChatView />
          </div>
        </div>
      </SocketProvider>
    </ChatProvider>
  );
}
