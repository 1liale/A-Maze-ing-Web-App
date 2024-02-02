<script lang="ts">
  import { Auth0LoginButton, isAuthenticated, userInfo } from '@dopry/svelte-auth0';
  import { Avatar } from '@skeletonlabs/skeleton';
</script>

{#if !$isAuthenticated}
  <Auth0LoginButton
    class="btn border-2 hover:border-primary-400/30 border-transparent rounded variant-soft"
    >Login
  </Auth0LoginButton>
{:else}
  <Avatar width="w-12" src={$userInfo.picture} initials="AL" />
  <!-- <span>
    <Auth0LogoutButton
      class="btn border-2 hover:border-primary-400/30 border-transparent rounded variant-soft"
      >Logout</Auth0LogoutButton
    >
  </span> -->
{/if}
