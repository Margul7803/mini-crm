package cmd

import (
	"fmt"
	"os"

	"mini-crm/store"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile   string
	dataStore store.Storer
)

var rootCmd = &cobra.Command{
	Use:   "mini-crm",
	Short: "A simple CLI CRM application",
	Long:  `A simple CLI CRM application to manage contacts.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.mini-crm.yaml)")
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name "config" (without extension).
		viper.AddConfigPath(home)
		viper.AddConfigPath(".") // add current directory to config search path
		viper.SetConfigType("yaml")
		viper.SetConfigName("config") // name of config file (without extension)
	}

	viper.AutomaticEnv()

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

	// Initialize the data store based on the configuration.
	storageType := viper.GetString("storage.type")
	storagePath := viper.GetString("storage.path")

	switch storageType {
	case "gorm":
		var err error
		dataStore, err = store.NewGORMStore(storagePath)
		cobra.CheckErr(err)
	case "json":
		dataStore = store.NewJSONStore(storagePath)
	case "memory":
		dataStore = store.NewMemoryStore()
	default:
		fmt.Println("Invalid storage type specified in configuration. Using in-memory store as default.")
		dataStore = store.NewMemoryStore()
	}
}