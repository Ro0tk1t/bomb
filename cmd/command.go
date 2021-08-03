package cmd

import (
    "os"
    "fmt"
    "github.com/spf13/cobra"

    log "github.com/sirupsen/logrus"
)


var (
    CF string
    Delay int
    Phone string
    AreaCode string
)

var rootCmd = &cobra.Command{
    Use: "bomb",
    Args: cobra.NoArgs,
}

var versionCmd = &cobra.Command{
    Use: "version",
    Short: "show current version",
    Run: showVersion,
}


func Execute(){
    if err := rootCmd.Execute(); err != nil {
        log.Fatal(err)
    }
    // the function MarkFlagRequired looks like dosen't work,
    // so judg it by hand
    if Phone == "" {
        log.Fatal("phone is empty")
    }
}

func init(){
    rootCmd.Flags().StringVarP(&CF, "config", "c", "", "data file")
    rootCmd.Flags().IntVarP(&Delay, "delay", "d", 0, "delay bomb between requests")
    rootCmd.Flags().StringVarP(&Phone, "phone", "p", "", "phone number")
    rootCmd.Flags().StringVarP(&AreaCode, "areacode", "a", "", "the phone number area code, example: +86")
    rootCmd.AddCommand(versionCmd)

    rootCmd.MarkFlagRequired("phone")
}


func showVersion(cmd *cobra.Command, args []string) {
    fmt.Println("version 1.0.0")
    os.Exit(0)
}
