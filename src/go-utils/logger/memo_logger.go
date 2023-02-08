package logger

import (
	"fmt"
	. "go-utils/colors"
	"io"
	"strings"
	"time"
)

type Memorandum interface {
	New(prefix ...any) MemoMessage
}

type memorandum struct {
	writer io.Writer
}

func (m *memorandum) New(prefix ...any) MemoMessage {
	date := time.Now().Format(DateTimeLayout)
	pref := anyJoinToStr(" ", prefix)

	memo := new(memoMessage)
	memo.Timestamps.start = time.Now().UnixMilli()
	memo.Timestamps.last = memo.Timestamps.start

	memo.message.WriteString(date)
	memo.message.WriteString(" ")
	memo.message.WriteString(WHITEBLACK)
	memo.message.WriteString(" ")
	memo.message.WriteString(pref)
	memo.message.WriteString(" ")
	memo.message.WriteString(RESET)

	return memo
}

func (m *memorandum) SetWriter(w io.Writer) {
	m.writer = w
}

type MemoMessage interface {
	Debug(format string, v ...any) MemoMessage
	Info(format string, v ...any) MemoMessage
	Warn(format string, v ...any) MemoMessage
	Error(format string, v ...any) MemoMessage
	Timestamp() MemoMessage
	Print()
	Ok()
	Failed()
}

type memoMessage struct {
	Timestamps struct {
		start int64
		last  int64
	}
	message strings.Builder
}

func (m *memoMessage) appendBlock(color string, spec string, format string, v ...any) MemoMessage {
	m.message.WriteString(" -> ")
	m.message.WriteString(color)
	m.message.WriteString(spec)
	m.message.WriteString("[")
	m.message.WriteString(RESET)
	m.message.WriteString(fmt.Sprintf(format, v...))
	m.message.WriteString(color)
	m.message.WriteString("]")
	m.message.WriteString(RESET)
	return m
}

func (m *memoMessage) Info(format string, v ...any) MemoMessage {
	return m.appendBlock(GREEN, "", format, v...)
}

func (m *memoMessage) Debug(format string, v ...any) MemoMessage {
	return m.appendBlock(CYAN, "~", format, v...)
}

func (m *memoMessage) Warn(format string, v ...any) MemoMessage {
	return m.appendBlock(YELLOW, "!", format, v...)
}

func (m *memoMessage) Error(format string, v ...any) MemoMessage {
	return m.appendBlock(RED, "?", format, v...)
}

func (m *memoMessage) Timestamp() MemoMessage {
	format := fmt.Sprintf("%.3fs", float64(time.Now().UnixMilli()-m.Timestamps.last)/1000)

	m.message.WriteString(" ")
	m.message.WriteString(DARKGREY)
	m.message.WriteString(format)
	m.message.WriteString(RESET)

	m.Timestamps.last = time.Now().UnixMilli()
	return m
}

func (m *memoMessage) appendStatus(color string, status string) {
	m.message.WriteString(" ")
	m.message.WriteString(color)
	m.message.WriteString(status)
	m.message.WriteString(RESET)
}

func (m *memoMessage) Print() {
	fmt.Fprintln(Memo.writer, m.message.String())
}

func (m *memoMessage) Ok() {
	m.appendStatus(GREEN, "OK")
	m.Print()
}

func (m *memoMessage) Failed() {
	m.appendStatus(RED, "FAILED")
	m.Print()
}
