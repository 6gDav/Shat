<div class="sidebar">
    <div class="sidebar-header">
        <h2>Dashboard</h2>
    </div>

    <nav class="sidebar-nav">
        <a href="#" class="nav-item active"> Főoldal</a>
    </nav>
</div>

<style>
    .sidebar {
        width: 260px;
        height: 100vh;
        background: linear-gradient(to bottom, #002169 20%, #ffffff 80%);
        color: #ffffff;
        display: flex;
        flex-direction: column;
        padding: 1.5rem 1rem;
        box-sizing: border-box;
        font-family:
            system-ui,
            -apple-system,
            sans-serif;
    }

    .sidebar-header {
        padding-bottom: 1.5rem;
        margin-bottom: 1rem;
        border-bottom: 1px solid #2d2d3f;
    }

    .sidebar-header h2 {
        margin: 0;
        font-size: 1.25rem;
        font-weight: 600;
        letter-spacing: 0.5px;
    }

    .sidebar-nav {
        display: flex;
        flex-direction: column;
        gap: 0.5rem;
        flex-grow: 1;
    }

    .nav-item {
        display: flex;
        align-items: center;
        gap: 0.75rem;
        padding: 0.75rem 1rem;
        color: #a2a3b7;
        text-decoration: none;
        border-radius: 8px;
        font-size: 0.95rem;
        transition: all 0.2s ease;
    }

    .nav-item:hover {
        box-shadow: 0px 10px 25px rgba(0, 0, 0, 0.519);
    }

    .nav-item.active {
        background-color: #6366f1;
        color: #ffffff;
        font-weight: 500;
    }
</style>
