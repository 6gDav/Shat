<script lang="ts">
    import "../app.css";
    import LoginPage from "./pages/Login/+page.svelte"
    import { deleteUserOnClose } from './components/removeUser';

    deleteUserOnClose();
</script>

<main>
    <LoginPage /> 
</main>

