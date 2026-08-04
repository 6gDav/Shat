<div>
    <div class="loader" />
</div>

<style>
    .loader {
        border: 10px solid #f3f3f3; 
        border-top: 10px solid #767676; 
        border-radius: 50%;
        width: 38px;
        height: 38px;
        animation: spin 2s linear infinite;
    }

    @keyframes spin {
        0% {
            transform: rotate(0deg);
        }
        100% {
            transform: rotate(360deg);
        }
    }
</style>