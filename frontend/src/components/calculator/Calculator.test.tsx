import {afterEach, describe, expect, it, vi} from "vitest";
import {userEvent} from "@testing-library/user-event";
import {render, screen, waitFor} from "@testing-library/react";
import {Calculator} from "./Calculator.tsx";

function mockFetchOk(result: number)
{
    vi.stubGlobal('fetch', vi.fn().mockResolvedValue({
        ok: true,
        json: () => Promise.resolve({operation: 'multiply', a: 8, b: 4, result})
    }));
}

function mockFetchError(message: string)
{
    vi.stubGlobal('fetch', vi.fn().mockResolvedValue({
        ok: false,
        json: () => Promise.resolve({error: {message}})
    }));
}

describe('Calculator', () =>
{
    afterEach(() =>
    {
        vi.unstubAllGlobals();
    });

    it('should calculate 8 * 4 and show result ', async () =>
    {
        mockFetchOk(32);
        const user = userEvent.setup();

        render(<Calculator/>);

        await user.click(screen.getByRole('button', {name: '8'}));
        await user.click(screen.getByRole('button', {name: '*'}));
        await user.click(screen.getByRole('button', {name: '4'}));
        await user.click(screen.getByRole('button', {name: '='}));

        await waitFor(() =>
        {
            expect(screen.getByText('32')).toBeInTheDocument();
        });
    });

    it('should show error message when backend rejects op', async () =>
    {
        mockFetchError('division by zero');
        const user = userEvent.setup();

        render(<Calculator/>);

        await user.click(screen.getByRole('button', {name: '5'}));
        await user.click(screen.getByRole('button', {name: '/'}));
        await user.click(screen.getByRole('button', {name: '0'}));
        await user.click(screen.getByRole('button', {name: '='}));

        await waitFor(() =>
        {
            expect(screen.getByText('division by zero')).toBeInTheDocument();
        });
    });

    it('should calculate root of 144 and get 12', async () =>
    {
        mockFetchOk(12);
        const user = userEvent.setup();

        render(<Calculator/>);

        await user.click(screen.getByRole('button', {name: '1'}));
        await user.click(screen.getByRole('button', {name: '4'}));
        await user.click(screen.getByRole('button', {name: '4'}));
        await user.click(screen.getByRole('button', {name: '√'}));

        await waitFor(() =>
        {
            expect(screen.getByTestId('display-value')).toHaveTextContent('12');
        });
    });

    it('should allow decimal input an clean with C', async () =>
    {
        const user = userEvent.setup();

        render(<Calculator/>);

        await user.click(screen.getByRole('button', {name: '1'}));
        await user.click(screen.getByRole('button', {name: '.'}));
        await user.click(screen.getByRole('button', {name: '5'}));

        expect(screen.getByTestId('display-value')).toHaveTextContent('1.5');

        await user.click(screen.getByRole('button', {name: 'C'}));

        expect(screen.getByTestId('display-value')).toHaveTextContent('0');
    });
});