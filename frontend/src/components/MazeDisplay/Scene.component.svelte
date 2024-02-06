<script lang="ts">
  import { T } from '@threlte/core';
  import { OrbitControls } from '@threlte/extras';
  import { BoxGeometry } from 'three';
</script>

<T.PerspectiveCamera makeDefault position={[20, 30, 0]}>
  <OrbitControls enablePan={false} maxDistance={40} minDistance={10} maxPolarAngle={1.56} />
</T.PerspectiveCamera>
<T.DirectionalLight position={[3, 10, 7]} />
<T.AmbientLight />
<!-- Make a box in every second cell to show aligment -->
{#each { length: 10 } as _h, x}
  {#each { length: 10 } as _v, y}
    {#if x % 3 == 0 && y % 3 == 0}
      <T.Group position={[x - 4.5, 0.5, y - 4.5]}>
        <T.Mesh>
          <T.BoxGeometry />
          <T.MeshBasicMaterial
            args={[
              {
                color: '#ffffff',
                opacity: 0.9,
                transparent: true,
              },
            ]}
          />
        </T.Mesh>
        <T.LineSegments>
          <T.EdgesGeometry args={[new BoxGeometry()]} />
          <T.MeshBasicMaterial
            args={[
              {
                color: 0x000000,
              },
            ]}
          />
        </T.LineSegments>
      </T.Group>
    {/if}
  {/each}
{/each}
