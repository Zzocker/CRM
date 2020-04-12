package main

import (
	"encoding/json"
	. "fmt"
	"time"

	"github.com/google/uuid"
	"github.com/hyperledger/fabric-protos-go/ledger/queryresult"
)

type DealOutput struct {
	Result []Deal `json:"result"`
	Count  uint   `json:"counter"`
}

func (c *Chaincode) CreatNewDeal(ctx CustomTransactionContextInterface, leadSourceID, input, requester string) (string, error) {
	existing := ctx.GetData()
	if existing == nil {
		return "", Errorf("Key with %v doesn't exists", leadSourceID)
	}
	var lead Lead
	json.Unmarshal(existing, &lead)
	if lead.Owner != requester { // new logic will be implemented when ecert is added
		return "", Errorf("Owner missmatch")
	}
	id := uuid.New().String()
	var deal Deal
	json.Unmarshal([]byte(input), &lead)
	deal.DocType = DEAL
	deal.DealID = id
	deal.DealOwner = requester
	deal.CreatedBy = requester
	deal.UpdatedBy = requester
	t := time.Now().Unix()
	deal.UpdatedDate = t
	deal.CreatedDate = t
	deal.DealLeadID = leadSourceID
	dealAsByte, _ := json.Marshal(deal)

	return id, ctx.GetStub().PutState(id, dealAsByte)
}

func (c *Chaincode) DeleteDeal(ctx CustomTransactionContextInterface, id, requester string) error {
	existing := ctx.GetData()
	if existing == nil {
		return Errorf("Key with %v doesn't exists", id)
	}
	var deal Deal
	json.Unmarshal(existing, &deal)
	if deal.DealOwner != requester {
		Errorf("Unauthorized to Delete this deal ")
	}
	return ctx.GetStub().DelState(id)
}

func (c *Chaincode) GetAllMyDeals(ctx CustomTransactionContextInterface, qtype, qvalue, requester string) (DealOutput, error) {
	var output DealOutput
	var query string
	if qtype == CREATED {
		query = `{"use_index": "DealOnCreated",`
	} else if qtype == UPDATED {
		query = `{"use_index": "DealOnUpdated",`
	} else {
		return DealOutput{}, Errorf("Error : No suc query type %v for deal", qtype)
	}
	query += `
			"selector": {
				"docTyp": "DEAL",
				"deal_owner": `
	query += "\"" + requester + "\""
	query += qvalue
	// to add into selector , ---new selector----}}
	// not to selector , --},new :----}
	result, err := ctx.GetStub().GetQueryResult(query)
	if err != nil {
		return DealOutput{}, err
	}
	for result.HasNext() {
		var resultKV *queryresult.KV

		resultKV, _ = result.Next()
		var deal Deal
		json.Unmarshal(resultKV.GetValue(), &deal)
		output.Result = append(output.Result, deal)
		output.Count++
	}
	return output, result.Close()
}

func (c *Chaincode) GetMyDeal(ctx CustomTransactionContextInterface, id, requester string) (Deal, error) {
	existing := ctx.GetData()
	if existing == nil {
		return Deal{}, Errorf("Key with %v doesn't exists", id)
	}
	var deal Deal
	json.Unmarshal(existing, &deal)
	if deal.DealOwner != requester { // new logic will be implemented when ecert is added
		return Deal{}, Errorf("Owner missmatch")
	}
	return deal, nil
}
