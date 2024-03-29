package upload

import (
	"reflect"
	"strconv"

	"github.com/rivo/tview"
	"github.com/shalldie/gog/gs"
	"github.com/shalldie/gog/sortedmap"
)

func struct2SM(sender any) *sortedmap.SortedMap[string, string] {
	sm := sortedmap.New[string, string]()

	typeOpt := reflect.TypeOf(sender)

	for i := 0; i < typeOpt.NumField(); i++ {
		field := typeOpt.Field(i)
		sm.Set(field.Name, field.Type.Name())
	}

	return sm
}

func Setup() {
	var filePath string
	options := &UploadOptions{
		PATH_STYLE: true,
	}

	sm := struct2SM(*options)

	app := tview.NewApplication().EnableMouse(true)
	form := tview.NewForm()

	sm.ForEach(func(fieldName, fieldType string) {
		field := reflect.ValueOf(options).Elem().FieldByName(fieldName)

		if fieldType == "string" {
			form.AddInputField(fieldName, "", 40, nil, func(text string) {
				field.SetString(text)
			})
		} else if fieldType == "bool" {
			boolValueList := []bool{true, false}

			boolLabelList := gs.Map(boolValueList, func(b bool, i int) string {
				return strconv.FormatBool(b)
			})

			form.AddDropDown(fieldName, boolLabelList, gs.IndexOf(boolValueList, field.Bool()), func(option string, optionIndex int) {
				field.SetBool(boolValueList[optionIndex])
			})
		}

	})

	form.AddInputField("要上传的文件（逗号分隔）", "", 50, nil, func(text string) {
		filePath = text
	}).AddButton("开始上传", func() {
		app.Stop()
		uploadFilesByFilePath(filePath, options)
	}).
		AddButton("退出", func() {
			app.Stop()
		})

	form.SetBorder(true).
		SetTitle(" S3 Upload ").
		SetTitleAlign(tview.AlignLeft).
		SetRect(0, 0, 80, 19)

	if err := app.SetRoot(form, false).SetFocus(form).Run(); err != nil {
		panic(err)
	}
}
