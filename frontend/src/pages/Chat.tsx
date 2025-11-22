import './Chat.css'
import ChatItem from '../components/ChatItem'
import { useState, useRef, useEffect } from 'react'
import Message from '../components/Message';

export interface WsMessage {
  username: string;    
  message: string;    
  time: string;   
  IsReceived: boolean;
}


function Chat() {
  const [typingText, setTypingText] = useState("");
  const [messages, setMessages] = useState<WsMessage[]>([]);
  const ws = useRef<WebSocket | null>(null);

  useEffect(()=>{

    ws.current = new WebSocket("ws://192.168.0.15:8003/ws");

    ws.current.onopen = () => {
      console.log("WS: подключено");
    };
  
    ws.current.onmessage = (event: MessageEvent) => {
      try{
        const message: WsMessage = JSON.parse(event.data);
        setMessages(prevMessages => [...prevMessages, {...message,
          IsReceived: true,
        }])
      } catch(error){
        console.error('Error parsing WebSocket message:', error);
      }
      console.log("WS: сообщение от сервера:", event.data);
    };
  
    ws.current.onerror = (err) => {
      console.log("WS: ошибка", err);
    };
  
    ws.current.onclose = () => {
      console.log("WS: соединение закрыто");
    };

    return () => {
      ws.current?.close();
    };
  }, [])


  const handleSendMessage = () => {
    ws.current?.send(JSON.stringify({
      username: "Имя",
      message: typingText,
    })
  )}


  return (
    <>
      <div className="chat-app">
        {/* Боковая панель */}
        <div className="chats-sidebar">
          <div className="sidebar-header">
            <div className="user-profile">
              <div className="user-avatar">АК</div>
              <div className="user-info">
                <h3>Алексей К.</h3>
                <p><span className="status-dot"></span>Online</p>
              </div>
            </div>
          </div>
    
          <div className="search-container">
            <input type="text" placeholder="Поиск чатов..."/>
          </div>

          <div className="chats-list">
            <ChatItem 
              avatarText='НП' 
              message='Присылаю много документов, eioiner erngoern ernfero'
              time='15:33'
              name='Иван Иванов Иван Иван Иван'
              isActive={false}
              isOnline={true}
            />

            <ChatItem 
              avatarText='МП'
              message='Присылаю документы'
              time='11:45'
              name='Дмиттрий Сидоров'
              isActive={true}
            />

            <ChatItem 
              avatarText='МП'
              message='Присылаю документы по проекту eee eroreooeri erofnreoin rnefioerno'
              time='11:45'
              name='ДмиттрийСидоровАнтонович'
              isActive={false}
              isOnline={false}
            />

            <ChatItem 
              avatarText='МП'
              message='Присылаю документы по проекту eee eroreooeri erofnreoin rnefioerno'
              time='11:45'
              name='Дмиттрий Сидоров'
              isOnline={true}
            />
          </div>
        </div>

        <div className="chat-area">
          <div className="chat-header">
            <div className="chat-title">
              <div className="chat-avatar">
                МП
                <div className="online-indicator"></div>
              </div>
              <div>
                <h3>Мария Петрова</h3>
                <span className="chat-status">online</span>
              </div>
            </div>
          </div>
          <div className="messages-container">
            {messages.map((message) => (
              <Message
              content={message.message}
              time={message.time}
              isReceived={message.IsReceived}
            />
            ))}
          </div>
          <div className="typing-indicator">
            Мария печатает
            <div className="typing-dots">
              <span></span>
              <span></span>
              <span></span>
            </div>
          </div>
          <div className="message-input-container">
            <div className="message-input-wrapper">
              <input 
                type="text" 
                className="message-input" 
                placeholder="Напишите сообщение..."
                value={typingText}
                onChange={(e)=>setTypingText(e.target.value)}
              />
              <div className="input-actions">
                <button 
                  className="send-btn"
                  onClick={handleSendMessage}
                >➤</button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </>
  );
}

export default Chat
