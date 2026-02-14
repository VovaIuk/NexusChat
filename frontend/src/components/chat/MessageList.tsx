import type { Message } from "../../types/chat";
import {
  MessageOwnWithTail,
  MessageOwnWithoutTail,
  MessageOtherWithTail,
  MessageOtherWithoutTail,
} from "./message";
import "./message/message.css";

const exampleMessageOwn: Message = {
  user_author: { id: 1, tag: "me", name: "Я" },
  message: { id: 1, text: "Привет! Как дела? ttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttterrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrtttttttttttttt", time: "12:00" },
};

const exampleMessageOther: Message = {
  user_author: { id: 2, tag: "friend", name: "Друг" },
  message: { id: 2, text: "Привет! Всё отлично.", time: "12:01" },
};

export default function MessageList() {
  return (
    <div className="chat-messages-wrapper">
      <MessageOtherWithoutTail message={exampleMessageOther}/>
      <MessageOtherWithTail message={exampleMessageOther}/>
      <MessageOwnWithoutTail message={exampleMessageOwn}/>
      <MessageOwnWithTail message={exampleMessageOwn}/>
    </div>
  );
}
