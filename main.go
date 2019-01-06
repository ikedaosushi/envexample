package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"time"

	"github.com/urfave/cli"
)

const defaultEnvPath = ".env"
const defaultEnvExamplePath = ".env.example"
const defaultFileMode = 0644

func Read(path string) string {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}

func Write(path string, content string) {
	err := ioutil.WriteFile(path, []byte(content), defaultFileMode)
	if err != nil {
		log.Fatal(err)
	}
}

func Run(filePath string, exampleFilePath string) {
	content := Read(filePath)
	// Note about regex: https://stackoverflow.com/questions/43723905/go-regex-to-match-all-lines-that-dont-start-with-timestamp
	rep := regexp.MustCompile(`(?m)\s*^([^=#]+=)([^=#\n]*).*$`)
	content = rep.ReplaceAllString(content, "\n$1")
	fmt.Println(string(content))
	Write(exampleFilePath, content)
}

func Action(c *cli.Context) error {
	filePath := c.Args().Get(0)
	if filePath == "" {
		filePath = defaultEnvPath
	}
	exampleFilePath := c.Args().Get(1)
	if exampleFilePath == "" {
		exampleFilePath = defaultEnvExamplePath
	}
	Run(filePath, exampleFilePath)
	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "envexample"
	app.Usage = "make .env.example"
	app.Version = "v0.0.1"
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "ikedaosushi",
			Email: "ikeda.yutaro@gmail.com",
		},
	}

	app.Action = Action

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
