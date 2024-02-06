<script lang="ts">
  import MazeControls from '@components/MazeControls/MazeControls.component.svelte';
  import ScoreBoard from '@components/ScoreBoard/ScoreBoard.component.svelte';
</script>

<div id="sidebar" class="container h-full w-96 p-3 bg-surface-200 dark:bg-surface-700">
  <div class="content h-full flex flex-col justify-between">
    <ScoreBoard />
    <MazeControls />
  </div>
</div>
