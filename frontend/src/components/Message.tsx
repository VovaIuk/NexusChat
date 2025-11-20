import "./Message.css"

interface MessageProps {
    content: string;
    time: string;
    isReceived: boolean;
}

function Message(props: MessageProps){
    return (
        <div className={`message ${props.isReceived ? "received" : "sent"}`}>
            <div className="message-content">{props.content}</div>
            <div className="message-time">{props.time}</div>
        </div>
    )
}

export default Message;