<script lang="ts">
  import { projectState } from '$lib/stores/project.svelte';
  import {goto} from '$app/navigation';
    import Button from './basic/button.svelte';
    import ProjectChip from './projectChip.svelte';

  async function handleAddProject() {
    goto('/projects/new');
  }
  
  const handleConfigClick = () => {
    goto('/config');
  }
</script>

<aside class="sidebar">
  <div class="header">
    <h2 class="header__title">Projects</h2>
    <Button subtype="icon" onClick={handleConfigClick}>&#x2699;</Button>
  </div>
  <div class="list">
    <nav class="project-list">
      {#if !projectState.data}
        <p>Loading Projects...</p>
      {:else if Object.keys(projectState.data.projects).length === 0}
        <p>No projects found. Use the button below to create your first project.</p>
      {:else}
        {#each Object.values(projectState.data.projects) as project}
          <ProjectChip data={project} />
        {/each}
      {/if}
    </nav>
    <Button subtype="add" onClick={handleAddProject}>Add Project</Button>
  </div>
</aside>

<style lang="scss">
.sidebar {
  display: flex;
  flex-direction: column;
  width: 320px;
  background-color: #484848;
  z-index: 100;

  .header {
    background-color: #282828;
    color: #ccc;
    padding: calc($padding * 2);
    margin: 0;
    display: flex;
    justify-content: space-between;
    align-items: center;

    &__title {
      margin: 0;
      font-size: 24px;
      font-weight: bold;
    }
  }

  .project-button {
    margin-bottom: 10px;
  }

  .project-link {
    text-decoration: underline;
  }

  .list {
    display: flex;
    flex-direction: column;
    gap: $padding;
    flex: 1;
    padding: $padding calc($padding * 2);
  }

  .bottom {
    background-color: #282828;
    color: #ccc;
    padding: $padding;
    justify-self: flex-end;
    font-size: 32px;
    width: calc(100% - ($padding * 2));
    display: flex;

    .config {
      justify-self: flex-end;
    }
  }
}
</style>