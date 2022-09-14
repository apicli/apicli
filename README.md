# apicli

## Usage

### Not working on Windows IDE integrated terminals?

> Windows has introduced a new "pseudo console" API that allows a terminal to be attached to a process.
> Sadly, `ConPTY` does not support mouse interaction, yet. There are workaoounds, to re-enable mouse interactions in integrated terminals.
>
> If you use `apicli` in an external terminal (e.g. Windows Terminal, PowerShell, CMD, etc.), it works as expected, without the need to modify settings.

There are two ways to run apicli in an IDE integrated terminal:

1. Using `apicli` in WSL (Windows Subsystem for Linux)
  - WSL is an official Windows feature to emulate a Linux environment on Windows. See more here: https://docs.microsoft.com/en-us/windows/wsl/install
2. Disabling ConPTY in your IDE.
  - ConPTY is the new "pseudo console", which unfortunately does not support mouse interaction, yet.

#### Disabling ConPTY
- VSCode: Set `terminal.integrated.windowsEnableConpty` to `false` in your settings.
- IntelliJ: Disable the registry key `terminal.use.conpty.on.windows`. (`Help` -> `Find Action` -> `Registry...` -> `terminal.use.conpty.on.windows` -> `uncheck`)