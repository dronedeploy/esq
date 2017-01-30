// Copyright © 2017 Joseph Schneider <https://github.com/astropuffin>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
    "context"
    "os"
    "strings"

    "time"
	"github.com/spf13/cobra"
    "github.com/spf13/viper"
    "gopkg.in/olivere/elastic.v5"
    "log"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get elasticsearch records",
	Long: `Execute a query on elasticsearch.

Note: Don't forget to escape special characters (e.g. * and $) with \ where needed`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
        fmt.Printf("Search query: %v\n", args)
        connection(args)
	},
}

func init() {
	RootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}


func connection(q []string) {

    ctx := context.Background()

    url := viper.GetString("url")
    options := []elastic.ClientOptionFunc{
        elastic.SetURL(url),
        elastic.SetSniff(false),
        elastic.SetHealthcheckTimeoutStartup(10 * time.Second),
        elastic.SetHealthcheckTimeout(2 * time.Second),
    }

    if viper.GetString("username") != "" {
        options = append(options,
            elastic.SetBasicAuth(viper.GetString("username"), viper.GetString("password")))
    }

    if verbose {
        options = append(options,
            elastic.SetTraceLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)))
    }

    client, err := elastic.NewClient(options...)
    if err != nil {
        log.Fatalf("Could not connect Elasticsearch client to %s: %s.", url, err)
    }

    result, _ := client.Search().
        Index(viper.GetString("index")).
        Query(elastic.NewQueryStringQuery(strings.Join(q, " "))).
        Sort(viper.GetString("timestamp"), false).
        From(0).
        Size(1).
        Do(ctx)
    //scroll := client.Scroll().
    //    Index(viper.GetString("index")).
    //    Query(elastic.NewQueryStringQuery(strings.Join(q, " "))).
    //    Sort(viper.GetString("timestamp"), false).
    //    From(0).
    //    Size(1)

    //pages := 0
    //docs := 0

    //for {
    //    res, err := scroll.Do(ctx)
    //    if err == io.EOF {
    //        return nil // all results retrieved
    //    }
    //    if err != nil {
    //            return err // something went wrong
    //    }
    //    

        
        
    }
    fmt.Printf("Query took %d milliseconds\n", result.TookInMillis)
}
