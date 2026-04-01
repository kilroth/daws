<script lang="ts">
    import { configState } from "$lib/stores/config.svelte";
    import {projectState} from "$lib/stores/project.svelte";
    import { onMount } from "svelte";

    onMount(async () => {
        await configState.refresh();
        await projectState.refresh();
    });
</script>

{#if !projectState.data}
    <p>Loading projects...</p>
{:else if Object.keys(projectState.data.projects).length === 0}
    <p>No projects found. Use the sidebar to create your first project.</p>
{:else}
    <p>Welcome to your projects dashboard! Use the sidebar to navigate between your projects or create a new one.</p>
{/if}