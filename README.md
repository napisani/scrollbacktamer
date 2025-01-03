# Scrollback Tamer

A command-line tool for efficiently managing and editing terminal scrollback buffer content.

## Overview

Scrollback Tamer allows you to capture, filter, and edit your terminal's scrollback buffer using your preferred text editor. It's particularly useful when you need to:

- Extract specific portions of your terminal history
- Edit and save command output
- Process terminal content in segments or lines
- Work with TMux scrollback content

Currently supports:
- TMux terminal multiplexer
- Various text editors (vim, emacs, nano, etc.)

## Building

To build from source:

```bash
make build
```

## Running

Basic usage:

```bash
scrollbacktamer [flags]
```

## CLI Usage

### Flags

- `--editor <EDITOR>` : Specify the editor to use (defaults to $EDITOR or auto-detected)
- `--file <FILEPATH>` : Read from a file instead of terminal scrollback
- `--last <X>` : Process only the last N units (lines or segments)
- `--units [segments|lines]` : Choose processing units ("lines" or "segments")
- `--terminator <REGEX>` : Regex pattern to split content into segments (required for segment mode)

### Environment Variables

All flags can be configured via environment variables with the `SBTAMER_` prefix:

- `SBTAMER_EDITOR` : Default editor
- `SBTAMER_FILE` : Default input file
- `SBTAMER_LAST` : Default number of units to process
- `SBTAMER_UNITS` : Default units type
- `SBTAMER_TERMINATOR` : Default segment terminator pattern

### Examples

1. Edit last 100 lines in vim:
```bash
# in a `tmux` pane 
scrollbacktamer --last 100 --editor vim
```

2. Process last 5 command segments using default editor:
```bash
# in a `tmux` pane
# open `neovim` with support for ANSI escape sequences and jump to the end of the file 
export SBTAMER_EDITOR='nvim +"term cat %s"  +"execute \":normal! G\""'
scrollbacktamer --units segments --last 5 --terminator ".*❯.*"
```

3. Edit content from a file:
```bash
# in any terminal emulator
scrollbacktamer --file scrollback_content.log
```
