import { MazeStatus, type MazeData } from 'types/maze.types';

const U = 1,
  D = 2,
  L = 4,
  R = 8;
const IN = 16;
const FRONTIER = 32;

// ############ Converts MazeData to Mappped Representation for Display ############

const _checkIsPerim = (x: number, y: number, data: MazeData) => {
  return (
    x == 0 || y == 0 || y == 2 * data.width || x == 2 * data.height || (x % 2 == 0 && y % 2 == 0)
  );
};

const _checkIsCell = (x: number, y: number, data: MazeData) => {
  return (x % 2 == 1 && y % 2 == 1)
}

const _checkIsPlaceWall = (x: number, y: number, data: MazeData) => {
  if (_checkIsPerim(x, y, data)) return true;
  if (x % 2 == 1 && y % 2 == 1) return false; // Cell position, do not place wall
  y = 2 * data.width - y;
  x = (x - 1) / 2;
  y = (y - 1) / 2;
  const checkBelow = x % 1 != 0; // x is a float
  const checkRight = y % 1 != 0; // y is a float 
  const ind = Math.floor(x) * data.width + Math.floor(y);
  if (checkBelow && (data.grid[ind] & D) == 0) {
    return true;
  }

  if (checkRight && (data.grid[ind] & R) == 0) {
    return true;
  }

  return false;
};

// for display only
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

const _convertToMappedPos = (index: number, mazeWidth: number) => {
  let x = Math.floor(index / mazeWidth);
  let y = index % mazeWidth;

  x = 2 * x + 1;
  y = 2 * y + 1;
  y = 2 * mazeWidth - y;

  return x * (2 * mazeWidth + 1) + y;
};

export const mapMazeDataFromSrc = (data: MazeData): MazeData | undefined => {
  const height = 2 * data.height + 1;
  const width = 2 * data.width + 1;

  const mappedGrid: MazeStatus[] = [];
  for (let x = 0; x < height; x++) {
    for (let y = 0; y < width; y++) {
      if (_checkIsPlaceWall(x, y, data)) mappedGrid.push(MazeStatus.WALL);
      else mappedGrid.push(MazeStatus.EMPTY);
    }
  }

  const mappedSolution = data.solution
    .filter(sol => sol !== data.start && sol !== data.end)
    .map(sol => _convertToMappedPos(sol, data.width));
  const start = _convertToMappedPos(data.start, data.width);
  mappedGrid[start] = MazeStatus.START;
  const end = _convertToMappedPos(data.end, data.width);
  mappedGrid[end] = MazeStatus.END;

  const mappedMazeData: MazeData = {
    start: start,
    end: end,
    grid: mappedGrid,
    width: width,
    height: height,
    solution: mappedSolution,
    history: data.history
  };

  console.log("maze data before", data)
  console.log("maze data after mapping", mappedMazeData)
  console.log("maze data reconvert to orig", fromMappedToMaze(mappedMazeData))

  return mappedMazeData;
};

// ############ Converts Mappped Representation to Condensed Format for storage / backend computation ############

const _convertToOriginalPos = (index: number, data: MazeData) => {
  let x = Math.floor(index / data.width);
  let y = index % data.width;

  x = (x - 1) / 2;
  y = (data.width - y - 1) / 2;

  return Math.floor(x) * Math.floor(data.width / 2) + Math.floor(y);
};

const _checkIsMappedCell = (x: number, y: number , mappedData: MazeData) => {
  return mappedData.grid[x * mappedData.width + y] === MazeStatus.EMPTY
}

export const fromMappedToMaze = (mappedData: MazeData): MazeData | undefined => {
  const originalHeight = (mappedData.height - 1) / 2;
  const originalWidth = (mappedData.width - 1) / 2;

  const originalGrid: MazeStatus[] = new Array(originalHeight * originalWidth).fill(MazeStatus.EMPTY);
  
  for (let x = 0; x < mappedData.height; x++) {
    for (let y = 0; y < mappedData.width; y++) {
      // if (_checkIsPlaceWallReversed(x, y, mappedData)) continue;
      // const originalIndex = _convertToOriginal(x * mappedData.width + y, mappedData);
      // if (mappedData.grid[x * mappedData.width + y] === MazeStatus.WALL) {
      //   originalGrid[originalIndex] = MazeStatus.WALL;
      // }
      if (_checkIsMappedCell(x, y, mappedData)) {
        // console.log("Mapped Cell", {x, y, state: mappedData.grid[x * mappedData.width + y]})
      }
    }
  }

  const originalStart = _convertToOriginalPos(mappedData.start, mappedData);
  originalGrid[originalStart] = MazeStatus.START;
  const originalEnd = _convertToOriginalPos(mappedData.end, mappedData);
  originalGrid[originalEnd] = MazeStatus.END;

  const originalMazeData: MazeData = {
    ...mappedData,
    start: originalStart,
    end: originalEnd,
    grid: originalGrid,
    width: originalWidth,
    height: originalHeight,
  };

  return originalMazeData;
};
