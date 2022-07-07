package upload

import (
	"fmt"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/manifoldco/promptui"
	"github.com/shalldie/gos3/internal/tool"
)

func printOptionTable() {

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"字段名", "字段类型"})

	fieldMap := tool.Struct2TypeTuples(UploadOptions{})

	fieldMap.ForEach(func(k, v string) {
		t.AppendRow([]any{k, v})
	})

	// t.SetStyle(table.StyleColoredBright)
	t.Render()

}

func Setup() {

	printOptionTable()

	pro := promptui.Select{
		Label: "选择【.env 文件】或者【自己输入】配置 ",
		Items: []string{".env 文件", "自己输入"},
	}
	_, result, err := pro.Run()

	if err != nil {
		panic(err)
	}

	var options *UploadOptions

	if result == ".env 文件" {
		options = ParseOptionsFromDotEnv(true)
	} else {
		options = ParseOptionsFromUI()
	}

	pro2 := promptui.Prompt{
		Label: "选择要上传的文件（空格分隔）",
	}
	content, _ := pro2.Run()

	for _, filePath := range strings.Split(content, " ") {
		filePath = strings.Trim(filePath, " ")
		file, err := os.Open(filePath)
		if err != nil {
			panic(err)
		}

		fmt.Println("uploading......" + filePath)
		UploadFile(file, options)
	}
	fmt.Println("Upload Done")

}
