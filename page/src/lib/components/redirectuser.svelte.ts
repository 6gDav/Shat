import { goto } from "$app/navigation";
import { user } from "$lib/components/userName.svelte"; 
import { port, mDNSname } from "$lib/components/port.svelte";

async function redirectUser() {
    try {
        let response = await fetch(`http://${mDNSname}:${port}/redirectuser`);
        let shouldRedirect = await response.json();

        if (shouldRedirect.redirect && shouldRedirect.userName) {
            user.userName = shouldRedirect.userName;
            user.initialIp = shouldRedirect.ipAddress;
            goto("/pages/Dashboard");
        }
    } catch (error) {
        console.error("Error occurred while trying to redirect.", error);
    }
}

export default redirectUser