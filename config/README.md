# Config

This folder contains TOML configuration files. The intention is that to run this program in a specific environment, one of these is selected by renaming it to `config.toml` (eg. in the CICD pipeline) and then it will be loaded by the `internal/common/epconfig` package into a ServiceConfig struct.

In the event that no `config.toml` is found, ServiceConfig will load `local.toml`.
