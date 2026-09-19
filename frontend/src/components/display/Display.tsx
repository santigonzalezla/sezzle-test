import styles from './display.module.css';
import {OPERATIONS} from "../../lib/operations.ts";
import type {Operation} from "../../types/calculator.ts";

interface DisplayProps
{
    value: string;
    previousValue: number | null;
    pendingOperation: Operation | null;
    error: string | null;
    loading: boolean;
}

export const Display = ({value, previousValue, pendingOperation, error, loading}: DisplayProps) =>
{
    const symbol = OPERATIONS.find((op) => op.operation === pendingOperation)?.label;

    return (
        <div className={styles.display}>
            <div className={styles.history}>
                {previousValue !== null && symbol ? `${previousValue} ${symbol}` : "\u00A0"}
            </div>
            <div className={styles.value} data-testid="display-value">{loading ? "..." : value}</div>
            {error && <div className={styles.error}>{error}</div>}
        </div>
    );
};