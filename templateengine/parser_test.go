package templateengine_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/myugen/kata-template-engine-go/log/mocks"
	"github.com/myugen/kata-template-engine-go/templateengine"
	"github.com/stretchr/testify/assert"
)

func TestParser_ParseShould(t *testing.T) {
	type args struct {
		text      string
		variables templateengine.Dictionary
	}
	type want struct {
		result string
		err    error
	}
	testcases := []struct {
		name string
		args args
		want want
	}{
		{
			name: "not replace any text if no variables",
			args: args{text: "Text without variables", variables: templateengine.EmptyDictionary},
			want: want{result: "Text without variables", err: nil},
		},
		{
			name: "not replace any text if no variables",
			args: args{
				text: "This is a text with a ${variable} to be replaced.\n" +
					"And this is another text with ${other-variable} to be replaced.\n" +
					"And this is another text with ${variable} to be replaced.",
				variables: templateengine.Dictionary{"variable": "value", "other-variable": "other-value"},
			},
			want: want{result: "This is a text with a value to be replaced.\n" +
				"And this is another text with other-value to be replaced.\n" +
				"And this is another text with value to be replaced.", err: nil,
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			parser := templateengine.Parser{}
			got, err := parser.Parse(testcase.args.text, testcase.args.variables)

			assert.Equal(t, testcase.want.err, err)
			assert.Equal(t, testcase.want.result, got)
		})
	}
}

func TestParser_ParseShouldWarnWhenVariableIsMissing(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLogger := mocks.NewMockLogger(ctrl)
	parser := templateengine.NewParser(mockLogger)

	mockLogger.EXPECT().Warn("variable 'other-variable' is missing").Times(1)

	text := "This is a text with a ${variable} to be replaced.\n" +
		"And this is another text with ${other-variable} to be replaced.\n" +
		"And this is another text with ${variable} to be replaced."
	variables := templateengine.Dictionary{"variable": "value"}

	got, err := parser.Parse(text, variables)

	assert.Equal(t, "This is a text with a value to be replaced.\n"+
		"And this is another text with ${other-variable} to be replaced.\n"+
		"And this is another text with value to be replaced.", got)
	assert.NoError(t, err)
}
