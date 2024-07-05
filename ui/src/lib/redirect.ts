import { get } from "svelte/store";
import type { AuthStore } from "./auth";
import { goto } from "$app/navigation";

export function checkAuth(authStore: AuthStore): void {
  const user = get(authStore);

  if (user === null) {
    goto("/");
  }
}

export function checkNotAuth(authStore: AuthStore): void {
  const user = get(authStore);

  if (user !== null) {
    goto("/todos");
  }
}
