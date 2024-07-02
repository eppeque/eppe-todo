<script>
  import Navlink from "./Navlink.svelte";

  let opened = false;

  function open() {
    opened = true;
  }

  function close() {
    opened = false;
  }
</script>

<nav class="py-10 md:flex md:justify-between md:items-center">
  <div class="flex justify-between items-center">
    <h2
      class="text-lg lg:text-xl font-bold bg-gradient-to-r from-sky-500 to-teal-600 bg-clip-text text-transparent"
    >
      Eppe Todo
    </h2>
    <button class="material-symbols-outlined md:hidden" on:click={open}
      >menu</button
    >
  </div>

  <div
    class="fixed inset-0 backdrop-blur-sm md:hidden"
    class:hidden={!opened}
  ></div>

  <ul
    class:hidden={!opened}
    class="p-10 md:p-0 fixed md:relative top-4 md:top-0 right-4 md:right-0 shadow-xl md:shadow-none rounded-lg md:rounded-none bg-white md:bg-transparent flex md:flex flex-col md:flex-row md:items-center gap-4 md:gap-8"
  >
    <button
      class="material-symbols-outlined fixed top-6 right-6 md:hidden"
      on:click={close}>close</button
    >
    <Navlink text="Home" url="/" />
    <Navlink text="Sign In" url="/auth/login" />
    <li>
      <a
        href="/auth/register"
        class="py-2 px-4 bg-gradient-to-r from-cyan-500 to-teal-500 hover:from-cyan-600 hover:to-teal-600 text-white font-semibold text-sm shadow-sm rounded-md"
        >Sign Up</a
      >
    </li>
  </ul>
</nav>
