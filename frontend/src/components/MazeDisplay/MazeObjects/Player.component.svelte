<script lang="ts">
  import { getMappedPosition } from '@services/display.service';
  import { mazeData } from '@stores/data.stores';
  import { playerState, sidebarState } from '@stores/state.stores';
  import { onMount } from 'svelte';
  import { MazeStatus } from 'types/maze.types';
  import type { Player } from 'types/player.types';
  import Rook from './Rook.component.svelte';

  export let color: string | undefined;
  export let initPosition: [x: number, y: number, z: number] | undefined;

  onMount(() => {
    playerState.set({
      mappedPos: initPosition!,
      relPos: $mazeData!.start,
      moves: 0,
      hasWon: false,
    });
  });

  const onKeyDown = async (e: KeyboardEvent) => {
    if (!$mazeData || !$playerState || $sidebarState !== 'started') return;
    let isActionTriggered = true;
    let delt_x = 0;
    let delt_y = 0;

    switch (e.key) {
      case 's':
      case 'ArrowDown':
        delt_x = 1;
        break;
      case 'w':
      case 'ArrowUp':
        delt_x = -1;
        break;
      case 'a':
      case 'ArrowLeft':
        delt_y = 1;
        break;
      case 'd':
      case 'ArrowRight':
        delt_y = -1;
        break;
      default:
        isActionTriggered = false;
        break;
    }

    // an expected key is triggered
    if (isActionTriggered && $playerState.relPos && $playerState.mappedPos) {
      // update player position
      let ind = $playerState.relPos;
      let x = Math.floor(ind / $mazeData.width);
      let y = ind % $mazeData.width;

      x = x + delt_x;
      y = y + delt_y;
      ind = x * $mazeData.width + y;
      if ($mazeData.grid[ind] !== MazeStatus.WALL) {
        const mappedPos = getMappedPosition(ind, $playerState.mappedPos[1], $mazeData);

        const newPlayerState: Player = {
          mappedPos: mappedPos,
          relPos: ind,
          moves: $playerState.moves + 1,
          hasWon: ind === $mazeData.end,
        };

        playerState.set(newPlayerState);
        if (newPlayerState.hasWon) $sidebarState = 'finished';
      }
    }
  };
</script>

<svelte:window on:keydown|preventDefault={onKeyDown} />
<Rook position={$playerState?.mappedPos} {color} />
