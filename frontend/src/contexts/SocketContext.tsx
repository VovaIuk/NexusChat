import { createContext, useCallback, useContext, useEffect, useMemo, useRef, type ReactNode } from "react";
import { useChat, } from "./ChatContext";
import {type ClientWsEnvelope, WS_MESSAGE_TYPE, broadcastToMessage,
    isServerBroadcast,} from "../types/websocket/protocol"

export interface SocketContextValue {
    sendMessage: (chatId: number, text: string) => void;
}

const SocketContext = createContext<SocketContextValue | null>(null);

function getWsUrl(): string {
    const apiBase = import.meta.env.VITE_API_BASE as string; // http://localhost:8004/api
    return apiBase.replace(/^http/, "ws") + "/v1/ws";        // ws://localhost:8004/api/v1/ws
  }

export function SocketProvider({children}: {children: ReactNode}){
    const {setChats} = useChat();
    const socketRef = useRef<WebSocket | null>(null);

    useEffect(() => {
        const token = localStorage.getItem("token");
        if(!token){
            console.error("empty token");
            return;
        }

        const ws = new WebSocket(getWsUrl());
        console.log("init ws");
        socketRef.current = ws;

        ws.onopen = ()=>{
            const authData: ClientWsEnvelope = {
                type: WS_MESSAGE_TYPE.AUTH,
                data: JSON.stringify({ token }),
            };
            ws.send(JSON.stringify(authData));
            console.log("auth");
        };

        ws.onmessage = (ev) => {
            try {
                const raw = JSON.parse(ev.data as string) as unknown;
                if (!isServerBroadcast(raw)) return;
                const message = broadcastToMessage(raw);
                setChats((prev) =>
                  prev.map((chat) =>
                    chat.id === raw.chat_id
                      ? { ...chat, messages: [...chat.messages, message] }
                      : chat
                  )
                );
              } catch {
                // игнор неизвестных сообщений
              }
        }

        ws.onerror = () => {
            console.error("WebSocket error");
        };

        ws.onclose = () => {
            socketRef.current = null;
        };
      
        return () => {
            ws.close();
            if (socketRef.current === ws) {
              socketRef.current = null;
            }
            console.log("close ws");
        };
    },[setChats]);

    const sendMessage = useCallback((chatId: number, text: string) => {
        console.log("start send message (ws)");
        const ws = socketRef.current;
        console.log(ws);
        console.log(ws?.readyState);
        if (ws?.readyState !== WebSocket.OPEN) return;
        console.log("form message");
        const envelope: ClientWsEnvelope = {
            type: WS_MESSAGE_TYPE.CHAT_MESSAGE,
            data: JSON.stringify({ chat_id: chatId, text }),
          };
          ws.send(JSON.stringify(envelope));
    }, []);

    const value = useMemo<SocketContextValue>(
        () => ({
            sendMessage,
        }), [sendMessage]
    );

    return (
        <SocketContext.Provider value={value}>{children}</SocketContext.Provider>
    );
}

export function useSocket(): SocketContextValue {
    const ctx = useContext(SocketContext);
    if (ctx == null) {
      throw new Error("useSocket must be used within SocketProvider");
    }
    return ctx;
  }
