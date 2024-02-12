export interface MazeConfig {
  materials?: { [key: string]: any };
  thickness?: number;
}

export enum MazeStatus {
  WALL,
  EMPTY,
  VISITED,
  START,
  END,
}

export interface MazeData {
  start: number;
  end: number;
  grid: number[];
  width: number;
  height: number;
}

export interface MazeMeta {
  solution: number[];
  history: number[][];
}

export interface MazeInput {
  width: number;
  height: number;
  generator: 'prim' | 'kruskal';
  solver: 'bfs' | 'bbfs' | 'dfs';
}

export interface MazeScore {
  name: string;
  time: number;
}
