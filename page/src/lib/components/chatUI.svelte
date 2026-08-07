<script lang="ts">
    let { userName } = $props();
    let messageText = $state<string>();

    function sendMessageText() {
        if (!messageText?.trim()) {
            alert("Write something if you want to send a message...");
        } else {
            console.log("Send message");
            messageText = "";
        }
    }
</script>

<main class="chat-main">
    <div class="title-container">
        <h1 id="name-title">{userName}</h1>
    </div>

    <div class="messages-list">
        <div class="message-bubble">Hi</div>
        <div class="message-bubble">Hi</div>
        <div class="message-bubble">Hi</div>
        <div class="message-bubble">Hi</div>
        <div class="message-bubble">Hi</div>
        <div class="message-bubble">Hi</div>
        <div class="message-bubble">Hi</div>
        <div class="message-bubble">Hi</div>
        <div class="message-bubble">Hi</div>
        <div class="message-bubble">Hi</div>
        <div class="message-bubble">Hi</div>
        <div class="message-bubble">Hi</div>
        <div class="message-bubble">Hi</div>
        <div class="message-bubble">Hi</div>
        <div class="message-bubble">Hi</div>
        <div class="message-bubble">Hi</div>
        <div class="message-bubble">Hi</div>
        <div class="message-bubble">Hi</div>
        <div class="message-bubble">Hi</div>
        <div class="message-bubble">Hi</div>
        <div class="message-bubble">Hi</div>
        <div class="message-bubble">Hi</div>
        <div class="message-bubble">Hi</div>
        <div class="message-bubble">Hi</div>
        <div class="message-bubble">Hi</div>
        <div class="message-bubble">Hi</div>
        <div class="message-bubble">Hi</div>
        <div class="message-bubble">Hi</div>
        <div class="message-bubble">Hi</div>
        <div class="message-bubble">Hi</div>
        <div class="message-bubble">Hi</div>
        <div class="message-bubble">Hi</div>
        <div class="message-bubble">Hi</div>
        <div class="message-bubble">Hi</div>
        <div class="message-bubble">Hi</div>
        <div class="message-bubble">Hi</div>
        <div class="message-bubble">Hi</div>
        <div class="message-bubble">Hi</div>
        <div class="message-bubble">Hi</div>
        <div class="message-bubble">Hi</div>
        <div class="message-bubble">Hi</div>
        <div class="message-bubble">Hi</div>
        <div class="message-bubble">Hi</div>
    </div>

    <div class="input-container">
        <input
            type="text"
            placeholder="Write something..."
            bind:value={messageText}
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
        gap: 1rem;
        flex: 1;
        overflow-y: auto;
        padding-right: 0.5rem;
    }

    .input-container {
        width: 100%;
        padding-top: 1rem;
        flex-shrink: 0;
        display: flex;
        gap: 0.75rem;
        align-items: center;
    }

    .message-bubble {
        width: 25%;
        border-radius: 15px;
        background-color: #6366f1;
        padding: 0.5rem 1rem;
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
