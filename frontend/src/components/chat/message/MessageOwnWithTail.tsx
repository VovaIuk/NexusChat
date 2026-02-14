import type { Message } from "../../../types/chat";
import CheckmarkIcon from "../CheckmarkIcon";

interface MessageItemProps {
  message: Message;
}

export default function MessageOwnWithTail({
  message,
}: MessageItemProps) {
  return (
    <div className="chat-message chat-message--separator">
      <div className="chat-message-row">
        <div className="chat-message__avatar"></div>
        <div className="chat-message-bubble chat-message-bubble--own chat-message-bubble--tail">
          {message.message.text.split("\n").map((line, i) => (
            <span key={i}>
              {line}
              {i < message.message.text.split("\n").length - 1 && <br />}
            </span>
          ))}
          <span className="chat-message-bubble__meta">
            <span className="chat-message-bubble__meta-time">
              {message.message.time}
            </span>
            <CheckmarkIcon />
          </span>
        </div>
      </div>
    </div>
  );
}
