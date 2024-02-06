<script lang="ts">
  import { Auth0LoginButton, isAuthenticated, userInfo } from '@dopry/svelte-auth0';
  import { Avatar, popup, type PopupSettings } from '@skeletonlabs/skeleton';

  import type { User } from '@auth0/auth0-spa-js';
  import UserCard from '@components/UserCard/UserCard.component.svelte';
  import type { UserCardProfile } from 'types/user.types';

  const popupFeatured: PopupSettings = {
    event: 'click',
    target: 'popupFeatured',
    placement: 'bottom-end',
  };

  const profile: UserCardProfile = {};

  userInfo.subscribe((value: User) => {
    if (Object.keys(value).length === 0) return;
    console.log(value);
    profile.name = value.name;
    profile.email = value.email;
    profile.picture = value.picture;
    profile.initials = (value.given_name || ' ')[0] + (value?.family_name || ' ')[0];
  });
</script>

{#if !$isAuthenticated}
  <Auth0LoginButton class="soft-action-button">Login</Auth0LoginButton>
{:else}
  <button use:popup={popupFeatured}>
    <Avatar
      border="border-2 border-surface-300-600-token hover:!border-primary-500/70"
      width="w-12"
      src={profile.picture}
      initials={profile.initials}
    />
  </button>
  <div data-popup="popupFeatured">
    <UserCard {profile} />
  </div>
{/if}
