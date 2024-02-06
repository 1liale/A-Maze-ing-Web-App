<script lang="ts">
  import { Canvas } from '@threlte/core';
  import { Grid } from '@threlte/extras';
  import Scene from './Scene.component.svelte';

  const plane = 'xz';
  const cellColor = '#cccccc';
  const cellSize = 1;
  const cellThickness = 1;
  const sectionColor = '#ff3e00';
  const sectionSize = 5;
  const sectionThickness = 2;
  const fadeDistance = 60;
  const fadeStrength = 1.5;
  const gridSize = 30;
  const followCamera = false;
  const infiniteGrid = false;
</script>

<Canvas>
  <Grid
    {plane}
    {cellColor}
    {cellSize}
    {cellThickness}
    {sectionColor}
    {sectionSize}
    {sectionThickness}
    {fadeDistance}
    {fadeStrength}
    gridSize={[gridSize, gridSize]}
    type="grid"
    {followCamera}
    {infiniteGrid}
  />
  <Scene />
</Canvas>
