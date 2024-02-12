<script lang="ts">
  import ReturnButton from '@components/Buttons/ReturnButton.component.svelte';
  import SaveButton from '@components/Buttons/SaveButton.component.svelte';
  import Timer from '@components/Timer/Timer.component.svelte';
  import { authToken, isAuthenticated, userInfo } from '@dopry/svelte-auth0';
  import { generateMaze, saveMaze } from '@services/maze.service';
  import { mazeInput, solveTime } from '@stores/data.stores';
  import { playerState, sidebarState as state } from '@stores/state.stores';

  const generators = ['prim', 'kruskal'];
  const solvers = ['bfs', 'bbfs', 'dfs'];
  const sliderRange = { min: 3, max: 15 };
  $: input = $mazeInput;

  const callGenerateMaze = async () => {
    mazeInput.set(input);
    await generateMaze(input).then(() => state.set('waiting'));
  };

  const callSaveMaze = async () => {
    await saveMaze($authToken, $userInfo).then(() => state.set('init'));
  };

  const setupGame = () => {
    state.set('setup');
  };

  const startGame = () => {
    state.set('started');
  };

  const stopGame = () => {
    state.set('finished');
  };

  const resetGame = () => {
    state.set('init');
    $solveTime = 0;
    $playerState = undefined;
  };
</script>

<div class="container h-full py-3">
  <div class="h-full flex flex-col justify-end gap-4">
    {#if $state === 'init'}
      <div id="base-input-group" class="flex flex-col gap-3">
        <span class="flex gap-6"
          >Width: <input
            type="range"
            bind:value={input.width}
            min={sliderRange.min}
            max={sliderRange.max}
          />
          {input.width}px</span
        >
        <span class="flex gap-6"
          >Height: <input
            type="range"
            bind:value={input.height}
            min={sliderRange.min}
            max={sliderRange.max}
          />
          {input.height}px</span
        >
        <span class="flex gap-6"
          >Generator:
          <select
            id="generator-select"
            class="select leading-tight p-1"
            bind:value={input.generator}
          >
            <option>{generators[0]}</option>
            <option>{generators[1]}</option>
          </select>
        </span>
        <span class="flex gap-6">
          Solver:
          <select id="solver-select" class="select leading-tight p-1" bind:value={input.solver}>
            <option>{solvers[0]}</option>
            <option>{solvers[1]}</option>
            <option>{solvers[2]}</option>
          </select>
        </span>
      </div>
      <div class="flex flex-col gap-2">
        <button on:click={callGenerateMaze} class="action-button"> Generate </button>
        <!-- TODO: Implement solve feature -->
        <button on:click={setupGame} class="action-button"> Solve </button>
      </div>
    {:else if $state === 'waiting' || $state === 'started' || $state === 'finished'}
      <Timer />
      <span class="flex gap-2">
        <ReturnButton className="w-20" disabled={$state === 'started'} onClick={resetGame} />
        {#if $state === 'waiting'}
          <button on:click={startGame} class="action-button grow"> Start </button>
        {:else if $state === 'finished' && $playerState?.hasWon}
          <SaveButton disabled={!$isAuthenticated} onClick={callSaveMaze} className="grow" />
        {:else if $state === 'finished' && !$playerState?.hasWon}
          <button on:click={stopGame} class="action-button grow"> Show Solution </button>
        {:else}
          <button on:click={stopGame} class="action-button grow"> Stop </button>
        {/if}
      </span>
    {:else if $state === 'setup'}
      <aside class="alert variant-filled-error">
        <i class="fa-solid fa-triangle-exclamation text-4xl"></i>
        <div class="alert-message">
          <p>Feature unavailable!</p>
        </div>
      </aside>
      <ReturnButton className="w-full" onClick={resetGame} />
    {/if}
  </div>
</div>
