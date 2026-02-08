# learn2code — Vision (v1)

learn2code is a minimal, opinionated collection of beginner-friendly Go exercises and a tiny CLI to run them. The goal is to provide:

- A clear learning path for newcomers to Go (focused, incremental exercises).
- A reproducible structure for exercises: template, tests, and a short README per exercise.
- A lightweight CLI to discover and run exercises locally.

Principles
- Learn by doing: each exercise contains a small task, a starter template, and tests that assert the expected behaviour.
- Automate feedback: tests are the primary feedback mechanism for learners and CI.
- Keep solutions out of public facing folders by default (gitignored `solutions/` directory).

Contents (v1)
- README.md — project vision and quick start
- exercises/ — individual exercises (template + tests)
- cmd/learn2code — CLI skeleton to run/list exercises

Quick start
1. Clone repository
2. cd learn2code-go
3. go test ./...

Contributing
See CONTRIBUTING.md for how to contribute and the CLA requirement.
