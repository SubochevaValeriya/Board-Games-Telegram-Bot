package internal

import (
	"context"
	"encoding/json"
	"github.com/fzerorubigd/gobgg"
)
import "fmt"

func (i *Info) BGGParsing(id int) error {
	bgg := ConnectToBGGClient()
	results, err := bgg.GetThings(context.Background(), gobgg.GetThingIDs(int64(id)))
	if err != nil {
		return err
	}
	data, err := json.Marshal(results[0])
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, i)
	if err != nil {
		return err
	}

	i.NumberOfPlayers = fmt.Sprintf("%v - %v", results[0].MinPlayers, results[0].MaxPlayers)
	i.RecommendedNumberOfPlayers = i.NumberOfPlayers

	return nil
}
