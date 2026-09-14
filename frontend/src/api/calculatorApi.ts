// Thin API client. Keeping this separate from the UI means the components
//don't know or care whether we're talking to REST, GraphQL, etc — and it
//gives us one place to mock in tests.

export type Operation =
  | "add"
  | "subtract"
  | "multiply"
  | "divide"
  | "power"
  | "sqrt"
  | "percentage";

export interface OperationResult {
  result: number;
}

// Thrown when the backend responds with a 4xx/5xx and an { error } body.
// Kept as a distinct class so callers can distinguish "the API told us
// this input is invalid" from "the network/request itself failed".
export class ApiError extends Error {
  constructor(message: string) {
    super(message);
    this.name = "ApiError";
  }
}

const BASE_URL = import.meta.env.VITE_API_BASE_URL ?? "http://localhost:8080";

async function post(path: string, body: object): Promise<OperationResult> {
  let response: Response;
  try {
    response = await fetch(`${BASE_URL}${path}`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(body),
    });
  } catch {
    throw new ApiError("No se pudo conectar con el servidor. ¿Está corriendo el backend?");
  }

  const data = await response.json().catch(() => null);

  if (!response.ok) {
    throw new ApiError(data?.error ?? "Ocurrió un error inesperado en el servidor.");
  }

  return data as OperationResult;
}

export const calculatorApi = {
  add: (a: number, b: number) => post("/api/add", { a, b }),
  subtract: (a: number, b: number) => post("/api/subtract", { a, b }),
  multiply: (a: number, b: number) => post("/api/multiply", { a, b }),
  divide: (a: number, b: number) => post("/api/divide", { a, b }),
  power: (a: number, b: number) => post("/api/power", { a, b }),
  sqrt: (a: number) => post("/api/sqrt", { a }),
  percentage: (a: number, b: number) => post("/api/percentage", { a, b }),
};

// Maps a UI operation id + operand pair to the right API call.
// Sqrt only uses `a`; every other operation uses both.
export function callOperation(
  operation: Operation,
  a: number,
  b: number
): Promise<OperationResult> {
  switch (operation) {
    case "add":
      return calculatorApi.add(a, b);
    case "subtract":
      return calculatorApi.subtract(a, b);
    case "multiply":
      return calculatorApi.multiply(a, b);
    case "divide":
      return calculatorApi.divide(a, b);
    case "power":
      return calculatorApi.power(a, b);
    case "sqrt":
      return calculatorApi.sqrt(a);
    case "percentage":
      return calculatorApi.percentage(a, b);
  }
}

// Does this operation need a second operand?
export function requiresSecondOperand(operation: Operation): boolean {
  return operation !== "sqrt";
}