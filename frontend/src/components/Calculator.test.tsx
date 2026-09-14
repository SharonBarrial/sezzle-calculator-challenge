import { render, screen, waitFor } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { describe, it, expect, vi, beforeEach } from "vitest";
import { Calculator } from "./Calculator";
import { callOperation, ApiError } from "../api/calculatorApi";

// The component only talks to the API through `callOperation`, so that's
// the single seam we need to mock to keep these tests network-free
vi.mock("../api/calculatorApi", async () => {
  const actual = await vi.importActual<typeof import("../api/calculatorApi")>(
    "../api/calculatorApi"
  );
  return {
    ...actual,
    callOperation: vi.fn(),
  };
});

describe("Calculator", () => {
  beforeEach(() => {
    vi.clearAllMocks();
  });

  it("renders with the add operation selected by default", () => {
    render(<Calculator />);
    expect(screen.getByRole("combobox")).toHaveValue("add");
    expect(screen.getByLabelText("Segundo valor")).toBeInTheDocument();
  });

  it("hides the second operand field for sqrt", async () => {
    const user = userEvent.setup();
    render(<Calculator />);

    await user.selectOptions(screen.getByRole("combobox"), "sqrt");

    expect(screen.queryByLabelText("Segundo valor")).not.toBeInTheDocument();
  });

  it("shows a validation error instead of calling the API when input is not a number", async () => {
    const user = userEvent.setup();
    render(<Calculator />);

    await user.type(screen.getByLabelText("Primer valor"), "abc");
    await user.type(screen.getByLabelText("Segundo valor"), "3");
    await user.click(screen.getByRole("button", { name: /calcular/i }));

    expect(
      await screen.findByText(/ingresa un número válido para el primer valor/i)
    ).toBeInTheDocument();
    expect(callOperation).not.toHaveBeenCalled();
  });

  it("calls the API and displays the result on valid input", async () => {
    (callOperation as ReturnType<typeof vi.fn>).mockResolvedValue({ result: 8 });
    const user = userEvent.setup();
    render(<Calculator />);

    await user.type(screen.getByLabelText("Primer valor"), "5");
    await user.type(screen.getByLabelText("Segundo valor"), "3");
    await user.click(screen.getByRole("button", { name: /calcular/i }));

    expect(await screen.findByText(/resultado/i)).toHaveTextContent("8");
    expect(callOperation).toHaveBeenCalledWith("add", 5, 3);
  });

  it("displays the backend error message when the API rejects", async () => {
    (callOperation as ReturnType<typeof vi.fn>).mockRejectedValue(
      new ApiError("division by zero")
    );
    const user = userEvent.setup();
    render(<Calculator />);

    await user.selectOptions(screen.getByRole("combobox"), "divide");
    await user.type(screen.getByLabelText("Primer valor"), "5");
    await user.type(screen.getByLabelText("Segundo valor"), "0");
    await user.click(screen.getByRole("button", { name: /calcular/i }));

    await waitFor(() => {
      expect(screen.getByRole("alert")).toHaveTextContent("division by zero");
    });
  });
});