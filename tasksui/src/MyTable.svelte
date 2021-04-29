<script>
  export let tasks;
  import { onMount } from 'svelte';

  import MyCell from './MyCell.svelte';
  
  const urlAPI = 'http://localhost:3000/api/todos/';
  const callAPI = async () => {
    await fetch(urlAPI)
    .then(async (response) => {
      let res = await response.json();
      tasks = res.todos.slice(0, 20);
    })
    .then(result => {
      console.log('Success:', result);
    })
    .catch(error => {
      console.error('Error:', error);
    });
  };

  onMount(callAPI);

  async function refreshTasks() {
    callAPI();
  }
</script>
<div class="column">
  <div class="columns is-vcentered">
    <div class="column">
      <h1 class="title is-size-4">List of tasks</h1>
    </div>
    <div class="column is-one-quarter">
      <button class="button is-primary" on:click={refreshTasks}>Refresh</button>
    </div>
  </div>
  <table class="table is-fullwidth is-hoverable is-striped">
    <thead>
      <tr>
        <th>Title</th>
        <th>Description</th>
      </tr>
    </thead>
    <tbody>
      {#each tasks as task}
        <MyCell task={task}></MyCell>
      {/each}
    </tbody>
  </table>
</div>
