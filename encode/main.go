package main

import (
	"io/ioutil"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

var aviutilPath = "Path/To/aviutl.exe"
var encordingTargetDirectoryPath = "Path/To/encodingTargetFolder/"
var outputDirectoryPath = "Path/To/encordedFilesFolder/"
var aucLibPath = "Path/To/lib"
func main() {
	result, _ := exec.Command(aucLibPath + "auc_exec.exe", aviutilPath).Output()
/*
	result, e = exec.Command("../lib/auc_findwnd.exe").Output()
	if e != nil {
		println(e)
	}
*/
	files, err := ioutil.ReadDir(encordingTargetDirectoryPath)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		fileName := file.Name()
		if !regexp.MustCompile(".ts").Match([]byte(fileName)) || regexp.MustCompile(".ts.lwi").Match([]byte(fileName)) {// tsファイル以外をエンコード対象外にする
			continue
		}
		println("エンコード対象 : " + file.Name())
		encordFilePath := strings.Replace(encordingTargetDirectoryPath + fileName, "/", "\\", -1)
		_ = exec.Command(aucLibPath+"/auc_open.exe", string(result), encordFilePath).Run()
		time.Sleep(1 * time.Minute)
		println("end import " + file.Name())
		outputFileName := strings.Replace(fileName, ".ts", ".mp4", 1)
		outputFilePath := strings.Replace(outputDirectoryPath + outputFileName, "/", "\\", -1)
		println(outputFilePath)
		_ = exec.Command(aucLibPath + "auc_plugout.exe", string(result), "1", outputFilePath).Run()
		_ = exec.Command(aucLibPath + "auc_wait.exe", string(result)).Run()
		println("end of encoded" + outputFilePath)
	}
}
