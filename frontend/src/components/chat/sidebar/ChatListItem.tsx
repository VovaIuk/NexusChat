import "./chatList.css";
import CheckmarkIcon from "../CheckmarkIcon";

export interface ChatListItemProps {
  isSelected: boolean;
  name: string;
  time: string;
  hasCheckmark: boolean;
  lastMessage: string;
  isOnline: boolean;
  onClick: () => void;
}

export function ChatListItem({
  isSelected,
  name,
  time,
  hasCheckmark,
  lastMessage,
  isOnline,
  onClick,
}: ChatListItemProps) {
  return (
    <li
      className={`chat-list-item ${isSelected ? "chat-list-item--selected" : ""}`}
      onClick={onClick}
      role="button"
      tabIndex={0}
      onKeyDown={(e) => {
        if (e.key === "Enter" || e.key === " ") {
          e.preventDefault();
          onClick();
        }
      }}
    >
      <div className="chat-list-item__avatar-wrap">
        <div className="chat-list-item__avatar" aria-hidden />
        {isOnline && (
          <span
            className="chat-list-item__avatar-online"
            aria-label="В сети"
            aria-hidden
          />
        )}
      </div>
      <div className="chat-list-item__body">
        <div className="chat-list-item__top">
          <span className="chat-list-item__name">{name}</span>
          <div className="chat-list-item__meta">
            {hasCheckmark && (
              <span className="chat-list-item__checks" aria-hidden>
                <CheckmarkIcon />
              </span>
            )}
            {time && <span className="chat-list-item__time">{time}</span>}
          </div>
        </div>
        <span className="chat-list-item__preview">{lastMessage}</span>
      </div>
    </li>
  );
}
