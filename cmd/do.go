package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/saurabh-sikchi/ytask/db"

	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "mark a task completed",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, n := range args {
			id, err := strconv.Atoi(n)
			if err != nil {
				fmt.Printf("Failed to parse: %s\n", n)
			} else {
				ids = append(ids, id)
			}
		}

		tasks, err := db.AllTasks()

		if err != nil {
			fmt.Println("Something went wrong: ", err.Error())
			os.Exit(1)
		}

		for _, id := range ids {
			if id > len(tasks) || id <= 0 {
				fmt.Printf("Invalid task %d\n", id)
				continue
			}

			task := tasks[id-1]
			err := db.DeleteTask(task.Key)
			if err != nil {
				fmt.Printf("Could not mark as done task \"%d\". Error: %s\n", id, err.Error())
			} else {
				fmt.Printf("marked as done task \"%d\"\n", id)
			}
		}

	},
}

func init() {
	RootCmd.AddCommand(doCmd)

}
