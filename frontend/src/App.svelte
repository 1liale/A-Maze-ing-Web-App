<script lang="ts">
  import GradientHeading from '@components/GradientHeading/GradientHeading.component.svelte';
  import GameWindow from './layouts/GameWindow/GameWindow.layout.svelte';
  import SideBar from './layouts/SideBar/SideBar.layout.svelte';

  import UserAction from '@components/UserAction/UserAction.component.svelte';
  import { Auth0Context } from '@dopry/svelte-auth0';
  import { arrow, autoUpdate, computePosition, flip, offset, shift } from '@floating-ui/dom';
  import {
    AppBar,
    AppShell,
    LightSwitch,
    Toast,
    initializeStores,
    modeCurrent,
    setModeCurrent,
    storePopup,
  } from '@skeletonlabs/skeleton';
  import { onMount } from 'svelte';
  initializeStores();
  storePopup.set({ computePosition, autoUpdate, offset, shift, flip, arrow });

  const DOMAIN = import.meta.env.VITE_AUTH0_DOMAIN;
  const CLIENT_ID = import.meta.env.VITE_AUTH0_CLIENT_ID;
  const AUDIENCE = import.meta.env.VITE_AUTH0_AUDIENCE;

  onMount(() => {
    setModeCurrent($modeCurrent);
  });
</script>

<Auth0Context
  callback_url={window.location.href}
  logout_url={window.location.href}
  domain={DOMAIN}
  client_id={CLIENT_ID}
  audience={AUDIENCE}
>
  <main style="display: contents" class="h-full overflow-hidden">
    <Toast />
    <AppShell>
      <AppBar
        slot="header"
        gridColumns="grid-cols-3"
        slotDefault="place-self-center"
        slotTrail="place-content-end"
      >
        <div class="h-full dark:bg-primary-500/80 p-1" slot="lead"><LightSwitch /></div>
        <GradientHeading className="h3">A-Maze-ing: Try some Mazes!</GradientHeading>
        <UserAction slot="trail" />
      </AppBar>
      <SideBar slot="sidebarLeft" />
      <GameWindow />
    </AppShell>
  </main>
</Auth0Context>
