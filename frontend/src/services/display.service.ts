import { MazeStatus, type MazeData } from 'types/maze.types';

const U = 1,
  D = 2,
  L = 4,
  R = 8;
const IN = 16;
const FRONTIER = 32;

const _checkIsPerim = (x: number, y: number, data: MazeData) => {
  return (
    x == 0 || y == 0 || y == 2 * data.width || x == 2 * data.height || (x % 2 == 0 && y % 2 == 0)
  );
};

const _checkIsPlaceWall = (x: number, y: number, data: MazeData) => {
  if (_checkIsPerim(x, y, data)) return true;
  if (x % 2 == 1 && y % 2 == 1) return false; // Cell position, do not place wall
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

export const getMappedPosition = (
  index: number,
  hOffset: number,
  data: MazeData,
): [x: number, y: number, z: number] | undefined => {
  const width = data.width;
  const height = data.height;
  let x = Math.floor(index / width);
  let y = index % width;
  return [x - Math.floor(height / 2), hOffset, y - Math.floor(width / 2)];
};

const _convertToMapped = (index: number, data: MazeData) => {
  let x = Math.floor(index / data.width);
  let y = index % data.width;

  x = 2 * x + 1;
  y = 2 * y + 1;
  y = 2 * data.width - y;

  return x * (2 * data.width + 1) + y;
};

export const fromMazeToMapped = (data: MazeData): MazeData | undefined => {
  const height = 2 * data.height + 1;
  const width = 2 * data.width + 1;

  const mappedGrid: MazeStatus[] = [];
  for (let x = 0; x < height; x++) {
    for (let y = 0; y < width; y++) {
      if (_checkIsPlaceWall(x, y, data)) mappedGrid.push(MazeStatus.WALL);
      else mappedGrid.push(MazeStatus.EMPTY);
    }
  }

  const start = _convertToMapped(data.start, data);
  mappedGrid[start] = MazeStatus.START;
  const end = _convertToMapped(data.end, data);
  mappedGrid[end] = MazeStatus.END;

  const mappedMazeData: MazeData = {
    start: start,
    end: end,
    grid: mappedGrid,
    width: width,
    height: height,
  };

  return mappedMazeData;
};
