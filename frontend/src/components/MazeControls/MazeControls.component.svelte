<script lang="ts">
  import ReturnButton from '@components/Buttons/ReturnButton.component.svelte';
  import Timer from '@components/Timer/Timer.component.svelte';
  import { generateMaze } from '@services/maze.service';
  import { mazeInput } from '@stores/data';
  import { sidebarState as state } from '@stores/state';

  const callGenerateMaze = async () => {
    await generateMaze($mazeInput).then(() => state.set('generate'));
  };

  const callSolveMaze = () => {
    state.set('solve');
  };

  const resetState = () => {
    state.set('init');
  };

  const generators = ['prim', 'kruskal'];
  const solvers = ['bfs', 'bbfs', 'dfs'];
  const sliderRange = { min: 3, max: 15 };
</script>

<div class="container h-full py-3">
  {#if $state === 'init'}
    <div class="h-full flex flex-col justify-end gap-4">
      <div id="base-input-group" class="flex flex-col gap-3">
        <span class="flex gap-6"
          >Width: <input
            type="range"
            bind:value={$mazeInput.width}
            min={sliderRange.min}
            max={sliderRange.max}
          />
          {$mazeInput.width}px</span
        >
        <span class="flex gap-6"
          >Height: <input
            type="range"
            bind:value={$mazeInput.height}
            min={sliderRange.min}
            max={sliderRange.max}
          />
          {$mazeInput.height}px</span
        >
        <span class="flex gap-6"
          >Generator:
          <select
            id="generator-select"
            class="select leading-tight p-1"
            bind:value={$mazeInput.generator}
          >
            <option>{generators[0]}</option>
            <option>{generators[1]}</option>
          </select>
        </span>
        <span class="flex gap-6">
          Solver:
          <select
            id="solver-select"
            class="select leading-tight p-1"
            bind:value={$mazeInput.solver}
          >
            <option>{solvers[0]}</option>
            <option>{solvers[1]}</option>
            <option>{solvers[2]}</option>
          </select>
        </span>
      </div>
      <div class="flex flex-col gap-2">
        <button on:click={callGenerateMaze} class="action-button"> Generate </button>
        <button on:click={callSolveMaze} class="action-button"> Solve </button>
      </div>
    </div>
  {:else if $state === 'generate'}
    <div class="h-full flex flex-col gap-2 justify-between">
      <Timer />
      <span class="flex gap-2">
        <ReturnButton onClick={resetState} />
        <button class="action-button grow"> Start </button>
      </span>
    </div>
  {:else if $state === 'solve'}
    <span class="flex gap-2">
      <ReturnButton onClick={resetState} />
      <button class="action-button grow"> Start </button>
    </span>
  {/if}
</div>
