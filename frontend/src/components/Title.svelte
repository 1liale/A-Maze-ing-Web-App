<script lang="ts">
  export let className = '';
</script>

<div class={className}>
  <span
    class="bg-gradient-to-br from-red-500 to-yellow-500 bg-clip-text text-transparent box-decoration-clone"
  >
    <slot />
  </span>
</div>
