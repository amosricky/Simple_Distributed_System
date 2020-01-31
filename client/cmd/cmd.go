package cmd

import (
"Simple_Distributed_System/pb"
"Simple_Distributed_System/setting"
"encoding/json"
"fmt"
"github.com/sirupsen/logrus"
"github.com/spf13/cobra"
"golang.org/x/net/context"
"google.golang.org/grpc"
"os"
)

var conn *grpc.ClientConn
var err error
var gameID string
var gameName string
var gameTeam int32
var gameRound int32
var gameAdd int32
var dbIP string
var dbPort int32

var rootCmd = &cobra.Command{Use: "",}

var gameCmd = &cobra.Command{
	Use: "game",
	Short: "Get & Modify game record",
	Long: `Get & Modify game record`,}

var scoreCmd = &cobra.Command{
	Use: "score",
	Short: "Get score by game ID.",
	Long: `Get score by game ID.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(gameID)!=0{
			connGetScore := pb.NewServiceServerClient(conn)
			resultGetScore, err := connGetScore.GetScore(context.Background(), &pb.GetScoreRequest{ID:gameID, DbIP:dbIP, DbPort:dbPort})
			if err != nil {
				logrus.Warnf("Can't execute [GetScore] function：%v", err.Error())
			}else {
				resultJson, _ := json.Marshal(resultGetScore)
				logrus.Printf("Reply [GetScore]：%s", resultJson)
			}
		}else {
			logrus.Printf("Incomplete command")
		}
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get game list. (Contain gameName & gameID)",
	Long: `Get game list. (Contain gameName & gameID)`,
	Run: func(cmd *cobra.Command, args []string) {
		connGetGameList := pb.NewServiceServerClient(conn)
		resultGetGameList, err := connGetGameList.GetGameList(context.Background(), &pb.GeneralRequest{DbIP:dbIP, DbPort:dbPort})
		if err != nil {
			logrus.Warnf("Can't execute [GetGameList] function：%v", err.Error())
		}else {
			resultJson, _ := json.Marshal(resultGetGameList)
			logrus.Printf("Reply [GetGameList]：%s", resultJson)
		}
	},
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add point by game ID.",
	Long: `Add point by game ID.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(gameID)!=0 && gameTeam!=-1 && gameRound!=0 && gameAdd!=0{
			connPutScore := pb.NewServiceServerClient(conn)
			gameTeamChangeType := pb.PutScoreRequest_TeamType(gameTeam)

			resultPutScore, err := connPutScore.PutScore(context.Background(), &pb.PutScoreRequest{ID:gameID, Team:gameTeamChangeType, Round:gameRound, Add:gameAdd})
			if err != nil {
				logrus.Warnf("Can't execute [PutScore] function：%v", err.Error())
			}else {
				resultJson, _ := json.Marshal(resultPutScore)
				logrus.Printf("Reply [PutScore]：%s", resultJson)
			}
		} else {
			logrus.Printf("Incomplete command")
		}
	},
}

var newGameCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new game.",
	Long: `Create a new game.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(gameName)
		if len(gameName)!=0{
			connPostNewGame := pb.NewServiceServerClient(conn)
			resultPostNewGame, err := connPostNewGame.PostNewGame(context.Background(), &pb.PostNewGameRequest{Game:gameName})
			if err != nil {
				logrus.Warnf("Can't execute [PostNewGame] function：%v", err.Error())
			}else {
				resultJson, _ := json.Marshal(resultPostNewGame)
				logrus.Printf("Reply [PostNewGame]：%s", resultJson)
			}
		}else {
			logrus.Printf("Incomplete command")
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logrus.Printf("Can't execute salve node：%v", err.Error())
	}
}

func Exit()  {
	defer conn.Close()
	os.Exit(1)
}

func init() {
	// Init Config
	setting.Setup()

	// Init Cli
	cobra.OnInitialize()
	scoreCmd.Flags().StringVarP(&gameID, "id", "i", "", "game id")
	scoreCmd.Flags().StringVarP(&dbIP, "dbIP", "d", setting.DatabaseSetting.ServerIP, "database ip")
	scoreCmd.Flags().Int32VarP(&dbPort, "dbPort", "p", int32(setting.DatabaseSetting.Port), "database port")
	listCmd.Flags().StringVarP(&dbIP, "dbIP", "d", setting.DatabaseSetting.ServerIP, "database ip")
	listCmd.Flags().Int32VarP(&dbPort, "dbPort", "p", int32(setting.DatabaseSetting.Port), "database port")
	addCmd.Flags().StringVarP(&gameID, "id", "i", "", "game id")
	addCmd.Flags().Int32VarP(&gameTeam, "team", "t", -1, "team ([0]home [1]visitor)")
	addCmd.Flags().Int32VarP(&gameRound, "round", "r", 0, "round ([min]1 [max]9)")
	addCmd.Flags().Int32VarP(&gameAdd, "add", "a", 0, "add point")
	newGameCmd.Flags().StringVarP(&gameName, "name", "n", "", "game name")
	rootCmd.AddCommand(gameCmd)
	gameCmd.AddCommand(scoreCmd, listCmd, addCmd, newGameCmd)

	// Init DB connection
	serverUrl := fmt.Sprintf("localhost:%v", setting.ServerSetting.Port)
	conn, err = grpc.Dial(serverUrl, grpc.WithInsecure())
	if err != nil {
		logrus.Warnf("Can't connect to server：%v", err.Error())
	}
}