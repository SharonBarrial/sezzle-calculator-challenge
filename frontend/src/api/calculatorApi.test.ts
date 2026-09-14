import { describe, it, expect, vi, beforeEach, afterEach } from "vitest";
import { calculatorApi, ApiError, callOperation, requiresSecondOperand } from "./calculatorApi";

// Mock the global fetch function provide a way to set the response for each test
function mockFetchOnce(status: number, body: unknown) {
  vi.stubGlobal(
    "fetch",
    vi.fn().mockResolvedValue({
      ok: status >= 200 && status < 300,
      status,
      json: async () => body,
    })
  );
}

describe("calculatorApi", () => {
  beforeEach(() => {
    vi.restoreAllMocks();
  });

  afterEach(() => {
    vi.unstubAllGlobals();
  });

  it("returns the result on a successful call", async () => {
    mockFetchOnce(200, { result: 5 });
    const { result } = await calculatorApi.add(2, 3);
    expect(result).toBe(5);
  });

  it("sends the operands as JSON in the request body", async () => {
    mockFetchOnce(200, { result: 5 });
    await calculatorApi.add(2, 3);

    const call = (fetch as unknown as ReturnType<typeof vi.fn>).mock.calls[0];
    const [, options] = call;
    expect(JSON.parse(options.body)).toEqual({ a: 2, b: 3 });
  });

  it("throws an ApiError with the backend message on a 4xx/5xx response", async () => {
    mockFetchOnce(422, { error: "division by zero" });
    await expect(calculatorApi.divide(10, 0)).rejects.toThrow(ApiError);
    await expect(calculatorApi.divide(10, 0)).rejects.toThrow("division by zero");
  });

  it("throws a friendly ApiError when the network request itself fails", async () => {
    vi.stubGlobal("fetch", vi.fn().mockRejectedValue(new Error("network down")));
    await expect(calculatorApi.add(1, 2)).rejects.toThrow(ApiError);
  });

  it("only sqrt does not require a second operand", () => {
    expect(requiresSecondOperand("sqrt")).toBe(false);
    expect(requiresSecondOperand("add")).toBe(true);
    expect(requiresSecondOperand("percentage")).toBe(true);
  });

  it("callOperation routes to the right API method", async () => {
    mockFetchOnce(200, { result: 4 });
    const { result } = await callOperation("sqrt", 16, 0);
    expect(result).toBe(4);

    const call = (fetch as unknown as ReturnType<typeof vi.fn>).mock.calls[0];
    expect(call[0]).toContain("/api/sqrt");
  });
});