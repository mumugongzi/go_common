package doc

import (
	"github.com/unidoc/unioffice/document"
	"strings"
)

// 读取doc文档中的纯文本
func ReadText(file string) ([]string, error) {
	doc, err := document.Open(file)
	var res []string
	if err != nil {
		return nil, err
	}

	for _, para := range doc.Paragraphs() {
		for _, run := range para.Runs() {
			if strings.TrimSpace(run.Text()) == "" {
				continue
			}
			res = append(res, run.Text())
		}
	}
	return res, err
}
