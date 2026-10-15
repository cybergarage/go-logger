// Copyright (C) 2026 The go-logger Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package logtest

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"sync"
	"testing"

	"github.com/cybergarage/go-logger/log"
	"github.com/cybergarage/go-logger/log/fmt"
	"github.com/cybergarage/go-logger/log/hexdump"
)

// syncBuffer is a bytes.Buffer the test reads while loggers write.
type syncBuffer struct {
	mu  sync.Mutex
	buf bytes.Buffer
}

func (b *syncBuffer) Write(p []byte) (int, error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.buf.Write(p)
}

func (b *syncBuffer) String() string {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.buf.String()
}

func TestWriterLoggerWritesToEveryWriter(t *testing.T) {
	t.Cleanup(func() { log.SetDefault(nil) })
	var a, b syncBuffer
	log.SetDefault(log.NewWriterLogger(io.MultiWriter(&a, &b), log.LevelInfo))
	if log.Infof("first\nsecond") <= 0 {
		t.Fatal("nothing written")
	}
	log.Debugf("below the level")
	for _, w := range []*syncBuffer{&a, &b} {
		lines := strings.Split(strings.TrimSuffix(w.String(), log.LF), log.LF)
		if len(lines) != 2 || !strings.HasSuffix(lines[0], "[INFO] first") || !strings.HasSuffix(lines[1], "[INFO] second") {
			t.Fatalf("unexpected lines: %q", lines)
		}
	}
}

func TestTimeLayoutMilliseconds(t *testing.T) {
	t.Cleanup(func() { log.SetDefault(nil) })
	var w syncBuffer
	log.SetDefault(log.NewWriterLogger(&w, log.LevelDebug, log.WithTimeLayout(log.TimeLayoutMilliseconds)))
	log.Infof("hello")
	line := strings.TrimSuffix(w.String(), log.LF)
	if !regexp.MustCompile(`^\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}\.\d{3} \[INFO\] hello$`).MatchString(line) {
		t.Fatalf("no milliseconds: %q", line)
	}
	if !regexp.MustCompile(fmt.LogPrefixRegex).MatchString(line) {
		t.Fatalf("LogPrefixRegex does not match %q", line)
	}
}

func TestHexdumpDecodesMillisecondLogs(t *testing.T) {
	t.Cleanup(func() { log.SetDefault(nil) })
	var w syncBuffer
	log.SetDefault(log.NewWriterLogger(&w, log.LevelDebug, log.WithTimeLayout(log.TimeLayoutMilliseconds)))
	msg := []byte("The quick brown fox jumps over the lazy dog")
	log.HexDebug(msg)
	lines := strings.Split(strings.TrimSuffix(w.String(), log.LF), log.LF)
	decoded, err := hexdump.DecodeHexdumpLogs(lines)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(decoded, msg) {
		t.Fatalf("%q != %q", decoded, msg)
	}
}

func TestOpenFileLogger(t *testing.T) {
	t.Cleanup(func() { log.SetDefault(nil) })
	path := filepath.Join(t.TempDir(), "test.log")
	if err := os.WriteFile(path, []byte("earlier"+log.LF), 0o644); err != nil {
		t.Fatal(err)
	}
	logger, err := log.OpenFileLogger(path, log.LevelInfo)
	if err != nil {
		t.Fatal(err)
	}
	log.SetDefault(logger)
	log.Infof("appended")
	if err := logger.Close(); err != nil {
		t.Fatal(err)
	}
	if n := log.Infof("after close"); n != 0 {
		t.Fatalf("%d bytes written after Close", n)
	}
	if err := logger.Close(); err != nil {
		t.Fatal("a second Close must be a no-op")
	}
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.HasPrefix(string(b), "earlier"+log.LF) || !strings.Contains(string(b), "[INFO] appended") || strings.Contains(string(b), "after close") {
		t.Fatalf("unexpected file: %q", b)
	}
	if runtime.GOOS != "windows" {
		info, err := os.Stat(path)
		if err != nil {
			t.Fatal(err)
		}
		if info.Mode().Perm() != 0o600 {
			t.Fatalf("file mode %v, want 0600", info.Mode().Perm())
		}
	}
}

func TestFileLoggerFileMode(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("no Unix file modes")
	}
	t.Cleanup(func() { log.SetDefault(nil) })
	dir := t.TempDir()
	for _, c := range []struct {
		name string
		opts []log.Option
		want os.FileMode
	}{
		{"default", nil, 0o644},
		{"private", []log.Option{log.WithFileMode(0o600)}, 0o600},
	} {
		path := filepath.Join(dir, c.name+".log")
		log.SetDefault(log.NewFileLogger(path, log.LevelInfo, c.opts...))
		log.Infof("hello")
		info, err := os.Stat(path)
		if err != nil {
			t.Fatal(err)
		}
		// The umask may only take permissions away.
		if got := info.Mode().Perm(); got&^c.want != 0 || got&0o600 != 0o600 {
			t.Fatalf("%s: file mode %v, want %v", c.name, got, c.want)
		}
	}
}

func TestSetLevelWhileLogging(t *testing.T) {
	t.Cleanup(func() { log.SetDefault(nil) })
	var w syncBuffer
	logger := log.NewWriterLogger(&w, log.LevelInfo)
	log.SetDefault(logger)
	var wg sync.WaitGroup
	for range 4 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 100 {
				log.Debugf("debug")
				log.Infof("info")
			}
		}()
	}
	for i := range 100 {
		if i%2 == 0 {
			logger.SetLevel(log.LevelDebug)
		} else {
			logger.SetLevel(log.LevelInfo)
		}
		_ = logger.Level()
	}
	wg.Wait()
	if !strings.Contains(w.String(), "[INFO] info") {
		t.Fatal("nothing logged")
	}
}
