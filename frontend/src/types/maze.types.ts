export interface MazeConfig {
  materials?: { [key: string]: any };
  renderer?: WebGL2RenderingContext;
  thickness?: number;
}

export interface MazeData {
  start: number;
  end: number;
  grid: number[];
  width: number;
  height: number;
}

export interface MazeInput {
  width: number;
  height: number;
  generator: 'prim' | 'kruskal';
  solver: 'bfs' | 'bbfs' | 'dfs';
}
