package main

import (
	"encoding/json"
	. "fmt"
	"time"

	"github.com/google/uuid"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/ledger/queryresult"
	"github.com/hyperledger/fabric-protos-go/peer"
)

type DealOutput struct {
	Bookmark string `json:"bookmark"`
	Result   []Deal `json:"result"`
	Count    uint   `json:"counter"`
}

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

func (c *Chaincode) DeleteDeal(ctx CustomTransactionContextInterface, id, requester string) error {
	existing := ctx.GetData()
	if existing == nil {
		return Errorf("Key with %v doesn't exists", id)
	}
	var deal Deal
	json.Unmarshal(existing, &deal)
	if deal.DealOwner != requester {
		Errorf("Unauthorized to Delete this lead ")
	}
	return ctx.GetStub().DelState(id)
}

func (c *Chaincode) GetAllMyDeals(ctx CustomTransactionContextInterface, bookmark, requester string) (DealOutput, error) {
	var pagesize int32 = 10
	if bookmark == "start" {
		bookmark = ""
	}
	var hasMorePages = true

	query := `{
			"selector": {
				"docTyp": "DEAL",
				"deal_owner": `
	query += "\"" + requester + "\""
	query += `
			},
			"use_index": "indexOnDealOwner"
		}`
	var output DealOutput
	var qryIterator shim.StateQueryIteratorInterface
	var queryMetaDate *peer.QueryResponseMetadata
	var err error

	lastBookmark := ""
	for hasMorePages {
		qryIterator, queryMetaDate, err = ctx.GetStub().GetQueryResultWithPagination(query, pagesize, bookmark)
		if err != nil {
			return DealOutput{}, Errorf("error connecting world state")
		}
		var resultKV *queryresult.KV

		if lastBookmark != queryMetaDate.Bookmark {
			for qryIterator.HasNext() {
				resultKV, err = qryIterator.Next()

				output.Count++

				dealAsByte := resultKV.GetValue()
				var deal Deal
				if requester != deal.DealOwner {
					return DealOutput{}, nil
				}
				json.Unmarshal(dealAsByte, &deal)
				output.Result = append(output.Result, deal)
			}
		}
		bookmark = queryMetaDate.Bookmark
		hasMorePages = (bookmark != "" && lastBookmark != bookmark)
		lastBookmark = bookmark
	}
	output.Bookmark = bookmark

	return output, nil
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
