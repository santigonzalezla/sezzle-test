import {describe, afterEach, expect, it, vi} from "vitest";
import {calculate} from "./useFetch.ts";

describe('calculate', () =>
{
    afterEach(() =>
    {
        vi.unstubAllGlobals();
    });

    it('returns result when response is succesful', async () =>
    {
        const mockResponse = {operation: 'add', a: 2, b: 3, result: 5};

        vi.stubGlobal('fetch', vi.fn().mockResolvedValue({
            ok: true,
            json: () => Promise.resolve(mockResponse)
        }));

        const result = await calculate('add', {a: 2, b: 3});

        expect(result).toEqual(mockResponse);
        expect(fetch).toHaveBeenLastCalledWith(
            expect.stringContaining('/api/calculate/add'),
            expect.objectContaining({method: 'POST'})
        );
    });

    it('throws error with message from backend when error occurs', async () =>
    {
        vi.stubGlobal('fetch', vi.fn().mockResolvedValue({
            ok: false,
            json: () => Promise.resolve({error: {message: 'division by zero'}})
        }));

        await expect(calculate('divide', {a: 10, b: 0})).rejects.toThrow('division by zero');
    });
});