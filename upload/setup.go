package upload

import (
	"reflect"

	"github.com/rivo/tview"
)

func struct2TupleList(sender any) [][]string {
	tuples := [][]string{}

	typeOpt := reflect.TypeOf(sender)

	for i := 0; i < typeOpt.NumField(); i++ {
		field := typeOpt.Field(i)
		tuples = append(tuples, []string{
			field.Name,
			field.Type.Name(),
		})
	}

	return tuples
}

func Setup() {
	var filePath string
	options := UploadOptions{
		PATH_STYLE: true,
	}

	tupleList := struct2TupleList(options)

	app := tview.NewApplication()
	form := tview.NewForm()

	for _, tuple := range tupleList {
		fieldName, fieldType := tuple[0], tuple[1]
		field := reflect.ValueOf(&options).Elem().FieldByName(fieldName)
		if fieldType == "string" {
			form.AddInputField(fieldName, "", 40, nil, func(text string) {
				field.SetString(text)
			})
		} else if fieldType == "bool" {
			boolLabelList := []string{"true", "false"}
			boolValueList := []bool{true, false}

			form.AddDropDown(fieldName, boolLabelList, 0, func(option string, optionIndex int) {
				field.SetBool(boolValueList[optionIndex])
			})
		}
	}

	form.AddInputField("要上传的文件（逗号分隔）", "", 50, nil, func(text string) {
		filePath = text
	}).AddButton("开始上传", func() {
		app.Stop()
		uploadFilesByFilePath(filePath, &options)
	}).
		AddButton("退出", func() {
			app.Stop()
		})

	form.SetBorder(true).SetTitle(" S3 upload ").SetTitleAlign(tview.AlignLeft)

	if err := app.SetRoot(form, true).SetFocus(form).Run(); err != nil {
		panic(err)
	}

}
