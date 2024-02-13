Pkl Config Adapter for Caddy
================================

This plugin allows you to use [Pkl](https://pkl-lang.org) to configure Caddy.

**REQUIREMENTS:** This package uses Apple's official Go bindings, which spawns a child process, so the `pkl` command must be [installed](https://pkl-lang.org/go/current/quickstart.html) on your system.

After plugging it in, run Caddy with `--adapter pkl --config ...` (specifying your `.pkl` file).
