<script lang="ts">
    import "../app.css";
    import ActivityIndicator from "./components/activityindicator.svelte";

    let nameText = $state<string>();
    let runActivityIndicator = $state<boolean>(false);

    async function registrateName() {
        runActivityIndicator = true;

        try {
            const response = await fetch("http://loginpage.local:3000/submit", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: nameText,
            });

            if (!response.ok) {
                throw new Error("POST resoult is not Ok");
            } else {
                console.log("Registration is succesfull");
                //The registration is succes full. We will navigate to a nother file
                //Navigation logic here
            }
        } catch (error) {
            console.error(
                "Error occured why trying to send your name: " + error,
            );
        } finally {
            runActivityIndicator = false;
        }
    }
</script>

<main>
    <div class="container login-container">
        <div class="container">
            <div class="form-box" class:loading={runActivityIndicator}>
                <h1>Registration page</h1>
                <input
                    type="text"
                    placeholder="Enter your user name here"
                    bind:value={nameText}
                    disabled={runActivityIndicator}
                />
                <button
                    onclick={registrateName}
                    disabled={runActivityIndicator}
                >
                    Registration
                </button>
            </div>
        </div>

        {#if runActivityIndicator}
            <div class="overlay">
                <ActivityIndicator />
            </div>
        {/if}
    </div>
</main>

<style>
    :global(body) {
        display: flex;
        justify-content: center;
        align-items: center;
        min-height: 100vh;
        margin: 0;
        background-color: #485e82;
    }

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
    }

    .form-box input:hover,
    .form-box button:hover {
        border-color: #485e82;
    }

    .login-container {
        background: linear-gradient(to bottom, #002169 20%, #ffffff 80%);
        padding: 100px;
        border-radius: 15px;
        box-shadow: 0 10px 25px rgba(0, 0, 0, 0.2);
        color: white;
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
</style>
