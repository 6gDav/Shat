<script lang="ts">
    import type { PageData } from './$types';


    let { data }: { data: PageData } = $props();
</script>
<h1>Hello in the Login page</h1>
<input type="text" placeholder="Enter your user name here"> <br>
<input type="text" disabled>
<p>Your IP adress <strong>{data.ip}</strong></p>