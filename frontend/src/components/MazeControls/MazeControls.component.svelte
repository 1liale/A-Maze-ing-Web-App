<script lang="ts">
  import ReturnButton from '@components/Buttons/ReturnButton.component.svelte';
  import Timer from '@components/Timer/Timer.component.svelte';
  import { sidebarState as state } from '@stores/store';

  const callGenerateMaze = () => {
    state.set('generate');
  };

  const callSolveMaze = () => {
    state.set('solve');
  };

  const resetState = () => {
    state.set('init');
  };

  const maze_input: any = {
    width: 3,
    height: 3,
    generator: 'prim',
    solver: 'dfs',
  };

  const generators = ['prim', 'kruskal'];
  const solvers = ['bfs', 'bbfs', 'dfs'];
</script>

<div class="container h-full py-3">
  {#if $state === 'init'}
    <div class="h-full flex flex-col justify-end gap-4">
      <div id="base-input-group" class="flex flex-col gap-3">
        <span class="flex gap-6"
          >Width: <input type="range" bind:value={maze_input.width} min="3" max="35" />
          {maze_input.width}px</span
        >
        <span class="flex gap-6"
          >Height: <input type="range" bind:value={maze_input.height} min="3" max="35" />
          {maze_input.height}px</span
        >
        <span class="flex gap-6"
          >Generator:
          <select id="generator-select" class="select h-10" bind:value={maze_input.generator}>
            <option>{generators[0]}</option>
            <option>{generators[1]}</option>
          </select>
        </span>
        <span class="flex gap-6"
          >Solver:
          <select id="solver-select" class="select h-10" bind:value={maze_input.solver}>
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
