import './ChatItem.css'

interface ChatItemProps {
    name: string;
    time: string;
    message: string;
    avatarText: string;
    isOnline?: boolean;
    isActive?: boolean;
}

function ChatItem(props: ChatItemProps){
    return (
        <div className={`chat-item ${props.isActive ? 'active': ''}`}>
            <div className="chat-avatar">
              {props.avatarText}
              {props.isOnline && <div className="online-indicator"></div>}
            </div>
            <div className="chat-info">
              <h4>
                <span className="chat-name">{props.name}</span>
                <span className="timestamp">{props.time}</span>
              </h4>
              <p className="last-message">{props.message}</p>
            </div>
        </div>
    )
}

export default ChatItem;