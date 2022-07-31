package upload

var OPT_KEYS = []string{"AK", "SK", "TOKEN", "BUCKET", "ENDPOINT", "PATH_STYLE"}

type UploadOptions struct {
	AK         string
	SK         string
	TOKEN      string
	BUCKET     string
	ENDPOINT   string
	PATH_STYLE bool
}
