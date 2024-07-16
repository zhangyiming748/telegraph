package soup

import (
	"fmt"
	"github.com/schollz/progressbar/v3"
	"log"
	"os"
	"strconv"
	"strings"
	"telegraph/constant"
	"telegraph/util"
)

func GetFromWeb(uri string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("未捕获的错误", err)
		} else {
			fmt.Println("完成")
		}
	}()
	resp, err := GetWithProxy(uri, constant.GetProxy())
	root := HTMLParse(resp)
	if err != nil {
		fmt.Println("get 出现错误", err)
	}
	file, err := os.OpenFile("exam.html", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return
	}
	defer file.Close()
	file.WriteString(resp)
	title := root.Find("title").Text()
	log.Printf("标题:%v\n", title)
	title = strings.Join([]string{constant.GetRoot(), title}, string(os.PathSeparator))
	os.Mkdir(title, 0777)
	imgs := root.FindAll("img")
	bar := progressbar.New(len(imgs))
	defer bar.Finish()
	for i, img := range imgs {
		bar.Set(i)
		src := img.Attrs()["src"]
		src = strings.Join([]string{"https://telegra.ph", src}, "")
		//fmt.Println(src)
		fname := strings.Join([]string{title, string(os.PathSeparator), strconv.Itoa(i), ".jpg"}, "")
		util.HttpGetProxy(constant.GetProxy(), src, fname)
	}
}
