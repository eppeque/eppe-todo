import { setContext } from "svelte";
import { authStore, type AuthStore } from "./auth";
import { type TodoStore, todoStore } from "./todos";

export class StoreProvider {
  public authStore: AuthStore;
  public todoStore: TodoStore;

  constructor() {
    this.authStore = authStore;
    this.todoStore = todoStore;
  }
}

export const PROVIDER_CTX = "provider";

export function createContext(): StoreProvider {
  const provider = new StoreProvider();
  setContext(PROVIDER_CTX, provider);
  return provider;
}
