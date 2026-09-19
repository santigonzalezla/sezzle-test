import type {Operation} from "../types/calculator.ts";

export interface OperationButton
{
    operation: Operation;
    label: string;
}

export const OPERATIONS: OperationButton[] = [
    {operation: "add", label: "+"},
    {operation: "subtract", label: "-"},
    {operation: "multiply", label: "*"},
    {operation: "divide", label: "/"},
    {operation: "power", label: "^"},
    {operation: "sqrt", label: "√"},
    {operation: "percentage", label: "%"},

]