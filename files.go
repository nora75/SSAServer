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
	err := checkFile(path)
	if err != nil {
		err = CreateUserDir(path)
		if err != nil {
			return fmt.Errorf("failed to open file: %s", err)
		}
	}
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
		return fmt.Errorf("failed to write file: %s", err)
	}
	return nil
}

// MoveUserDir move user directory in server
func MoveUserDir(oldpath string, newpath string) error {
	err := checkFile(oldpath)
	if err != nil {
		return fmt.Errorf("failed to open dir: %s", err)
	}
	err = os.Rename(oldpath,newpath)
	if err != nil {
		return fmt.Errorf("failed to move dir: %s", err)
	}
	return nil
}

// MoveUserDirToTmp move files in server
func MoveUserDirToTmp(path string) error {
	err := checkFile(path)
	if err != nil {
		return fmt.Errorf("failed to open dir: %s", err)
	}
	err = os.Rename(path,"/tmp/"+path)
	if err != nil {
		return fmt.Errorf("failed to move dir: %s", err)
	}
	return nil
}

// PickUpFile pick up file from server
func PickUpFile(name string) error {
	if err := checkFile(name); err != nil {
		return fmt.Errorf("failed to open file: %s", err)
	}
	return nil
}

// DeleteUserDir delete user dir
func DeleteUserDir(path string) error {
	err := checkFile(path)
	if err != nil {
		return fmt.Errorf("failed to open user dir: %s", err)
	}
	err = os.RemoveAll(path)
	if err != nil {
		return fmt.Errorf("failed to delete user dir: %s", err)
	}
	return nil
}

// DeleteGroupDir delete user dir
func DeleteGroupDir(path string) error {
	err := checkFile(path)
	if err != nil {
		return nil
	}
	err = os.Remove(path)
	if err != nil {
		return fmt.Errorf("failed to delete group dir: %s", err)
	}
	return nil
}

// GetUserDirPath return current dir path
func GetUserDirPath(gr string, id int) (path string) {
	path, _ = os.Getwd()
	slash := getSlash()
	path = path + slash + gr + slash + strconv.Itoa(id)
	return path
}

// GetPickUpPath return data's path	from group_name, user_id and data_name
func GetPickUpPath(gr string, id int, name string) (path string) {
	slash := getSlash()
	path = GetUserDirPath(gr,id) + slash + name
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

// checkFile is check exists file or dir
func checkFile(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return err
	}
	return nil
}

// CreateUserDir  is create new user dir
func CreateUserDir(path string) error {
	if err := os.MkdirAll(path, 0700); err !=  nil {
		return err
	}
	return nil
}
