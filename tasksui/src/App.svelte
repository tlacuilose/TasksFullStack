<script>
  import MyForm from './MyForm.svelte';
  import MyTable from './MyTable.svelte';

  let tasks = [];
</script>

<main>
  <nav class="navbar is-primary" role="navigation" aria-label="main navigation">
    <div class="navbar-brand">
      <a class="navbar-item" href="https://bulma.io">
        <h1>TasksUI</h1>
      </a>
    </div>
  </nav>
  <div class="section">
    <div class="container">
      <div class="columns">
        <MyForm bind:tasks={tasks}></MyForm>
        <MyTable bind:tasks={tasks}></MyTable>
      </div>
    </div>
  </div>
</main>
<svelte:head>
  <title>TasksUI</title>
  <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bulma@0.9.2/css/bulma.min.css">
</svelte:head>

