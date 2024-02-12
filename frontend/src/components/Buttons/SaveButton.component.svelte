<script lang="ts">
  import { popup, type PopupSettings } from '@skeletonlabs/skeleton';

  export let onClick: () => (Promise<void> | void) | undefined;
  export let disabled: boolean | undefined = undefined;
  export let className: string = '';

  const popupHover: PopupSettings = {
    event: 'hover',
    target: 'popupHover',
    placement: 'top',
  };
</script>

{#if disabled}
  <button
    {disabled}
    on:click={onClick}
    class="action-button [&>*]:pointer-events-none {className}"
    use:popup={popupHover}
  >
    Save
  </button>
  <div class="card p-4 variant-filled-error" data-popup="popupHover">
    <p>Login first to save!</p>
    <div class="arrow variant-filled-error" />
  </div>
{:else}
  <button on:click={onClick} class="action-button {className}"> Save </button>
{/if}
