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

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get elasticsearch records",
	Long: `Execute a query on elasticsearch.

Note: Don't forget to escape special characters (e.g. *, $, ", and ') with \ where needed`,
	Run: func(cmd *cobra.Command, args []string) {
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

//func connection(q []string) {
//
//    url := viper.GetString("url")
//    options := []elastic.ClientOptionFunc{
//        elastic.SetURL(url),
//        elastic.SetSniff(false),
//        elastic.SetHealthcheckTimeoutStartup(10 * time.Second),
//        elastic.SetHealthcheckTimeout(2 * time.Second),
//    }
//
//    if viper.GetString("username") != "" {
//        options = append(options,
//            elastic.SetBasicAuth(viper.GetString("username"), viper.GetString("password")))
//    }
//
//    if verbose {
//        options = append(options,
//            elastic.SetTraceLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)))
//    }
//
//    client, err := elastic.NewClient(options...)
//    if err != nil {
//        log.Fatalf("Could not connect Elasticsearch client to %s: %s.", url, err)
//    }
//
//    //result, _ := client.Search().
//    //    Index(viper.GetString("index")).
//    //    Query(elastic.NewQueryStringQuery(strings.Join(q, " "))).
//    //    Sort(viper.GetString("timestamp"), false).
//    //    From(0).
//    //    Size(1).
//    //    Do(context.Background())
//    //fmt.Printf("Query took %d milliseconds\n", result.TookInMillis)
//
//    // Count total and setup progress
//    total, err := client.Count().
//        Index(viper.GetString("index")).
//        Query(elastic.NewQueryStringQuery(strings.Join(q, " "))).
//        Do(context.Background())
//    if err != nil {
//        panic(err)
//    }
//    fmt.Printf("records: %d\n", total)
//
//    //bar := pb.StartNew(int(total))
//
//    // 1st goroutine sends individual hits to channel.
//    //hits := make(chan json.RawMessage)
//
//    //g, ctx := errgroup.WithContext(context.Background())
//    //g.Go(func() error {
//        //defer close(hits)
//        // Initialize scroller. Just don't call Do yet.
//        scroll := client.Scroll().
//            Index(viper.GetString("index")).
//            Query(elastic.NewQueryStringQuery(strings.Join(q, " "))).
//            Sort(viper.GetString("timestamp"), false).
//            Size(1000)
//
//        for {
//            results, err := scroll.Do(context.Background())
//            if err == io.EOF {
//                //return nil // all results retrieved
//                return
//            }
//            if err != nil {
//                fmt.Println("something went wrong")
//                //return err // something went wrong
//                return
//            }
//
//            // print the results
//            // convience function that needs a type
//            //for _, item := range results.Each() {
//            //    fmt.Printf("%v\n", item)
//            //}
//
//
//            for _, hit := range results.Hits.Hits {
//                var l map[string]interface{}
//                err := json.Unmarshal(*hit.Source, &l)
//                if err != nil {
//                    // Deserialization failed
//                }
//                fmt.Printf("%s",l["log"])
//            }
//
//            //for _, hit := range results.Hits.Hits {
//            //    fmt.Println(*hit.Source)
//            //}
//
//            //// Check if we need to terminate early
//            //select {
//            //default:
//            //case <-ctx.Done():
//            //    return ctx.Err()
//            //}
//        }
//    //})
//    //// 2nd goroutine receives hits and deserializes them.
//    //for i := 0; i < 10; i++ {
//    //    g.Go(func() error {
//    //        for _ = range hits {
//    //        //for hit := range hits {
//    //            //// Deserialize
//    //            //var p Product
//    //            //err := json.Unmarshal(hit, &p)
//    //            //if err != nil {
//    //            //    return err
//    //            //}
//
//    //            /// Do something with the product here, e.g. send it to another channel
//    //            /// for further processing.
//    //            //_ = p
//
//    //            bar.Increment()
//
//    //            // Terminate early?
//    //            select {
//    //            default:
//    //            case <-ctx.Done():
//    //                return ctx.Err()
//    //            }
//    //        }
//    //        return nil
//    //    })
//    //}
//
//    //// Check whether any goroutines failed.
//    //if err := g.Wait(); err != nil {
//    //    panic(err)
//    //}
//
//    //// Done.
//    //bar.FinishPrint("Done")
//}
//
