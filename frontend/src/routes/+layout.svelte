<script lang="ts">
    import "../style.scss";
    import Sidebar from "$lib/components/sidebar.svelte";

    let {children} = $props();
</script>

<div class="daws">
    <Sidebar />
    <main class="main-content">
        {@render children()}
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