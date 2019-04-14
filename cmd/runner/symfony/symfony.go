package symfony

import (
	"encoding/xml"
	"io/ioutil"
)

type Symfony struct {
	CommandList []string
	CommandMap  map[string]Command
}

func (s *Symfony) GetCommand(commandName string) Command {
	return s.CommandMap[commandName]
}

func (s *Symfony) Contains(commandName string) bool {
	for _, name := range s.CommandList {
		if name == commandName {
			return true
		}
	}
	return false
}

func NewSymfony(path string) *Symfony {
	xmlBytes := readXML(path)
	symfonyXML := parseSymfonyXML(xmlBytes)

	s := new(Symfony)
	s.CommandList = []string{}
	s.CommandMap = map[string]Command{}

	for _, command := range symfonyXML.Commands.Command {
		s.CommandMap[command.ID] = command
		s.CommandList = append(s.CommandList, command.ID)
	}

	return s
}

func readXML(path string) []byte {
	xmlBytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return xmlBytes
}

func parseSymfonyXML(xmlBytes []byte) *SymfonyXml {
	data := new(SymfonyXml)
	if err := xml.Unmarshal(xmlBytes, data); err != nil {
		panic(err)
	}

	return data
}
