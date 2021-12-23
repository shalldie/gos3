package s3upload

import (
	"fmt"
	"os"
	"strings"

	"github.com/manifoldco/promptui"
)

func Setup() {

	pro := promptui.Select{
		Label: "从 【自己输入】或者【环境变量】配置 ",
		Items: []string{"自己输入", "环境变量"},
	}
	_, result, err := pro.Run()

	if err != nil {
		panic(err)
	}

	var options *UploadOptions

	if result == "环境变量" {
		options = ParseOptionsFromEnv()
	} else {
		options = ParseOptionsFromUI()
	}

	pro2 := promptui.Prompt{
		Label: "选择要上传的文件（逗号分隔）",
	}
	content, _ := pro2.Run()

	for _, filePath := range strings.Split(content, ",") {
		file, err := os.Open(filePath)
		if err != nil {
			panic(err)
		}

		fmt.Println("uploading......" + filePath)
		UploadFile(file, options)
	}
	fmt.Println("Upload Done")

	println(options.AK, options.SK, options.PATH_STYLE)
}
