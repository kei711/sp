package runner

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/kei711/symfony-console-commands-prompt/cmd/runner/symfony"
	"github.com/kei711/symfony-console-commands-prompt/config"
	"github.com/kei711/symfony-console-commands-prompt/util"
	"os"
	"os/exec"
)

var selectedCommand string
var verbose bool
var outputPath string

var s *symfony.Symfony

const (
	cmdExit       = "sp:exit"
	cmdQuit       = "sp:quit"
	cmdClearCache = "sp:clear-cache"
)

func Run(c string, v bool) {
	selectedCommand = c
	verbose = v

	saveCommandXML()
	loadSymfonyXML()

	printUsage()

	p := prompt.New(
		Executor,
		Completer,
		prompt.OptionPrefix(">>> "),
		prompt.OptionSelectedDescriptionBGColor(prompt.LightGray),
		prompt.OptionSelectedSuggestionBGColor(prompt.LightGray),
		prompt.OptionSelectedDescriptionTextColor(prompt.DarkGray),
		prompt.OptionSelectedSuggestionTextColor(prompt.DarkGray),
	)
	p.Run()
}

func printUsage() {
	fmt.Printf("  %-17s: Terminating.\n", cmdExit+"|"+cmdQuit)
	fmt.Printf("  %-17s: Regenerate Command Cache.\n", cmdClearCache)
}

func getCachePath() string {
	if outputPath != "" {
		return outputPath
	}

	cacheDir := config.GetCacheDir()

	d := md5.New()
	d.Write([]byte(selectedCommand))
	hash := hex.EncodeToString(d.Sum(nil))

	outputPath = cacheDir + "/" + hash + ".xml"

	return outputPath
}

func loadSymfonyXML() {
	s = symfony.NewSymfony(getCachePath())
}

func reloadSymfonyXML() {
	loadSymfonyXML()
	clearSuggestsCache()
}

func saveCommandXML() {
	if !util.FileExists(selectedCommand) {
		fmt.Println("selectedCommand not found.")
		return
	}

	cacheDir := config.GetCacheDir()
	if !util.FileExists(cacheDir) {
		os.Mkdir(cacheDir, 0755)
	}

	if !util.FileExists(getCachePath()) {
		cmd := exec.Command("php", selectedCommand, "list", "--format=xml")
		out, err := cmd.Output()
		if err != nil {
			panic(err)
		}

		file, err := os.OpenFile(getCachePath(), os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		fmt.Fprintln(file, string(out))
	}
}
