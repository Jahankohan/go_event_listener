package model

type CollectionCreatedEvent struct {
	Collection string `json:"collection"`
	Name       string `json:"name"`
	Symbol     string `json:"symbol"`
}

type TokenMintedEvent struct {
	Collection string `json:"collection"`
	Recipient  string `json:"recipient"`
	TokenId    uint   `json:"tokenId"`
	TokenUri   string `json:"tokenUri"`
}

type AllEvents struct {
    CollectionCreatedEvents []CollectionCreatedEvent `json:"collectionCreatedEvents"`
    TokenMintedEvents       []TokenMintedEvent       `json:"tokenMintedEvents"`
}

