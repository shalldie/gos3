package upload

import (
	"os"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/vbauerster/mpb/v7"
)

type CustomReader struct {
	fp      *os.File
	size    int64
	read    int64
	signMap map[int64]struct{}
	mux     sync.Mutex
	bar     *mpb.Bar
}

func (r *CustomReader) Read(p []byte) (int, error) {
	return r.fp.Read(p)
}

func (r *CustomReader) ReadAt(p []byte, off int64) (int, error) {
	n, err := r.fp.ReadAt(p, off)
	if err != nil {
		return n, err
	}

	r.mux.Lock()
	// Ignore the first signature call
	if _, ok := r.signMap[off]; ok {
		// Got the length have read( or means has uploaded), and you can construct your message
		r.read += int64(n)
		// fmt.Printf("\rtotal read:%d    progress:%d%%", r.read, int(float32(r.read*100)/float32(r.size)))
		r.bar.SetCurrent(r.read)
	} else {
		r.signMap[off] = struct{}{}
	}
	r.mux.Unlock()
	return n, err
}

func (r *CustomReader) Seek(offset int64, whence int) (int64, error) {
	return r.fp.Seek(offset, whence)
}

func uploadFile(file *os.File, options *UploadOptions, bar *mpb.Bar) {

	sess := session.Must(session.NewSession(&aws.Config{
		S3ForcePathStyle: aws.Bool(options.PATH_STYLE),
		Region:           aws.String(endpoints.ApSoutheast1RegionID),
		Endpoint:         aws.String(options.ENDPOINT),
		Credentials:      credentials.NewStaticCredentials(options.AK, options.SK, options.TOKEN),
	}))

	fileInfo, err := file.Stat()
	if err != nil {
		panic(err)
	}

	// bar.SetTotal(fileInfo.Size(), false)

	reader := &CustomReader{
		fp:      file,
		size:    fileInfo.Size(),
		signMap: map[int64]struct{}{},
		bar:     bar,
	}

	uploader := s3manager.NewUploader(sess, func(u *s3manager.Uploader) {
		u.PartSize = 5 * 1024 * 1024
		u.LeavePartsOnError = true
	})

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(options.BUCKET),
		Key:    aws.String(file.Name()),
		Body:   reader,
	})
	if err != nil {
		panic(err)
	}

}
