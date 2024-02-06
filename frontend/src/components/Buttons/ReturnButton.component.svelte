<script lang="ts">
  import arrowBackDark from '@assets/arrow_back_dark.svg';
  import arrowBackLight from '@assets/arrow_back_light.svg';
  import { modeCurrent } from '@skeletonlabs/skeleton';

  export let onClick: () => void | undefined;
</script>

<button on:click={onClick} class="action-button w-20">
  <img class="w-5 h-5" src={$modeCurrent ? arrowBackLight : arrowBackDark} alt="Return Button" />
</button>
