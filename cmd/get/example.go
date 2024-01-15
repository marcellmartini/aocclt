package get

import (
	"github.com/spf13/cobra"

	"github.com/marcellmartini/aocclt/pkg/puzzle"
)

var cmdExample = &cobra.Command{
	Use:   "example [options]",
	Short: "Download example text.",
	Long: `Download example text and save it into a txt file in the path 
indicate by -p. download the day (-d)  and the year (-y)
ex.:

aocclt get example -d 05 -y 2023 -p $(pwd)
`,
	// Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := "https://adventofcode.com/2023/day/1"
		puzzle.GetExample(url)
	},
}

func init() {
	// fileCmd.PersistentFlags().String("foo", "", "A help for foo")
	// fileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	cmdExample.Flags().IntVarP(&day, "day", "d", 1, "The day to download the puzzle example")

	CmdGet.AddCommand(cmdExample)
}
