package file

import (
	"io"
	"os"
)

func ReverseRead(fileName string, num int) (lines []string, err error) {
	file, err := os.Open(fileName)
	if err != nil {
		return
	}
	l := readLastLines(file, num)
	if len(l) < num {
		num = len(l)
	}
	for i := num; i > 0; i-- {
		lines = append(lines, l[i-1])
	}
	return lines, nil
}

func readLastLines(file *os.File, num int) []string {
	stats, statsErr := file.Stat()
	if statsErr != nil {
		return nil
	}

	size := stats.Size()
	lines := make([]string, 0, num)
	buf := make([]byte, size)
	for {
		readSize, err := file.ReadAt(buf, size-int64(len(buf)))
		if err != nil && err != io.EOF {
			break
		}

		size -= int64(readSize)
		for i := readSize - 1; i >= 0; i-- {
			if buf[i] == '\n' {
				lines = append(lines, string(buf[i+1:readSize]))
				readSize = i
				if len(lines) == num {
					return lines
				}
			}
		}
		if size == 0 {
			lines = append(lines, string(buf[:readSize]))
			break
		}
		if len(lines) >= num {
			break
		}
		if size < int64(len(buf)) {
			buf = make([]byte, size)
		}
	}
	return lines
}
