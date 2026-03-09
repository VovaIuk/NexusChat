import {
  createContext,
  useContext,
  useMemo,
  useState,
  useEffect,
  useCallback,
  type ReactNode,
} from "react";
import type { Chat } from "../types/chat";
import { getChats, getChatMessages } from "../api/chat";

const MESSAGES_PAGE_SIZE = 50;

export interface ChatContextValue {
  chats: Chat[];
  setChats: React.Dispatch<React.SetStateAction<Chat[]>>;
  selectedChatIndex: number | null;
  setSelectedChatIndex: (index: number | null) => void;
  loadOlderMessages: (chatId: number, beforeMessageId: number) => Promise<void>;
  isLoadingOlder: boolean;
  hasMoreOlderMessages: (chatId: number) => boolean;
}

export const ChatContext = createContext<ChatContextValue | null>(null);

export function ChatProvider({ children }: { children: ReactNode }) {
  const [chats, setChats] = useState<Chat[]>([]);
  const [selectedChatIndex, setSelectedChatIndex] = useState<number | null>(null);
  const [isLoadingOlder, setIsLoadingOlder] = useState(false);
  const [chatsWithNoMoreOlder, setChatsWithNoMoreOlder] = useState<Set<number>>(
    () => new Set()
  );

  const hasMoreOlderMessages = useCallback(
    (chatId: number) => !chatsWithNoMoreOlder.has(chatId),
    [chatsWithNoMoreOlder]
  );

  const loadOlderMessages = useCallback(
    async (chatId: number, beforeMessageId: number) => {
      const token = localStorage.getItem("token");
      if (!token || isLoadingOlder || chatsWithNoMoreOlder.has(chatId)) return;
      setIsLoadingOlder(true);
      try {
        const { messages: older } = await getChatMessages(
          token,
          chatId,
          MESSAGES_PAGE_SIZE,
          beforeMessageId
        );
        if (older.length === 0) {
          setChatsWithNoMoreOlder((prev) => new Set(prev).add(chatId));
          return;
        }
        if (older.length < MESSAGES_PAGE_SIZE) {
          setChatsWithNoMoreOlder((prev) => new Set(prev).add(chatId));
        }
        setChats((prev) =>
          prev.map((c) =>
            c.id === chatId
              ? { ...c, messages: [...older, ...c.messages] }
              : c
          )
        );
      } catch (err) {
        console.error(err);
      } finally {
        setIsLoadingOlder(false);
      }
    },
    [isLoadingOlder, chatsWithNoMoreOlder]
  );

  useEffect(() => {
    console.log("Текущее время:", new Date().toLocaleString());
    const token = localStorage.getItem("token");
    if (!token) return;
    getChats(token, 1)
      .then((data) => {
        setChats(data.chats);
        setChatsWithNoMoreOlder(new Set());
        console.log(data);
      })
      .catch((err) => console.error(err));
  }, []);

  const value = useMemo<ChatContextValue>(
    () => ({
      chats,
      setChats,
      selectedChatIndex,
      setSelectedChatIndex,
      loadOlderMessages,
      isLoadingOlder,
      hasMoreOlderMessages,
    }),
    [chats, selectedChatIndex, loadOlderMessages, isLoadingOlder, hasMoreOlderMessages]
  );

  return (
    <ChatContext.Provider value={value}>{children}</ChatContext.Provider>
  );
}

export function useChat(): ChatContextValue {
  const ctx = useContext(ChatContext);
  if (ctx == null) {
    throw new Error("useChat must be used within ChatProvider");
  }
  return ctx;
}
