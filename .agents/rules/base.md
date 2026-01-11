# AI Agent Development Rules  

This document contains all development rules and guidelines for this project, applicable to all AI agents (Claude, Gemini, etc.).

## 1. Core Principles

- **Baby Steps**: Always work in baby steps, one at a time. Never go forward more than one step.
- **Test-Driven Development**: Start with a failing test for any new functionality (TDD).
- **Progressive Revelation**: Never show all the code at once; only the next step.
- **Type Safety**: All code must be fully typed.
- **Simplicity First**: Use the simplest working solution; avoid unnecessary abstractions.
- **Small Components**: Classes and methods should be small (10–20 lines max).
- **Clear Naming**: Use clear, descriptive names for all variables and functions.
- **Incremental Changes**: Prefer incremental, focused changes over large, complex modifications.
- **Question Assumptions**: Always question assumptions and inferences.
- **Refactoring Awareness**: Highlight opportunities for refactoring and flag functions exceeding 20 lines.
- **Pattern Detection**: Detect and highlight repeated code patterns.

## 2. Code Quality & Coverage

- **MANDATORY Validation**: Before EVERY commit, run the project's validation task and fix ALL errors. Zero tolerance.
- **Quality Requirements**: The project has strict requirements for code quality and maintainability.
- **High Coverage**: All code must have very high test coverage; strive for 100% where practical.
- **Pre-commit Checks**: All code must pass the project's quality checks before any commit (typically: typing, formatting, style/linting).
- **TDD Workflow**: Test-Driven Development (TDD) is the default workflow: always write tests first.
- **Idiomatic Go Design**: Use composition over inheritance, small interfaces, and clear package boundaries.

## 3. Style Guidelines

- **Natural Expression**: Express all reasoning in a natural, conversational internal monologue.
- **Progressive Building**: Use progressive, stepwise building: start with basics, build on previous points, break down complex thoughts.
- **Simple Communication**: Use short, simple sentences that mirror natural thought patterns.
- **Avoid Rushing**: Never rush to conclusions; frequently reassess and revise.
- **Seek Clarification**: If in doubt, always ask for clarification before proceeding.
- **Self-Documenting Code**: Avoid comments in code; rely on self-documenting names. Eliminate superficial comments (Arrange/Act/Assert, describing obvious code behavior, historical references that Git already manages).

## 4. Communication Style

- Think through problems before responding
- Show reasoning when solving complex issues
- Ask for clarification when requirements are unclear
- Keep responses focused and professional

## 5. Process & Key Requirements

- **Show Work**: Show reasoning when solving complex issues.
- **Embrace Uncertainty**: Embrace uncertainty and revision.
- **Persistence**: Persist through multiple attempts until resolution.
- **Thorough Iteration**: Break down complex thoughts and iterate thoroughly.
- **Sequential Questions**: Only one question at a time; each question should build on previous answers.

## 6. Language Standards

- **Communication Flexibility**: Conversations with AI agents can be in any language—respond in the same language the user writes in.
- **English-Only Artifacts**: All technical artifacts must always use English, including:
  - Code (variables, functions, types, comments)
  - Documentation (README, guides, API docs)
  - GitHub Issues and Pull Requests (titles, descriptions, comments)
  - Data schemas and type definitions
  - Configuration files and scripts
  - Git commit messages and branch names
  - Test names and descriptions
- **Never Mix Languages in Artifacts**: Even if the conversation is in Spanish, all generated code, commits, docs, and issues must be in English.
- **Professional Consistency**: This ensures global collaboration, tool compatibility, and industry best practices.

## 7. Documentation Standards

- **User-Focused README**: README.md must be user-focused, containing only information relevant to end users.
- **Separate Dev Docs**: All developer, CI, and infrastructure documentation must be placed in a separate development guide (e.g., docs/development_guide.md), with a clear link from the README.
- **Error Examples**: User-facing documentation should include example error messages for common validation failures to help users quickly resolve issues.

## 8. Development Best Practices

### Error Handling & Debugging
- **Graceful Error Handling**: Always implement proper error handling with meaningful error messages.
- **Debugging First**: When encountering issues, use debugging tools and logging before asking for help.
- **Error Context**: Provide sufficient context in error messages to enable quick problem resolution.
- **Fail Fast**: Design code to fail fast and fail clearly when errors occur.

### Code Review & Collaboration  
- **Pair Programming**: Prefer pairing sessions for complex features and knowledge sharing.
- **Small Pull Requests**: Keep changes small and focused for easier review and faster integration.
- **Code Review Standards**: All code must be reviewed before merging, following project quality standards.
- **Knowledge Sharing**: Document decisions and share context with team members.

### Security Considerations
- **Security by Design**: Consider security implications in all design decisions.
- **Input Validation**: Always validate and sanitize user inputs and external data.
- **Secrets Management**: Never hardcode secrets; use environment variables or secure secret management.
- **Dependency Security**: Regularly update dependencies and monitor for security vulnerabilities.

### Testing Strategy Distinction
- **Unit Tests**: Fast, isolated tests for individual components (majority of test suite).
- **Integration Tests**: Test interactions between components and external systems (limited, focused).
- **E2E Tests**: Full system validation (minimal, critical user paths only).
- **Test Pyramid**: Follow the test pyramid - many unit tests, some integration tests, few E2E tests.

## 9. Test-Driven Development Rules

### TDD Approach
- **Failing Test First**: Always start with a failing test before implementing new functionality.
- **Single Test**: Write only one test at a time; never create more than one test per change.
- **Complete Coverage**: Ensure every new feature or bugfix is covered by a test.

### Test Structure & Style
- **Consistent Tooling**: Use the project's configured test runner, assertion library, and mocking framework consistently.
- **Explicit Types**: All function signatures must have explicit types; avoid `interface{}` where possible.
- **Focused Tests**: Keep each test focused and under 20 lines.
- **Clear Naming**: Use clear, descriptive names for test functions and variables.
- **No Comments**: Avoid comments; make code self-documenting through naming.
- **Simple Helpers**: Use helper methods (e.g., object mothers/factories) for repeated setup.
- **Strategic Mocking**: Use standard library mocking for system/platform modules. Use the project's mocking framework for application code.

### Test Simplicity & Maintainability
- **Simplest Setup**: Prefer the simplest test setup that covers the requirement.
- **Refactor Tests**: Refactor tests to remove duplication and improve readability.
- **Consistent Assertions**: Use one assertion style consistently throughout the suite.
- **Extract Helpers**: If a test setup is repeated, extract a helper or fixture.
- **Readable Tests**: Always keep tests readable and easy to modify.

### Test Process & Output
- **Single Test Display**: Only show one test at a time; never present multiple tests in a single step.
- **Single File Display**: Never show more than one file at a time.
- **Self-Contained Tests**: Each test should be self-contained and not depend on the order of execution.
- **Clarify Requirements**: If in doubt about requirements, ask for clarification before writing the test.
- **Verify Failure**: After writing a test, run it to ensure it fails before implementing the feature.
- **Automatic Test Running**: After every code or test change, always run the relevant tests using the project's task runner. Do not ask for permission to run tests—just do it.

### Test Naming & Coverage
- **Descriptive Names**: Test function names should clearly describe the scenario and expected outcome.
- **Purpose-Driven Variables**: Use descriptive variable names that reflect their purpose in the test.
- **Incremental Coverage**: Ensure all code paths and edge cases are eventually covered by tests, but add them incrementally.

### Test Review & Refactoring
- **Post-Pass Review**: After a test passes, review for opportunities to simplify or clarify.
- **Helper Refactoring**: Refactor test helpers and fixtures as needed to keep the suite DRY and maintainable.

## 10. Task Runner Usage

### Core Rule
**NEVER** call tools like `go test`, `golangci-lint`, `gofmt`, or similar directly. Always use the project's task runner (Makefile).

### Discovering Available Tasks
Before starting work on a project, run `make help` to see all available tasks.

### Common Task Categories
Projects typically provide tasks for:
- **Testing**: Unit tests, integration tests, e2e tests
- **Formatting**: Code formatting and format checking
- **Linting/Style**: Style checking and static analysis
- **Type Checking**: Static type analysis
- **Building**: Build and package the application
- **Running**: Start the application locally

### Usage Rules
1. **Discover first**: Always check what tasks are available before running any tool.
2. **Use task runner**: Never call underlying tools directly; use the configured task.
3. **Add new tasks**: If a new operation is needed, prefer adding a new task rather than running a tool directly.

### Good vs Bad Examples
```sh
# Good: Use task runner
make test             # Run tests
make lint             # Run linter
make validate         # Run all checks
make fmt              # Format code

# Bad: Call tools directly
go test ./...
golangci-lint run
gofmt -w .
```

## 11. Pre-Commit Validation (MANDATORY)

Before ANY commit:
1. Run the project's validation task (e.g., `make validate`, `npm run validate`, `./gradlew check`)
2. If errors exist: fix them and re-run
3. Only commit when validation passes with ZERO errors

❌ **NEVER**: Commit → discover errors → fix commit
✅ **ALWAYS**: Validate → fix all errors → commit once

## 12. Quick Reference for All AI Agents

When working on this project:

1. **Take baby steps** - one test, one file, one change at a time 👣
2. **Always write the failing test first** (TDD) ❌➡️✅
3. **Use task runner** - never call tools directly 🔧
4. **Keep code small and typed** - max 20 lines per method 📏
5. **Show reasoning for complex issues** - be clear and professional 💭
6. **Question everything** - assumptions, requirements, design choices ❓
7. **Validate before EVERY commit** - zero tolerance ✅
8. **Run tests automatically** after every change 🧪
9. **Focus on simplicity** over cleverness ✨
10. **Ask for clarification** when in doubt 🤔

Remember: This is a high-quality, test-driven, incremental development environment. Quality over speed, clarity over cleverness, baby steps over big leaps. 