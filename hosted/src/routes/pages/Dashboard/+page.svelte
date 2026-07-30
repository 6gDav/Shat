<script lang="ts">
    import { port } from "$lib/components/port.svelte";
    import { user } from "$lib/components/userName.svelte";

    let isOpen = $state(false);
    let newName = $state(user.userName ?? "");

    $effect(() => {
        if (user.userName && !newName) {
            newName = user.userName;
        }
    });

    function toggleSidebar() {
        isOpen = !isOpen;
    }

    function closeSidebar() {
        isOpen = false;
    }

    async function changeName() {
        try {
            if (newName && newName != user.userName) {
                const response = await fetch(
                    `http://loginpage.local:${port}/submitnewusername`,
                    {
                        method: "POST",
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
                alert("Succresfull name change");
            } else {
                alert(
                    "Please give a valid input if you want to cahnge your user name.",
                );
            }
        } catch (error) {
            alert("Error occured while trying to change the username " + error)
        }
    }
</script>

<div class="layout-container">
    <button
        class="hamburger"
        onclick={toggleSidebar}
        aria-label="Toggle navigation"
    >
        <span class="bar"></span>
        <span class="bar"></span>
        <span class="bar"></span>
    </button>

    {#if isOpen}
        <div class="overlay" onclick={toggleSidebar}></div>
    {/if}

    <aside class="sidebar" class:open={isOpen}>
        <h1 id="title">Dashboard</h1>
        <div class="username-div">
            <input
                type="text"
                placeholder="No name typed here"
                defaultValue={user.userName}
                bind:value={newName}
            />
            <button onclick={changeName}>Change Usermame</button>
        </div>
        <div class="sidebar-header">
            <button
                class="close-btn"
                onclick={closeSidebar}
                aria-label="Close menu">&times;</button
            >
        </div>

        <nav class="sidebar-nav">
            <a href="#" class="nav-item active" onclick={toggleSidebar}>User</a>
        </nav>
    </aside>
</div>

<style>
    .layout-container {
        display: flex;
        width: 100vw;
        height: 100vh;
        overflow: hidden;
        position: relative;
    }

    .sidebar {
        width: 350px;
        height: 100vh;
        background: linear-gradient(to bottom, #002169 20%, #ffffff 80%);
        color: #ffffff;
        display: flex;
        flex-direction: column;
        padding: 1.5rem 1rem;
        box-sizing: border-box;
        transition: transform 0.3s ease-in-out;
        z-index: 100;
    }

    .sidebar-header {
        padding-bottom: 1.5rem;
        margin-bottom: 1rem;
        border-bottom: 1px solid #2d2d3f;
        display: flex;
        justify-content: space-between;
        align-items: center;
    }

    .sidebar-header h2 {
        margin: 0;
        font-size: 1.25rem;
        font-weight: 600;
        letter-spacing: 0.5px;
    }

    .sidebar-nav {
        display: flex;
        flex-direction: column;
        gap: 0.5rem;
        flex-grow: 1;
    }

    .nav-item {
        display: flex;
        align-items: center;
        gap: 0.75rem;
        padding: 0.75rem 1rem;
        color: #a2a3b7;
        text-decoration: none;
        border-radius: 8px;
        font-size: 0.95rem;
        transition: all 0.2s ease;
    }

    .nav-item:hover {
        box-shadow: 0px 10px 25px rgba(0, 0, 0, 0.519);
    }

    .nav-item.active {
        background-color: #6366f1;
        color: #ffffff;
        font-weight: 500;
    }

    .hamburger {
        display: none;
        position: fixed;
        top: 1rem;
        left: 1rem;
        z-index: 90;
        background: #002169;
        border: none;
        padding: 0.6rem;
        border-radius: 6px;
        cursor: pointer;
        flex-direction: column;
        gap: 4px;
    }

    .hamburger .bar {
        width: 22px;
        height: 2px;
        background-color: #ffffff;
        border-radius: 2px;
    }

    .close-btn {
        display: none;
        background: none;
        border: none;
        color: #ffffff;
        font-size: 1.75rem;
        cursor: pointer;
        line-height: 1;
    }

    .overlay {
        display: none;
    }

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

    #title {
        font-size: 50px;
        display: flex;
        justify-content: center;
    }

    @media (max-width: 765px) {
        .hamburger {
            display: flex;
        }

        .close-btn {
            display: block;
        }

        .sidebar {
            position: fixed;
            top: 0;
            left: 0;
            transform: translateX(-100%);
            box-shadow: 2px 0 10px rgba(0, 0, 0, 0.3);
        }

        .sidebar.open {
            transform: translateX(0);
        }

        .overlay {
            display: block;
            position: fixed;
            inset: 0;
            background: rgba(0, 0, 0, 0.5);
            z-index: 95;
        }
    }
</style>
