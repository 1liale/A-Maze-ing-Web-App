<script lang="ts">
  import MazeControls from '@components/MazeControls/MazeControls.component.svelte';
  import ScoreBoard from '@components/ScoreBoard/ScoreBoard.component.svelte';
  import { getScoreboard } from '@services/data.service';
  import { mazeScores } from '@stores/data.stores';
</script>

<div id="sidebar" class="container h-full w-96 p-3 bg-surface-200 dark:bg-surface-700">
  <div class="content h-full flex flex-col justify-between">
    {#await getScoreboard() then}
      <ScoreBoard scores={$mazeScores} />
    {/await}
    <MazeControls />
  </div>
</div>
