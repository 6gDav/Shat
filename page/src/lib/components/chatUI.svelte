<script lang="ts">
    import { messages, sendChatMessage } from "./manageWebSocketConn.svelte"; 

    let { userName, selectedIp } = $props<{ userName: string; selectedIp: string }>();
    let messageText = $state<string>("");
    let chatContainer = $state<HTMLDivElement>();

    function sendMessageText() {
        if (!messageText?.trim()) {
            alert("Write something if you want to send a message...");
            return;
        }
        sendChatMessage(selectedIp, messageText.trim());
        messageText = "";
    }

    function handleKeyDown(e: KeyboardEvent) {
        if (e.key === "Enter") {
            sendMessageText();
        }
    }

    $effect(() => {
        if (messages.length && chatContainer) {
            chatContainer.scrollTop = chatContainer.scrollHeight;
        }
    });
</script>

<main class="chat-main">
    <div class="title-container">
        <h1 id="name-title">{userName} <span><i>{selectedIp}</i></span></h1>
    </div>

    <div class="messages-list" bind:this={chatContainer}>
        {#each messages as msg}
            <div 
                class="message-bubble" 
                class:sent={msg.from !== selectedIp} 
                class:received={msg.from === selectedIp}
            >
                <div class="message-text">{msg.text}</div>
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

    #name-title {
        font-size: 52px;
        margin: 0;
        color: #ffffff;
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
</style>