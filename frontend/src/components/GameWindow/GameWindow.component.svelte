<script lang="ts">
  import { Canvas, T } from '@threlte/core';
  import type { MazeConfig, MazeData } from 'types/maze.types';

  export let className: string = '';
  export let data: MazeData | undefined = undefined;
  export let config: MazeConfig | undefined = undefined;
</script>

<div class={className}>
  <Canvas>
    <T.Camera></T.Camera>
    <T.Mesh>
      <T.BoxGeometry />
      <T.MeshBasicMaterial />
    </T.Mesh>
  </Canvas>
</div>
