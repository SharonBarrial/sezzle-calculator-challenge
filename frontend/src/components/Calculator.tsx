import { useState, type FormEvent } from "react";
import {
  callOperation,
  requiresSecondOperand,
  ApiError,
  type Operation,
} from "../api/calculatorApi";
import "./Calculator.css";

const OPERATIONS: { value: Operation; label: string; symbol: string }[] = [
  { value: "add", label: "Suma", symbol: "+" },
  { value: "subtract", label: "Resta", symbol: "−" },
  { value: "multiply", label: "Multiplicación", symbol: "×" },
  { value: "divide", label: "División", symbol: "÷" },
  { value: "power", label: "Potencia", symbol: "^" },
  { value: "sqrt", label: "Raíz cuadrada", symbol: "√" },
  { value: "percentage", label: "Porcentaje (a es qué % de b)", symbol: "%" },
];

// A blank or partially-typed number field should not trigger a validation 
// error while the user is still typing they try to submit
function parseOperand(raw: string): number | null {
  if (raw.trim() === "") return null;
  const value = Number(raw);
  return Number.isFinite(value) ? value : null;
}

export function Calculator() {
  const [operation, setOperation] = useState<Operation>("add");
  const [aInput, setAInput] = useState("");
  const [bInput, setBInput] = useState("");
  const [result, setResult] = useState<number | null>(null);
  const [error, setError] = useState<string | null>(null);
  const [isLoading, setIsLoading] = useState(false);

  const needsB = requiresSecondOperand(operation);

  async function handleSubmit(e: FormEvent) {
    e.preventDefault();
    setResult(null);
    setError(null);

    const a = parseOperand(aInput);
    if (a === null) {
      setError("Ingresa un número válido para el primer valor.");
      return;
    }

    let b = 0;
    if (needsB) {
      const parsedB = parseOperand(bInput);
      if (parsedB === null) {
        setError("Ingresa un número válido para el segundo valor.");
        return;
      }
      b = parsedB;
    }

    setIsLoading(true);
    try {
      const { result } = await callOperation(operation, a, b);
      setResult(result);
    } catch (err) {
      setError(err instanceof ApiError ? err.message : "Ocurrió un error inesperado.");
    } finally {
      setIsLoading(false);
    }
  }

  const selectedOp = OPERATIONS.find((op) => op.value === operation)!;

  return (
    <div className="calculator-card">
      <h1>Calculadora</h1>

      <form onSubmit={handleSubmit} noValidate>
        <label className="field">
          <span>Operación</span>
          <select
            value={operation}
            onChange={(e) => {
              setOperation(e.target.value as Operation);
              setResult(null);
              setError(null);
            }}
          >
            {OPERATIONS.map((op) => (
              <option key={op.value} value={op.value}>
                {op.label}
              </option>
            ))}
          </select>
        </label>

        <div className="operands">
          <label className="field">
            <span>{needsB ? "Valor A" : "Número"}</span>
            <input
              type="text"
              inputMode="decimal"
              value={aInput}
              onChange={(e) => setAInput(e.target.value)}
              placeholder="0"
              aria-label="Primer valor"
            />
          </label>

          {needsB && (
            <>
              <span className="operator-symbol" aria-hidden="true">
                {selectedOp.symbol}
              </span>
              <label className="field">
                <span>Valor B</span>
                <input
                  type="text"
                  inputMode="decimal"
                  value={bInput}
                  onChange={(e) => setBInput(e.target.value)}
                  placeholder="0"
                  aria-label="Segundo valor"
                />
              </label>
            </>
          )}
        </div>

        <button type="submit" disabled={isLoading}>
          {isLoading ? "Calculando..." : "Calcular"}
        </button>
      </form>

      {error && (
        <p className="message message-error" role="alert">
          {error}
        </p>
      )}

      {result !== null && !error && (
        <p className="message message-result" role="status">
          Resultado: <strong>{result}</strong>
        </p>
      )}
    </div>
  );
}