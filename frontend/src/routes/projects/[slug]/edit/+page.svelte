<script lang="ts">
    import { page } from "$app/state";
    import { projectState } from "$lib/stores/project.svelte";
    import { onMount } from "svelte";

    const slug = decodeURIComponent(page.params.slug);
    let project = $derived(projectState.data?.projects[slug]);

    onMount(async () => {
        await projectState.refresh();
    });
</script>

{#if !projectState.data}
    <p>Loading project...</p>
{:else if !project}
    <p>Project not found.</p>
{:else}
    <h1>Edit {project.name}</h1>
    <p>{project.description}</p>
{/if}