package main

import (
	"os"
	"strings"
	"telegraph/constant"
	"telegraph/log"
	"telegraph/soup"
	"telegraph/util"
)

func main() {
	constant.SetRoot("C:\\Users\\zen\\Github\\telegraph")
	log.SetLog()
	fp := strings.Join([]string{constant.GetRoot(), "urls.txt"}, string(os.PathSeparator))
	urls := util.ReadByLine(fp)
	for _, url := range urls {
		soup.GetFromWeb(url)
	}
}
