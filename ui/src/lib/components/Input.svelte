<script lang="ts">
  export let label: string;
  export let type: string = "text";
  export let name: string;
  export let value: string = "";
  export let required: boolean;
</script>

<label class="pb-2" for={name}>{label}</label>
<input
  class="p-3 mb-4 outline-none border-2 border-gray-200 focus:border-teal-500 rounded-md"
  {type}
  id={name}
  {name}
  {value}
  {required}
/>
