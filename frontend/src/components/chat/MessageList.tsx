import type { Message } from "../../types/chat";
import { useUser } from "../../contexts/UserContext";
import {
  MessageOwnWithTail,
  MessageOwnWithoutTail,
  MessageOtherWithTail,
  MessageOtherWithoutTail,
} from "./message";
import "./message/message.css";
import { useEffect, useRef } from "react";
import { useChat } from "../../contexts/ChatContext";

interface MessageListProps {
  messages: Message[];
  chatId: number;
}

export default function MessageList({ messages, chatId }: MessageListProps) {
  const { user } = useUser();
  const currentUserId = user?.id;
  const { loadOlderMessages, isLoadingOlder, hasMoreOlderMessages } = useChat();
  const scrollContainerRef = useRef<HTMLDivElement>(null);
  const firstMessageRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    const container = scrollContainerRef.current;
    const first = firstMessageRef.current;
    if (
      !container ||
      !first ||
      messages.length === 0 ||
      isLoadingOlder ||
      !hasMoreOlderMessages(chatId)
    )
      return;

    const observer = new IntersectionObserver(
      (entries) => {
        if (!entries[0].isIntersecting) return;
        if (!hasMoreOlderMessages(chatId)) return;
        const firstMsg = messages[0];
        if (firstMsg) loadOlderMessages(chatId, firstMsg.message.id);
      },
      { root: container, rootMargin: "0px", threshold: 0 }
    );
    observer.observe(first);
    return () => observer.disconnect();
  }, [messages, chatId, loadOlderMessages, isLoadingOlder, hasMoreOlderMessages]);

  return (
    <div ref={scrollContainerRef} className="chat-messages-wrapper">
      {messages.map((message, index) => {
        const isFirst = index === 0;
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

        let content;
        if (isOwn) {
          content = hasTail ? (
            <MessageOwnWithTail key={message.message.id} message={message} />
          ) : (
            <MessageOwnWithoutTail key={message.message.id} message={message} />
          );
        } else {
          content = hasTail ? (
            <MessageOtherWithTail key={message.message.id} message={message} />
          ) : (
            <MessageOtherWithoutTail key={message.message.id} message={message} />
          );
        }

        return isFirst ? (
          <div key={message.message.id} ref={firstMessageRef}>
            {content}
          </div>
        ) : (
          content
        );
      })}
    </div>
  );
}
