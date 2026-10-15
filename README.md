# go-logger

![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/cybergarage/go-logger)
[![test](https://github.com/cybergarage/go-logger/actions/workflows/make.yml/badge.svg)](https://github.com/cybergarage/go-logger/actions/workflows/make.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/cybergarage/go-logger.svg)](https://pkg.go.dev/github.com/cybergarage/go-logger)

`go-logger` is a small leveled logger for Go, shared by CyberGarage's Go projects such as [go-matter](https://github.com/cybergarage/go-matter) and [go-mdns](https://github.com/cybergarage/go-mdns). It depends on the standard library only.

- **Leveled, package-level functions**: `log.Infof`, `log.Debugf` and the like write to one default logger, and write nothing until one is set, so a library can log freely and leave the choice to the application.
- **Hex dumps**: `log.HexDebug` and the like write bytes as offset, hex and ASCII columns, one log line per 16 bytes, for protocol messages; the `hexdump` package turns such a log back into bytes, and the `hexdump2bin` and `hexdump2ascii` commands a hex dump file.
- **Outputs**: standard output, any `io.Writer`, a file, or a buffer.
- **Safe to reconfigure while logging**: lines are written one at a time, and the default logger and its level can be changed while other goroutines log.

```
2026-10-15 20:37:30.990 [INFO] listening on :5540
2026-10-15 20:37:30.990 [DEBUG] 0000 48 65 6C 6C 6F 2C 20 67   6F 2D 6C 6F 67 67 65 72     Hello, g o-logger
2026-10-15 20:37:30.990 [DEBUG] 0010 21 00 01 02                                           !...
```

## Install

```
go get github.com/cybergarage/go-logger
```

The hexdump commands are installed with:

```
go install github.com/cybergarage/go-logger/cmd/hexdump2bin@latest
go install github.com/cybergarage/go-logger/cmd/hexdump2ascii@latest
```

## Usage

The package-level functions write to the default logger, which `log.SetDefault` sets; nothing is written while it is nil.

```go
import "github.com/cybergarage/go-logger/log"

log.SetDefault(log.NewStdoutLogger(log.LevelInfo))
log.Infof("listening on %s", addr)
log.HexDebug(msg) // not written at the info level
```

`log.EnableStdoutDebug(true)` and `log.EnableStdoutTrace(true)` are shortcuts for tests.

### Levels

A logger writes the messages of its level and of every level above it:

| Level | Writes |
| --- | --- |
| `LevelTrace` | Everything |
| `LevelDebug` | Debug and above |
| `LevelInfo` | Info and above |
| `LevelWarn` | Warn and above |
| `LevelError` | Error and fatal |
| `LevelFatal` | Fatal only |

`log.NewLevelFromString("debug")` parses a level, and `Level.String()` names one. `Logger.SetLevel` changes it at any time.

### Loggers

| Logger | Writes each line to |
| --- | --- |
| `NewStdoutLogger(level)` | Standard output |
| `NewWriterLogger(w, level)` | Any `io.Writer`, such as an `io.MultiWriter` of a file and a terminal pane |
| `OpenFileLogger(path, level)` | A file opened once, created or appended to with mode 0600, until `Close` |
| `NewFileLogger(path, level)` | A file opened and closed for each line, so that it can be rotated, created with mode 0644 |
| `NewBufferLogger(buf, level)` | A `bytes.Buffer` |

Every constructor takes options:

- `WithTimeLayout(layout)` sets the timestamp layout; `TimeLayoutMilliseconds` adds milliseconds to the default `TimeLayoutSeconds`.
- `WithFileMode(mode)` sets the permissions of the file of a file logger.

```go
logger, err := log.OpenFileLogger("app.log", log.LevelDebug, log.WithTimeLayout(log.TimeLayoutMilliseconds))
if err != nil {
	return err
}
defer logger.Close()
log.SetDefault(logger)
```

### Reading hex dumps back

The `hexdump` package reads the bytes of hex dump log lines back, removing their timestamps and levels:

```go
import "github.com/cybergarage/go-logger/log/hexdump"

b, err := hexdump.DecodeHexdumpLogs(lines)
```

From a terminal, the commands read a file of hex dump lines, without the timestamps and levels of a log:

```
hexdump2bin <HEX FILE> <BIN FILE>       # writes the bytes
hexdump2ascii <HEX FILE> <OUTPUT FILE>  # writes their ASCII
```

## License

Apache License 2.0. See [LICENSE](LICENSE).
