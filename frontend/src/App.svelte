<script lang="ts">
  import { Auth0Context } from '@dopry/svelte-auth0';
  import { arrow, autoUpdate, computePosition, flip, offset, shift } from '@floating-ui/dom';
  import {
    AppBar,
    AppShell,
    LightSwitch,
    modeCurrent,
    setModeCurrent,
    storePopup,
  } from '@skeletonlabs/skeleton';
  import { onMount } from 'svelte';
  import Title from './components/Title/Title.component.svelte';
  import UserAction from './components/UserAction/UserAction.components.svelte';
  storePopup.set({ computePosition, autoUpdate, offset, shift, flip, arrow });

  const DOMAIN = import.meta.env.VITE_AUTH0_DOMAIN;
  const CLIENT_ID = import.meta.env.VITE_AUTH0_CLIENT_ID;
  const AUDIENCE = import.meta.env.VITE_AUTH0_AUDIENCE;

  onMount(() => {
    setModeCurrent($modeCurrent);
  });
</script>

<Auth0Context domain={DOMAIN} client_id={CLIENT_ID} audience={AUDIENCE}>
  <main style="display: contents" class="h-full overflow-hidden">
    <AppShell>
      <AppBar
        slot="header"
        gridColumns="grid-cols-3"
        slotDefault="place-self-center"
        slotTrail="place-content-end"
      >
        <div class="h-full dark:bg-primary-400 p-1" slot="lead"><LightSwitch /></div>
        <span class="flex gap-2 items-center">
          <Title className="h3">A-Maze-ing: Try some Mazes</Title>😎
        </span>
        <UserAction slot="trail">(actions)</UserAction>
      </AppBar>
    </AppShell>
  </main>
</Auth0Context>
