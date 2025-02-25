package cmd

import (
	"fmt"
	"os"

	"github.com/golrice/deadlink/internal/checker"
	"github.com/golrice/deadlink/internal/crawler"
	"github.com/golrice/deadlink/internal/models"
	"github.com/golrice/deadlink/internal/reporter"
	"github.com/golrice/deadlink/internal/utils"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "deadlink",
	Short: "check dead link",
	Long:  "check all links in a web whether is dead or not",
	Run:   rootRun,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	fmt.Println("Parsing input arguments")
}

func rootRun(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Println("we want only 1 start url")
		return
	}

	// user input -> crawler -> checker -> checkResult -> reporter
	u := args[0]
	webcrawler := crawler.NewCrawler(u, 2)
	totUrls, err := webcrawler.Crawl()
	if err != nil {
		utils.Error("crawling...", err)
	}

	timer := utils.StartTimer()
	webchecker := checker.NewLinkChecker()
	links := webchecker.Check(totUrls)
	crawlDuration := timer.Elapsed()
	checkResult := models.NewCheckResult(links, crawlDuration)

	webreporter := reporter.NewReportGenerator()
	if err := webreporter.Generate(checkResult, "output.csv"); err != nil {
		utils.Error("saving...", err)
	}

	fmt.Println("finish!")
}
