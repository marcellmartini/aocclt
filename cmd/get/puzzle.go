package get

import (
	"github.com/spf13/cobra"

	"github.com/marcellmartini/aocclt/pkg/puzzle"
)

var day int

var puzzleCmd = &cobra.Command{
	Use:   "puzzle [options]",
	Short: "Download puzzle text.",
	Long: `Download puzzle text and save it into a txt file in the path 
indicate by -p. download the day (-d)  and the year (-y)
ex.:

aocclt get puzzle -d 05 -y 2023 -p $(pwd)
`,
	// Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := "https://adventofcode.com/2023/day/1"
		puzzle.GetPuzzle(url)
	},
}

func init() {
	// fileCmd.PersistentFlags().String("foo", "", "A help for foo")
	// fileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	puzzleCmd.Flags().IntVarP(&day, "day", "d", 1, "The day to download the puzzle")

	CmdGet.AddCommand(puzzleCmd)
}
