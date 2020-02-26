# Contributing

Cool, that you are interested in contributing. We like pull request and suggestions from everyone.

If you are planning in making a major contribution we appreciate your opening an issue.

Some things that will increase the chance that your pull request is accepted:

Write clean code.
Write tests.
Write a good commit message.

We will *NOT* accept any pull requests that don't follow the Contributor Covenant [code of conduct](CODE_OF_CONDUCT.md).

## Naming conventions

### Commit Messages
To unify the appearance of all commit messages we only accept commit messages using the following principles:

- A commit contains the following, while the `<body-description>` is optional:
  ```
  <subject-line> #<issue-id>
  
  [<body-description>]
  ```
- Separate subject-line from optional body-description with a blank line
- Use the body-description to explain what, why or how within max 72 characters
- Limit the subject line to 50 characters and add the **GitHub Issue Number** with the `#`
- **Capitalize** the subject line
- Do not end the subject line with a full stop
- The subject line always uses the **imperative mood** and is able to **complete the following sentence**:
  > If applied, this commit will ...

#### Correct commits
- ... Add new function to do X #123
- ... Add test for X #123
- ... Refactor subsystem X for readability #123
- ... Update getting started documentation #123
- ... Remove deprecated methods #123
    
#### Wrong commits
- *adding new function to do X for Y* #123
- *added service X* #123
- *adds tests for X* #123
- *fixed bug with Y* #123
- *changing behavior of Y* #123
- *more fixes for broken stuff* #123
- *sweet new API methods* #123
