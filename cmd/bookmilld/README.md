<!-- category: Bookmill -->

# bookmilld

The bookmill daemon: a small web server holding the pipeline dashboard's
place in the cross-app navigation.

## Usage

```bash
bookmilld --addr :8082
```

bookmilld is one of the repo's daemons (with `arbiterd` and `plated`), built
by the root Makefile into `~/.local/share/trueblocks/bookmilld/` and run
under launchd rather than by hand. It serves:

- `/` — the Bookmill Dashboard page, currently a placeholder ("Pipeline
  status and controls will appear here"); the live per-cycle dashboard still
  belongs to the `pipeline` command itself.
- `/health` — `{"status":"ok"}` for the process monitor.
- the shared `/__tb__/` navigation assets, registered from `daemons.json`
  (`--apps-config`), so the app bar can hop between this and the other
  daemons' pages.

It shuts down cleanly on SIGINT/SIGTERM, giving in-flight requests ten
seconds. After a rebuild, restart it with launchd bootout + bootstrap, not
kickstart.

## Options

Run `bookmilld --help` for the full option list:

```
bookmilld dev
Web server for the bookmill pipeline dashboard.

Usage:
  bookmilld [flags]

Flags:
  --addr         listen address
  --apps-config  path to daemons.json for cross-app nav
  -v, --verbose  enable verbose (debug) logging
  -q, --quiet    suppress info logging (warnings/errors only)
  -h, --help     show this help and exit
      --version  print version and exit
```

## Example

```bash
curl http://127.0.0.1:8082/health
```

## Infographic

![How bookmilld holds the mill's place on the shared dashboard rail — a front page, a heartbeat, and the navigation joining it to the other daemons](README-infographic.jpg)
