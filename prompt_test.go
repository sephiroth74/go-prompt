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

func TestPrompt_Input(t *testing.T) {
	in := Input(">>> ", completer,
		OptionTitle("sql-prompt"),

		OptionPrefixTextColor(Yellow),
		OptionCompletionOnDown(),
		OptionPreviewSuggestionTextColor(Blue),
		OptionSelectedSuggestionBGColor(LightGray),
		OptionSuggestionBGColor(DarkGray),
		OptionScrollbarThumbColor(Purple),
		OptionScrollbarBGColor(DarkGray),
	)
	fmt.Println("Your input: " + in)
}
