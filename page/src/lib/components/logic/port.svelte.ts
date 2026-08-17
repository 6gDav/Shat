import { browser } from '$app/environment';

export let port = browser ? window.location.port : "3000";
export let mDNSname = browser ? window.location.hostname : "loginpage.local";