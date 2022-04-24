package files

import (
	"fmt"
	"os"
	"strings"
)

func IsDir(dir string) bool {
	f, err := os.Stat(dir)
	if err != nil {
		return false
	}

	return f.IsDir()
}

func IsFile(filePath string) bool {
	f, err := os.Stat(filePath)
	if err != nil {
		return false
	}

	return !f.IsDir()
}

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	if err != nil {
		return false
	}

	return os.IsExist(err)
}

func SizeFormat(size float64) string {
	units := []string{"Byte", "KB", "MB", "GB", "TB"}
	n := 0
	for size > 1024 {
		size /= 1024
		n += 1
	}

	return fmt.Sprintf("%.2f %s", size, units[n])
}

func WriteString(path string, content string, append bool) error {
	flag := os.O_RDWR | os.O_CREATE
	if append {
		flag = flag | os.O_APPEND
	}

	file, err := os.OpenFile(path, flag, 0644)
	if err != nil {
		return err
	}

	defer func() {
		_ = file.Close()
	}()

	_, err = file.WriteString(content)

	return err
}

func WriteStringAppendLine(path string, content string) error {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}

	defer func() {
		_ = file.Close()
	}()

	content = strings.Join([]string{content, "\n"}, "")
	_, err = file.WriteString(content)

	return err
}
