package cmd

import (
	"fmt"
	"log/slog"
	"path"
	"slices"
	"strings"

	"github.com/spf13/cobra"
	"github.com/xuri/excelize/v2"
)

type Result struct {
	Headers []string
	Values  []string
}

var file string

var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: parseJson,
}

func init() {
	parseCmd.Flags().StringVarP(&file, "file", "f", "", "Json file to parse")
	rootCmd.AddCommand(parseCmd)
}

func parseJson(cmd *cobra.Command, args []string) {
	if file == "" {
		fmt.Println("You must specify a file to parse")
		return
	}

	extension := path.Ext(file)
	if extension != ".json" {
		fmt.Println("The file must be json")
		return
	}

	fmt.Println(file)
}

func createFile() {
	file := excelize.NewFile()
	defer func() {
		if err := file.Close(); err != nil {
			slog.Error("Error", err)
		}
	}()

	index, err := file.NewSheet("Sheet1")
	if err != nil {
		slog.Error("Error", err)
		return
	}
	file.SetActiveSheet(index)
	if err := file.SaveAs("json.xlsx"); err != nil {
		fmt.Println(err)
	}
}

func parseString(str string) Result {
	res := Result{
		Headers: []string{},
		Values:  []string{},
	}
	parsedArr := []string{}
	arrStr := strings.Split(str, ",")
	replacer := strings.NewReplacer(
		"[", "",
		"]", "",
		"{", "",
		"}", "",
		"\"", "",
	)

	for _, item := range arrStr {
		modifiedString := replacer.Replace(item)
		s := strings.Split(modifiedString, ":")

		for i, subString := range s {
			subString = strings.TrimSpace(subString)
			parsedArr = append(parsedArr, subString)

			if i == 0 && !slices.Contains(res.Headers, subString) {
				res.Headers = append(res.Headers, subString)
			}

			if i == 1 {
				res.Values = append(res.Values, subString)
			}
		}
	}

	return res
}
