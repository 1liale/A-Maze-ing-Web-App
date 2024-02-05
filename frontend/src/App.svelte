<script lang="ts">
  import GameSideBar from '@components/GameSideBar/GameSideBar.component.svelte';
  import GameWindow from '@components/GameWindow/GameWindow.component.svelte';
  import GradientHeading from '@components/GradientHeading/GradientHeading.component.svelte';
  import UserAction from '@components/UserAction/UserAction.components.svelte';
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
        <div class="h-full dark:bg-primary-500/80 p-1" slot="lead"><LightSwitch /></div>
        <GradientHeading className="h3">A-Maze-ing: Try some Mazes!</GradientHeading>
        <UserAction slot="trail">(actions)</UserAction>
      </AppBar>
      <GameSideBar slot="sidebarLeft" />
      <div class="h-full flex flex-col">
        <GameWindow className="h-5/6" config={undefined} data={undefined} />
        <div class="flex-1 p-3">
          <div class="h-full variant-ghost rounded-lg p-3">
            <p class="dark:text-primary-500/70 text-surface-500/70">
              <strong>Instructions: </strong>Use 'wasd' or arrow keys to move and mouse to control
              the view rotations.
            </p>
          </div>
        </div>
      </div>
    </AppShell>
  </main>
</Auth0Context>
