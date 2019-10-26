package file

import (
	"bufio"
	"io/ioutil"
	"os"
	"strings"
)

func ReadFile(fName string) ([]string, error) {
	res := []string{}
	file, err := os.Open(fName)
	if err != nil {
		return res, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			continue
		}
		res = append(res, strings.TrimSpace(line))
		if err := scanner.Err(); err != nil {
			return []string{}, err
		}
	}
	return res, nil
}

func Mkdir(path string) error {
	return os.MkdirAll(path, 0777)
}

// 如果目录不存在, 自动创建目录, 如果存在同名文件, 会覆盖
func WriteFile(fileName string, content []string) error {

	if strings.Contains(fileName, "/") {
		dir := fileName[:strings.LastIndex(fileName, "/")]
		dir = strings.TrimSpace(dir)
		if dir != "" && dir != "./" && dir != "." && dir != "../" && dir != ".." {
			err := Mkdir(dir)
			if err != nil {
				return err
			}
		}
	}
	return ioutil.WriteFile(fileName, []byte(strings.Join(content, "\n")), 0666)
}
