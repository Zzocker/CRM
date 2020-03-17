package main

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

func (c *Chaincode) CreateNewLead(ctx CustomTransactionContextInterface, id, salutation, fname, lname, mobile, requester string) string {
	id = uuid.New().String()
	createTime := time.Now().Unix()
	lead := Lead{
		DocType:  LEAD,
		ID:       id,
		Update:   make(map[int64]string),
	}

	lead.Update[createTime] = requester

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
		DocType:  LEAD,
		ID:       id,
		Update:   make(map[int64]string),
	}
	lead.Update[createTime] = requester

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
