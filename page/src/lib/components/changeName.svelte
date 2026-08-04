<script lang="ts">
    import { port, mDNSname } from "$lib/components/port.svelte";
    import { user } from "$lib/components/userName.svelte";

    let newName = $state(user.userName ?? "");

    $effect(() => {
        if (user.userName && !newName) {
            newName = user.userName;
        }
    });

    async function changeName() {
        try {
            if (newName && newName != user.userName) {
                const response = await fetch(
                    `http://${mDNSname}:${port}/submitnewusername`,
                    {
                        method: "PATCH",
                        headers: {
                            "Content-Type": "application/json",
                        },
                        body: JSON.stringify({ name: newName }),
                    },
                );

                if (!response.ok) {
                    throw new Error("Cannot change username");
                }
                user.userName = newName;
                alert("Successful name change");
            } else {
                alert(
                    "Please give a valid input if you want to cahnge your user name.",
                );
            }
        } catch (error) {
            alert("Error occured while trying to change the username " + error);
        }
    }
</script>

<div class="username-div">
    <input
        type="text"
        placeholder="No name typed here"
        defaultValue={user.userName}
        bind:value={newName}
    />
    <button onclick={changeName}>Change Usermame</button>
</div>

<style>
    .username-div {
        background-color: #ffffff;
        color: #2d2d3f;
        padding: 2%;
        border-radius: 25px;
        display: flex;
        flex-direction: column;
        justify-content: space-between;
        align-items: center;
        gap: 15px;
        margin-bottom: 20px;
    }

    .username-div > input {
        color: #2d2d3f;
        font-size: 25px;
        width: 55%;
        border-radius: 25px;
    }
    .username-div > button {
        border-radius: 25px;
        background-color: #002169;
        color: #ffffff;
        height: 50px;
        padding: 15px;
    }
</style>
