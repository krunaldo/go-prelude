package fs

import (
    "fmt"
    "os"
)

func EnsureDirOpen(path string) (*os.File, error) {
    // open it
    dirFile, err := os.OpenFile(path, os.O_RDONLY, os.ModeDir)
    if os.IsNotExist(err) {
        // create it
        err = os.Mkdir(path, os.ModeDir | 0755)
        dirFile, err = os.OpenFile(path, os.O_RDONLY, os.ModeDir)
    } else {
        // check that it is a directory
        stat, err := dirFile.Stat()
        if err != nil {
            dirFile.Close()
            return nil, err
        }
        if !stat.IsDir() {
            return nil, fmt.Errorf("%v is not a directory", dirFile)
        }
    }
    return dirFile, err
}

func EnsureDir(path string) (bool, error) {
    file, err := EnsureDirOpen(path)
    if file != nil {
        file.Close()
        return true, err
    } else {
        return false, err
    }
}
