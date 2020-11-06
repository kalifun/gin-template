package utils

import (
	"os"
)

func CreateFile(filePath string) error {
	f, err := os.Create(filePath)
	defer f.Close()
	if err != nil {
		return err
	}
	return nil
}

func CreateDateDir(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// 必须分成两步
		// 先创建文件夹
		err := os.Mkdir(path, 0766)
		if err != nil {
			return err
		}
		// 再修改权限
		err1 := os.Chmod(path, 0766)
		if err1 != nil {
			return err1
		}
	}
	return nil
}

/*
文件是否存在
*/
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, err
	}
	return false, err
}