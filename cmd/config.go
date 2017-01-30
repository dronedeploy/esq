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
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "creates a config file for esq",
	Long: `
Creates a config file for you based on provided flags. You can also just make one manually.
You can overwrite a subset of the config by providing only the flags you want to change.

The config file must be located at: ` + os.Getenv("HOME") + `/.esq.yml`,
	Run: func(cmd *cobra.Command, args []string) {
		writeConfig()

	},
}

func init() {
	RootCmd.AddCommand(configCmd)

	// Local flags which will only run when this command is called
	configCmd.Flags().String("url", "http://127.0.0.1:9200", "Elasticsearch url, with protocol and port specified")
	viper.BindPFlag("url", configCmd.Flags().Lookup("url"))

	configCmd.Flags().StringP("user", "u", "", "Username for http basic auth.")
	viper.BindPFlag("username", configCmd.Flags().Lookup("user"))

	configCmd.Flags().StringP("password", "p", "", "Password for http basic auth.")
	viper.BindPFlag("password", configCmd.Flags().Lookup("password"))

	configCmd.Flags().StringP("timestamp", "t", "@timestamp", "Timestamp field name used for sorting entries")
	viper.BindPFlag("timestamp", configCmd.Flags().Lookup("timestamp"))

	configCmd.Flags().StringP("index", "i", "logstash-*", "Index to query")
	viper.BindPFlag("index", configCmd.Flags().Lookup("index"))

	configCmd.Flags().StringP("fields", "f", "message", "Comma dilimited list of fields to query by default")
	viper.BindPFlag("fields", configCmd.Flags().Lookup("fields"))
}

func writeConfig() {
	//configFilePath := os.Getenv("HOME") + "/.esq.yml"

	var C config

	err := viper.Unmarshal(&C)
	if err != nil {
		fmt.Println("unable to decode into struct, %v", err)
	}

	cfgFile, err := yaml.Marshal(&C)
	if err != nil {
		fmt.Println("error: %v", err)
	}

	configFilePath := viper.ConfigFileUsed()
	fmt.Println("writing config to " + configFilePath)
	err = ioutil.WriteFile(configFilePath, cfgFile, 0600)
}

type config struct {
	Url       string
	Username  string
	Password  string
	Timestamp string
	Index     string
	Fields    string
}
