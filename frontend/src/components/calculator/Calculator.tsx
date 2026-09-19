import styles from './calculator.module.css';
import {useCalculator} from "../../hooks/useCalculator.ts";
import {Display} from "../display/Display.tsx";
import {Keypad} from "../keypad/Keypad.tsx";

export const Calculator = () =>
{
    const {
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
    } = useCalculator();

    return (
        <div className={styles.calculator}>
            <Display
                value={display}
                previousValue={previousValue}
                pendingOperation={pendingOperation}
                error={error}
                loading={loading}
            />
            <Keypad
                onDigit={inputDigit}
                onDecimal={inputDecimal}
                onClear={clear}
                onOperation={chooseOperation}
                onEquals={equals}
                disabled={loading}
            />
        </div>
    );
};