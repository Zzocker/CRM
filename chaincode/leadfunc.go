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

type LeadOutput struct {
	Bookmark string `json:"bookmark"`
	Result   []Lead `json:"result"`
	Count    uint   `json:"counter"`
}

func (c *Chaincode) CreateNewLead(ctx CustomTransactionContextInterface, id, salutation, fname, lname, mobile, requester string) string {
	id = uuid.New().String()
	createTime := time.Now().Unix()
	lead := Lead{
		DocType: LEAD,
		ID:      id,
	}

	lead.UpdatedBy = requester
	lead.UpdatedDate = createTime

	lead.Saluation = salutation
	lead.FirstName = fname
	lead.LastName = lname
	lead.Mobile = mobile

	lead.CreateBy = requester
	lead.CreateDate = createTime
	lead.Owner = requester

	leadAsByte, _ := json.Marshal(lead)

	ctx.GetStub().PutState(id, leadAsByte)

	// return id
	return id
}
func (c *Chaincode) CreateLeadFromContact(ctx CustomTransactionContextInterface, id, contactid, title, salutation, fname, lname, mobile, country, state, city, pincode, email, requester string) string {
	id = uuid.New().String()
	createTime := time.Now().Unix()
	lead := Lead{
		DocType: LEAD,
		ID:      id,
	}
	lead.UpdatedBy = requester
	lead.UpdatedDate = createTime

	lead.JobTitle = title
	lead.Saluation = salutation
	lead.FirstName = fname
	lead.LastName = lname
	lead.Mobile = mobile
	lead.Country = country
	lead.State = state
	lead.City = city
	lead.Zipcode = pincode
	lead.Email = email

	lead.CreateBy = requester
	lead.CreateDate = createTime
	lead.Owner = requester
	lead.ContactID = contactid

	leadAsByte, _ := json.Marshal(lead)

	ctx.GetStub().PutState(id, leadAsByte)

	// return id
	return id

}

func (c *Chaincode) DeleteLead(ctx CustomTransactionContextInterface, id, requester string) error {
	existing := ctx.GetData()
	if existing == nil {
		return Errorf("Key with %v doesn't exists", id)
	}
	var lead Lead
	json.Unmarshal(existing, &lead)
	if lead.Owner != requester {
		Errorf("Unauthorized to Delete this lead ")
	}
	return ctx.GetStub().DelState(id)
}

func (c *Chaincode) GetAllMyLeads(ctx CustomTransactionContextInterface, bookmark, requester string) (LeadOutput, error) {
	var pagesize int32 = 10
	if bookmark == "start" {
		bookmark = ""
	}
	var hasMorePages = true

	query := `{
			"selector": {
				"docTyp": "LEAD",
				"lead_owner": `
	query += "\"" + requester + "\""
	query += `
			},
			"use_index": "indexOnLeadOwner"
		}`
	var output LeadOutput
	var qryIterator shim.StateQueryIteratorInterface
	var queryMetaDate *peer.QueryResponseMetadata
	var err error
	
	lastBookmark := ""
	for hasMorePages {
		qryIterator, queryMetaDate, err = ctx.GetStub().GetQueryResultWithPagination(query, pagesize, bookmark)
		if err != nil {
			return LeadOutput{}, Errorf("error connecting world state")
		}
		var resultKV *queryresult.KV

		if lastBookmark != queryMetaDate.Bookmark {
			for qryIterator.HasNext() {
				resultKV, err = qryIterator.Next()

				output.Count++

				leadAsByte := resultKV.GetValue()
				var lead Lead
				if requester != lead.Owner {
					return LeadOutput{}, nil
				}
				json.Unmarshal(leadAsByte, &lead)
				output.Result = append(output.Result, lead)
			}
		}
		bookmark = queryMetaDate.Bookmark
		hasMorePages = (bookmark != "" && lastBookmark != bookmark)
		lastBookmark = bookmark
	}
	output.Bookmark = bookmark

	return output, nil
}

func (c *Chaincode) GetMyLead(ctx CustomTransactionContextInterface, id, requester string) (Lead, error) {
	existing := ctx.GetData()
	if existing == nil {
		return Lead{}, Errorf("Key with %v doesn't exists", id)
	}
	var lead Lead
	json.Unmarshal(existing, &lead)
	if lead.Owner != requester { // new logic will be implemented when ecert is added
		return Lead{}, Errorf("Owner missmatch")
	}
	return lead, nil
}
