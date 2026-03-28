<script lang="ts">
  import { projectState } from '$lib/stores/project.svelte';
  import {goto} from '$app/navigation';

  async function handleAddProject() {
    const projectName = prompt("Enter a name for the new project:");
    if (projectName) {
      try {
        await projectState.createProject({ Name: projectName });
        goto(`/project/${encodeURIComponent(projectName)}`);
      } catch (error) {
        alert(`Failed to create project: ${error.message}`);
      }
    }
  }
</script>

<aside class="sidebar">
  <h2>Projects</h2>
  <nav class="project-list">
    {#if projectState.data}
      {#each Object.values(projectState.data.projects) as project}
        <div class="project-button">
          <a href={`/project/${encodeURIComponent(project.Name)}`} class="project-link">{project.Name}</a>
          <button on:click={() => projectState.deleteProject(project.Name)}>Delete</button>
        </div>
      {/each}
    {:else}
      <p>Loading Projects...</p>
    {/if}
  </nav>
  <button on:click={handleAddProject}>Add Project</button>
</aside>

<style lang="scss">
.sidebar {
  display: flex;
  flex-direction: column;
  width: 256px;
  padding: 20px;
  background-color: #484848;
  z-index: 100;

  .project-button {
    margin-bottom: 10px;
  }

  .project-link {
    text-decoration: underline;
  }
}
</style>