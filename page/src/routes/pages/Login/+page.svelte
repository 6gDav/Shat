<script lang="ts">
    import redirectUser from "$lib/components/logic/redirectuser.svelte";
    import Registration from "$lib/components/ui/registration.svelte";
    import { onMount } from "svelte";
    
    
    onMount(() => {
        redirectUser()
    })
</script>

<div class="container login-container">
    <Registration />
</div>

<style>
    :global(html, body) {
        height: 100%;
        margin: 0;
    }

    :global(body) {
        display: flex;
        justify-content: center;
        align-items: center;
        min-height: 100vh;
        background-color: #485e82;
    }

    .container {
        position: relative;
    }

    .login-container {
        background: linear-gradient(to bottom, #002169 20%, #ffffff 80%);
        padding: 100px;
        border-radius: 15px;
        box-shadow: 0 10px 25px rgba(0, 0, 0, 0.2);
        color: white;
    }

    @media (min-width: 768px) {
        .container {
            transform: scale(1.3);
        }
    }

    @media (max-width: 768px) {
        .container {
            transform: scale(0.9);
        }
    }
</style>
