/*
Copyright © 2024 Adam B Kaplan adam@adambkaplan.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/adambkaplan/kidbase10/encoding"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	decode  bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kidbase10 data",
	Short: "Simple encoding a child can do!",
	Long: `kidbase10 is a simple encoding protocol that can be written with paper and
pencil. It uses a short encoding table to turn words and phrases into a series
of base-10 numbers. The default implementation produces an output of 8-digit
numbers, separated by newlines.

Examples:

  kidbase10 "LANE" // encodes "LANE" to 51620000

  echo "I STARE" | kidbase10 - // encodes data received from STDIN.

  kidbase10 -d 51620000 // decodes "51620000" to "LANE"
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return cmd.Help()
		}
		var input io.Reader
		data := args[0]
		if data == "-" {
			input = cmd.InOrStdin()
		} else {
			input = strings.NewReader(data)
		}

		if decode {
			decoder := encoding.NewDecoder(input)
			_, err := decoder.DecodeTo(cmd.OutOrStdout())
			return err
		}

		encoder := encoding.NewEncoder(cmd.OutOrStdout())
		return encoder.Encode(input)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.kidbase10.yaml)")

	// flag for decoding
	rootCmd.Flags().BoolVarP(&decode, "decode", "d", false, "decode data")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".kidbase10" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".kidbase10")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
