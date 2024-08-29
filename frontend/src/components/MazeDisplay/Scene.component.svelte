<script>
  import { getMappedPosition } from '@services/display.service';
  import { mazeData } from '@stores/data.stores';
  import { sidebarState as state } from '@stores/state.stores';
  import { T } from '@threlte/core';
  import { OrbitControls } from '@threlte/extras';
  import { MazeStatus } from 'types/maze.types';
  import { SideBarState } from 'types/sidebar.types';
  import Block from './MazeObjects/Block.component.svelte';
  import Player from './MazeObjects/Player.component.svelte';
  import { interpolateRainbow } from "d3-scale-chromatic"; // D3 library for color interpolation

  const wallHeight = 0.25;
  const pathHeight = 0.1;

  // Function to get a rainbow color based on index
  const getRainbowColor = (index, total) => {
    return interpolateRainbow(index / total);
  };
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
  {:else if $state === SideBarState.SETUP}
      <Block
      color="#ff3e00"
      height={pathHeight}
      position={getMappedPosition($mazeData.end, pathHeight / 2, $mazeData)}
    />
  {:else if $state === SideBarState.SHOW_SOLUTION}
    {#each $mazeData.solution as item, index (index)}
      <Block color={getRainbowColor(index, $mazeData.solution.length)} height={pathHeight} position={getMappedPosition(item, pathHeight / 2, $mazeData)} />
    {/each}

    <Block
      color="#ff3e00"
      height={pathHeight}
      position={getMappedPosition($mazeData.start, pathHeight / 2, $mazeData)}
    />
    <Block
      color="#ff3e00"
      height={pathHeight}
      position={getMappedPosition($mazeData.end, pathHeight / 2, $mazeData)}
    />
  {/if}
{/if}


