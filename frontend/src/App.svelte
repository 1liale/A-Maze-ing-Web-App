<script>
  import { onMount } from 'svelte';
  import { createScene } from "./scene";
  let el;
  onMount(() => {
    createScene(el)
  });
</script>

<canvas bind:this={el}></canvas>

<h1 class="text-3xl font-bold underline">
  A-Maze-ing: Generate and Solve Mazes :) 
</h1>

<style>
</style>
