<script lang="ts">
  import { getMappedPosition } from '@services/display.service';
  import { mazeData } from '@stores/data.stores';
  import { sidebarState as state } from '@stores/state.stores';
  import { T } from '@threlte/core';
  import { OrbitControls } from '@threlte/extras';
  import { MazeStatus } from 'types/maze.types';
  import { SideBarState } from 'types/sidebar.types';
  import Block from './MazeObjects/Block.component.svelte';
  import Player from './MazeObjects/Player.component.svelte';

  const wallHeight = 0.25;
  const pathHeight = 0.1;
</script>

<T.PerspectiveCamera makeDefault position={[20, 34, 0]}>
  <OrbitControls enablePan={false} maxDistance={40} minDistance={8} maxPolarAngle={1.56} />
</T.PerspectiveCamera>
<T.DirectionalLight position={[3, 10, 7]} />
<T.AmbientLight />
{#if $mazeData}
  {#each $mazeData.grid as item, index (index)}
    {#if item === MazeStatus.WALL}
      <Block height={wallHeight} position={getMappedPosition(index, wallHeight / 2, $mazeData)} />
    {/if}
  {/each}
  {#if $state === SideBarState.STARTED || $state === SideBarState.FINISHED}
    <Block
      color="#ff3e00"
      height={pathHeight}
      position={getMappedPosition($mazeData.end, pathHeight / 2, $mazeData)}
    />
    <Player
      color="cyan"
      initPosition={getMappedPosition($mazeData.start, 2 * pathHeight, $mazeData)}
    />
  {/if}
{/if}
