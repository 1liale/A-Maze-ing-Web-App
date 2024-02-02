<script lang="ts">
  import { Auth0LoginButton, isAuthenticated, userInfo } from '@dopry/svelte-auth0';
  import { Avatar, popup, type PopupSettings } from '@skeletonlabs/skeleton';

  const popupFeatured: PopupSettings = {
    event: 'click',
    target: 'popupFeatured',
    placement: 'bottom-end',
  };
</script>

{#if !$isAuthenticated}
  <Auth0LoginButton
    class="btn border-2 hover:border-primary-400/30 border-transparent rounded variant-soft"
    >Login
  </Auth0LoginButton>
{:else}
  <!-- <Avatar width="w-12" src={$userInfo.picture} initials="AL" /> -->
  <!-- <span>
    <Auth0LogoutButton
      class="btn border-2 hover:border-primary-400/30 border-transparent rounded variant-soft"
      >Logout</Auth0LogoutButton
    >
  </span> -->
  <button use:popup={popupFeatured}>
    <Avatar
      border="border-2 border-surface-300-600-token hover:!border-primary-500/70"
      width="w-12"
      src={$userInfo.picture}
      initials="AL"
    />
  </button>

  <div class="card p-4 w-72 shadow-xl" data-popup="popupFeatured">
    <div class="space-y-4 m-6">
      <Avatar
        src="https://pbs.twimg.com/profile_images/1587479781544759297/TINbbJLC_400x400.png"
        width="w-16"
      />
      <div>
        <p class="font-bold">Skeleton</p>
        <p class="opacity-50">@SkeletonUI</p>
      </div>
      <p>Skeleton is a fully featured UI component library using the power of Svelte + Tailwind.</p>
      <div class="flex gap-4">
        <small><strong>100</strong> <span class="opacity-50">Following</span></small>
        <small><strong>1000</strong> <span class="opacity-50">Followers</span></small>
      </div>
      <a
        class="btn variant-soft w-full"
        href="https://twitter.com/SkeletonUI"
        target="_blank"
        rel="noreferrer"
      >
        View on Twitter
      </a>
    </div>
    <div class="arrow bg-surface-100-800-token" />
  </div>
{/if}
