import { BASE_URL } from './config.js';

export const WS_URL = BASE_URL.replace(/^http/, 'ws') + '/v1/ws';

export function authEnvelope(token) {
  return JSON.stringify({
    type: 'auth',
    data: JSON.stringify({ token: token })
  });
}

export function chatMessageEnvelope(chatId, text) {
  return JSON.stringify({
    type: 'chat_message',
    data: JSON.stringify({ chat_id: chatId, text: text })
  });
}

export function parseBroadcast(data) {
  try {
    const msg = JSON.parse(data);
    if (typeof msg.message_id === 'number' && typeof msg.chat_id === 'number') {
      return msg;
    }
  } catch (e) {
    // ignore
  }
  return null;
}
