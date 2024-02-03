<script lang="ts">
  import { Auth0LoginButton, isAuthenticated, userInfo } from '@dopry/svelte-auth0';
  import { Avatar, popup, type PopupSettings } from '@skeletonlabs/skeleton';
  import UserCard from './UserCard.svelte';
  import type { User, UserCardProfile } from './user.types.ts';

  const popupFeatured: PopupSettings = {
    event: 'click',
    target: 'popupFeatured',
    placement: 'bottom-end',
  };

  const profile: UserCardProfile = {};

  userInfo.subscribe((value: User) => {
    if (!value) return;
    console.log(value);
    profile.name = value.name;
    profile.email = value.email;
    profile.picture = value.picture;
  });

  // isAuthenticated.subscribe(() => )
</script>

{#if !$isAuthenticated}
  <Auth0LoginButton
    class="btn border-2 hover:border-primary-400/30 border-transparent rounded variant-soft"
    >Login
  </Auth0LoginButton>
{:else}
  <button use:popup={popupFeatured}>
    <Avatar
      border="border-2 border-surface-300-600-token hover:!border-primary-500/70"
      width="w-12"
      src={$userInfo.picture}
      initials="AL"
    />
  </button>
  <div data-popup="popupFeatured">
    <UserCard {profile} />
  </div>
{/if}
