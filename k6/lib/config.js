let BASE_URL_DEFAULT = 'http://localhost:8004/api';
try {
  // Single source of truth for default API URL.
  BASE_URL_DEFAULT = open('k6/base-url.txt').trim();
} catch (e) {
  // Keep local fallback if file is unavailable.
}

export const BASE_URL = __ENV.BASE_URL || BASE_URL_DEFAULT;
export const CHAT_ID = parseInt(__ENV.CHAT_ID || '1', 10);
export const HISTORY_LIMIT = parseInt(__ENV.HISTORY_LIMIT || '50', 10);
export const SCROLL_PAGES = parseInt(__ENV.SCROLL_PAGES || '8', 10);

export const USERS = [
  { tag: 'alice_dev', password: 'alice1234' },
  { tag: 'bob_codes', password: 'bob12345' },
  { tag: 'charlie_js', password: 'char1234' },
  { tag: 'diana_py', password: 'diana1234' },
];
