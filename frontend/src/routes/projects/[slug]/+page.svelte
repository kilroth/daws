<script lang="ts">
    import { page } from "$app/state";
    import { projectState } from "$lib/stores/project.svelte";
    import { onMount } from "svelte";

    const slug = decodeURIComponent(page.params.slug);
    let project = $derived(projectState.data?.projects[slug]);

    onMount(async () => {
        await projectState.refresh();
        console.log($state.snapshot(projectState.data));
    });
</script>

{#if !projectState.data}
    <p>Loading project...</p>
{:else if !project}
    <p>Project not found.</p>
{:else}
    <h1>{project.name}</h1>
    <p>{project.description}</p>
{/if}