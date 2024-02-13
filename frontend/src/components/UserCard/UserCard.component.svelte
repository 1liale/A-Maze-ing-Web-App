<script lang="ts">
  import type { UserCardProfile } from 'types/user.types';

  import { Auth0LogoutButton, authToken, userInfo } from '@dopry/svelte-auth0';
  import { getScoreboard, removeUser } from '@services/data.service';
  import { Avatar } from '@skeletonlabs/skeleton';

  export let profile: UserCardProfile;

  export const callRemoveUser = async () => {
    await removeUser($authToken, $userInfo);
    await getScoreboard();
  };
</script>

<div class="card border border-white/30 m-4 p-4 w-72 shadow-xl">
  <div class="space-y-4 m-4">
    {#if profile.picture}
      <Avatar width="w-12" src={profile.picture} initials="AL" />
    {/if}
    <div>
      {#if profile.name}
        <p class="font-bold">{profile.name}</p>
      {/if}
    </div>
    <div class="flex gap-4">
      {#if profile.email}
        <small><span class="opacity-50">{profile.email}</span></small>
      {/if}
    </div>
    {#if profile.bio}
      <p>{profile.bio}</p>
    {/if}
    <button
      on:click={callRemoveUser}
      class="btn variant-soft-error border-2 border-transparent hover:border-primary-400/30 w-full rounded"
    >
      Remove User
    </button>
    <Auth0LogoutButton class="soft-action-button w-full">Logout</Auth0LogoutButton>
  </div>
</div>
