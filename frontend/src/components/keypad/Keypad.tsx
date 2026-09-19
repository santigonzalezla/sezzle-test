import styles from './keypad.module.css';
import type {Operation} from "../../types/calculator.ts";

interface KeypadProps
{
    onDigit: (digit: string) => void;
    onDecimal: () => void;
    onClear: () => void;
    onOperation: (operation: Operation) => void;
    onEquals: () => void;
    disabled: boolean;
}

export const Keypad = ({onDigit, onDecimal, onEquals, onClear, onOperation, disabled}: KeypadProps) =>
{
    return (
        <div className={styles.keypad}>
            <button className={styles.func} onClick={onClear} disabled={disabled}>C</button>
            <button className={styles.func} onClick={() => onOperation("sqrt")} disabled={disabled}>√</button>
            <button className={styles.func} onClick={() => onOperation("percentage")} disabled={disabled}>%</button>
            <button className={styles.op} onClick={() => onOperation("power")} disabled={disabled}>x^</button>

            <button className={styles.digit} onClick={() => onDigit("7")} disabled={disabled}>7</button>
            <button className={styles.digit} onClick={() => onDigit("8")} disabled={disabled}>8</button>
            <button className={styles.digit} onClick={() => onDigit("9")} disabled={disabled}>9</button>
            <button className={styles.op} onClick={() => onOperation("divide")} disabled={disabled}>/</button>

            <button className={styles.digit} onClick={() => onDigit("4")} disabled={disabled}>4</button>
            <button className={styles.digit} onClick={() => onDigit("5")} disabled={disabled}>5</button>
            <button className={styles.digit} onClick={() => onDigit("6")} disabled={disabled}>6</button>
            <button className={styles.op} onClick={() => onOperation("multiply")} disabled={disabled}>*</button>

            <button className={styles.digit} onClick={() => onDigit("1")} disabled={disabled}>1</button>
            <button className={styles.digit} onClick={() => onDigit("2")} disabled={disabled}>2</button>
            <button className={styles.digit} onClick={() => onDigit("3")} disabled={disabled}>3</button>
            <button className={styles.op} onClick={() => onOperation("subtract")} disabled={disabled}>-</button>

            <button className={styles.digit} onClick={() => onDigit("0")} disabled={disabled}>0</button>
            <button className={styles.digit} onClick={onDecimal} disabled={disabled}>.</button>
            <button className={styles.equals} onClick={onEquals} disabled={disabled}>=</button>
            <button className={styles.op} onClick={() => onOperation("add")} disabled={disabled}>+</button>
        </div>
    );
};