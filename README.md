# go-logger

![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/cybergarage/go-logger)
[![test](https://github.com/cybergarage/go-logger/actions/workflows/make.yml/badge.svg)](https://github.com/cybergarage/go-logger/actions/workflows/make.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/cybergarage/go-logger.svg)](https://pkg.go.dev/github.com/cybergarage/go-logger)

`go-logger` is a common logger package for CyberGarage's go-lang projects.

## Usage

The package-level functions, such as `log.Infof` and `log.HexDebug`, write to the default logger, which `log.SetDefault` sets; nothing is written while it is nil.

```go
log.SetDefault(log.NewStdoutLogger(log.LevelInfo))
log.Infof("hello %s", "world")
```

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

Lines are written one at a time, so a writer needs no locking of its own, and `SetDefault` and `Logger.SetLevel` can be called while other goroutines log.
