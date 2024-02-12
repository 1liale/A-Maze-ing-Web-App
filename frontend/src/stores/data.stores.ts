import { fromMazeToMapped } from '@services/display.service';
import type { Readable } from 'svelte/motion';
import { derived, writable, type Writable } from 'svelte/store';
import type { MazeData, MazeInput, MazeMeta, MazeScore } from 'types/maze.types';

export const apiData: Writable<{ [key: string]: any }> = writable({});

const defaultInput: MazeInput = {
  width: 3,
  height: 3,
  generator: 'prim',
  solver: 'bfs',
};
export const mazeInput: Writable<MazeInput> = writable(defaultInput);

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
      };

      return fromMazeToMapped(data);
    }
    return undefined;
  },
);

export const mazeMeta: Readable<MazeMeta | undefined> = derived([apiData], ([$apiData]) => {
  if ($apiData.solution && $apiData.data) {
    const meta: MazeMeta = {
      solution: $apiData.solution,
      history: $apiData.data.history,
    };
    return meta;
  }
  return undefined;
});

const defaultMazeScores: MazeScore[] = [
  { name: 'Test', time: 1 },
  { name: 'Test', time: 2 },
  { name: 'Test', time: 3 },
  { name: 'Test', time: 4 },
  { name: 'Test', time: 5 },
  { name: 'Test', time: 1 },
  { name: 'Test', time: 2 },
  { name: 'Test', time: 3 },
  { name: 'Test', time: 4 },
  { name: 'Test', time: 5 },
];
export const mazeScores: Writable<MazeScore[]> = writable(defaultMazeScores);

export const solveTime: Writable<number> = writable(0);
