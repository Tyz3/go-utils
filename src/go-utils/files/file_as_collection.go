package files

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

type FileAsList struct {
	fileName    string
	updateEvery time.Duration
	closed      bool

	list []string

	sync.Mutex
}

func NewFileAsList(fileName string, updateEvery time.Duration) *FileAsList {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	file.Close()

	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	steamIDs := make(map[string]interface{})
	lines := strings.Split(string(data), "\n")
	for _, e := range lines {
		steamID := strings.TrimSpace(e)
		if steamID == "" {
			continue
		}

		steamIDs[steamID] = nil
	}

	f := &FileAsList{
		fileName:    fileName,
		updateEvery: updateEvery,
	}

	for k := range steamIDs {
		f.list = append(f.list, k)
	}

	return f
}

func (f *FileAsList) Close() {
	f.closed = true
	f.updateState()
}

func (f *FileAsList) Run() {
	for !f.closed {
		f.updateState()
		time.Sleep(20 * time.Second)
	}
}

func (f *FileAsList) updateState() {
	f.Lock()
	defer f.Unlock()
	file, err := os.OpenFile(f.fileName, os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0666)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		return
	}
	defer file.Close()

	var sb strings.Builder
	for _, line := range f.list {
		sb.WriteString(line)
		sb.WriteString("\n")
	}
	_, err = file.WriteString(sb.String())
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
	}
}

func (f *FileAsList) Pop() string {
	f.Lock()
	defer f.Unlock()

	pop := f.list[0]
	f.list = f.list[1:]

	return pop
}

func (f *FileAsList) Push(line string) {
	f.Lock()
	defer f.Unlock()

	f.list = append(f.list, line)
}

func (f *FileAsList) Exclude(removeSet map[string]interface{}) {
	f.Lock()
	defer f.Unlock()

	for k := range removeSet {
		for i, line := range f.list {
			if k == line {
				f.list = append(f.list[:i], f.list[i+1:]...)
				break
			}
		}
	}
}

func (f *FileAsList) Length() int {
	return len(f.list)
}
