<script lang="ts">
    import { configState } from "$lib/stores/config.svelte";
    import { projectState } from "$lib/stores/project.svelte";
    import { onMount } from "svelte";

    // get all projects (not implemented)

    onMount(async () => {
        // load projects from backend (not implemented)
        await configState.refresh();
        await projectState.refresh();
    });
</script>

<!-- Projects page content -->
<div class="projects-page">
    <h1>Projects</h1>
    <!--<p>Here you can view all your projects.</p>-->
    <p>No projects available.</p>
</div>