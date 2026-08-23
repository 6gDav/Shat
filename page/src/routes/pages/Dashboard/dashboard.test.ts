import { render, screen, fireEvent } from '@testing-library/svelte';
import { describe, it, expect, vi, beforeEach } from 'vitest';
// @ts-ignore
import Dashboard from './+page.svelte';

vi.mock('$lib/components/logic/manageWebSocketConn.svelte', () => ({
  manageConnection: vi.fn(),
  usersList: [
    { name: 'Alice', ip: '192.168.1.10' },
    { name: 'Bob', ip: '192.168.1.11' }
  ]
}));

vi.mock('$lib/components/logic/userName.svelte', () => ({
  user: { userName: 'TesztUser' }
}));

vi.mock('$lib/components/ui/chatUI.svelte', () => ({
  default: vi.fn()
}));

vi.mock('$lib/components/ui/activityindicator.svelte', () => ({
  default: vi.fn()
}));

vi.mock('$lib/components/ui/changeName.svelte', () => ({
  default: vi.fn()
}));

vi.mock('$lib/assets/user-not-found.png', () => ({
  default: 'mock-user-not-found.png'
}));

describe('Dashboard Components', () => {
  beforeEach(() => {
    vi.clearAllMocks();
  });

  it('start the websocket', async () => {
    const { manageConnection } = await import(
      '$lib/components/logic/manageWebSocketConn.svelte'
    );
    render(Dashboard);
    expect(manageConnection).toHaveBeenCalledTimes(1);
  });

  it('the hamburger menu opens the sidebar.', async () => {
    render(Dashboard);
    const hamburgerBtn = screen.getByRole('button', { name: 'Toggle navigation' });
    const sidebar = document.querySelector('.sidebar');

    expect(sidebar?.classList.contains('open')).toBe(false);

    await fireEvent.click(hamburgerBtn);
    expect(sidebar?.classList.contains('open')).toBe(true);

    const closeBtn = screen.getByRole('button', { name: 'Close menu' });
    await fireEvent.click(closeBtn);
    expect(sidebar?.classList.contains('open')).toBe(false);
  });

  it('clicking on the user activates the button', async () => {
    render(Dashboard);
    const aliceBtn = screen.getByRole('button', { name: 'Alice' });

    expect(aliceBtn.classList.contains('active')).toBe(false);
    await fireEvent.click(aliceBtn);
    expect(aliceBtn.classList.contains('active')).toBe(true);
  });
});