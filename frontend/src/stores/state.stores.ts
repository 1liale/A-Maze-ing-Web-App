import { writable, type Writable } from 'svelte/store';
import type { Player } from 'types/player.types';
import type { SideBarState } from 'types/sidebar.types';

export const sidebarState: Writable<SideBarState> = writable('init');

export const playerState: Writable<Player | undefined> = writable(undefined);
