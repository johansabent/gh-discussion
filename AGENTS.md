# GitHub Copilot Agent Instructions

This repository contains custom instructions for GitHub Copilot coding agents in the `.github/agents` directory.

## What are Agent Instructions?

Agent instruction files help GitHub Copilot understand how to work effectively in this repository. They provide context about:
- The technology stack and frameworks used
- How to build, test, and run the code
- Code style and conventions
- Project structure and organization
- Boundaries and best practices

## Available Agents

### 🔧 Coding Agent (`coding-agent.md`)
Expert Go developer specialized in CLI development with Cobra and Bubble Tea frameworks. Use this agent for:
- Implementing new commands
- Working with the GitHub API
- Building UI components
- General code changes

### 🧪 Test Agent (`test-agent.md`)
Testing and quality assurance specialist. Use this agent for:
- Writing unit tests
- Improving test coverage
- Setting up mocks and fixtures
- Quality checks

### 📚 Docs Agent (`docs-agent.md`)
Technical documentation specialist. Use this agent for:
- Updating README
- Writing code comments
- Creating usage examples
- Maintaining documentation consistency

## How to Use

When working with GitHub Copilot coding agents, these instruction files are automatically read to provide context-aware assistance. The agent will understand:
- How to build and test the project
- What coding patterns to follow
- What boundaries not to cross
- How to structure documentation

## Best Practices

The agent instructions in this repository follow GitHub's best practices for agent configuration:
- Each agent has a specific, well-defined role
- Instructions include concrete code examples
- Clear boundaries are established
- Tech stack is precisely documented
- Commands are provided with actual examples

## Learn More

- [GitHub Copilot coding agent documentation](https://docs.github.com/en/copilot/tutorials/coding-agent)
- [How to write great agent instructions](https://github.blog/ai-and-ml/github-copilot/how-to-write-a-great-agents-md-lessons-from-over-2500-repositories/)
- [Community examples](https://github.com/github/awesome-copilot)
