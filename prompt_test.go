package prompt

import (
	"fmt"
	"testing"
)

func completer(in Document) []Suggest {
	s := make([]Suggest, 20)

	for i := 0; i < 20; i++ {
		s[i] = Suggest{Text: fmt.Sprintf("users %d", i)}
	}

	return FilterFuzzy(s, in.Text, false)
}

func ExitCheckerFn(key Key, in string, breakline bool) bool {
	switch key {
	case Escape:
		return true
	}
	return false
}

func TestPrompt_Input(t *testing.T) {
	in := Input(">>> ", completer,
		OptionTitle("sql-prompt"),
		OptionPrefixTextColor(Yellow),
		OptionPrefixTextBold(true),
		OptionInputTextColor(Yellow),
		OptionInputTextBold(true),
		OptionCompletionOnDown(),
		OptionPreviewSuggestionTextColor(Blue),
		OptionSelectedSuggestionBGColor(LightGray),
		OptionSuggestionBGColor(DarkGray),
		OptionScrollbarThumbColor(Purple),
		OptionScrollbarBGColor(DarkGray),
		OptionSetExitCheckerOnInput(ExitCheckerFn),
	)
	fmt.Println("Your input: " + in)
}
