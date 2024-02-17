package oas

import (
	"go/ast"
	"testing"

	"github.com/stretchr/testify/assert"
)

type parseInfosTestCase struct {
	description         string
	gofiles             []*ast.File
	expectedVersion     string
	expectedTitle       string
	expectedDescription string
}

func Test_validatePath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Vendor path",
			args{
				path: "/foo/bar/test/vendor/test.go",
			},
			false,
		},
		{
			"No go file path",
			args{
				path: "/foo/bar/test/test.py",
			},
			false,
		},
		{
			"Dot File",
			args{
				path: ".DS_STORE",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validatePath(tt.args.path, []string{}); got != tt.want {
				t.Errorf("validatePath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseInfos(t *testing.T) {
	commentsWithoutInfoTag := &ast.CommentGroup{
		List: []*ast.Comment{
			{Text: "// Some useless comment"},
			{Text: "// Another useless comment"},
		},
	}

	commentsWithInfoTagOnly := &ast.CommentGroup{
		List: []*ast.Comment{
			{Text: "// @openapi:info"},
			{Text: "// Another useless comment"},
		},
	}

	commentsWithInfoTagAndData := &ast.CommentGroup{
		List: []*ast.Comment{
			{Text: "// @openapi:info"},
			{Text: `// version: "1.0.1"`},
			{Text: `// title: "some cool title"`},
			{Text: `// description: "some cool description tho"`},
		},
	}

	commentsWithInfoDataOnly := &ast.CommentGroup{
		List: []*ast.Comment{
			{Text: `// version: "1.0.1"`},
			{Text: `// title: "some cool title"`},
			{Text: `// description: "some cool description tho"`},
		},
	}

	commentsWithInfoTagAndDataDuplicated := &ast.CommentGroup{
		List: []*ast.Comment{
			{Text: "// @openapi:info"},
			{Text: `// version: "1.4"`},
			{Text: `// title: "another cool title"`},
			{Text: `// description: "another cool description tho"`},
		},
	}

	testCases := []parseInfosTestCase{
		{
			description: "Parse comments not containing info tags shouldn't change Info data",
			gofiles: []*ast.File{
				{Comments: []*ast.CommentGroup{commentsWithoutInfoTag}},
			},
			expectedVersion:     "",
			expectedTitle:       "",
			expectedDescription: "",
		},
		{
			description: "Parse comments containing only info tag shouldn't change Info data",
			gofiles: []*ast.File{
				{Comments: []*ast.CommentGroup{commentsWithInfoTagOnly}},
			},
			expectedVersion:     "",
			expectedTitle:       "",
			expectedDescription: "",
		},
		{
			description: "Parse comments containing info tag and infos data should set Info data",
			gofiles: []*ast.File{
				{Comments: []*ast.CommentGroup{commentsWithInfoTagAndData}},
			},
			expectedVersion:     "1.0.1",
			expectedTitle:       "some cool title",
			expectedDescription: "some cool description tho",
		},
		{
			description: "Parse comments containing only infos data shouldn't change Info data",
			gofiles: []*ast.File{
				{Comments: []*ast.CommentGroup{commentsWithInfoDataOnly}},
			},
			expectedVersion:     "",
			expectedTitle:       "",
			expectedDescription: "",
		},
		{
			description: "Parse comments containing info tag and duplicated infos data should set Info data with the first values parsed",
			gofiles: []*ast.File{
				{Comments: []*ast.CommentGroup{commentsWithInfoTagAndData}},
				{Comments: []*ast.CommentGroup{commentsWithInfoTagAndDataDuplicated}},
			},
			expectedVersion:     "1.0.1",
			expectedTitle:       "some cool title",
			expectedDescription: "some cool description tho",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			spec := NewOpenAPI()
			for _, gofile := range tc.gofiles {
				spec.parseInfos(gofile)
			}
			assert.Equal(t, tc.expectedVersion, spec.Info.Version)
			assert.Equal(t, tc.expectedTitle, spec.Info.Title)
			assert.Equal(t, tc.expectedDescription, spec.Info.Description)
		})
	}
}

func Test_parseImportContentPath(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test correct string",
			args: args{
				"import(readme.md)",
			},
			want:    "readme.md",
			wantErr: false,
		},
		{
			name: "test not an import file",
			args: args{
				"some random content for the test",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseImportContentPath(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseImportContentPath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("parseImportContentPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
