package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
)


var (
    CF string
    Phone string
    AreaCode string
)

var rootCmd = &cobra.Command{
    Use: "",
}

var versionCmd = &cobra.Command{
    Use: "version",
    Short: "show current version",
    Run: showVersion,
}


func Execute(){
    rootCmd.Execute()
}

func init(){
    rootCmd.Flags().StringVarP(&CF, "config", "c", "", "data file")
    rootCmd.Flags().StringVarP(&Phone, "phone", "p", "", "phone number")
    rootCmd.Flags().StringVarP(&AreaCode, "areacode", "a", "", "the phone number area code, example: +86")
    rootCmd.AddCommand(versionCmd)
}


func showVersion(cmd *cobra.Command, args []string) {
    fmt.Println("version 1.0.0")
}
