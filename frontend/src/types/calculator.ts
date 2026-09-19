export type Operation =
    | "add"
    | "subtract"
    | "multiply"
    | "divide"
    | "power"
    | "sqrt"
    | "percentage";

export interface CalculateRequest {
    a: number;
    b: number;
}

export interface CalculateResponse {
    operation: Operation;
    a: number;
    b: number;
    result: number;
}

export interface ApiErrorDetail {
    code: string;
    message: string;
    path: string;
    timestamp: string;
}

export interface ApiErrorBody {
    error: ApiErrorDetail
}