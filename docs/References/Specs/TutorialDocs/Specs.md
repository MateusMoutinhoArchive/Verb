# TutorialDocs Specification

## Description
Defines the required shape of a **Tutorial** page — any page under `docs/Tutorials/`, which guides a workflow. A tutorial covers a single goal (e.g. `HandleLibrarySamples.md`, `HandleLibElements.md`), built from actionable numbered steps.

### Rules
- Every page must comply with [GeneralDoc](/docs/References/Specs/GeneralDoc/Specs.md).
- **One subject per page**: a tutorial covers one subject, and each goal it holds is a `##` section of its own carrying that goal's `### Workflow` — a page with a single goal keeps one `## Workflow` and no goal sections.
- The title names the subject as an action (e.g. `Handle Library Samples`, `Parse an Option`).
- A tutorial page must contain a `## Description` and, per goal, **actionable**, numbered steps.
- An optional `### Rules` section states constraints specific to that tutorial, or to one of its goals.
- Steps prescribe actions, not descriptions; use fenced code blocks when a step involves writing or running code.
- **Full Code at the end**: when the workflow builds a program across multiple steps (code snippets that the reader assembles into one file), the page must end with a `## Full Code` section containing the complete, copy-pasteable program the steps produce. A tutorial whose code already appears whole in a single step is exempt.
- Background explanations belong in their own explanation page — link to them instead of embedding them.
- Every new page must be registered in its theme index under `docs/Index/`.

## Structure
1. **Title** (H1): the goal, phrased as an action.
2. **`## Description`**: one short paragraph on what the tutorial covers, linking to the neighboring tutorials it does *not* cover.
3. **`### Rules`** *(optional)*: constraints for this tutorial.
4. **`---`**: horizontal rule separating the header from the workflow.
5. **`## Workflow`**: numbered, actionable steps, with fenced code blocks where a step involves code, and links to other tutorials for any step that is itself a separate goal. A multi-goal page carries one `##` section per goal instead, each holding its own `### Rules` *(optional)* and `### Workflow`.
6. **`## Full Code`** *(required when the workflow assembles a program across steps)*: the complete resulting code in a single fenced block, ready to copy and run — see [ParseOption.md](/docs/Tutorials/ParseOption.md) for a concrete example.

> **Note**: For a concrete example, refer to [sample.md](/docs/References/Specs/TutorialDocs/sample.md).
