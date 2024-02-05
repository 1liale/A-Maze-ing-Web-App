<script lang="ts">
  import GradientHeading from '@components/GradientHeading/GradientHeading.component.svelte';
  import MazeControls from '@components/MazeControls/MazeControls.component.svelte';
  import { Table, tableMapperValues, type TableSource } from '@skeletonlabs/skeleton';

  const sourceData = [
    { name: 'Hydrogen', time: 'H' },
    { name: 'Helium', time: 'He' },
    { name: 'Lithium', time: 'Li' },
    { name: 'Beryllium', time: 'Be' },
    { name: 'Boron', time: 'B' },
  ];

  const scoreboard: TableSource = {
    head: ['Player', 'Solve Time'],
    body: tableMapperValues(sourceData, ['name', 'time']),
  };
</script>

<div id="sidebar" class="container h-full w-96 p-3 bg-surface-200 dark:bg-surface-700">
  <div class="content h-full flex flex-col justify-between">
    <div id="scoreboard">
      <GradientHeading className="h3 py-3 text-center">Scoreboard</GradientHeading>
      <Table interactive source={scoreboard} />
    </div>
    <MazeControls />
  </div>
</div>
