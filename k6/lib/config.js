// Default must match k6/config.ps1 ($K6_BASE_URL)
export const BASE_URL = __ENV.BASE_URL || 'http://localhost:8003/api';
export const CHAT_ID = parseInt(__ENV.CHAT_ID || '1', 10);
export const HISTORY_LIMIT = parseInt(__ENV.HISTORY_LIMIT || '50', 10);
export const SCROLL_PAGES = parseInt(__ENV.SCROLL_PAGES || '8', 10);

export const USERS = [
  { tag: 'alice_dev', password: 'alice1234' },
  { tag: 'bob_codes', password: 'bob12345' },
  { tag: 'charlie_js', password: 'char1234' },
  { tag: 'diana_py', password: 'diana1234' },
];
