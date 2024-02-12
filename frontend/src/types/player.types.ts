export interface Player {
  mappedPos?: [x: number, y: number, z: number];
  relPos?: number;
  moves: number;
  hasWon: boolean;
}
