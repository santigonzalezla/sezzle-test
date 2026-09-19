export const formatNumber = (value: number) => {
    if (!Number.isFinite(value)) return "Error";
    return String(Number(value.toFixed(10)));
}