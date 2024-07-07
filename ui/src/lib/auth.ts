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
    run: Subscriber<AuthState>,
    invalidate?: Invalidator<AuthState> | undefined
  ) => Unsubscriber;
  register: (
    username: string,
    email: string,
    password: string
  ) => Promise<void>;
  login: (email: string, password: string) => Promise<void>;
  init: () => Promise<void>;
  signOut: () => void;
}

export interface AuthState {
  initializing: boolean;
  user: User | null;
}

export const authStore = createStore();

function createStore(): AuthStore {
  const { subscribe, set, update } = writable<AuthState>({
    initializing: true,
    user: null,
  });

  return {
    subscribe,
    init: async () => {
      const token = localStorage.getItem("auth");

      if (token === null) {
        set({ initializing: false, user: null });
        return;
      }

      const refreshedToken = await _refreshToken(token);
      localStorage.setItem("auth", refreshedToken);

      const user = await _accountData(refreshedToken);
      set({ initializing: false, user: user });
    },
    register: async (username: string, email: string, password: string) => {
      const token = await _register(username, email, password);

      localStorage.setItem("auth", token);
      update((old) => {
        old.user = { username, email };
        return old;
      });
    },
    login: async (email: string, password: string) => {
      const token = await _login(email, password);
      const user = await _accountData(token);

      localStorage.setItem("auth", token);
      update((old) => {
        old.user = user;
        return old;
      });
    },
    signOut: () => {
      localStorage.removeItem("auth");

      update((old) => {
        old.user = null;
        return old;
      });
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

async function _login(email: string, password: string): Promise<string> {
  const res = await fetch("http://localhost:8080/api/login", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
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

async function _refreshToken(token: string): Promise<string> {
  const res = await fetch("http://localhost:8080/api/refresh", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      token,
    }),
  });

  const json = await res.json();

  if (!res.ok) {
    throw new Error(json.message);
  }

  return json.token;
}
