// Copyright © 2016 Netuitive, Inc.
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
	"errors"
	"fmt"
	"os"

	"github.com/Netuitive/netuitive-event-handler/netuitive"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var debug bool
var cfgFile string
var APIurl string
var APIkey string

var element string
var level string
var message string
var title string
var eventType string
var source string
var tags string

// This represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "netuitive-event-handler",
	Short: "Post events to Netuitive from the command line",
	Long: Name + " (" + Version + ")\n\n" + `Netuitive Event Handler is a CLI for posting external events
to Netuitive. It is designed to be used from Nagios, Icinga, Sensu, and other monitoring systems.`,

	RunE: func(cmd *cobra.Command, args []string) error {

		if viper.GetString("apikey") == "" {
			return errors.New("missing apikey\n")
		}

		if viper.GetString("url") == "" {
			return errors.New("missing url\n")
		}

		if element == "" {
			return errors.New("missing element\n")
		}

		if title == "" {
			return errors.New("missing title\n")
		}

		if message == "" {
			return errors.New("missing message\n")
		}

		if level == "" {
			return errors.New("missing level\n")
		}

		err := netuitive.PostEvent(viper.GetString("apikey"), viper.GetString("url"), source, element, eventType, title, message, level, tags, debug)

		return err

	},
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

	RootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is /etc/netuitive/netuitive-event-handler.yaml)")
	RootCmd.PersistentFlags().StringVarP(&APIkey, "apikey", "a", "", " API Key if not otherwise specified (optional)")
	RootCmd.PersistentFlags().StringVarP(&APIurl, "url", "u", "https://api.app.netuitive.com/ingest/events", " API URL if not otherwise specified (optional)")

	RootCmd.Flags().StringVarP(&element, "element", "e", "", "Element FQN for the event")
	RootCmd.Flags().StringVarP(&title, "title", "t", "", "Title of the event")
	RootCmd.Flags().StringVarP(&message, "message", "m", "", "Message text of the event")
	RootCmd.Flags().StringVarP(&level, "level", "l", "", "Level of the event")

	//	RootCmd.PersistentFlags().StringVarP(&eventType, "type", "", "INFO", "Type of the event (optional)")
	eventType = "INFO"

	RootCmd.PersistentFlags().StringVarP(&source, "source", "s", "netuitive-event-handler", "Source of the event (optional)")

	RootCmd.Flags().StringVarP(&tags, "tags", "", "", "Tags for the event (optional) Example: tag1:value1,tag2:value2")

	RootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "enable debug")

	viper.BindPFlag("apikey", RootCmd.PersistentFlags().Lookup("apikey"))
	viper.BindPFlag("url", RootCmd.PersistentFlags().Lookup("url"))

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetEnvPrefix("NETUITIVE_EVENT_HANDLER")

	viper.SetConfigName("netuitive-event-handler") // name of config file (without extension)
	viper.AddConfigPath("/etc/netuitive")          // adding home directory as first search path
	viper.AutomaticEnv()                           // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		if debug {
			fmt.Println("Using config file:", viper.ConfigFileUsed())
		}
	}
}
