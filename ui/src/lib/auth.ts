import {
  writable,
  type Invalidator,
  type Subscriber,
  type Unsubscriber,
} from "svelte/store";
import type { User } from "./User";

export interface AuthStore {
  subscribe: (
    this: void,
    run: Subscriber<User | null>,
    invalidate?: Invalidator<User | null> | undefined
  ) => Unsubscriber;
  register: (
    username: string,
    email: string,
    password: string
  ) => Promise<void>;
}

export const authStore = createStore();

function createStore(): AuthStore {
  const { subscribe, set } = writable<User | null>(null);

  return {
    subscribe,
    register: async (username: string, email: string, password: string) => {
      const token = await _register(username, email, password);
      const user = await _accountData(token);

      localStorage.setItem("auth", token);
      set(user);
    },
  };
}

async function _register(
  username: string,
  email: string,
  password: string
): Promise<string> {
  const res = await fetch("http://localhost:8080/api/register", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      username,
      email,
      password,
    }),
  });

  const json = await res.json();

  if (!res.ok) {
    throw new Error(json.message);
  }

  return json.token;
}

async function _accountData(token: string): Promise<User> {
  const res = await fetch("http://localhost:8080/api/account", {
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${token}`,
    },
  });

  const json = await res.json();

  if (!res.ok) {
    throw new Error(json.message);
  }

  return json;
}
