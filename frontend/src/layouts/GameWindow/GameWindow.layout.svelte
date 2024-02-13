<script lang="ts">
  import InstructionBox from '@components/InstructionBox/InstructionBox.component.svelte';
  import MazeDisplay from '@components/MazeDisplay/MazeDisplay.component.svelte';
</script>

<div class="container h-full flex flex-col">
  <div class="h-5/6">
    <MazeDisplay />
  </div>
  <InstructionBox />
</div>
