<script lang="ts">
  import { solveTime, sidebarState as state } from '@stores/state.stores';
  import { SideBarState } from 'types/sidebar.types';
  const interval = 1000; // Interval in milliseconds

  let timerInterval: any;

  $: elapsedSeconds = $solveTime;

  const startTimer = () => {
    const timerInterval = setInterval(() => {
      elapsedSeconds += 1;
    }, interval);

    return timerInterval;
  };

  const stopTimer = (timerInterval: string | number | NodeJS.Timeout | undefined) => {
    clearInterval(timerInterval);
  };

  state.subscribe((val: SideBarState) => {
    if (timerInterval) {
      stopTimer(timerInterval);
      $solveTime = elapsedSeconds;
    }
    if (val !== SideBarState.FINISHED) {
      $solveTime = 0;
      if (val === SideBarState.STARTED) {
        timerInterval = startTimer();
      }
    }
  });
</script>

<div>
  <p><strong>Time</strong>: {elapsedSeconds} seconds</p>
</div>
