# Demo: Intentional Vulnerable Dependency

This repository is a demonstration app that intentionally depends on a vulnerable version of a library for educational and testing purposes only.

- Vulnerable module added: `github.com/dgrijalva/jwt-go` v3.2.0

Notes:
- Do not use these versions or patterns in production.
- This is for local testing, scanning, or educational demos only.

Endpoints:
- `GET /token` - creates a short-lived demo JWT using the vulnerable library.

Use responsibly.
