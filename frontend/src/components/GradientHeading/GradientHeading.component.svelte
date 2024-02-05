<script lang="ts">
  export let className: string;
</script>

<div class={className}>
  <span class="gradient-heading bg-clip-text text-transparent box-decoration-clone">
    <slot />
  </span>
</div>
