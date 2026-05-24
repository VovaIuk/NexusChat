import http from 'k6/http';
import { check } from 'k6';
import { BASE_URL } from './config.js';

export function login(tag, password) {
  const res = http.post(
    `${BASE_URL}/v1/login`,
    JSON.stringify({ tag: tag, password: password }),
    {
      headers: { 'Content-Type': 'application/json' },
      tags: { name: 'login' }
    }
  );

  const ok = check(res, {
    'login status 201': function (r) { return r.status === 201; },
    'login has token': function (r) {
      try {
        return r.json('token.refresh') !== undefined;
      } catch (e) {
        return false;
      }
    }
  });

  if (!ok) {
    return null;
  }

  return res.json('token.refresh');
}

export function authHeaders(token) {
  return {
    'Content-Type': 'application/json',
    Authorization: 'Bearer ' + token
  };
}
