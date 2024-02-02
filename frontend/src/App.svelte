<script lang="ts">
  import { Auth0Context, userInfo } from '@dopry/svelte-auth0';
  import { AppBar, AppShell, LightSwitch } from '@skeletonlabs/skeleton';
  import Title from './components/Title.svelte';
  import UserAction from './components/UserAction.svelte';

  const DOMAIN = import.meta.env.VITE_AUTH0_DOMAIN;
  const CLIENT_ID = import.meta.env.VITE_AUTH0_CLIENT_ID;
</script>

<!-- <pre>isLoading: {$isLoading}</pre>
<pre>isAuthenticated: {$isAuthenticated}</pre>
<pre>authToken: {$authToken}</pre>
<pre>idToken: {$idToken}</pre>
<pre>userInfo: {JSON.stringify($userInfo, null, 2)}</pre>
<pre>authError: {$authError}</pre>  -->
<Auth0Context domain={DOMAIN} client_id={CLIENT_ID}>
  <main style="display: contents" class="h-full overflow-hidden">
    <AppShell>
      <AppBar
        slot="header"
        gridColumns="grid-cols-3"
        slotDefault="place-self-center"
        slotTrail="place-content-end"
      >
        <div class="h-full dark:bg-primary-400 p-1" slot="lead">
          <LightSwitch />
        </div>
        <span class="flex gap-2 items-center">
          <Title className="h3">A-Maze-ing: Try some Mazes</Title>😎
        </span>
        <UserAction slot="trail">(actions)</UserAction>
      </AppBar>
      <pre>userInfo: {JSON.stringify($userInfo, null, 2)}</pre>
      <!-- <svelte:fragment slot="pageFooter">Page Footer</svelte:fragment> -->
    </AppShell>
  </main>
</Auth0Context>
