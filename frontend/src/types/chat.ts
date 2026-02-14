export interface User {
  id: number;
  tag: string;
  name: string;
}

export interface MessageContent {
  id: number;
  text: string;
  time: string;
}

export interface Message {
  user_author: User;
  message: MessageContent;
}

export interface Chat {
  id: number;
  name: string;
  users: User[];
  messages: Message[];
}
