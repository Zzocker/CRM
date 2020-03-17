package main

import (
	"encoding/json"
	"time"
	. "fmt"
	"github.com/google/uuid"
)

func (c *Chaincode) CreatNewDeal(ctx CustomTransactionContextInterface,leadId,requester string) (string,error) {
	existing := ctx.GetData()
	if existing == nil {
		return "",Errorf("Key with %v doesn't exists", leadId)
	}
	json.Unmarshal(existing, &lead)
	if lead.Owner != requester { // new logic will be implemented when ecert is added
		return Errorf("Owner missmatch")
	}
	id := uuid.New().String()
	deal:=Deal{
		DocType: DEAL,
		DealID: id,
		DealLeadID: leadId,
		DealOwner: requester,
		CreatedBy: requester,
		CreatedDate: time.Now().Unix(),
		UpdatedBy: requester,
		UpdatedDate: time.Now().Unix(),
	}

	dealAsByte,_:=json.Marshal(deal)

	
	return id,ctx.GetStub().PutState(id,dealAsByte)
}