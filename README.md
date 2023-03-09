# ESbuild TS HTML Fixture

## Motivation

Sometimes a simple HTML+CSS+JS toolchain is needed for rapid experimentation.
However, speed comes at a cost, and [TypeScript](https://www.typescriptlang.org/) +
[esbuild](https://esbuild.github.io/) might help.

## Trying out the Process

### No Node Tooling Necessary (at first)

- prerequisite: Go installed
- `go run .`
- &rarr; http://localhost:8080/ui
- edit files in [ui-src](ui-src)
  - reload the browser and observe the result

### IDEs

In order for Typescript tooling to work in IDEs, run `npm i` to fetch type dependencies.

Reformat the code using `npm run prettier`.

## Using the Pattern

Simply, copy all non-ignored files from this repo into your project and continue from there.

## Background

Initial experiment: https://github.com/d-led/esbuild-test
