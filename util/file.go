package util

import "os"

func FileExists(path string) (bool, error) {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func RemoveIfExists(path string) error {
	exists, err := FileExists(path)
	if err != nil {
		return err
	}

	if exists {
		if err = os.Remove(path); err != nil {
			return err
		}
	}

	return nil
}

func IsDir(path string) (bool, error) {
	fileinfo, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return fileinfo.IsDir(), nil
}

func IsSymlink(path string) (bool, error) {
	fileInfo, err := os.Lstat(path)
	if err != nil {
		return false, err
	}

	return fileInfo.Mode()&os.ModeSymlink == os.ModeSymlink, nil
}
