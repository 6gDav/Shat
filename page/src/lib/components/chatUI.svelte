<script lang="ts">
    let { userName } = $props();
</script>

<div>
    <div class="title-container">
        <h1 id="name-title">{userName}</h1>
    </div>
</div>

<style>
    .title-container {
        display: flex;
        justify-content: center;
        align-items: center;
        width: 100%;
        height: 100%;
    }

    #name-title {
        font-size: 52px;
        text-align: center;
    }
</style>
