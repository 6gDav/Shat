import { browser } from '$app/environment';

export function deleteUserOnClose() {
    if (!browser) return;

    $effect(() => {
        const handleUnload = () => {
            navigator.sendBeacon('http://loginpage.local:3000/disconnect');
        };

        window.addEventListener('pagehide', handleUnload);

        return () => {
            window.removeEventListener('pagehide', handleUnload);
        };
    });
}