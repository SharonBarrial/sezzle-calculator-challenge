# AI Prompts Used

This project was developed with the assistance of AI (Claude) following an **interactive, learning-driven approach**. I used AI differently depending on the technology involved: as a code reviewer and architecture discussion partner for React/TypeScript, as a technical mentor while learning Go, and as a troubleshooting and best-practice assistant for Docker.

The purpose was not to generate the complete application automatically, but to discuss implementation decisions, validate my approach, investigate errors, and better understand areas where I had less experience.

## 1. Requirement Analysis & Architecture

* **Prompt intent:** Shared the assessment requirements and discussed how to divide responsibilities between the React/TypeScript frontend and the Go backend.

* **Example questions:**

  * *"Given these requirements, which parts of this architecture are actually necessary for the scope of the project, and which would be overengineering?"*
  * *"How should the API contract be designed so the React application does not become tightly coupled to the backend implementation?"*

* **Key takeaway:** Established a clear separation between presentation, API communication, business logic, and HTTP concerns.

## 2. React & TypeScript — Advanced Implementation

* **Prompt intent:** Used AI primarily as a **code reviewer and architectural discussion partner**. Since I have stronger experience with React and TypeScript, the questions focused less on basic syntax and more on maintainability, performance, component design, asynchronous state, and testing.

* **Example questions:**

  * *"Review this React component as if it were going into a production codebase. Where are the responsibilities becoming too coupled?"*
  * *"Would you keep this API logic inside the component or extract it into a custom hook/service? Compare the trade-offs."*
  * *"Is this state actually necessary, or can this value be derived from existing state and props?"*
  * *"Could this implementation cause unnecessary re-renders? Walk me through the render cycle and explain why."*
  * *"How would you test this component without coupling the tests to its implementation details?"*
  * *"Review the component hierarchy and suggest improvements without introducing unnecessary abstractions."*
  
* **Key takeaway:** Used AI mainly to challenge and validate my existing React/TypeScript knowledge. The discussions focused on component architecture, state management, rendering behavior, reusable abstractions, API integration, type safety, accessibility, and testing.

## 3. Go — Learning the Language and Backend Fundamentals

* **Prompt intent:** Used AI more extensively as a **mentor while learning Go**, particularly to understand concepts that differ from languages and frameworks I was already familiar with.

* **Example questions:**

  * *"I'm new to Go. Can you explain how `http.HandlerFunc` works instead of just showing me the implementation?"*
  * *"Why are errors handled this way in Go instead of using exceptions?"*
  * *"What's the difference between a struct, interface, and type alias in Go, and which one makes sense for this case?"*
  * *"Why would I separate the calculator logic from the HTTP handler? Can you show me what becomes easier to test?"*
  * *"What does `json.NewDecoder(r.Body).DisallowUnknownFields()` actually do, and what problem does it solve?"*
  * *"I'm getting this compiler error. Before fixing it, can you explain what Go is telling me?"*
  * *"Am I writing Go as if it were TypeScript?"*

* **Key takeaway:** The backend became an opportunity to learn Go fundamentals, particularly `net/http`, structs, interfaces, error handling, JSON decoding, validation, package organization, and Go's testing conventions.

## 4. Go Testing & Error Handling

* **Prompt intent:** Asked for guidance on writing tests while learning Go's testing philosophy and table-driven testing approach.

* **Example questions:**

  * *"Can you explain table-driven tests in Go and why they are preferred in many Go projects?"*
  * *"Should this error be handled by the calculator package or by the HTTP handler?"*
  * *"Why does my business logic have 100% coverage while my HTTP handler has lower coverage?"*
  * *"Which branches are realistically testable and which ones represent defensive error handling?"*

* **Key takeaway:** Learned to separate unit tests for pure business logic from HTTP-level tests and to use table-driven tests to cover multiple cases consistently.

## 5. Docker — Intermediate-Level Usage

* **Prompt intent:** Used AI to understand and troubleshoot Docker while working with multiple services. The focus was on understanding how containers communicate, how images are built, and how development differs from production.

* **Example questions:**

  * *"Can you explain what is happening in each stage of this multi-stage Dockerfile?"*
  * *"What should be a build-time environment variable versus a runtime environment variable in this setup?"*
  * *"How can I reduce the final Go image size without making the Dockerfile unnecessarily complicated?"*
  * *"The container is running, but the frontend cannot reach the backend. What networking issues should I check?"*
  * *"How would you organize Docker Compose for a React frontend, Go API, and their respective development/production configurations?"*

* **Key takeaway:** Developed an intermediate understanding of Docker, including multi-stage builds, Docker Compose, service-to-service networking, port mapping, build arguments, environment variables, image optimization, and debugging containers.

## 6. Debugging & Problem Solving

* **Prompt intent:** Used AI to investigate problems encountered during development instead of immediately replacing the implementation.

* **Example questions:**

  * *"Here is the error and the code that caused it. Can you help me identify which layer is actually responsible?"*
  * *"What information from the compiler/runtime output should I look at first?"*
  * *"The frontend works locally but fails inside Docker. How can I systematically determine whether the problem is Vite, Nginx, Docker networking, or the Go API?"*
  * *"Can you give me a debugging checklist instead of immediately giving me the final solution?"*

* **Key takeaway:** Practiced debugging systematically by identifying whether an issue originated in the frontend, backend, configuration, or container environment.

## 7. Code Review & Engineering Decisions

* **Prompt intent:** Used AI to challenge implementation decisions and identify unnecessary complexity.

* **Example questions:**

  * *"Review this implementation and tell me which parts demonstrate good engineering practices and which parts could be simplified."*
  * *"Would an experienced React developer consider this abstraction useful, or is it premature abstraction?"*
  * *"Is this approach idiomatic Go even though I'm coming from a TypeScript background?"*
  * *"Which Docker optimizations are actually worth implementing for a project of this size?"*
  * *"If you were reviewing this as a technical assessment, what implementation decisions would you question?"*

* **Key takeaway:** AI was used not only for fixing problems but also for challenging technical decisions and understanding the trade-offs behind different approaches.

## Summary of AI Usage

* **React/TypeScript:** AI was primarily used as a **code reviewer, architecture discussion partner, and optimization assistant**, allowing me to validate more advanced frontend decisions around component design, state, rendering, type safety, API integration, and testing.

* **Go:** AI was primarily used as a **technical mentor while learning the language**, helping me understand Go's syntax, idioms, HTTP handling, error management, testing patterns, and backend organization.

* **Docker:** AI was used at an **intermediate level** for containerization, networking, multi-stage builds, environment configuration, image optimization, and troubleshooting.

* **Role of Developer:** I defined the requirements, made the final implementation decisions, wrote and modified the code, ran the application locally, executed tests, investigated errors, designed the UI, configured the project environment, and managed the repository with Git.

Overall, AI functioned as a **learning and development partner rather than an autonomous code generator**. The interaction varied according to my experience with each technology: deeper architectural discussions for React/TypeScript, guided learning for Go, and practical troubleshooting and optimization for Docker.
