import {
  createContext,
  useContext,
  useMemo,
  useState,
  useEffect,
  type ReactNode,
} from "react";
import type { Chat } from "../types/chat";
import { getChats } from "../api/chat";

export interface ChatContextValue {
  chats: Chat[];
  setChats: React.Dispatch<React.SetStateAction<Chat[]>>;
  selectedChatIndex: number | null;
  setSelectedChatIndex: (index: number | null) => void;
}

export const ChatContext = createContext<ChatContextValue | null>(null);

export function ChatProvider({ children }: { children: ReactNode }) {
  const [chats, setChats] = useState<Chat[]>([]);
  const [selectedChatIndex, setSelectedChatIndex] = useState<number | null>(
    null
  );

  useEffect(() => {
    console.log("Текущее время:", new Date().toLocaleString());
    const token = localStorage.getItem("token");
    if (!token) return;
    getChats(token, 10)
      .then((data) => {setChats(data.chats); console.log(data)})
      .catch((err) => console.error(err));
  }, []);

  const value = useMemo<ChatContextValue>(
    () => ({
      chats,
      setChats,
      selectedChatIndex,
      setSelectedChatIndex,
    }),
    [chats, selectedChatIndex]
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
