import {useCallback, useState} from "react";
import type {CalculateResponse, Operation} from "../types/calculator.ts";
import {calculate, useFetch} from "./useFetch.ts";
import {formatNumber} from "../lib/format.ts";

export const useCalculator = () =>
{
    const [display, setDisplay] = useState("0");
    const [previousValue, setPreviousValue] = useState<number | null>(null);
    const [pendingOperation, setPendingOperation] = useState<Operation | null>(null);
    const [overwrite, setOverwrite] = useState(false);

    const {loading, error, execute} = useFetch<CalculateResponse>();

    const inputDigit = useCallback((digit: string) =>
    {
        if (overwrite) {
            setDisplay(digit);
            setOverwrite(false);
            return;
        }
        setDisplay((prev) => (prev === "0" ? digit : prev + digit));
    }, [overwrite]);

    const inputDecimal = useCallback(() =>
    {
        if (overwrite) {
            setDisplay("0.");
            setOverwrite(false);
            return;
        }
        setDisplay((prev) => (prev.includes(".") ? prev : prev + "."));
    }, [overwrite]);

    const clear = useCallback(() =>
    {
        setDisplay("0");
        setPreviousValue(null);
        setPendingOperation(null);
        setOverwrite(false);
    }, []);

    const chooseOperation = useCallback(async (operation: Operation) =>
    {
        if (operation == "sqrt") {
            try {
                const result = await execute(() => calculate("sqrt", {a: Number(display), b: 0}));
                setDisplay(formatNumber(result.result));
            } catch (e) {
            } finally {
                setOverwrite(true);
            }
            return;
        }
        setPreviousValue(Number(display));
        setPendingOperation(operation);
        setOverwrite(true);
    }, [display, execute]);

    const equals = useCallback(async () =>
    {
        if (previousValue === null || pendingOperation === null) return;

        try {
            const result = await execute(() =>
            {
                return calculate(pendingOperation, {a: previousValue, b: Number(display)});
            });

            setDisplay(formatNumber(result.result));
        } catch (e) {
        } finally {
            setPreviousValue(null);
            setPendingOperation(null);
            setOverwrite(true);
        }
    }, [previousValue, pendingOperation, display, execute]);

    return {
        display,
        previousValue,
        pendingOperation,
        loading,
        error,
        inputDigit,
        inputDecimal,
        clear,
        chooseOperation,
        equals
    };
};