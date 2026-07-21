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
    <div class="container">
        <div class:loading={runActivityIndicator}>
            <h1>Hello in the Login page</h1>
            <input
                type="text"
                placeholder="Enter your user name here"
                bind:value={nameText}
                disabled={runActivityIndicator}
            />
            <br /> <br />
            <button onclick={registrateName} disabled={runActivityIndicator}
                >Registrate</button
            >
        </div>

        {#if runActivityIndicator}
            <div class="overlay">
                <ActivityIndicator />
            </div>
        {/if}
    </div>
</main>

<style>
    .container {
        position: relative;
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
