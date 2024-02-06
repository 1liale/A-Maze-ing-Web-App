import { apiData } from '@stores/data';
import type { MazeInput } from 'types/maze.types';
import { requestAPI } from './api.service';

const apiUrl = import.meta.env.VITE_API_URL;

const getMazesList = async () => {};

const generateMaze = async (mazeInput: MazeInput) => {
  const options = {
    url: `${apiUrl}/maze/generate`,
    method: 'POST',
    data: mazeInput,
    headers: {
      'content-type': 'application/json',
    },
  };
  const { data, error } = await requestAPI(options);

  if (data) {
    apiData.set(data.response);
  }

  if (error) {
    console.error(error);
  }
};

const solveMaze = async () => {};

export { generateMaze, getMazesList, solveMaze };
