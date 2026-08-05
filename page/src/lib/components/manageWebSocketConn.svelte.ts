import { onDestroy, onMount } from "svelte";
import { mDNSname, port } from "./port.svelte";

interface WSResponse<T> {
    type: string;
    data: T;
}

export let usersList = $state<string[]>([]);

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
                
                usersList.length = 0;
                usersList.push(...payload.data);
            }
        } catch (err) {
            alert("Cannot fetch user list.");
        }
    };

    chat.onerror = (error) => {
        console.error("WebSocket error:", error);
    };
}

export function manageConnection() {
    onMount(async () => {
        //chat and name list
        $effect(() => {
            startHandShake();
        },)
    });

    onDestroy(() => {
        chat?.close();
    });
}