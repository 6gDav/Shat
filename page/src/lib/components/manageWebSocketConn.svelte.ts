import { onDestroy, onMount } from "svelte";
import { mDNSname, port } from "./port.svelte";

interface UserObj {
    ip: string,
    name: string
};

interface ChatMessage {
    from: string;
    to: string;
    text: string;
    roomId: string;
}

export let usersList = $state<UserObj[]>([]);
export let messages = $state<ChatMessage[]>([]);

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
                const rawData = res.data as Record<string, string>;

                const parsedUsers: UserObj[] = Object.entries(rawData).map(([ip, name]) => ({
                    ip: ip,
                    name: String(name)
                }));

                usersList.splice(0, usersList.length, ...parsedUsers);
            } else if (res.text && res.from) {
                messages.push({
                    from: res.from,
                    to: res.to,
                    text: res.text,
                    roomId: res.room_id
                });
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
    onMount(() => {
        //chat and name list
        startHandShake();
        return () => {
            chat?.close();
        };
    });

    // onDestroy(() => {
    //     chat?.close();
    // });
}

function getRoomId(ip1: string, ip2: string): string {
    return [ip1, ip2].sort().join("_");
}

export function sendChatMessage(senderIp: string, targetIp: string, text: string) {
    if (chat && chat.readyState === WebSocket.OPEN) {
        const payload = JSON.stringify({
            from: senderIp,
            target_ip: targetIp,
            text: text,
            room_id: getRoomId(senderIp, targetIp)
        });

        chat.send(payload);
    } else {
        alert("The connection is inactive");
    }
}