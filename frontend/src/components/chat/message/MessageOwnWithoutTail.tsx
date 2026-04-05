import type { Message } from "../../../types/chat";
import { formatMessageTime } from "../../../utils/formatMessageTime";
import CheckmarkIcon from "../CheckmarkIcon";

interface MessageItemProps {
  message: Message;
  animate: boolean;
}

export default function MessageOtherWithoutTail({
  message,
  animate=false,
}: MessageItemProps) {
  return (
    <div className="chat-message">
      <div className={`chat-message-row ${animate ? "chat-message-row--animate" : ""}`}>
        <div className="chat-message__spacer"></div>
        <div className="chat-message-bubble chat-message-bubble--own chat-message-bubble--no-tail">
          {message.message.text.split("\n").map((line, i) => (
            <span key={i}>
              {line}
              {i < message.message.text.split("\n").length - 1 && <br />}
            </span>
          ))}
          <span className="chat-message-bubble__meta">
            <span className="chat-message-bubble__meta-time">
              {formatMessageTime(message.message.time)}
            </span>
            <CheckmarkIcon />
          </span>
        </div>
      </div>
    </div>
  );
}
