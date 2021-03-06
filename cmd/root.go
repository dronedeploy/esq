// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"
	"regexp"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var verbose bool

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "esq",
	Short: "A cli for querying elasticsearch",
	Long: `
Kibana is awesome as a search interface, but isn't that useful for 
scanning through a long stream of logs, and doesn't integrate with 
the myriad cli tools available. Esq is an opinionated way to query 
from the command line and pipe the output to other tools.
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is "+os.Getenv("HOME")+"/.esq.yml)")
	RootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "enable for extra logs")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and flags if set.
func initConfig() {
	var validFile = regexp.MustCompile(`^.*\.yml$`)

	if cfgFile != "" { // enable ability to specify config file via flag
		if validFile.MatchString(cfgFile) {
			viper.SetConfigFile(cfgFile)
		} else {
			fmt.Println("Config file must use the '.yml' extension")
			os.Exit(1)
		}
	} else {
		viper.SetConfigName(".esq")   // name of config file (without extension)
		viper.AddConfigPath("$HOME/") // adding home directory as first search path
		//viper.AutomaticEnv()          // read in environment variables that match
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Unable to read config file:", viper.ConfigFileUsed())
	}
}

func validateFlags() {
	//TODO all the flags
}
