<script lang="ts">
    import {
        manageConnection,
        usersList,
    } from "$lib/components/logic/manageWebSocketConn.svelte";
    import { user } from "$lib/components/logic/userName.svelte";
    import Chat from "$lib/components/ui/chatUI.svelte";
    import Activityindicator from "$lib/components/ui/activityindicator.svelte";
    import ChangeUserName from "$lib/components/ui/changeName.svelte"
    import UserNotFoundImage from "$lib/assets/user-not-found.png";

    let runActivityIndicator = $state<boolean>(true);
    let isOpen = $state(false);

    function toggleSidebar() {
        isOpen = !isOpen;
    }

    function closeSidebar() {
        isOpen = false;
    }

    let combinedIpAdresses = $derived(
        "Group_" + usersList.map((u) => u.ip).join("_"),
    );

    let selectedUser = $state<string>();
    let selectedIp = $state<string>();

    function selectUser(userParam: { name: string; ip: string }) {
        selectedUser = userParam.name;
        selectedIp = userParam.ip;
    }

    $effect(() => {
        if (usersList && usersList.length > 0) {
            runActivityIndicator = false;
        } else {
            runActivityIndicator = true;
        }
    });

    $effect(() => {
        if (!selectedUser && user.userName) {
            selectedUser = user.userName;
        }
    });

    manageConnection();
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
            {#if runActivityIndicator}
                <div id="activity-indicator">
                    <Activityindicator />
                </div>
            {:else}
                <button
                    class="nav-item"
                    class:active={selectedUser === "Group"}
                    onclick={() =>
                        selectUser({
                            name: "Group",
                            ip: combinedIpAdresses,
                        })}
                    title="Group"
                >
                    Group
                </button>
                {#each usersList as userEL}
                    <button
                        class="nav-item"
                        class:active={selectedUser === userEL.name}
                        onclick={() => selectUser(userEL)}
                        title={userEL.ip}
                    >
                        {userEL.name}
                    </button>
                {/each}
            {/if}
        </nav>
    </aside>

    <div class="chat-wrapper">
        {#if selectedUser && selectedIp}
            <Chat userName={selectedUser} {selectedIp} />
        {:else}
            <img
                id="user-not-found-img"
                src={UserNotFoundImage}
                alt="User not found"
            />
        {/if}
    </div>
</div>

<style>
    #activity-indicator {
        display: flex;
        justify-content: center;
        align-items: center;
        width: 100%;
        padding: 1.5rem 0;
    }

    #user-not-found-img {
        width: 650px;
        height: 650px;
        display: block;
        margin: 15% auto 0 auto;
    }

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
        border: none;
        outline: none;
        font-family: inherit;
        cursor: pointer;
        width: 100%;
        text-align: left;
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

    .chat-wrapper {
        flex: 1;
        height: 100vh;
        display: flex;
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

    @media (max-width: 765px) {
        #user-not-found-img {
            width: 350px;
            height: 350px;
        }

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
