<script>
  export let tasks;

  const urlAPI = 'http://localhost:3000/api/todos/'
  const getUrlAPI = 'http://localhost:3000/api/todos/';
  const callAPI = async () => {
    await fetch(getUrlAPI)
    .then(async response => {
      let res = await response.json();
      console.log(res);
      tasks = res.todos.slice(0, 20);
    })
    .then(result => {
      console.log('Success:', result);
    })
    .catch(error => {
      console.error('Error:', error);
    });
  };

  var newtitle = '';
  var newdescription = '';
  var errorTitle = false;
  var errorDescription = false;

  async function handleSubmit() {
    errorTitle = newtitle == "";
    errorDescription = newdescription == "";
    if (!errorTitle && !errorDescription) {
      let newTask = {
        "title": newtitle,
        "description": newdescription
      }
      let jsonVal = JSON.stringify(newTask);
      await fetch(urlAPI, {
        method: 'POST',
        headers: {
          'Accept': 'application/json',
          'Content-Type': 'application/json'
        },
        body: jsonVal
      })
      .then(response => console.log(response.json()))
      .then(result => {
        callAPI()
        newtitle = "";
        newdescription = "";
        console.log('Success:', result);
      })
      .catch(error => {
        console.error('Error:', error);
      });
    }
  }
</script>

<div class="column is-one-quarter">
  <form id="studentsForm">
    <div class="field">
      <label class="label">Title</label>
      <div class="control">
        <input class="input {errorTitle ? 'is-danger' : ''}" type="text" name="title" placeholder="Title of task" bind:value={newtitle}>
      </div>
      {#if errorDescription }
        <p class="help is-danger">This title is empty</p>
      {/if}
    </div>
    <div class="field">
      <label class="label">Description</label>
      <div class="control">
        <input class="input {errorDescription ? 'is-danger' : ''}" type="text" name="description" placeholder="Task description" bind:value={newdescription}>
      </div>
      {#if errorDescription }
        <p class="help is-danger">This description is empty</p>
      {/if}
    </div>
    <div class="field">
      <div class="control">
        <button type="button" class="button is-primary" on:click={handleSubmit}>Submit</button>
      </div>
    </div>
  </form>
</div>
