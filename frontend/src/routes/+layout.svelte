<script lang="ts">
    import "../style.css";
    import {onMount} from "svelte";
    import {projectState} from "$lib/stores/project.svelte";
    import Sidebar from "$lib/components/sidebar.svelte";

    let {children} = $props();

    onMount(async () => {
        await projectState.refresh();
    });
</script>

<div class="daws">
    <Sidebar />
    <main class="main-content">
        {#if !projectState.data}
            <p>Loading projects...</p>
        {:else if Object.keys(projectState.data.projects).length === 0}
            <p>No projects found. Use the sidebar to create your first project.</p>
        {:else}
             {@render children()}
        {/if}
    </main>
</div>

<style lang="scss">
.daws {
  display: flex;
  height: 100vh;

  .main-content {
    flex: 1;
    padding: 20px;
    overflow-y: auto;
  }
}
</style>