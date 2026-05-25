import ws from 'k6/ws';
import { check, group, sleep } from 'k6';
import { login } from '../lib/auth.js';
import { CHAT_ID } from '../lib/config.js';

/** Участники chat_id=1 (см. db_init). Остальные пользователи не в этом чате. */
const WS_USERS = [
  { tag: 'alice_dev', password: 'alice1234' },
  { tag: 'bob_codes', password: 'bob12345' }
];
import {
  WS_URL,
  authEnvelope,
  chatMessageEnvelope,
  parseBroadcast
} from '../lib/ws.js';

const AUTH_WAIT_MS = parseInt(__ENV.AUTH_WAIT_MS || '800', 10);
const CLOSE_AFTER_MS = parseInt(__ENV.CLOSE_AFTER_MS || '4500', 10);
const MESSAGES_PER_SESSION = parseInt(__ENV.MESSAGES_PER_SESSION || '2', 10);
const MESSAGE_INTERVAL_MS = parseInt(__ENV.MESSAGE_INTERVAL_MS || '400', 10);

/**
 * Нагрузка через WebSocket: логин → auth → отправка chat_message → broadcast с message_id.
 * Сообщения пишутся в БД (см. wsserver.runBroadcast).
 *
 * Запуск: k6 run k6/scenarios/ws-messages.js
 * Grafana: WebSocket соединения, Messages created, Ws broadcast.
 */
export const options = {
  scenarios: {
    ws_messages: {
      executor: 'ramping-vus',
      startVUs: 0,
      stages: [
        { duration: '20s', target: 10 },
        { duration: '1m', target: 25 },
        { duration: '1m', target: 25 },
        { duration: '20s', target: 0 }
      ],
      gracefulRampDown: '10s'
    }
  },
  thresholds: {
    checks: ['rate>0.85'],
    ws_connecting: ['p(95)<5000'],
    ws_session_duration: ['p(95)<10000']
  }
};

function pickUser() {
  return WS_USERS[(__VU - 1) % WS_USERS.length];
}

function runWsSession(token, chatId) {
  const res = ws.connect(WS_URL, {}, function (socket) {
    let broadcasts = 0;
    let messagesSent = 0;

    socket.on('open', function () {
      socket.send(authEnvelope(token));
    });

    socket.on('message', function (data) {
      if (parseBroadcast(data) !== null) {
        broadcasts += 1;
      }
    });

    for (let i = 0; i < MESSAGES_PER_SESSION; i++) {
      const delay = AUTH_WAIT_MS + i * MESSAGE_INTERVAL_MS;
      socket.setTimeout(function () {
        const text = 'k6 ws vu=' + __VU + ' iter=' + __ITER + ' n=' + i;
        socket.send(chatMessageEnvelope(chatId, text));
        messagesSent += 1;
      }, delay);
    }

    socket.setTimeout(function () {
      check(null, {
        'ws: got broadcast': function () {
          return broadcasts >= 1;
        },
        'ws: sent messages': function () {
          return messagesSent >= 1;
        }
      });
      socket.close();
    }, CLOSE_AFTER_MS);
  });

  check(res, {
    'ws: status 101': function (r) {
      return r && r.status === 101;
    }
  });
}

export default function () {
  const user = pickUser();

  group('auth + ws messages', function () {
    const token = login(user.tag, user.password);
    if (!token) {
      sleep(1);
      return;
    }

    runWsSession(token, CHAT_ID);
    sleep(0.5);
  });
}
