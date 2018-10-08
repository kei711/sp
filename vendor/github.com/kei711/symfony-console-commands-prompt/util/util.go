package util

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/mitchellh/go-homedir"
	"github.com/renstrom/fuzzysearch/fuzzy"
	"os"
	"sort"
)

func ArrayContains(a []string, sub string) bool {
	for _, str := range a {
		if str == sub {
			return true
		}
	}
	return false
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

var homeDirCache = ""

func GetHomeDir() string {
	if homeDirCache != "" {
		return homeDirCache
	}

	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	homeDirCache = home

	return homeDirCache
}

func FilterFuzzy(completions []prompt.Suggest, sub string) []prompt.Suggest {
	if sub == "" {
		return completions
	}

	texts := make([]string, 0, len(completions))
	descMap := make(map[string]string)
	for _, v := range completions {
		texts = append(texts, v.Text)
		descMap[v.Text] = v.Description
	}

	matches := fuzzy.RankFind(sub, texts)
	sort.Sort(matches)

	ret := make([]prompt.Suggest, 0, len(matches))
	for _, v := range matches {
		ret = append(ret, prompt.Suggest{
			Text: v.Target,
			Description: descMap[v.Target],
		})
	}

	return ret
}
