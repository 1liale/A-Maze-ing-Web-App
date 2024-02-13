<script lang="ts">
  import GradientHeading from '@components/GradientHeading/GradientHeading.component.svelte';
  import {
    Paginator,
    Table,
    tableMapperValues,
    type PaginationSettings,
  } from '@skeletonlabs/skeleton';
  import type { MazeScore } from 'types/maze.types';

  export let scores: MazeScore[] = [];

  let paginationSettings = {
    page: 0,
    limit: 5,
    size: scores.length,
    amounts: [1, 2, 5],
  } satisfies PaginationSettings;
  $: paginatedSource = tableMapperValues(scores, ['name', 'score']).slice(
    paginationSettings.page * paginationSettings.limit,
    paginationSettings.page * paginationSettings.limit + paginationSettings.limit,
  );
</script>

<div id="scoreboard">
  <GradientHeading className="h3 py-3 text-center">Scoreboard</GradientHeading>
  <div id="score-table" class="flex flex-col gap-2">
    <Table
      interactive
      source={{
        head: ['Player', 'Score'],
        body: paginatedSource,
      }}
      regionBody="variant-ghost rounded"
    />
    <Paginator
      select="leading-tight variant-ringed h-8 rounded p-1"
      controlVariant="variant-ringed h-8 rounded"
      justify="justify-between"
      bind:settings={paginationSettings}
      showFirstLastButtons={false}
      showPreviousNextButtons={true}
    />
  </div>
</div>
