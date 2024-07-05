import {
  writable,
  type Invalidator,
  type Subscriber,
  type Unsubscriber,
} from "svelte/store";
import type { Todo } from "./Todo";

export interface TodoStore {
  subscribe: (
    this: void,
    run: Subscriber<Todo[]>,
    invalidate?: Invalidator<Todo[]> | undefined
  ) => Unsubscriber;
  init: (token: string) => Promise<void>;
  add: (title: string) => Promise<void>;
}

export const todoStore = createStore();

function createStore(): TodoStore {
  const { subscribe, update, set } = writable<Todo[]>();

  return {
    subscribe,
    init: async (token) => {
      const todos = await _userTodos(token);
      set(todos);
    },
    add: async (title) => {
      const newTodo = await _addTodo(title);
      update((todos) => [...todos, newTodo]);
    },
  };
}

async function _userTodos(token: string): Promise<Todo[]> {
  const res = await fetch("http://localhost:8080/api/todos", {
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${token}`,
    },
  });

  const json = await res.json();

  if (!res.ok) {
    throw new Error(json.message);
  }

  return json.todos;
}

async function _addTodo(title: string): Promise<Todo> {
  const res = await fetch("http://localhost:8080/api/todos", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      title,
    }),
  });

  const json = await res.json();

  if (!res.ok) {
    throw new Error(json.message);
  }

  return {
    id: json.id,
    title: json.title,
    done: false,
  };
}
