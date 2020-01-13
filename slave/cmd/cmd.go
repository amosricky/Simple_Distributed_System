package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var gameID string
var gameName string
var gameTeam int32
var gameRound int32
var gameAdd int32

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{Use: "",}

var gameCmd = &cobra.Command{Use: "game",}

var scoreCmd = &cobra.Command{
	Use: "score",
	Short: "Get score by game ID.",
	Long: `Get score by game ID.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(gameID)!=0{
			fmt.Printf("MY gameID  %s",gameID)
		}else {
			fmt.Println("scoreCmd")
		}
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get game list. (Contain gameName & gameID)",
	Long: `Get game list. (Contain gameName & gameID)`,
	//Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("listCmd")
	},
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add point by game ID.",
	Long: `Add point by game ID.`,
	//Args: cobra.MinimumNArgs(4),
	Run: func(cmd *cobra.Command, args []string) {
		if len(gameID)!=0{
			fmt.Printf("MY gameID %v",gameID)
		}
		if gameTeam!=0{
			fmt.Printf("MY gameTeam %v",gameTeam)
		}
		if gameRound!=0{
			fmt.Printf("MY gameRound %v",gameRound)
		}
		if gameAdd!=0{
			fmt.Printf("MY gameAdd %v",gameAdd)
		}
	},
}

var newGameCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new game.",
	Long: `Create a new game.`,
	//Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(gameName)!=0{
			fmt.Printf("MY gameName %v",gameName)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.TempProject.yaml)")
	scoreCmd.Flags().StringVarP(&gameID, "id", "i", "", "game id")
	addCmd.Flags().StringVarP(&gameID, "id", "i", "", "game id")
	addCmd.Flags().Int32VarP(&gameTeam, "team", "t", 0, "team ([1]home [2]visitor)")
	addCmd.Flags().Int32VarP(&gameRound, "round", "r", 0, "round ([min]1 [max]9)")
	addCmd.Flags().Int32VarP(&gameAdd, "add", "a", 0, "add point")
	newGameCmd.Flags().StringVarP(&gameName, "name", "n", "", "game name")
	RootCmd.AddCommand(gameCmd)
	gameCmd.AddCommand(scoreCmd, listCmd, addCmd, newGameCmd)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".TempProject" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".TempProject")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}