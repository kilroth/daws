<script>
    let {} = $props();
</script>

<div class="project-card">
    <h3>Project Name</h3>
    <p>Created: 2024-01-01</p>
    <p>Last Modified: 2024-01-02</p>
    <button>Edit</button>
    <button>Delete</button>
</div>

<style lang="scss">
.project-card {
  background-color: #fff;
  border: 1px solid #ccc;
  border-radius: 8px;
  padding: 16px;
  margin-bottom: 16px;

  h3 {
    margin-top: 0;
    margin-bottom: 8px;
  }

  p {
    margin: 4px 0;
    color: #666;
  }

  button {
    margin-right: 8px;
    padding: 8px 12px;
    border: none;
    border-radius: 4px;
    cursor: pointer;

    &:first-of-type {
      background-color: #007bff;
      color: #fff;

      &:hover {
        background-color: #0056b3;
      }
    }

    &:last-of-type {
      background-color: #dc3545;
      color: #fff;

      &:hover {
        background-color: #c82333;
      }
    }
  }
}
</style>