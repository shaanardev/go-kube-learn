# Go code quality checks

Pull requests that change Go code run [`golangci-lint`](https://golangci-lint.run/) through [reviewdog](https://github.com/reviewdog/reviewdog). The initial linter set is deliberately small and focuses on correctness:

- `govet` finds suspicious Go constructs.
- `staticcheck` catches likely bugs and ineffective code. Its style-only checks are left at their defaults rather than adding extra naming or comment rules.
- `errcheck` points out ignored errors.
- `ineffassign` and `unused` find assignments and declarations that do not affect the program.

`govet` also explicitly enables its `nilness` and `waitgroup` analyzers. Together, these checks catch some concrete pointer and concurrency mistakes: nil dereferences, copying lock-bearing values, lost context cancellation, suspicious `WaitGroup` usage, and certain mutex or test-goroutine errors. Staticcheck also has concurrency checks such as adding to a `WaitGroup` from inside the goroutine it is meant to track and suspicious lock patterns.

Reviewdog comments on added lines only. Findings are warnings and do not fail a pull request, so they can guide the next edit without turning CI into a grading gate. Tool setup and configuration errors still fail the workflow. The workflow skips analysis while the scaffold has no Go source files.

Run the same configured checks locally with:

```sh
golangci-lint run
```

When a finding is unclear, use its linter name to look up the rule and understand the underlying Go behavior before changing the code. If a finding is intentionally not applicable, add a narrow suppression with a reason, for example `//nolint:errcheck // closing this response body is best-effort during shutdown`. Avoid disabling a whole linter to silence one case.

The CI version is pinned so new upstream rules do not appear unexpectedly. Update it deliberately and review any new findings as part of that change.

Static analysis cannot prove that concurrent code is race-free or decide whether a struct's ownership and pointer/value design is right for the program. Once concurrency tests exist, run the relevant unit-test packages with `go test -race`; the race detector only catches races that the tests actually exercise. Keep cluster-dependent integration tests separate from the fast local test set.
