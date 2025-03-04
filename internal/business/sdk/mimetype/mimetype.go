// Disclaimer: This package is copied from [go-mime](https://github.com/qingstor/go-mime/tree/master)

package mimetype

import (
	"path"
	"strings"
)

// defaultMIMEType will be returned while mime detect failed
const defaultMIMEType = "application/octet-stream"

// DetectFileExt will detect mime with file's extension name.
//
// Input string SHOULD NOT have leading "."
//
// Valid example: pdf, gz
func DetectFileExt(s string) string {
	if v, ok := extensionToMIME[s]; ok {
		return v
	}
	return defaultMIMEType
}

// DetectFilePath will detect mime with files' path.
func DetectFilePath(s string) string {
	ext := path.Ext(s)
	return DetectFileExt(strings.TrimPrefix(ext, "."))
}
