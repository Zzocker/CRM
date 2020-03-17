package main

import (
	"encoding/json"
	. "fmt"
	"time"

	"github.com/google/uuid"
)

func (c *Chaincode) CreatNewDeal(ctx CustomTransactionContextInterface, leadID, requester string) (string, error) {
	existing := ctx.GetData()
	if existing == nil {
		return "", Errorf("Key with %v doesn't exists", leadID)
	}
	var lead Lead
	json.Unmarshal(existing, &lead)
	if lead.Owner != requester { // new logic will be implemented when ecert is added
		return "", Errorf("Owner missmatch")
	}
	id := uuid.New().String()
	deal := Deal{
		DocType:     DEAL,
		DealID:      id,
		DealLeadID:  leadID,
		DealOwner:   requester,
		CreatedBy:   requester,
		CreatedDate: time.Now().Unix(),
		UpdatedBy:   requester,
		UpdatedDate: time.Now().Unix(),
	}

	dealAsByte, _ := json.Marshal(deal)

	return id, ctx.GetStub().PutState(id, dealAsByte)
}
