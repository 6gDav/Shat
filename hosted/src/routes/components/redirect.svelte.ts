import { onMount } from "svelte";
import { port } from "./port.svelte";
import { goto } from "$app/navigation";

let loading = true;

onMount(async () => {
    try {
        let response = await fetch(`http://loginpage.local:${port}/redirectuser`);

        let shouldRedirect = await response.json();

        if (shouldRedirect) {
            //window.location.href = `http://loginpage.local:${port}/dasboard`;
            goto("/pages/Dasboard")
        }
    } catch (error) {
        console.error("Hiba az átirányítás ellenőrzésekor:", error);
    } finally {
        loading = false;
    }
});