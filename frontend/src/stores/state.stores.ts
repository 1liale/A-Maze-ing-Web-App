import { localStorageStore } from '@skeletonlabs/skeleton';
import { type Writable } from 'svelte/store';
import type { Player } from 'types/player.types';
import { SideBarState } from 'types/sidebar.types';

export const sidebarState: Writable<SideBarState> = localStorageStore('state', SideBarState.INIT);

export const defaultPlayer: Player = {
  mappedPos: undefined,
  relPos: undefined,
  moves: 0,
  hasWon: false,
};
export const playerState: Writable<Player> = localStorageStore('playerState', defaultPlayer);

export const solveTime: Writable<number> = localStorageStore('solveTime', 0);

const setupGame = () => {
  sidebarState.set(SideBarState.SETUP);
};

const readyGame = () => {
  sidebarState.set(SideBarState.WAITING)
}

const startGame = () => {
  sidebarState.set(SideBarState.STARTED);
};

const stopGame = () => {
  sidebarState.set(SideBarState.FINISHED);
};

const resetGame = () => {
  solveTime.set(0);
  playerState.set(defaultPlayer);
  sidebarState.set(SideBarState.INIT);
};

const showSolution = () => {
  sidebarState.set(SideBarState.SHOW_SOLUTION)
}

export { resetGame, setupGame, startGame, stopGame, showSolution, readyGame };
