package manipulate

import (
	"net/http"
	"os"
)

var kMaxImageDimension = 3000
var kMinImageDimension = 30

// taken from https://golangcode.com/check-if-a-file-exists/
// fileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func validImageDimensions(width int, height int) bool {

	// reject anything wih negative values
	if width < 0 || height < 0 {
		return false
	}

	// reject anything above max, zero is allowed (means "not specified")
	if width > kMaxImageDimension || height > kMaxImageDimension {
		return false
	}

	// both specified: both must be at least min
	if width > 0 && height > 0 {
		return width >= kMinImageDimension && height >= kMinImageDimension
	}

	// exactly one specified: the specified one must be at least min (we checked for kMaxImageDimension above)
	if width > 0 {
		return width >= kMinImageDimension
	}
	if height > 0 {
		return height >= kMinImageDimension
	}

	// neither specified
	return true
}

func GetFileContentType(file *os.File) (string, error) {

	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)
	_, err := file.Read(buffer)
	if err != nil {
		return "", err
	}

	// Reset the read pointer if necessary.
	_, _ = file.Seek(0, 0)

	// Use the net/http package's handy DectectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)

	return contentType, nil
}
