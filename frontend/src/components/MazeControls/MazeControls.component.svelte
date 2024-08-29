<script lang="ts">
  import ReturnButton from '@components/Buttons/ReturnButton.component.svelte';
  import SaveButton from '@components/Buttons/SaveButton.component.svelte';
  import Timer from '@components/Timer/Timer.component.svelte';
  import { authToken, isAuthenticated, userInfo } from '@dopry/svelte-auth0';
  import { checkUserCreate, generateMaze, getScoreboard, saveMaze } from '@services/data.service';
  import { computeScore } from '@services/misc.service';
  import { getToastStore } from '@skeletonlabs/skeleton';
  import { apiData, mazeData, mazeInput } from '@stores/data.stores';
  import {
    playerState,
    resetGame,
    setupGame,
    solveTime,
    startGame,
    showSolution,
    sidebarState as state,
    stopGame,
    readyGame
  } from '@stores/state.stores';
  import type { MazeSaveFormat } from 'types/maze.types';
  import { SideBarState } from 'types/sidebar.types';

  const toastStore = getToastStore();
  const generators = ['prim', 'kruskal'];
  const solvers = ['bfs', 'bbfs', 'dfs'];
  const sliderRange = { min: 3, max: 15 };
  $: input = $mazeInput;

  const callGenerateMaze = async () => {
    mazeInput.set(input);
    await generateMaze(input).then(readyGame);
  };

  const callSaveMaze = async () => {
    const score = computeScore(
      $solveTime,
      $mazeData!.width * $mazeData!.height,
      $playerState.moves,
    );

    const data: MazeSaveFormat = {
      maze: $apiData.data.maze,
      solution: $apiData.solution,
      score: score,
    };
    try {
      await checkUserCreate($authToken, $userInfo);
      await saveMaze($authToken, $userInfo, data).then(() => {
        toastStore.trigger({
          timeout: 10000,
          message: 'Successfully saved maze record! Your score is: ' + score,
        });
        resetGame();
      });
      await getScoreboard(); // update scoreboard after saving
    } catch (e) {
      console.error(e);
    }
  };
</script>

<div class="container h-full py-3">
  <div class="h-full flex flex-col justify-end gap-4">
    {#if $state === SideBarState.INIT}
      <div id="base-input-group" class="flex flex-col gap-3">
        <span class="flex gap-6">
          Width:
          <input
            type="range"
            bind:value={input.width}
            min={sliderRange.min}
            max={sliderRange.max}
          />
          <p class="w-12">{input.width} x</p>
        </span
        >
        <span class="flex gap-6">
          Height: 
          <input
            type="range"
            bind:value={input.height}
            min={sliderRange.min}
            max={sliderRange.max}
          />
          <p class="w-12">{input.height} x</p>
        </span>
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
        <!-- TODO: Implement solve feature in V2 -->
        <!-- <button on:click={setupGame} class="action-button"> Setup Maze </button> -->
      </div>
    {:else if $state === SideBarState.SETUP}
      <aside class="alert variant-filled-error">
        <i class="fa-solid fa-triangle-exclamation text-4xl"></i>
        <div class="alert-message">
          <p>Feature currently unavailable!</p>
        </div>
      </aside>
      <span class="flex gap-2">
        <ReturnButton className="w-20" onClick={resetGame} />
        <button on:click={readyGame} class="action-button grow"> Done </button>
      </span>
      
    {:else}
      <Timer />
      <span class="flex gap-2">
        <ReturnButton
          className="w-20"
          disabled={$state === SideBarState.STARTED}
          onClick={resetGame}
        />
        {#if $state === SideBarState.WAITING}
          <button on:click={startGame} class="action-button grow"> Start </button>
        {:else if $state === SideBarState.FINISHED || $state === SideBarState.SHOW_SOLUTION }
          {#if $playerState.hasWon}
            <SaveButton disabled={!$isAuthenticated} onClick={callSaveMaze} className="grow" />
          {:else}
            <button on:click={$state === SideBarState.FINISHED ? showSolution : stopGame} class="action-button grow">
              {$state === SideBarState.FINISHED ? "Show Solution" : "Hide Solution"} 
            </button>
          {/if}
        {:else}
          <button on:click={stopGame} class="action-button grow"> Stop </button>
        {/if}
      </span>
    {/if}
  </div>
</div>
