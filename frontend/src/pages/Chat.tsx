import '../components/chat/chat.css';
import '../components/chat/theme.css';
import ChatView from '../components/chat/ChatView';
import ChatSidebar from '../components/chat/sidebar/ChatSidebar';


export default function Chat(){
    return (
      <div className="theme-provider" data-theme="light">
        <div className="chat-page">
          <ChatSidebar/>
          <ChatView/>
        </div>
      </div>
    );
}

//TODO: добавить анимацию сообщениям