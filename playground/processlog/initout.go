package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

func initOut(logDir, commandName string) (stdout, stderr io.Writer, err error) {
	if logDir == "" {
		stdout = os.Stdout
		stderr = os.Stderr
	} else {
		ts := time.Now().Unix()

		// 定義する
		stdoutFileName := fmt.Sprintf("%s-%v-stdout.log", commandName, ts)
		// ファイルを指定のパスに作成
		stdoutFile, err := os.Create(filepath.Join(logDir, stdoutFileName))
		if err != nil {
			return nil, nil, err
		}
		// 書き込む
		stdout = io.MultiWriter(os.Stdout, stdoutFile)

		stderrFilename := fmt.Sprintf("%s-%v-stderr.log", commandName, ts)
		stderrFile, err := os.Create(filepath.Join(logDir, stderrFilename))
		if err != nil {
			return nil, nil, err
		}
		stderr = io.MultiWriter(os.Stderr, stderrFile)
	}
	return
}
