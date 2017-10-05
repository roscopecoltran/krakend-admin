package download

import (
	"fmt"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func GetFilePathFromURL(remoteURL string) (string, string, string) {
	u, err := url.Parse(remoteURL)
	if err != nil {
		fmt.Printf("Invalid URL %s, skipping.\n", remoteURL)
		return "", "", ""
	}
	ext := path.Ext(u.Path)
	// fmt.Printf("ext: %s\n", ext)
	if IsSanitizeUrl {
		sanitizedName := SanitizeUrl(u.Scheme + u.Host + TrimSuffix(u.Path, ext) + "?" + u.RawQuery)
		fmt.Printf("sanitizedName: %s\n", sanitizedName)
	}
	filePath, err := GetPrefixPath(u.Path)
	if err != nil {
		fmt.Println("Error while getting the destination file prefix path, skipping:", err)
		return "", "", ""
	}

	fileName := strings.TrimRight(u.Path, "\n") // filePath + ext
	fullPath := strings.TrimRight(fmt.Sprintf("%s%s", DownloadDir, filePath), "\n")
	fullName := strings.TrimRight(fmt.Sprintf("%s%s", DownloadDir, fileName), "\n") // parentDir + fileName
	fmt.Println("GetFilePathFromURL(...): |- fileName: ", fileName, ", fullPath: ", fullPath, ", fullName: ", fullName)
	return fileName, fullPath, fullName
}

func GetPrefixPath(filePath string) (string, error) {
	prefixPath, err := filepath.Abs(filepath.Dir(filePath))
	if err != nil {
		return "", err
	}
	return prefixPath, nil
}

// Exists reports whether the named file or directory exists.
func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
