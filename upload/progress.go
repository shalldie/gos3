package upload

import (
	"os"
	"strings"
	"sync"

	"github.com/shalldie/gog/gs"
	"github.com/vbauerster/mpb/v7"
	"github.com/vbauerster/mpb/v7/decor"
)

func formatNames(fileList []*os.File) []string {

	nameList := gs.Map(fileList, func(f *os.File, i int) string {
		return f.Name()
	})

	maxLen := len(gs.Sort(nameList, func(item1, item2 string) bool {
		return len(item1) > len(item2)
	})[0])

	return gs.Map(nameList, func(item string, index int) string {
		offLen := maxLen - len(item)
		for i := 0; i < offLen; i++ {
			item = item + " "
		}
		return item
	})
}

func uploadFilesByFilePath(filePath string, options *UploadOptions) {

	// 文件列表
	fileList := gs.Map(strings.Split(filePath, ","), func(fp string, i int) *os.File {
		file, err := os.Open(fp)

		if err != nil {
			panic(err)
		}

		return file
	})

	var wg sync.WaitGroup
	// passed wg will be accounted at p.Wait() call
	p := mpb.New(mpb.WithWaitGroup(&wg), mpb.WithWidth(38))

	wg.Add(len(fileList))

	fileNames := formatNames(fileList)

	gs.ForEach(fileList, func(f *os.File, i int) {
		fileInfo, err := f.Stat()
		if err != nil {
			panic(err)
		}
		bar := p.AddBar(fileInfo.Size(),
			mpb.PrependDecorators(
				// 名称
				decor.Name(fileNames[i]),
				// decor.DSyncWidth bit enables column width synchronization
				decor.Percentage(decor.WCSyncSpace),
			),
		)

		go func() {
			defer wg.Done()
			uploadFile(f, options, bar)
		}()
	})

	p.Wait()

}
