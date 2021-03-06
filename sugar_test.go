package log

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"net"
	"testing"
	"time"
)

func TestSugarLoggerNil(t *testing.T) {
	logger := Logger{
		Level: noLevel,
	}

	sugar := logger.Sugar(NewContext(nil).Str("tag", "hi nil sugar").Value())
	sugar.Print("hello from sugar Print")
	sugar.Println("hello from sugar Println")
	sugar.Printf("hello from sugar %s", "Printf")
	sugar.Debug("hello from sugar", "Debug")
	sugar.Debugf("hello from sugar %s", "Debugf")
	sugar.Debugw("a Debugw message", "foo", "bar", "number", 42)
	sugar.Info("hello from sugar", "Info")
	sugar.Infof("hello from sugar %s", "Infof")
	sugar.Infow("a Infow message", "foo", "bar", "number", 42)
	sugar.Warn("hello from sugar", "Warn")
	sugar.Warnf("hello from sugar %s", "Warnf")
	sugar.Warnw("a Warnw message", "foo", "bar", "number", 42)
	sugar.Error("hello from sugar", "Error")
	sugar.Errorf("hello from sugar %s", "Errorf")
	sugar.Errorw("a Errorw message", "foo", "bar", "number", 42)
	sugar.Fatal("hello from sugar", "Fatal")
	sugar.Fatalf("hello from sugar %s", "Fatalf")
	sugar.Fatalw("a Fatalw message", "foo", "bar", "number", 42)
	sugar.Panic("hello from sugar", "Panic")
	sugar.Panicf("hello from sugar %s", "Panicf")
	sugar.Panicw("a Panicw message", "foo", "bar", "number", 42)
	sugar.Log("foo", "bar", "number", 42)
}

func TestSugarLoggerPrintf(t *testing.T) {
	logger := Logger{
		Level:  InfoLevel,
		Caller: 1,
		Writer: &ConsoleWriter{ColorOutput: true, EndWithMessage: true},
	}

	sugar := logger.Sugar(NewContext(nil).Str("tag", "hi sugar").Value())
	logger.Level = InfoLevel
	sugar.Print("hello from sugar Print")
	sugar.Println("hello from sugar Println")
	sugar.Printf("hello from sugar %s", "Printf")

	sugar = logger.Sugar(NewContext(nil).Str("tag", "hi sugar").Value())
	sugar.Print("hello from sugar Print")
	sugar.Println("hello from sugar Println")
	sugar.Printf("hello from sugar %s", "Printf")
}

func TestSugarLoggerDebug(t *testing.T) {
	logger := Logger{
		Level:  DebugLevel,
		Caller: 1,
		Writer: &ConsoleWriter{ColorOutput: true, EndWithMessage: true},
	}

	sugar := logger.Sugar(NewContext(nil).Str("tag", "hi sugar").Value())
	sugar.Debug("hello from sugar", "Debug")
	sugar.Debugf("hello from sugar %s", "Debugf")
	sugar.Debugw("a Debugw message", "foo", "bar", "number", 42)
}

func TestSugarLoggerInfo(t *testing.T) {
	logger := Logger{
		Level:  InfoLevel,
		Caller: 1,
		Writer: &ConsoleWriter{ColorOutput: true, EndWithMessage: true},
	}

	sugar := logger.Sugar(NewContext(nil).Str("tag", "hi sugar").Value())
	sugar.Info("hello from sugar", "Info")
	sugar.Infof("hello from sugar %s", "Infof")
	sugar.Infow("a Infow message", "foo", "bar", "number", 42)
}

func TestSugarLoggerWarn(t *testing.T) {
	logger := Logger{
		Level:  InfoLevel,
		Caller: 1,
		Writer: &ConsoleWriter{ColorOutput: true, EndWithMessage: true},
	}

	sugar := logger.Sugar(NewContext(nil).Str("tag", "hi sugar").Value())
	sugar.Warn("hello from sugar", "Warn")
	sugar.Warnf("hello from sugar %s", "Warnf")
	sugar.Warnw("a Warnw message", "foo", "bar", "number", 42)
}

func TestSugarLoggerError(t *testing.T) {
	logger := Logger{
		Level:  InfoLevel,
		Caller: 1,
		Writer: &ConsoleWriter{ColorOutput: true, EndWithMessage: true},
	}

	sugar := logger.Sugar(NewContext(nil).Str("tag", "hi sugar").Value())
	sugar.Error("hello from sugar", "Error")
	sugar.Errorf("hello from sugar %s", "Errorf")
	sugar.Errorw("a Errorw message", "foo", "bar", "number", 42)
}

func TestSugarLoggerFatal(t *testing.T) {
	notTest = false

	logger := Logger{
		Level:  InfoLevel,
		Caller: 1,
		Writer: &ConsoleWriter{ColorOutput: true, EndWithMessage: true},
	}

	sugar := logger.Sugar(NewContext(nil).Str("tag", "hi sugar").Value())
	sugar.Fatal("hello from sugar", "Fatal")
	sugar.Fatalf("hello from sugar %s", "Fatalf")
	sugar.Fatalw("a Fatalw message", "foo", "bar", "number", 42)
}

func TestSugarLoggerPanic(t *testing.T) {
	notTest = false

	logger := Logger{
		Level:  InfoLevel,
		Caller: 1,
		Writer: &ConsoleWriter{ColorOutput: true, EndWithMessage: true},
	}

	sugar := logger.Sugar(NewContext(nil).Str("tag", "hi sugar").Value())
	sugar.Panic("hello from sugar", "Panic")
	sugar.Panicf("hello from sugar %s", "Panicf")
	sugar.Panicw("a Fatalw message", "foo", "bar", "number", 42)
}

func TestSugarLoggerLog(t *testing.T) {
	ipv4Addr, ipv4Net, err := net.ParseCIDR("192.0.2.1/24")
	if err != nil {
		t.Fatalf("net.ParseCIDR error: %+v", err)
	}

	logger := Logger{
		Level:  InfoLevel,
		Caller: 1,
		Writer: &ConsoleWriter{ColorOutput: true, EndWithMessage: true},
	}

	sugar := logger.Sugar(NewContext(nil).Str("tag", "hi sugar").Value())
	logger.Level = InfoLevel
	sugar.Log("foo", "bar")

	sugar = sugar.Level(DebugLevel)
	sugar.Log(
		"bool", true,
		"bools", []bool{false},
		"bools", []bool{true, false},
		"1_hour", time.Hour,
		"hour_minute_second", []time.Duration{time.Hour, time.Minute, time.Second},
		"error", errors.New("test error"),
		"an_error", fmt.Errorf("an %w", errors.New("test error")),
		"an_nil_error", nil,
		"dict", NewContext(nil).Str("foo", "bar").Int("no", 1).Value(),
		"float32", float32(1.111),
		"float32", []float32{1.111},
		"float32", []float32{1.111, 2.222},
		"float64", float64(1.111),
		"float64", []float64{1.111, 2.222},
		"int64", int64(1234567890),
		"int32", int32(123),
		"int16", int16(123),
		"int8", int8(123),
		"int", int(123),
		"uint64", uint64(1234567890),
		"uint32", uint32(123),
		"uint16", uint16(123),
		"uint8", uint8(123),
		"uint", uint(123),
		"raw_json", []byte("{\"a\":1,\"b\":2}"),
		"hex", []byte("\"<>?'"),
		"bytes1", []byte("bytes1"),
		"bytes2", []byte("\"<>?'"),
		"foobar", "\"\\\t\r\n\f\b\x00<>?'",
		"strings", []string{"a", "b", "\"<>?'"},
		"stringer_1", nil,
		"stringer_2", ipv4Addr,
		"gostringer_1", nil,
		"gostringer_2", binary.BigEndian,
		"now_1", timeNow(),
		"ip_str", ipv4Addr,
		"big_edian", binary.BigEndian,
		"ip6", net.ParseIP("2001:4860:4860::8888"),
		"ip4", ipv4Addr,
		"ip_prefix", *ipv4Net,
		"mac", net.HardwareAddr{0x00, 0x00, 0x5e, 0x00, 0x53, 0x01},
		"errors", []error{errors.New("error1"), nil, errors.New("error3")},
		"console_writer", ConsoleWriter{ColorOutput: true},
		"time.Time", timeNow(),
		"buffer", bytes.NewBuffer([]byte("a_bytes_buffer")),
		"message", "this is a test")
}
