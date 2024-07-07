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
  init: () => Promise<void>;
  add: (title: string) => Promise<void>;
  update: (id: number, title: string, done: boolean) => Promise<void>;
  delete: (id: number) => Promise<void>;
}

export const todoStore = createStore();

function createStore(): TodoStore {
  const { subscribe, update, set } = writable<Todo[]>([]);

  return {
    subscribe,
    init: async () => {
      const token = localStorage.getItem("auth");

      if (token === null) return;

      const todos = await _userTodos(token);
      set(todos);
    },
    add: async (title) => {
      const token = localStorage.getItem("auth");

      if (token === null) return;

      const newTodo = await _addTodo(token, title);
      update((todos) => [...todos, newTodo]);
    },
    update: async (id, title, done) => {
      const token = localStorage.getItem("auth");

      if (token === null) return;

      await _updateTodo(token, id, title, done);
      update((todos) => {
        const todo = todos.find((t) => t.id === id);

        if (!todo) return todos;

        todo.title = title;
        todo.done = done;
        return todos;
      });
    },
    delete: async (id) => {
      const token = localStorage.getItem("auth");

      if (token === null) return;

      await _deleteTodo(token, id);
      update((todos) => todos.filter((t) => t.id !== id));
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

async function _addTodo(token: string, title: string): Promise<Todo> {
  const res = await fetch("http://localhost:8080/api/todos", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${token}`,
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

async function _updateTodo(
  token: string,
  id: number,
  title: string,
  done: boolean
) {
  const res = await fetch("http://localhost:8080/api/todos", {
    method: "PUT",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${token}`,
    },
    body: JSON.stringify({
      id,
      title,
      done,
    }),
  });

  const json = await res.json();

  if (!res.ok) {
    throw new Error(json.message);
  }
}

async function _deleteTodo(token: string, id: number) {
  const res = await fetch("http://localhost:8080/api/todos", {
    method: "DELETE",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${token}`,
    },
    body: JSON.stringify({
      id,
    }),
  });

  const json = await res.json();

  if (!res.ok) {
    throw new Error(json.message);
  }
}
