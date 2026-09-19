import {useCallback, useState} from "react";
import type {ApiErrorBody, CalculateRequest, CalculateResponse, Operation} from "../types/calculator.ts";

const BASE_URL = import.meta.env.VITE_API_BASE_URL ?? "http://localhost:8080";

interface UseFetchState<T>
{
    data: T | null;
    loading: boolean;
    error: string | null;
}
export function useFetch<T>()
{

    const [state, setState] = useState<UseFetchState<T>>({
        data: null,
        loading: false,
        error: null
    });

    const execute = useCallback(async (request: () => Promise<T>) =>
    {
        setState({data: null, loading: true, error: null});

        try {
            const data = await request();
            setState({data, loading: false, error: null});
            return data;
        } catch (e) {
            const message = e instanceof Error ? e.message : "Unexpected error";
            setState({data: null, loading: false, error: message});
            throw e;
        }
    }, []);

    return {...state, execute};
}

export async function calculate(operation: Operation, payload: CalculateRequest): Promise<CalculateResponse> {
    const res = await fetch(`${BASE_URL}/api/calculate/${operation}`, {
        method: "POST",
        headers: {"Content-Type": "application/json"},
        body: JSON.stringify(payload),
    });

    if (!res.ok) {
        const body : ApiErrorBody = await res.json();
        throw new Error(body.error.message);
    }

    return res.json();
}