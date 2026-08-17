import { onMount } from "svelte";
import { mDNSname, port } from "./port.svelte";

interface UserObj {
    ip: string,
    name: string
};

interface ChatMessage {
    type: string,
    from: string;
    to?: string;
    text: string;
    roomId: string;
    userName: string;
}

export let usersList = $state<UserObj[]>([]);
export let messages = $state<ChatMessage[]>([]);

let chat: WebSocket;

function startHandShake() {
    chat = new WebSocket(`ws://${mDNSname}:${port}/setws`);

    chat.onopen = async () => {
        console.log("WebSocket is active.");

        try {
            const response = await fetch(`http://${mDNSname}:${port}/restorechathistory`)

            if (!response.ok) {
                throw new Error("Error occured while trying to fetch the chat history.")
            }

            const data = await response.json()

            if (Array.isArray(data)) {
                const historyMessages: ChatMessage[] = data.map((element: any) => ({
                    type: element.type,
                    from: element.ip,
                    text: element.message,
                    roomId: element.roomId,
                    userName: element.username,
                }));
                console.log(historyMessages)

                messages.length = 0;
                messages.push(...historyMessages);
            }
        } catch (err) {
            console.error(err)
        }
    };

    chat.onmessage = async (event) => {
        try {
            const res = await fetchingMessageEvent(event)

            //User name List
            if (res.type === "USER_NAME_LIST") {
                const rawData = res.data as Record<string, string>;

                const parsedUsers: UserObj[] = Object.entries(rawData).map(([ip, name]) => ({
                    ip: ip,
                    name: String(name)
                }));

                usersList.splice(0, usersList.length, ...parsedUsers);
            }
            //private chat history
            // else if (res.type === "private_history" && res.text && res.from) {
            //     messages.push({
            //         type: res.type,
            //         from: res.from,
            //         text: res.text,
            //         roomId: res.room_id,
            //         userName: res.username,
            //     });
            // }
            // //group chat history
            // else if (res.type === "group_history" && res.text && res.from) {
            //     messages.push({
            //         type: res.type,
            //         from: res.from,
            //         text: res.text,
            //         roomId: res.room_id,
            //         userName: res.username,
            //     });
            // }
            //private chat 
            else if (res.type === "private" && res.text && res.from) {
                messages.push({
                    type: res.type,
                    from: res.from,
                    to: res.to,
                    text: res.text,
                    roomId: res.room_id,
                    userName: res.username,
                });
            }
            //group chat
            else if (res.type === "group" && res.text && res.from) {
                messages.push({
                    type: res.type,
                    from: res.from,
                    to: res.to,
                    text: res.text,
                    roomId: res.room_id,
                    userName: res.username
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

export function sendChatMessage(senderIp: string, targetIp: string, text: string, type: string, userName: string) {
    if (chat && chat.readyState === WebSocket.OPEN) {
        const payload = JSON.stringify({
            type: type,
            from: senderIp,
            target_ip: targetIp,
            text: text,
            room_id: getRoomId(senderIp, targetIp),
            userName: userName
        });

        chat.send(payload);
    } else {
        alert("The connection is inactive");
    }
}

async function fetchingMessageEvent(event: MessageEvent<any>) {
    const textData = typeof event.data === "string"
        ? event.data
        : await event.data.text();

    return JSON.parse(textData);
}

function getRoomId(ip1: string, ip2: string): string {
    return [ip1, ip2].sort().join("_");
}