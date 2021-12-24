package s3upload

import (
	"fmt"
	"go-cli/tool"
	"reflect"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"github.com/manifoldco/promptui"
)

var OPT_KEYS = []string{"AK", "SK", "TOKEN", "BUCKET", "ENDPOINT", "PATH_STYLE"}

type UploadOptions struct {
	AK         string `env:"AK"`
	SK         string `env:"SK"`
	TOKEN      string `env:"TOKEN" envDefault:""`
	BUCKET     string `env:"BUCKET"`
	ENDPOINT   string `env:"ENDPOINT"`
	PATH_STYLE bool   `env:"PATH_STYLE" envDefault:"true"`
}

func ParseOptionsFromDotEnv(withUI bool) *UploadOptions {
	var fileName string = ".env"
	var err error

	if withUI {
		pro := promptui.Prompt{
			Label:   "配置环境变量的文件",
			Default: ".env",
		}

		fileName, err = pro.Run()
		if err != nil {
			panic(err)
		}
	}

	err = godotenv.Load(fileName)
	if err != nil {
		panic(err)
	}

	options := UploadOptions{}

	if err := env.Parse(&options, env.Options{RequiredIfNoDef: true}); err != nil {
		panic(err)
	}

	return &options
}

func ParseOptionsFromUI() *UploadOptions {
	options := UploadOptions{}
	typeTuples := tool.Struct2TypeTuple(options)

	for _, tuple := range typeTuples {
		fieldName, fieldType := tuple[0], tuple[1]
		field := reflect.ValueOf(&options).Elem().FieldByName(fieldName)
		if fieldType == "string" {
			pro := promptui.Prompt{
				Label: fmt.Sprintf("输入 %s，类型 %s", fieldName, fieldType),
			}
			result, _ := pro.Run()
			field.SetString(result)

		} else if fieldType == "bool" {
			label := fmt.Sprintf("选择 %s，类型 %s", fieldName, fieldType)
			fmt.Println(label)
			pro := promptui.Select{
				Label: label,
				Items: []string{"true", "false"},
			}
			_, result, _ := pro.Run()
			field.SetBool(result == "true")

		}
	}

	return &options
}
