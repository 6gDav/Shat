import { render, screen } from '@testing-library/svelte';
import { describe, it, expect, vi, beforeEach } from 'vitest';
// @ts-ignore
import Page from './+page.svelte';

vi.mock('$lib/components/logic/redirectuser.svelte', () => ({
    default: vi.fn()
}));

vi.mock('$lib/components/ui/registration.svelte', () => ({
    default: vi.fn()
}));

import redirectUser from '$lib/components/logic/redirectuser.svelte';

describe('Registration Page', () => {
    beforeEach(() => {
        vi.clearAllMocks();
    });
    it('call redirect user via onMount', () => {
        render(Page);
        expect(redirectUser).toHaveBeenCalledTimes(1);
    });

    it('show the proper containers', () => {
        const { container } = render(Page);

        const wrapperDiv = container.querySelector('.container.login-container');
        expect(wrapperDiv).not.toBeNull();
    });
});