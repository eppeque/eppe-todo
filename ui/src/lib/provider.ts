import { setContext } from "svelte";
import { authStore, type AuthStore } from "./auth";

export class StoreProvider {
  public authStore: AuthStore;

  constructor() {
    this.authStore = authStore;
  }
}

export const PROVIDER_CTX = "provider";

export function createContext(): StoreProvider {
  const provider = new StoreProvider();
  setContext(PROVIDER_CTX, provider);
  return provider;
}
