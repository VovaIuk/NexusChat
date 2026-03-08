import { useEffect } from 'react';
import { login } from '../api/auth';
import '../components/chat/chat.css';
import '../components/chat/theme.css';
import ChatView from '../components/chat/ChatView';
import ChatSidebar from '../components/chat/sidebar/ChatSidebar';
import { ChatProvider } from '../contexts/ChatContext';
import { UserProvider } from '../contexts/UserContext';

export default function Chat() {
  useEffect(() => {
    async function initLogin() {
      try {
        const data = await login("alice_dev", "alice1234");
        localStorage.setItem("token", data.token.refresh);
      } catch (error) {
        console.error(error);
      }
    }
    initLogin();
  }, []);

  return (
    <UserProvider>
      <ChatProvider>
        <div className="theme-provider" data-theme="light">
          <div className="chat-page">
            <ChatSidebar />
            <ChatView />
          </div>
        </div>
      </ChatProvider>
    </UserProvider>
  );
}

//TODO: добавить анимацию сообщениям
