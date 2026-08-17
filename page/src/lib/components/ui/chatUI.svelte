<script lang="ts">
    import { messages, sendChatMessage } from "../manageWebSocketConn.svelte";
    import { user } from "../userName.svelte";

    let { userName: selectedUserName, selectedIp } = $props<{
        userName: string;
        selectedIp: string;
    }>();

    let messageText = $state<string>("");
    let chatContainer = $state<HTMLDivElement>();
    //get chat
    let isGroupView = $derived(selectedUserName === "Group");

    let currentChatMessages = $derived(
        messages.filter((msg) => {
            if (isGroupView) {
                return (
                    msg.type === "group" && selectedUserName.startsWith("Group")
                );
            }
            if (msg.type === "private") {
                const expectedRoomId = [user.initialIp, selectedIp]
                    .sort()
                    .join("_");
                return msg.roomId === expectedRoomId;
            }

            return false;
        }),
    );

    //send chat
    function sendMessageText() {
        if (!messageText?.trim()) {
            alert("Write something if you want to send a message...");
            return;
        }

        sendChatMessage(
            user.initialIp,
            selectedIp,
            messageText.trim(),
            selectedUserName === "Group" ? "group" : "private",
            user.userName,
        );
        messageText = "";
    }

    function handleKeyDown(e: KeyboardEvent) {
        if (e.key === "Enter") {
            sendMessageText();
        }
    }

    $effect(() => {
        if (currentChatMessages.length && chatContainer) {
            chatContainer.scrollTop = chatContainer.scrollHeight;
        }
    });
</script>

<main class="chat-main">
    <div class="title-container">
        <h1 id="name-title">{selectedUserName}</h1>
        <h2 id="ip-title">{selectedIp}</h2>
    </div>

    <div class="messages-list" bind:this={chatContainer}>
        {#each currentChatMessages as msg}
            <div
                class="message-bubble"
                class:sent={msg.from === user.initialIp}
                class:received={msg.from !== user.initialIp}
            >
                <div class="message-text">
                    <strong class="sender-name">{msg.userName}:</strong>
                    {msg.text}
                </div>
            </div>
        {/each}
    </div>

    <div class="input-container">
        <input
            type="text"
            placeholder="Write something..."
            bind:value={messageText}
            onkeydown={handleKeyDown}
        />
        <button onclick={sendMessageText}>Send</button>
    </div>
</main>

<style>
    #name-title {
        font-size: 52px;
        margin: 0;
        color: #ffffff;
    }

    #ip-title {
        font-size: 25px;
        margin: 0;
        color: #ffffff;
        text-shadow: 4px 4px 6px rgba(0, 0, 0, 0.5);
        font-style: italic;
    }

    .chat-main {
        display: flex;
        flex-direction: column;
        height: 100vh;
        width: 100%;
        padding: 2rem;
        box-sizing: border-box;
        overflow: hidden;
    }

    .title-container {
        text-align: center;
        margin-bottom: 2rem;
        flex-shrink: 0;
    }

    .messages-list {
        display: flex;
        flex-direction: column;
        gap: 0.75rem;
        flex: 1;
        overflow-y: auto;
        padding-right: 0.5rem;
    }

    .message-bubble {
        max-width: 60%;
        border-radius: 15px;
        padding: 0.75rem 1rem;
        word-break: break-word;
        color: white;
    }

    .sent {
        align-self: flex-end;
        background-color: #6366f1;
        border-bottom-right-radius: 2px;
    }

    .received {
        align-self: flex-start;
        background-color: #374151;
        border-bottom-left-radius: 2px;
    }

    .input-container {
        width: 100%;
        padding-top: 1rem;
        flex-shrink: 0;
        display: flex;
        gap: 0.75rem;
        align-items: center;
    }

    .sender-name {
        display: inline-block;
        font-size: 0.85rem;
        opacity: 0.8;
        margin-right: 0.25rem;
    }

    input {
        flex: 1;
        padding: 0.75rem 1rem;
        border-radius: 8px;
        border: 1px solid rgba(255, 255, 255, 0.15);
        color: #000000;
        font-size: 0.95rem;
        font-family: inherit;
        outline: none;
        transition: all 0.2s ease;
        box-sizing: border-box;
    }

    button {
        display: flex;
        align-items: center;
        justify-content: center;
        gap: 0.75rem;
        padding: 0.75rem 1.5rem;
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
        white-space: nowrap;
    }

    button:hover {
        box-shadow: 0px 10px 25px rgba(0, 0, 0, 0.519);
    }

    button:active {
        transform: translateY(0);
    }

    @media (max-width: 765px) { 
        input {
            width: 65%;
        }
    }
</style>
