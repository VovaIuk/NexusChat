import { useState, type SubmitEvent } from 'react'
import "./messageInput.css"

function SendIcon() {
  return (
    <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" aria-hidden>
      <g clipPath="url(#clip0_send)">
        <path d="M3.4 20.4L20.85 12.92C21.66 12.57 21.66 11.43 20.85 11.08L3.4 3.60002C2.74 3.31002 2.01 3.80002 2.01 4.51002L2 9.12002C2 9.62002 2.37 10.05 2.87 10.11L17 12L2.87 13.88C2.37 13.95 2 14.38 2 14.88L2.01 19.49C2.01 20.2 2.74 20.69 3.4 20.4V20.4Z" fill="currentColor"/>
      </g>
      <defs>
        <clipPath id="clip0_send">
          <rect width="24" height="24" fill="white"/>
        </clipPath>
      </defs>
    </svg>
  );
}

export default function MessageInput() {
  const [text, setText] = useState("");

  const handleSubmit = (e: SubmitEvent<HTMLFormElement>) => {
    e.preventDefault();
    //TODO: добавить функцию отправки сообщения
    setText("");
  }

  const hasText = text.trim().length > 0; 

  return (
    <div className="chat-input-wrap">
      <form 
      className={`chat-input-row ${hasText ? 'chat-input-row--has-text' : ''}`}
      onSubmit={handleSubmit}
      >
        <input
          type="text"
          className="chat-input"
          placeholder="Введите сообщение..."
          value={text}
          onChange={(e) => setText(e.target.value)}
        />
        <button
          type="submit"
          className="chat-send"
          disabled={false}
        >
          <SendIcon />
        </button>
      </form>
    </div>
  );
}