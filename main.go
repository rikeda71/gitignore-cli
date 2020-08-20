package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func requestGitignore(options []string) ([]byte, error) {
	const GitignoreURL = "https://www.toptal.com/developers/gitignore/api/"

	resp, err := http.Get(GitignoreURL + strings.Join(options, ","))
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}
	return body, nil

}

func main() {
	fileSaveFlag := flag.Bool("f", false, "")
	flag.Parse()
	options := flag.Args()
	if len(options) == 0 {
		fmt.Println("invalid args")
		fmt.Println("./gitignore-cli arg1, arg2, ...")
		fmt.Println("Please set operating systems, IDEs, or programming languages to args")
		return
	}
	ignoreBytes, err := requestGitignore(options)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	if *fileSaveFlag {
		ioutil.WriteFile(".gitignore", ignoreBytes, 644)
		fmt.Println("write ignore list to `.gitignore`")
		return
	}
	fmt.Println(string(ignoreBytes))
}
