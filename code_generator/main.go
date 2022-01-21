package main

import (
	"github.com/urfave/cli"
	"io/ioutil"
	"os"
	"strings"
)

var filePath string

func main() {
	app := cli.NewApp()
	app.Name = "XimCodeGenerator"
	app.Usage = "为使用Xim框架的程序生成代码"
	app.Author = "Kogic"
	app.Version = "0.0.1"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "path, p",
			Usage:       "需要生成的源文件的路径",
			Value:       "",
			Destination: &filePath,
		},
	}
	app.Action = func(c *cli.Context) error {
		if filePath == "" {
			os.Exit(2)
		}
		return Process(filePath)
	}
	err := app.Run(os.Args)
	if err != nil {
		os.Exit(1)
	}
}

func Process(filePath string) error {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	lines := strings.Split(string(bytes), "\n")
	marksLineIndex := FindMarks(lines)
	err = AnalyzeFuncSignature(marksLineIndex)
	if err != nil {
		return err
	}
	return nil
}

func FindMarks(lines []string) (marksLineIndex []int) {
	for lineIndex, line := range lines {
		strIndex := strings.Index(line, "//xim:handler")
		if strIndex == -1 {
			continue
		}
		for _, char := range line[:strIndex] {
			if char != ' ' && char != '\t' {
				goto endForLines
			}
		}
		marksLineIndex = append(marksLineIndex, lineIndex)
	endForLines:
	}
	return marksLineIndex
}

//goland:noinspection GoUnusedParameter
func AnalyzeFuncSignature(marksLineIndex []int) error {
	return nil
}
