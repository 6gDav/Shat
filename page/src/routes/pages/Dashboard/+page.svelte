<script lang="ts">
    import { user } from "$lib/components/userName.svelte";
    import ChangeUserName from "$lib/components/changeName.svelte";
    import { onDestroy, onMount } from "svelte";
    import { port, mDNSname } from "$lib/components/port.svelte";

    let isOpen = $state(false);

    function toggleSidebar() {
        isOpen = !isOpen;
    }

    function closeSidebar() {
        isOpen = false;
    }

    interface WSResponse<T> {
        type: string;
        data: T;
    }

    let usersList = $state<string[]>([]);

    let chat: WebSocket;

    function startHandShake() {
        chat = new WebSocket(`ws://${mDNSname}:${port}/setws`);

        chat.onopen = () => {
            console.log("WebSocket is active.");
        };

        chat.onmessage = async (event) => {
            try {
                const textData = typeof event.data === "string" 
                ? event.data 
                : await event.data.text();

            const res = JSON.parse(textData);

                if (res.type === "USER_NAME_LIST") {
                    const payload = res as WSResponse<string[]>;
                    usersList = payload.data;
                }
            } catch (err) {
                alert("Cannot fetch user list.");
            }
        };

        chat.onerror = (error) => {
            console.error("WebSocket error:", error);
        };
    }

    onMount(async () => {
        //chat and name list
        $effect(() => {
            startHandShake();
        },)
    });

    onDestroy(() => {
        chat?.close();
    });
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
        <div class="sidebar-header">
            <button
                class="close-btn"
                onclick={closeSidebar}
                aria-label="Close menu">&times;</button
            >
        </div>
        <h1 id="title">Dashboard</h1>
        <ChangeUserName />

        <nav class="sidebar-nav">
            {#each usersList as userEL}
                <a href="#" class="nav-item" onclick={toggleSidebar}>
                    {userEL}
                </a>
            {/each}
        </nav>
    </aside>
    <main>{user.userName}</main>
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
        gap: 0.8rem;
        flex-grow: 1;
    }

    .nav-item {
        display: flex;
        align-items: center;
        gap: 0.75rem;
        padding: 0.75rem 1rem;
        color: #ffffff;
        text-decoration: none;
        border-radius: 8px;
        font-size: 0.95rem;
        transition: all 0.2s ease;
        background-color: #002169;
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
        margin-left: auto;
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
