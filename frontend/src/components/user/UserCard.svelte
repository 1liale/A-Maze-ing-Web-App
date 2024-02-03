<script lang="ts">
  import { Auth0LogoutButton } from '@dopry/svelte-auth0';
  import { Avatar } from '@skeletonlabs/skeleton';
  import type { UserCardProfile } from './user.types';

  export let profile: UserCardProfile;
</script>

<div class="card border border-white/30 m-4 p-4 w-72 shadow-xl">
  <div class="space-y-4 m-6">
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
    <Auth0LogoutButton
      class="btn w-full border-2 hover:border-primary-400/30 border-transparent rounded variant-soft"
      >Logout</Auth0LogoutButton
    >
  </div>
  <div class="arrow bg-surface-100-800-token" />
</div>
