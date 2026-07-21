<script lang="ts">
    import { Spinner } from "$lib/components/ui/spinner/index.js";
    import "../app.css";

    let nameText = $state<string>();

    async function registrateName() {
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
        }
    }
</script>

<h1>Hello in the Login page</h1>
<input
    type="text"
    placeholder="Enter your user name here"
    bind:value={nameText}
/>
<br /> <br />
<button onclick={registrateName}>Registrate</button>

<div class="loader"></div>

<style>
    .loader {
        border: 16px solid #f3f3f3; /* Light grey */
        border-top: 16px solid #3498db; /* Blue */
        border-radius: 50%;
        width: 120px;
        height: 120px;
        animation: spin 2s linear infinite;
    }

    @keyframes spin {
        0% {
            transform: rotate(0deg);
        }
        100% {
            transform: rotate(360deg);
        }
    }
</style>
