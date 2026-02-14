import type { Message } from "../../../types/chat";

import "../message/message.css";


interface MessageItemProps {
  message: Message;
}

export default function MessageOtherWithoutTail({
  message,
}: MessageItemProps) {
  return (
    <div className="chat-message">
      <div className="chat-message-row">
        <div className="chat-message__spacer"></div>
        <div className="chat-message-bubble chat-message-bubble--other chat-message-bubble--no-tail">
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
          </span>
        </div>
      </div>
    </div>
  );
}
