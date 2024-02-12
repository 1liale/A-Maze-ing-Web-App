import type { User } from '@auth0/auth0-spa-js';
import { apiData } from '@stores/data.stores';
import type { MazeInput } from 'types/maze.types';
import { requestAPI } from './api.service';

const apiUrl = import.meta.env.VITE_API_URL;

export const getMazesList = async () => {};

export const generateMaze = async (mazeInput: MazeInput) => {
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

export const saveMaze = async (authToken: string, userInfo: User) => {
  const options = {
    url: `${apiUrl}/maze/save/${userInfo.sub}`,
    method: 'POST',
    data: {},
    headers: {
      'content-type': 'application/json',
      Authorization: `Bearer ${authToken}`,
    },
  };
  const { error } = await requestAPI(options);

  if (error) {
    console.error(error);
  }
};

export const solveMaze = async () => {};
