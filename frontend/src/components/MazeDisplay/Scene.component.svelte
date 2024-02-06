<script lang="ts">
  import { mazeData } from '@stores/data';
  import { T } from '@threlte/core';
  import { OrbitControls } from '@threlte/extras';
  import type { MazeData } from 'types/maze.types';
  import Wall from './MazeObjects/Wall.component.svelte';
  const checkIsPerim = (x: number, y: number, data: MazeData) => {
    return (
      x == 0 || y == 0 || y == 2 * data.width || x == 2 * data.height || (x % 2 == 0 && y % 2 == 0)
    );
  };

  const checkIsPlaceWall = (x: number, y: number, data: MazeData) => {
    // Cell position, do not place wall
    if (x % 2 == 1 && y % 2 == 1) return false;
    y = 2 * data.width - y;
    x = (x - 1) / 2;
    y = (y - 1) / 2;
    const checkBelow = x % 1 != 0;
    const checkRight = y % 1 != 0;
    const ind = Math.floor(x) * data.width + Math.floor(y);
    if (checkBelow && (data.grid[ind] & D) == 0) {
      return true;
    }

    if (checkRight && (data.grid[ind] & R) == 0) {
      return true;
    }

    return false;
  };
  const wallHeight = 0.2;

  const U = 1,
    D = 2,
    L = 4,
    R = 8;
  const IN = 16;
  const FRONTIER = 32;
</script>

<T.PerspectiveCamera makeDefault position={[20, 34, 0]}>
  <OrbitControls enablePan={false} maxDistance={40} minDistance={10} maxPolarAngle={1.56} />
</T.PerspectiveCamera>
<T.DirectionalLight position={[3, 10, 7]} />
<T.AmbientLight />
{#if $mazeData}
  {#each { length: 2 * $mazeData.height + 1 } as _, x}
    {#each { length: 2 * $mazeData.width + 1 } as _, y}
      {#if checkIsPerim(x, y, $mazeData)}
        <Wall
          height={wallHeight}
          position={[x - $mazeData.height, wallHeight, y - $mazeData.width]}
        />
      {:else if checkIsPlaceWall(x, y, $mazeData)}
        <Wall
          height={wallHeight}
          position={[x - $mazeData.width, wallHeight, y - $mazeData.height]}
        />
      {/if}
    {/each}
  {/each}
{/if}
