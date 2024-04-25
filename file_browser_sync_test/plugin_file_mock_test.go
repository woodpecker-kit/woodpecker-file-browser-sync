package file_browser_sync_test

import (
	"fmt"
	"github.com/sinlov-go/unittest-kit/unittest_file_kit"
	"path/filepath"
)

func initTestDataForWorkspace(mockRootPath string) error {

	rootLevCnt := 3

	err := addTextFileByTry(mockRootPath, "data", "json", rootLevCnt)
	if err != nil {
		return err
	}

	err = addTextFileByTry(mockRootPath, "out", "apk", 1)
	if err != nil {
		return err
	}

	innerLev1JsonCnt := 5
	innerLev1Folder := filepath.Join(mockRootPath, "inner_1")
	err = addTextFileByTry(innerLev1Folder, "data", "json", innerLev1JsonCnt)
	if err != nil {
		return err
	}

	innerLev11JsonCnt := 4
	innerLev11TxtCnt := 3
	innerLev11Folder := filepath.Join(innerLev1Folder, "inner_1_1")
	err = addTextFileByTry(innerLev11Folder, "data", "json", innerLev11JsonCnt)
	if err != nil {
		return err
	}
	err = addTextFileByTry(innerLev11Folder, "log", "txt", innerLev11TxtCnt)
	if err != nil {
		return err
	}

	innerLev111JsonCnt := 4
	innerLev111TxtCnt := 3
	innerLev111Folder := filepath.Join(innerLev1Folder, "inner_1_1_1")
	err = addTextFileByTry(innerLev111Folder, "data", "json", innerLev111JsonCnt)
	if err != nil {
		return err
	}
	err = addTextFileByTry(innerLev111Folder, "log", "txt", innerLev111TxtCnt)
	if err != nil {
		return err
	}

	innerLev12JsonCnt := 4
	innerLev12TxtCnt := 3
	innerLev12Folder := filepath.Join(innerLev1Folder, "inner_1_2")
	err = addTextFileByTry(innerLev12Folder, "data", "json", innerLev12JsonCnt)
	if err != nil {
		return err
	}
	err = addTextFileByTry(innerLev12Folder, "log", "txt", innerLev12TxtCnt)
	if err != nil {
		return err
	}
	return nil
}

func addTextFileByTry(targetDir, fileHead, suffix string, cnt int) error {

	if !unittest_file_kit.PathExistsFast(targetDir) {
		err := unittest_file_kit.Mkdir(targetDir)
		if err != nil {
			return err
		}
	}

	var foo struct {
		Foo int    `json:"foo"`
		Bar string `json:"bar"`
	}

	for i := 0; i < cnt; i++ {
		fName := fmt.Sprintf("%s_%d.%s", fileHead, i, suffix)
		newJsonPath := filepath.Join(targetDir, fName)
		foo.Foo = i
		errJsonWrite := unittest_file_kit.WriteFileAsJsonBeauty(newJsonPath, foo, true)
		if errJsonWrite != nil {
			return errJsonWrite
		}
	}
	return nil
}
