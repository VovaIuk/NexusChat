import type { Message } from "../../types/chat";
import { useUser } from "../../contexts/UserContext";
import {
  MessageOwnWithTail,
  MessageOwnWithoutTail,
  MessageOtherWithTail,
  MessageOtherWithoutTail,
} from "./message";
import "./message/message.css";

interface MessageListProps {
  messages: Message[];
}

export default function MessageList({ messages }: MessageListProps) {
  const { user } = useUser();
  const currentUserId = user?.id;
  return (
    <div className="chat-messages-wrapper">
      {messages.map((message, index) => {
        const isOwn = message.user_author.id === currentUserId;
        const nextMessage = messages[index + 1];
        const timeDiffMs =
          nextMessage != null
            ? new Date(nextMessage.message.time).getTime() -
              new Date(message.message.time).getTime()
            : 0;
        const hasTail =
          nextMessage == null ||
          nextMessage.user_author.id !== message.user_author.id ||
          timeDiffMs > 2 * 60 * 1000;

        if (isOwn) {
          return hasTail ? (
            <MessageOwnWithTail key={message.message.id} message={message} />
          ) : (
            <MessageOwnWithoutTail key={message.message.id} message={message} />
          );
        }
        return hasTail ? (
          <MessageOtherWithTail key={message.message.id} message={message} />
        ) : (
          <MessageOtherWithoutTail key={message.message.id} message={message} />
        );
      })}
    </div>
  );
}
