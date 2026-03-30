<script lang="ts">
    import {projectState} from "$lib/stores/project.svelte";
    import { onMount } from "svelte";

    onMount(async () => {
        await projectState.refresh();
    });
</script>

{#if !projectState.data}
    <p>Loading projects...</p>
{:else if Object.keys(projectState.data.projects).length === 0}
    <p>No projects found. Use the sidebar to create your first project.</p>
{:else}

{/if}