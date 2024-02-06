import axios, { AxiosError, type AxiosRequestConfig } from 'axios';

export const requestAPI = async (options: AxiosRequestConfig) => {
  try {
    const { data } = await axios(options);
    return {
      data,
      error: null,
    };
  } catch (e) {
    if (axios.isAxiosError(e)) {
      return {
        response: null,
        error: (e as AxiosError).message,
      };
    }

    return {
      data: null,
      error: {
        message: (e as Error).message,
      },
    };
  }
};
