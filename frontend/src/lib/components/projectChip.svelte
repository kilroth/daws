<script lang="ts">
    import { goto } from "$app/navigation";
    import { projectState } from "$lib/stores/project.svelte";
    import { onMount } from "svelte";

    let {data} = $props();

    let updated = $derived.by(() => {
      console.log("Updated timestamp:", data.timeData.updatedAt); // Debug log
      let tmpUpdated = new Date(data.timeData.updatedAt);
      return tmpUpdated.toLocaleDateString(undefined, { month: 'numeric', day: 'numeric', year: 'numeric' });
    });

    let deployed = $derived.by(() => {
      console.log("Deployed timestamp:", data.timeData.deployedAt); // Debug log
      let tmpDeployed = new Date(data.timeData.deployedAt);
      if (isNaN(tmpDeployed.getTime())) {
          return "Never";
      }
      return tmpDeployed.toLocaleDateString(undefined, { month: 'numeric', day: 'numeric', year: 'numeric' });
    });

    onMount(() => {
      projectState.refresh();
    });

    const handleDelete = (e: Event) => {
        e.preventDefault();
        const confirmed = confirm(`Are you sure you want to delete project "${data.name}"? This action cannot be undone.`);
        if (confirmed) {
          try {
            // Call the delete function from the project store
            // You may want to add error handling and feedback to the user here
            projectState.deleteProject(data.name);
            goto('/projects'); // Redirect to projects list after deletion
          } catch (error) {
            console.error("Failed to delete project:", error);
            alert("An error occurred while trying to delete the project. Please try again.");
          }
        }
    };
</script>

<div class="project-card__wrapper">
  <a class="project-card" href={`/projects/${data.slug}`}>
    <div class="name">
      <h3>{data.name}</h3>
      <div class="row">
        <p>Updated: {updated}</p>
        <p>Deployed: {deployed}</p>
      </div>
    </div>
  </a>
  <div class="project-card__btn-group">
    <button class="btn btn-delete" onclick={handleDelete}>&#x2716;</button>
  </div>
</div>

<style lang="scss">
.project-card {
  display: flex;
  flex: 1;
  text-decoration: none;
  background-color: $clr_dark;
  border: 1px solid transparent;
  border-radius: $border-radius 0 0 $border-radius;
  cursor: pointer;
  transition: all 0.2s ease;

  &__wrapper {
    display: flex;
    justify-content: space-between;
    align-items: center;
    width: 100%;
    gap: 0;
  }

  &:hover {
    background-color: $clr_dark_hover;
  }

  .name {
    padding: $padding;
    margin: $padding;
    color: $clr_light;
    h3 {
      font-size: 24px;
      margin-top: 0;
      margin-bottom: $padding;
      border-bottom: 1px solid $clr_secondary;
      padding-bottom: $padding;
    }
  
    p {
      margin: 0;
      padding: 0;
      font-size: 10px;
    }
  }

  &__btn-group {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    height: -webkit-fill-available;

    .btn {
      border-radius: 0 $border-radius $border-radius 0;
      cursor: pointer;
      height: -webkit-fill-available;
      
      &-delete {
        background-color: $clr_error;
        font-size: 18px;
        padding: 0 calc($padding * 1.5);
        height: 100%;
        margin-right: -1px;
        border: 1px solid $clr_dark;
        box-shadow: inset 0 0 4px 2px rgba($clr_dark, 0.5);
        transition: all 0.2s ease;
        
        &:hover {
          background-color: $clr_error_hover;
          border: 1px solid transparent;
          box-shadow: none;
          color: $clr_light;
        }
      }
    }
  }


}
</style>