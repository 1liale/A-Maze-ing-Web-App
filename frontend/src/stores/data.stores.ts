import { mapMazeDataFromSrc, } from '@services/display.service';
import { localStorageStore } from '@skeletonlabs/skeleton';
import type { Readable } from 'svelte/motion';
import { derived, writable, type Writable } from 'svelte/store';
import type { MazeData, MazeInput, MazeScore } from 'types/maze.types';

export const apiData: Writable<{ [key: string]: any }> = localStorageStore('apiData', {});

export const defaultInput: MazeInput = {
  width: 3,
  height: 3,
  generator: 'prim',
  solver: 'bfs',
};
export const mazeInput: Writable<MazeInput> = localStorageStore('mazeInput', defaultInput);

export const mazeData: Readable<MazeData | undefined> = derived(
  [apiData, mazeInput],
  ([$apiData, $mazeInput]) => {
    if ($apiData.data && $apiData.data.maze && $mazeInput) {
      const data: MazeData = {
        start: $apiData.data.maze.start,
        end: $apiData.data.maze.end,
        grid: $apiData.data.maze.grid,
        width: $mazeInput.width,
        height: $mazeInput.height,
        solution: $apiData.solution,
        history: $apiData.data.history,
      };

      return mapMazeDataFromSrc(data);
    }
    return undefined;
  },
);

export const mazeScores: Writable<MazeScore[]> = writable([]);
