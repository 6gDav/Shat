<script lang="ts">
    import { goto } from "$app/navigation";
    import { onMount } from "svelte";
    import { port } from "$lib/components/port.svelte";
    import { user } from "$lib/components/userName.svelte";

    import Registration from "$lib/components/registration.svelte";

    onMount(async () => {
        try {
            let response = await fetch(
                `http://loginpage.local:${port}/redirectuser`,
            );
            let shouldRedirect = await response.json();

            if (shouldRedirect.redirect && shouldRedirect.userName) {
                user.userName = shouldRedirect.userName;
                goto("/pages/Dashboard");
            }
        } catch (error) {
            console.error("Error occured while trying to redirect.", error);
        }
    });
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
