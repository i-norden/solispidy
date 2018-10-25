package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/i-norden/solispidy/config"
)

var (
	cfgFile     string
	fileConfig  config.Config
	sourceFiles map[string]string
)

var rootCmd = &cobra.Command{
	Use:              "solispidy",
	PersistentPreRun: configure,
}

func Execute() {

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func configure(cmd *cobra.Command, args []string) {

	fileConfig = config.Config{
		Input:  viper.GetString("solispidy.input"),
		Output: viper.GetString("solispidy.output"),
	}

	viper.Set("solispidy.config", fileConfig)

	loadSourceFiles()
}

func init() {

	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "environments/example.toml", "config file location")
	rootCmd.PersistentFlags().String("input", "./examples", "solispidy input")
	rootCmd.PersistentFlags().String("output", "./output_files", "solispidy output")

	viper.BindPFlag("solispidy.input", rootCmd.PersistentFlags().Lookup("input"))
	viper.BindPFlag("solispidy.output", rootCmd.PersistentFlags().Lookup("output"))
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".solispidy")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Printf("Using config file: %s\n\n", viper.ConfigFileUsed())
	}
}

func loadSourceFiles() {

	sourceFiles = make(map[string]string)

	inputFiles, err := ioutil.ReadDir(fileConfig.Input)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range inputFiles {
		fileName := file.Name()
		fileName = fmt.Sprintf("%s/%s", fileConfig.Input, fileName)
		fmt.Fprintf(os.Stderr, "file name: %v\r\n", fileName)

		text, err := ioutil.ReadFile(fileName)
		if err != nil {
			log.Fatal(err)
		}

		//sourceFiles = append(sourceFiles, string(text))
		sourceFiles[fileName] = string(text)
	}

}
