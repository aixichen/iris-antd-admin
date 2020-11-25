package libs

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
	"time"
)

// 当前目录
func CWD() string {
	// 兼容 travis 集成测试
	if os.Getenv("TRAVIS_BUILD_DIR") != "" {
		return os.Getenv("TRAVIS_BUILD_DIR")
	}

	path, err := os.Executable()
	path = "G:/go/car-tms/"
	if err != nil {
		return ""
	}
	return filepath.Dir(path)
}

func LogDir() string {
	dir := filepath.Join(CWD(), "logs")
	EnsureDir(dir)
	return dir
}

func UploadDir(perPath string) string {
	dir := filepath.Join(CWD(), Config.UploadDir, perPath)
	EnsureDir(dir)
	return dir
}

func EnsureDir(dir string) (err error) {
	if _, err = os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return
		}
	}
	return
}

func IsPortInUse(port int) bool {
	if conn, err := net.DialTimeout("tcp", net.JoinHostPort("", fmt.Sprintf("%d", port)), 3*time.Second); err == nil {
		conn.Close()
		return true
	}
	return false
}
