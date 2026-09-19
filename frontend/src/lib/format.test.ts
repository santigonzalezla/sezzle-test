import {describe, expect, it} from "vitest";
import {formatNumber} from "./format.ts";

describe('formatNumber', () =>
{
    it.each([
        [5, '5'],
        [5.1, '5.1'],
        [0.1 + 0.2, '0.3'],
        [Infinity, 'Error'],
        [-Infinity, 'Error'],
        [NaN, 'Error']
    ])('formatNumber(%s) devuelve %s', (input, expected) =>
    {
        expect(formatNumber(input)).toBe(expected);
    });
});