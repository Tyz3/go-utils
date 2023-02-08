package logger

import (
	"fmt"
	"go-utils/colors"
	"os"
	"strings"
)

var (
	Log = &log{
		INFO:  &info{prefix: colors.GREEN + "INFO" + colors.RESET},
		WARN:  &warn{prefix: colors.YELLOW + "WARN" + colors.RESET},
		ERROR: &err{prefix: colors.RED + "ERROR" + colors.RESET},

		writer: os.Stdout,
	}
	Memo = &memorandum{
		writer: os.Stdout,
	}
)

const DateTimeLayout = "2006-01-02 15:04:05 Z07"

func anyJoinToStr(sep string, a ...any) string {
	var sb strings.Builder
	for i := 0; i < len(a); i++ {
		sb.WriteString(fmt.Sprint(a[i]))
		if i != len(a)-1 {
			sb.WriteString(sep)
		}
	}

	return sb.String()
}
