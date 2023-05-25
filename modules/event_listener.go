package modules

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	nftCollection "github.com/jahankohan/go_event_listener/artifacts"
	models "github.com/jahankohan/go_event_listener/model"
)

type ClientHandler struct {
	*ethclient.Client
	DeployedAddress	string
}

func (ch ClientHandler) PullEvents() models.AllEvents{
	allEvents := models.AllEvents{}
	contractAddress := common.HexToAddress("0xdC4491Dca714E27dE9A245d5c02Ad7b5F2d6c5c7")
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(22249888),
		ToBlock:   big.NewInt(22251395),
		Addresses: []common.Address{
		  contractAddress,
		},
	}

	logs, err := ch.Client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(nftCollection.NftCollectionABI)))
	if err != nil {
        log.Fatal(err)
    }

	collectionCreatedSig := []byte("CollectionCreated(address,string,string)")
	tokenMintedSig := []byte("TokenMinted(address,address,uint256,string)")

	collectionCreatedSigHash := crypto.Keccak256Hash(collectionCreatedSig)
	tokenMintedSigHash := crypto.Keccak256Hash(tokenMintedSig)

	for _, vLog := range logs {


		switch vLog.Topics[0].Hex() {
		case collectionCreatedSigHash.Hex():
			fmt.Println("Log Name: Collection Created")
			
			event := make(map[string]interface{})
			err := contractAbi.UnpackIntoMap(event, "CollectionCreated", vLog.Data)
		
            if err != nil {
                log.Fatal(err)
            }

			_collectionCreatedEvent := models.CollectionCreatedEvent{
				Collection: (common.HexToAddress(vLog.Topics[1].Hex())).String(),
				Name:       event["name"].(string),
				Symbol:     event["symbol"].(string),
			}
			
			allEvents.CollectionCreatedEvents = append(allEvents.CollectionCreatedEvents, _collectionCreatedEvent)

		case tokenMintedSigHash.Hex():
			fmt.Println("Log Name: Token Minted")

			event := make(map[string]interface{})
			err := contractAbi.UnpackIntoMap(event, "TokenMinted", vLog.Data)
			
			if err != nil {
                log.Fatal(err)
            }

			_tokenMintedEvent := models.TokenMintedEvent {
				Collection:	(common.HexToAddress(vLog.Topics[1].Hex())).String(),
				Recipient:	(common.HexToAddress(vLog.Topics[2].Hex())).String(),
				TokenId:	uint((common.HexToAddress(vLog.Topics[3].Hex())).Big().Uint64()),
				TokenUri: 	event["tokenUri"].(string),
			}

			allEvents.TokenMintedEvents = append(allEvents.TokenMintedEvents, _tokenMintedEvent)
		}
	}
	return allEvents
}