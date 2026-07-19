<script lang="ts">
    
    let nameText = $state<string>()

    async function registrateName() {

    }
</script>

<h1>Hello in the Login page</h1>
<input type="text" placeholder="Enter your user name here" bind:value={nameText}> 
<br> <br>
<button onclick={registrateName}>Registrate</button>
