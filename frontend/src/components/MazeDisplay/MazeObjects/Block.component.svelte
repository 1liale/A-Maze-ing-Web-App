<script lang="ts">
  import { modeCurrent } from '@skeletonlabs/skeleton';
  import { T } from '@threlte/core';
  import { BoxGeometry } from 'three';
  export let position: number | [x: number, y: number, z: number] | undefined;
  export let height: number;
  export let color: string | undefined = undefined;
  const thickness = 0.99;
</script>

<T.Group {position}>
  <T.Mesh>
    <T.BoxGeometry args={[thickness, height, thickness]} />
    <T.MeshBasicMaterial
      args={[
        {
          color: color ? color : $modeCurrent ? 'grey' : 'lightgrey',
          opacity: 0.75,
          transparent: true,
        },
      ]}
    />
  </T.Mesh>
  <T.LineSegments>
    <T.EdgesGeometry args={[new BoxGeometry(thickness, height, thickness)]} />
    <T.MeshBasicMaterial
      args={[
        {
          color: '#000000',
        },
      ]}
    />
  </T.LineSegments>
</T.Group>
