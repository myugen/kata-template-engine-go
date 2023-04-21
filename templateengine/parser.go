package templateengine

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/myugen/kata-template-engine-go/log"
)

type Dictionary = map[string]string

var (
	EmptyDictionary = make(map[string]string)
)

type Parser struct {
	logger log.Logger
}

func (p Parser) Parse(text string, variables Dictionary) (string, error) {
	result := text
	for _, placeholders := range p.findAllVPlaceholdersIn(text) {
		if value, ok := variables[placeholders]; ok {
			variableToSubstitute := fmt.Sprintf("${%s}", placeholders)
			result = strings.ReplaceAll(result, variableToSubstitute, value)
		} else {
			p.logger.Warn(fmt.Sprintf("variable '%s' is missing", placeholders))
		}
	}
	return result, nil
}

func (p Parser) findAllVPlaceholdersIn(text string) []string {
	variables := make([]string, 0)
	regex := regexp.MustCompile("\\${(.*)}")
	for _, foundGroup := range regex.FindAllStringSubmatch(text, -1) {
		variable := foundGroup[1]
		variables = append(variables, variable)
	}
	return variables
}

func NewParser(logger log.Logger) *Parser {
	return &Parser{logger: logger}
}
