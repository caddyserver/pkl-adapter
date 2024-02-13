Pkl Config Adapter for Caddy
================================

This plugin allows you to use [Pkl](https://pkl-lang.org) to configure Caddy.

**REQUIREMENTS:** This package uses Apple's official Go bindings, which spawns a child process, so the `pkl` command must be [installed](https://pkl-lang.org/go/current/quickstart.html) on your system.

## Install

Install with [xcaddy](https://github.com/caddyserver/xcaddy).

```
xcaddy build \
    --with github.com/caddyserver/pkl-adapter
```

## Usage

Specify with the `--adapter` flag for `caddy run`.

```
caddy run --config /path/to/pkl/config.pkl --adapter pkl
```
