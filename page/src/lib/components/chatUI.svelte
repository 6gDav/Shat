<script lang="ts">
    let { userName } = $props();
    let messageText = $state<string>();

    function sendMessageText() {
        if (!messageText?.trim()) {
            alert("Write something if you want to send a message...")
        } else {
            console.log("Send message")
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
        <input type="text" placeholder="Write something..." bind:value={messageText}/>
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
    }

    .message-bubble {
        width: 25%;
        border-radius: 15px;
        background-color: #6366f1;
        padding: 0.5rem 1rem; 
    }

    input {
        width: 80%;
        padding: 12px 16px;
        border-radius: 8px;
        border: 1px solid #ccc;
        box-sizing: border-box;
    }
    button {
        background-color: #6366f1;
        border-radius: 15px;
        width: 20%;
    }
</style>