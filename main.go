package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// These variables are set in build step
const (
	Version  = "0.1.0"
	Revision = "unset"
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
	flag.Usage = func() {
		usageTxt := `usage: gitignore-cli [-f] [operating systems, IDEs, or programming languages]
		This command is get the gitignore list.
		gitignore is specifies intentionally untracked files to ignore.

		-f, --f : if this flag is true, save ignore list to .gitignore
		`
		fmt.Fprintf(os.Stderr, "%s\n", usageTxt)
	}
	fileSaveFlag := flag.Bool("f", false, "")
	flag.Parse()
	options := flag.Args()
	if len(options) == 0 {
		fmt.Println(
			`invalid args
./gitignore-cli arg1, arg2,
Please set operating systems, IDEs, or programming languages to args`)
		return
	}
	ignoreBytes, err := requestGitignore(options)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	if *fileSaveFlag {
		ioutil.WriteFile(".gitignore", ignoreBytes, 0644)
		fmt.Println("write ignore list to `.gitignore`")
		return
	}
	fmt.Println(string(ignoreBytes))
}
