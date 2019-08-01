package ssa

import (
	"os"
	"fmt"
	"runtime"
	"strconv"
)


// GetLogPath return log path
func GetLogPath() (string) {
	logpath, _ := os.Getwd()
	slash := getSlash()
	logpath += slash + "log" + slash + "ssa.log";
	return logpath
}

// SaveFile save file in server
func SaveFile(data []byte,path string, name string) error {
	name = path + getSlash() + name
	f, err := os.Create(name)
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()
	if err != nil {
		return fmt.Errorf("failed to open file: %s", err)
	}
	_, err = f.Write(data)
	if err != nil {
		return fmt.Errorf("failed to open file: %s", err)
	}
	return nil
}

// GetSavePath return current dir path
func GetSavePath(gr string, id int) (path string) {
	path, _ = os.Getwd()
	slash := getSlash()
	path = path + slash + gr + slash + strconv.Itoa(id)
	return path
}

// GetPickUpPath return data's path	from group_name, user_id and data_name
func GetPickUpPath(gr string, id int, name string) (path string) {
	path, _ = os.Getwd()
	slash := getSlash()
	path = path + slash + gr + slash + strconv.Itoa(id) + slash + name
	fmt.Println(path)
	return path
}

// GetSlash return backslash or slash on server's environment
func getSlash() (slash string) {
	if runtime.GOOS == "windows" {
		slash = "\\\\"
	} else {
		slash = "/"
	}
	return slash
}
