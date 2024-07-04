package anticafe

import (
	"context"
	"fmt"
	"github.com/SevereCloud/vksdk/api"
	"golang.org/x/oauth2/jwt"
	"gopkg.in/Iwark/spreadsheet.v2"
	"log"
	"os"
	swagger "pwd/pkg/service/game/tesera-swagger"
	"strings"
	"sync"
)

const (
	confFile = "client_secret.json"
)

var (
	spreadSheetAnticafe = []anticafeConf{
		{
			anticafeName: "Арена Эмоций",
			link:         "https://vk.com/emarena",
			googleSheetsInfo: &googleSheetsInfo{
				id:         "1ln_rHGOEsGSwXSKSsa_SBytObh8NKbEmzF5chT-Y8mk",
				sheetTitle: "База игр",
				column:     1,
			},
		},
		{
			anticafeName: "Москвалогия",
			link:         "https://vk.com/moskvalogiya",
			googleSheetsInfo: &googleSheetsInfo{
				id:         "1f2sytGOg9FYtgMyzqNsg2-8RlBmC_27zhV0pcw4rpkI",
				sheetTitle: "Список игр",
				column:     1,
			},
		},
		{
			anticafeName: "Терра",
			link:         "https://vk.com/terra_games",
			vkTopicInfo: &vkTopicInfo{
				groupID: "156570572",
				topicID: "38983587",
			},
		},
	}
)

type Service struct {
	googleSheetClient *spreadsheet.Service
	vkClient          *api.VK
}

func New() *Service {
	// Create a context
	ctx := context.Background()

	//data, err := os.ReadFile(confFile)

	//if err != nil {
	//	log.Fatalf("Unable to parse client secret file to config: %v", err)
	//}

	//conf, err := google.JWTConfigFromJSON(data, spreadsheet.Scope)
	//if err != nil {
	//	log.Fatalf("Unable to parse client secret file to config: %v", err)
	//}

	googleTokenEnv := os.Getenv("GOOGLE_TOKEN")
	if googleTokenEnv == "" {
		log.Fatalf("Unable to get google token from environment variables")
	}

	privateKey := []byte(googleTokenEnv)
	conf := jwt.Config{
		Email:        "boardgame@sturdy-device-427311-d5.iam.gserviceaccount.com",
		PrivateKey:   privateKey,
		PrivateKeyID: "ef9a60f3df13a7c61899ed5b3138669c4c40358c",
		Scopes:       []string{spreadsheet.Scope},
		TokenURL:     "https://oauth2.googleapis.com/token",
	}

	client := conf.Client(ctx)

	// Create a new spreadsheet service
	googleSheetClient := spreadsheet.NewServiceWithClient(client)

	vkToken := os.Getenv("VK_TOKEN")
	if vkToken == "" {
		log.Fatalf("Unable to get vk token from environment variables")
	}

	vk := api.NewVK(vkToken)

	return &Service{googleSheetClient: googleSheetClient, vkClient: vk}
}

func (s *Service) isGameInStockGoogleSheets(conf googleSheetsInfo, game swagger.GameInfo) (bool, error) {
	spreadsheet, err := s.googleSheetClient.FetchSpreadsheet(conf.id)
	if err != nil {
		return false, fmt.Errorf("unable to fetch spreadsheet: %w", err)
	}

	sheet, err := spreadsheet.SheetByTitle(conf.sheetTitle)
	if err != nil {
		return false, fmt.Errorf("unable to fetch sheet: %w", err)
	}

	for _, cell := range sheet.Columns[conf.column] {
		if isGameIsFounded(cell.Value, game) {
			return true, nil
		}
	}

	return false, nil
}

func (s *Service) isGameInStockVKTopic(conf vkTopicInfo, game swagger.GameInfo) (bool, error) {
	params := api.Params{"group_id": conf.groupID, "topic_id": conf.topicID}

	// get information about the group
	comments, err := s.vkClient.BoardGetComments(params)
	if err != nil {
		return false, fmt.Errorf("unable to fetch comments: %w", err)
	}

	for _, comment := range comments.Items {
		if isGameIsFounded(comment.Text, game) {
			return true, nil
		}
	}

	return false, nil
}

func isGameIsFounded(text string, game swagger.GameInfo) bool {
	if game.Title != "" && strings.Contains(text, game.Title) {
		return true
	}
	if game.Title2 != "" && strings.Contains(text, game.Title2) {
		return true
	}
	if game.Title3 != "" && strings.Contains(text, game.Title3) {
		return true
	}

	return false
}

func (s *Service) IsGameInStockInAnticafe(game swagger.GameInfo) []AnticafeGameInStock {
	isInStockAnticafe := []AnticafeGameInStock{}
	wg := sync.WaitGroup{}
	wg.Add(len(spreadSheetAnticafe))
	mu := sync.Mutex{}
	for _, conf := range spreadSheetAnticafe {
		go func(conf anticafeConf) {
			var isInStock bool
			var err error
			if conf.googleSheetsInfo != nil {
				isInStock, err = s.isGameInStockGoogleSheets(*conf.googleSheetsInfo, game)
				if err != nil {
					log.Println("failed to check game in stock", err)
				}
			}

			if conf.vkTopicInfo != nil {
				isInStock, err = s.isGameInStockVKTopic(*conf.vkTopicInfo, game)
				if err != nil {
					log.Println("failed to check game in stock", err)
				}
			}

			if isInStock {
				mu.Lock()
				isInStockAnticafe = append(isInStockAnticafe, AnticafeGameInStock{
					AnticafeName: conf.anticafeName,
					IsInStock:    true,
					Link:         conf.link,
				})
				mu.Unlock()
			}
			wg.Done()
		}(conf)
	}
	wg.Wait()

	return isInStockAnticafe
}
