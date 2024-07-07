import { get } from "svelte/store";
import type { AuthStore } from "./auth";
import { goto } from "$app/navigation";

export function checkNotAuth(authStore: AuthStore): void {
  const state = get(authStore);

  if (state.user !== null) {
    goto("/todos");
  }
}
