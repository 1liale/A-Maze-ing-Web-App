import type { User } from '@auth0/auth0-spa-js';
import { apiData, mazeScores } from '@stores/data.stores';
import type { MazeInput, MazeSaveFormat, MazeScore } from 'types/maze.types';
import { requestAPI } from './api.service';

const apiUrl = import.meta.env.VITE_API_URL;

export const checkUserCreate = async (authToken: string, userInfo: User) => {
  const userId = userInfo.sub!.split('|')[1];

  const options = {
    url: `${apiUrl}/user/${userId}?name=${userInfo.name}`,
    method: 'GET',
    headers: {
      'content-type': 'application/json',
      Authorization: `Bearer ${authToken}`,
    },
  };

  const { data, error } = await requestAPI(options);
  if (data) console.log(data.response);
  if (error) {
    console.error(error);
  }
};

export const removeUser = async (authToken: string, userInfo: User) => {
  const userId = userInfo.sub!.split('|')[1];

  const options = {
    url: `${apiUrl}/user/${userId}`,
    method: 'DELETE',
    headers: {
      'content-type': 'application/json',
      Authorization: `Bearer ${authToken}`,
    },
  };

  const { data, error } = await requestAPI(options);
  if (data) console.log(data.response);
  if (error) console.error(error);
};

export const getMazesList = async () => {};

export const generateMaze = async (mazeInput: MazeInput) => {
  console.log('api_url', apiUrl);
  const options = {
    url: `${apiUrl}/maze/generate`,
    method: 'POST',
    data: mazeInput,
    headers: {
      'content-type': 'application/json',
    },
  };
  const { data, error } = await requestAPI(options);

  if (data) apiData.set(data.response);
  if (error) console.error(error);
};

export const saveMaze = async (authToken: string, userInfo: User, data: MazeSaveFormat) => {
  const userId = userInfo.sub!.split('|')[1];
  const options = {
    url: `${apiUrl}/maze/${userId}`,
    method: 'PUT',
    data: data,
    headers: {
      'content-type': 'application/json',
      Authorization: `Bearer ${authToken}`,
    },
  };
  const { error } = await requestAPI(options);
  if (error) console.error(error);
};

export const solveMaze = async () => {};

export const getScoreboard = async () => {
  const options = {
    url: `${apiUrl}/maze/scoreboard`,
    method: 'GET',
    headers: {
      'content-type': 'application/json',
    },
  };

  const { data, error } = await requestAPI(options);

  if (data) {
    const response = data.response;
    const scores: MazeScore[] = response.map((val: any) => {
      const mazeScore: MazeScore = {
        name: val.Name,
        score: val.Highscore,
      };
      return mazeScore;
    });

    mazeScores.set(scores);
  }

  if (error) {
    console.error(error);
  }
};
