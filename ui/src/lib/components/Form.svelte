<script lang="ts">
  import type { InputData } from "$lib/InputData";
  import Button from "./Button.svelte";
  import Input from "./Input.svelte";

  export let inputs: InputData[];
  export let submitText: string;
</script>

<form class="flex flex-col py-4" on:submit>
  {#each inputs as input (input.name)}
    <Input {...input} />
  {/each}

  <Button text={submitText} />
</form>
