<script lang="ts">
import { onMount } from 'svelte';
    
    let userIP = $state("Betöltés...");

    onMount(async () => {
        try {
            const res = await fetch('https://api.ipify.org?format=json');
            const data = await res.json();
            userIP = data.ip;
        } catch (e) {
            userIP = "Error occured";
        }
    });
</script>
<h1>Hello in the Login page</h1>
<input type="text" placeholder="Enter your user name here"> <br>
<input type="text" disabled>
<p>Your IP adress <strong>{userIP}</strong></p>