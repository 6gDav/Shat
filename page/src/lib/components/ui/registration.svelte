<script lang="ts">
    import { goto } from "$app/navigation";
    
    import { port, mDNSname } from "$lib/components/logic/port.svelte";
    import { user } from "$lib/components/logic/userName.svelte";
    
    import ActivityIndicator from "$lib/components/ui/activityindicator.svelte";

    let nameText = $state<string>();
    let runActivityIndicator = $state<boolean>(false);

    async function registrateName() {
        runActivityIndicator = true;

        if (!nameText?.trim()) {
            alert("Please add a valid user name.");
            runActivityIndicator = false;
            return;
        }

        try {
            const response = await fetch(
                `http://${mDNSname}:${port}/user/username`,
                {
                    method: "POST",
                    headers: {
                        "Content-Type": "application/json",
                    },
                    body: JSON.stringify({ name: nameText }),
                },
            );

            if (!response.ok) {
                throw new Error("Submit user name is unsuccesfull.");
            }

            const responseData = await response.json();

            user.userName = nameText;
            user.initialIp = responseData.ip;

            console.log("Registration is succesfull");
            goto("/pages/Dashboard");
        } catch (error) {
            alert("Error occured why trying to send your name: " + error);
        } finally {
            runActivityIndicator = false;
        }
    }
</script>

<div class="container">
    <div class="form-box" class:loading={runActivityIndicator}>
        <h1>Registration page</h1>
        <input
            type="text"
            placeholder="Enter your user name here"
            bind:value={nameText}
            disabled={runActivityIndicator}
        />
        <button onclick={registrateName} disabled={runActivityIndicator}>
            Registration
        </button>
    </div>
</div>

{#if runActivityIndicator}
    <div class="overlay">
        <ActivityIndicator />
    </div>
{/if}

<style>
    .container {
        position: relative;
    }

    .form-box {
        display: flex;
        flex-direction: column;
        gap: 20px;
        align-items: stretch;
    }

    .form-box input,
    .form-box button {
        padding: 10px 14px;
        font-size: 16px;
        box-sizing: border-box;
        border-radius: 25px;
        border: 2px solid #e2e8f0;
        transition:
            border-color 0.3s ease,
            box-shadow 0.3s ease;
        border: 2px solid black;
    }

    .form-box input:hover,
    .form-box button:hover {
        box-shadow: 0px 10px 25px rgba(0, 0, 0, 0.519);
    }
    div.loading {
        filter: blur(4px);
        opacity: 0.6;
        pointer-events: none;
        transition:
            filter 0.3s ease,
            opacity 0.3s ease;
    }

    .overlay {
        position: absolute;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
        display: flex;
        justify-content: center;
        align-items: center;
        z-index: 10;
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
