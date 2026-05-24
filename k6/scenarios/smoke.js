import http from 'k6/http';
import { check, sleep } from 'k6';
import { login, authHeaders } from '../lib/auth.js';
import { BASE_URL, CHAT_ID, HISTORY_LIMIT } from '../lib/config.js';

/** Быстрая проверка перед полным load-тестом (5 VU, 30 сек). */
export const options = {
  vus: 5,
  duration: '30s',
  thresholds: {
    http_req_failed: ['rate<0.1']
  }
};

export default function () {
  const token = login('alice_dev', 'alice1234');
  if (!token) {
    sleep(1);
    return;
  }

  const chats = http.get(`${BASE_URL}/v1/chats`, {
    headers: authHeaders(token)
  });
  check(chats, { 'chats ok': function (r) { return r.status === 200; } });

  const history = http.get(
    `${BASE_URL}/v1/chats/${CHAT_ID}/messages?limit=${HISTORY_LIMIT}`,
    { headers: authHeaders(token) }
  );
  check(history, { 'history ok': function (r) { return r.status === 200; } });

  sleep(0.5);
}
