<script lang="ts">
  import InstructionGroup from '@components/InstructionGroup/InstructionGroup.component.svelte';
  import MazeDisplay from '@components/MazeDisplay/MazeDisplay.component.svelte';
</script>

<div class="container h-full flex flex-col">
  <div class="h-5/6">
    <MazeDisplay />
  </div>
  <InstructionGroup />
</div>
