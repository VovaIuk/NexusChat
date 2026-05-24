import http from 'k6/http';
import { check, group, sleep } from 'k6';
import { login, authHeaders } from '../lib/auth.js';
import {
  BASE_URL,
  CHAT_ID,
  HISTORY_LIMIT,
  SCROLL_PAGES,
  USERS
} from '../lib/config.js';

/**
 * Нагрузочный сценарий для курсовой:
 * - логин нескольких пользователей
 * - список чатов
 * - постраничная загрузка истории (имитация скролла в большом чате)
 *
 * Запуск: k6 run k6/scenarios/load.js  (из корня проекта)
 * Grafana: http://localhost:13000 → NexusChat Overview
 */
export const options = {
  scenarios: {
    chat_history_load: {
      executor: 'ramping-vus',
      startVUs: 0,
      stages: [
        { duration: '30s', target: 15 },
        { duration: '2m', target: 40 },
        { duration: '2m', target: 40 },
        { duration: '30s', target: 0 }
      ],
      gracefulRampDown: '15s'
    }
  },
  thresholds: {
    http_req_failed: ['rate<0.05'],
    http_req_duration: ['p(95)<3000'],
    'http_req_duration{name:history_page}': ['p(95)<2500']
  }
};

function pickUser() {
  return USERS[(__VU - 1) % USERS.length];
}

function fetchChats(token) {
  const res = http.get(`${BASE_URL}/v1/chats?limit_messages=5`, {
    headers: authHeaders(token),
    tags: { name: 'chats_list' }
  });
  check(res, { 'chats 200': function (r) { return r.status === 200; } });
  return res;
}

function scrollChatHistory(token, chatId) {
  let beforeMessageId = null;

  for (let page = 0; page < SCROLL_PAGES; page++) {
    let url = `${BASE_URL}/v1/chats/${chatId}/messages?limit=${HISTORY_LIMIT}`;
    if (beforeMessageId !== null) {
      url += `&before_message_id=${beforeMessageId}`;
    }

    const res = http.get(url, {
      headers: authHeaders(token),
      tags: { name: 'history_page' }
    });

    const pageOk = check(res, {
      'history 200': function (r) { return r.status === 200; },
      'history has messages': function (r) {
        try {
          const body = r.json();
          return Array.isArray(body.messages);
        } catch (e) {
          return false;
        }
      }
    });

    if (!pageOk) {
      break;
    }

    const messages = res.json('messages');
    if (!messages || messages.length === 0) {
      break;
    }

    beforeMessageId = messages[0].message.id;
    sleep(0.15);
  }
}

export default function () {
  const user = pickUser();

  group('auth', function () {
    const token = login(user.tag, user.password);
    if (!token) {
      sleep(1);
      return;
    }

    group('chats', function () {
      fetchChats(token);
      sleep(0.2);
    });

    group('history_pagination', function () {
      scrollChatHistory(token, CHAT_ID);
      sleep(0.3);
    });
  });
}
