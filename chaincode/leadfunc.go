package main

import (
	"encoding/json"
	. "fmt"
	"time"

	"github.com/google/uuid"
	"github.com/hyperledger/fabric-protos-go/ledger/queryresult"
)

type LeadOutput struct {
	Result []Lead `json:"result"`
	Count  uint   `json:"counter"`
}

func (c *Chaincode) CreateNewLead(ctx CustomTransactionContextInterface, input, requester string) string {
	id := uuid.New().String()
	createTime := time.Now().Unix()
	var lead Lead
	json.Unmarshal([]byte(input),&lead)
	// lead := Lead{
	// 	DocType: LEAD,
	// 	ID:      id,
	// }
	lead.DocType = LEAD
	lead.ID = id
	lead.UpdatedBy = requester
	lead.UpdatedDate = createTime

	lead.CreateBy = requester
	lead.CreateDate = createTime
	lead.Owner = requester

	leadAsByte, _ := json.Marshal(lead)

	ctx.GetStub().PutState(id, leadAsByte)

	// return id
	return id
}
func (c *Chaincode) CreateLeadFromContact(ctx CustomTransactionContextInterface, input , requester string) string {
	id := uuid.New().String()
	createTime := time.Now().Unix()
	var lead Lead
	json.Unmarshal([]byte(input),&lead)
	lead.DocType =LEAD
	lead.ID=id 
	lead.UpdatedBy = requester
	lead.UpdatedDate = createTime
	lead.CreateBy = requester
	lead.CreateDate = createTime
	lead.Owner = requester

	leadAsByte, _ := json.Marshal(lead)

	ctx.GetStub().PutState(id, leadAsByte)

	// return id
	return id

}
func (c *Chaincode) UpdateLead(ctx CustomTransactionContextInterface, leadID ,input, requester string) error {
	leadAsByte := ctx.GetData()
	if leadAsByte == nil{
		return Errorf("Lead with ID %v doesn't exists",leadID)
	}
	var lead Lead
	json.Unmarshal(leadAsByte,&lead)
	if lead.Owner != requester{
		return Errorf("Owner missmatch")
	}
	json.Unmarshal([]byte(input),&lead)
	
	lead.UpdatedBy = requester
	lead.UpdatedDate = time.Now().Unix()
	lead.Owner = requester
	leadAsByte,_= json.Marshal(lead)

	return ctx.GetStub().PutState(leadID,leadAsByte)
}



func (c *Chaincode) DeleteLead(ctx CustomTransactionContextInterface, id, requester string) error {
	existing := ctx.GetData()
	if existing == nil {
		return Errorf("Key with %v doesn't exists", id)
	}
	var lead Lead
	json.Unmarshal(existing, &lead)
	if lead.Owner != requester {
		return Errorf("Unauthorized to Delete this lead ")
	}
	return ctx.GetStub().DelState(id)
}

func (c *Chaincode) GetAllMyLeads(ctx CustomTransactionContextInterface, qtype, qvalue, requester string) (LeadOutput, error) {
	var output LeadOutput
	var query string
	if qtype == CREATED {
		query = `{"use_index": "LeadOnCreated",`
	} else if qtype == UPDATED {
		query = `{"use_index": "LeadOnUpdated",`
	} else {
		return LeadOutput{}, Errorf("Error : No suc query type %v for lead", qtype)
	}
	query += `
			"selector": {
				"docTyp": "LEAD",
				"lead_owner": `
	query += "\"" + requester + "\""
	query += qvalue
	// to add into selector , ---new selector----}}
	// not to selector , --},new :----}
	result, err := ctx.GetStub().GetQueryResult(query)
	if err != nil {
		return LeadOutput{}, err
	}
	for result.HasNext() {
		var resultKV *queryresult.KV

		resultKV, _ = result.Next()
		var lead Lead
		json.Unmarshal(resultKV.GetValue(), &lead)
		output.Result = append(output.Result, lead)
		output.Count++
	}
	return output, result.Close()
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
