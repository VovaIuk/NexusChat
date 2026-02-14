import { useState, type FormEvent } from "react";
import "./chatSearch.css";

export default function ChatSearch() {
  const [query, setQuery] = useState("");

  const handleSubmit = (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    // Логика поиска не реализована — только отправка по Enter
  };

  return (
    <div className="chat-search-wrap">
      <form className="chat-search-row" onSubmit={handleSubmit}>
        <input
          type="text"
          className="chat-search-input"
          placeholder="Поиск..."
          value={query}
          onChange={(e) => setQuery(e.target.value)}
          aria-label="Поиск"
        />
      </form>
    </div>
  );
}
