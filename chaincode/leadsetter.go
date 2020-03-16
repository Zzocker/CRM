package main

import (
	"encoding/json"
	. "fmt"
	"time"
)

func (c *Chaincode) updateName(ctx CustomTransactionContextInterface, id, Saluation, fname, lname, requester string) error {
	existing := ctx.GetData()
	if existing == nil {
		return Errorf("Key with %v doesn't exists", id)
	}
	var lead Lead
	json.Unmarshal(existing, &lead)
	if lead.General.Owner != requester { // new logic will be implemented when ecert is added
		return Errorf("Owner missmatch")
	}
	lead.Update[time.Now().Unix()] = requester

	lead.Personal.Saluation = Saluation
	lead.Personal.FirstName = fname
	lead.Personal.LastName = lname

	existing, _ = json.Marshal(lead)

	return ctx.GetStub().PutState(id, lead)
}

func (c *Chaincode) updateJobTitle(ctx CustomTransactionContextInterface, id, NewUp, requester string) error {
	existing := ctx.GetData()
	if existing == nil {
		return Errorf("Key with %v doesn't exists", id)
	}
	var lead Lead
	json.Unmarshal(existing, &lead)
	if lead.General.Owner != requester { // new logic will be implemented when ecert is added
		return Errorf("Owner missmatch")
	}
	lead.Update[time.Now().Unix()] = requester

	lead.Personal.JobTitle = NewUp

	existing, _ = json.Marshal(lead)

	return ctx.GetStub().PutState(id, lead)
}

func (c *Chaincode) updatePhone(ctx CustomTransactionContextInterface, id, NewUp, requester string) error {
	existing := ctx.GetData()
	if existing == nil {
		return Errorf("Key with %v doesn't exists", id)
	}
	var lead Lead
	json.Unmarshal(existing, &lead)
	if lead.General.Owner != requester { // new logic will be implemented when ecert is added
		return Errorf("Owner missmatch")
	}
	lead.Update[time.Now().Unix()] = requester

	lead.Personal.Phone = NewUp

	existing, _ = json.Marshal(lead)

	return ctx.GetStub().PutState(id, lead)
}

func (c *Chaincode) updateMobile(ctx CustomTransactionContextInterface, id, NewUp, requester string) error {
	existing := ctx.GetData()
	if existing == nil {
		return Errorf("Key with %v doesn't exists", id)
	}
	var lead Lead
	json.Unmarshal(existing, &lead)
	if lead.General.Owner != requester { // new logic will be implemented when ecert is added
		return Errorf("Owner missmatch")
	}
	lead.Update[time.Now().Unix()] = requester

	lead.Personal.Mobile = NewUp

	existing, _ = json.Marshal(lead)

	return ctx.GetStub().PutState(id, lead)
}

func (c *Chaincode) updateEmail(ctx CustomTransactionContextInterface, id, NewUp, requester string) error {
	existing := ctx.GetData()
	if existing == nil {
		return Errorf("Key with %v doesn't exists", id)
	}
	var lead Lead
	json.Unmarshal(existing, &lead)
	if lead.General.Owner != requester { // new logic will be implemented when ecert is added
		return Errorf("Owner missmatch")
	}
	lead.Update[time.Now().Unix()] = requester

	lead.Personal.Email = NewUp

	existing, _ = json.Marshal(lead)

	return ctx.GetStub().PutState(id, lead)
}

func (c *Chaincode) updateSecondaryEmail(ctx CustomTransactionContextInterface, id, NewUp, requester string) error {
	existing := ctx.GetData()
	if existing == nil {
		return Errorf("Key with %v doesn't exists", id)
	}
	var lead Lead
	json.Unmarshal(existing, &lead)
	if lead.General.Owner != requester { // new logic will be implemented when ecert is added
		return Errorf("Owner missmatch")
	}
	lead.Update[time.Now().Unix()] = requester

	lead.Personal.SecondaryEmail = NewUp

	existing, _ = json.Marshal(lead)

	return ctx.GetStub().PutState(id, lead)
}

func (c *Chaincode) updateSkypid(ctx CustomTransactionContextInterface, id, NewUp, requester string) error {
	existing := ctx.GetData()
	if existing == nil {
		return Errorf("Key with %v doesn't exists", id)
	}
	var lead Lead
	json.Unmarshal(existing, &lead)
	if lead.General.Owner != requester { // new logic will be implemented when ecert is added
		return Errorf("Owner missmatch")
	}
	lead.Update[time.Now().Unix()] = requester

	lead.Personal.Skypid = NewUp

	existing, _ = json.Marshal(lead)

	return ctx.GetStub().PutState(id, lead)
}

func (c *Chaincode) updateAddress(ctx CustomTransactionContextInterface, id, street, arealoction, city, state, zipcode, country, requester string) error {
	existing := ctx.GetData()
	if existing == nil {
		return Errorf("Key with %v doesn't exists", id)
	}
	var lead Lead
	json.Unmarshal(existing, &lead)
	if lead.General.Owner != requester { // new logic will be implemented when ecert is added
		return Errorf("Owner missmatch")
	}
	lead.Update[time.Now().Unix()] = requester

	lead.Personal.Street = street
	lead.Personal.Arealoction = arealoction
	lead.Personal.City = city
	lead.Personal.State = state
	lead.Personal.Zipcode = zipcode
	lead.Personal.Country = country

	existing, _ = json.Marshal(lead)

	return ctx.GetStub().PutState(id, lead)
}

func (c *Chaincode) updateEmailoptout(ctx CustomTransactionContextInterface, id, NewUp, requester string) error {
	existing := ctx.GetData()
	if existing == nil {
		return Errorf("Key with %v doesn't exists", id)
	}
	var lead Lead
	json.Unmarshal(existing, &lead)
	if lead.General.Owner != requester { // new logic will be implemented when ecert is added
		return Errorf("Owner missmatch")
	}
	lead.Update[time.Now().Unix()] = requester

	lead.Personal.Emailoptout = NewUp

	existing, _ = json.Marshal(lead)

	return ctx.GetStub().PutState(id, lead)
}

func (c *Chaincode) updateFax(ctx CustomTransactionContextInterface, id, NewUp, requester string) error {
	existing := ctx.GetData()
	if existing == nil {
		return Errorf("Key with %v doesn't exists", id)
	}
	var lead Lead
	json.Unmarshal(existing, &lead)
	if lead.General.Owner != requester { // new logic will be implemented when ecert is added
		return Errorf("Owner missmatch")
	}
	lead.Update[time.Now().Unix()] = requester

	lead.Personal.Fax = NewUp

	existing, _ = json.Marshal(lead)

	return ctx.GetStub().PutState(id, lead)
}

func (c *Chaincode) updateCompany(ctx CustomTransactionContextInterface, id, cname, industry, requester string) error {
	existing := ctx.GetData()
	if existing == nil {
		return Errorf("Key with %v doesn't exists", id)
	}
	var lead Lead
	json.Unmarshal(existing, &lead)
	if lead.General.Owner != requester { // new logic will be implemented when ecert is added
		return Errorf("Owner missmatch")
	}
	lead.Update[time.Now().Unix()] = requester

	lead.Company.Company = cname
	lead.Company.Industry = industry
	lead.Company.AnnualRevenue = 0
	lead.Company.Website = ""
	lead.Company.NoofEmployees = 0

	existing, _ = json.Marshal(lead)

	return ctx.GetStub().PutState(id, lead)
}

func (c *Chaincode) updateAnnualRevenue(ctx CustomTransactionContextInterface, id, NewUp int64, requester string) error {
	existing := ctx.GetData()
	if existing == nil {
		return Errorf("Key with %v doesn't exists", id)
	}
	var lead Lead
	json.Unmarshal(existing, &lead)
	if lead.General.Owner != requester { // new logic will be implemented when ecert is added
		return Errorf("Owner missmatch")
	}
	lead.Update[time.Now().Unix()] = requester

	lead.Company.AnnualRevenue = NewUp

	existing, _ = json.Marshal(lead)

	return ctx.GetStub().PutState(id, lead)
}

func (c *Chaincode) updateWebsite(ctx CustomTransactionContextInterface, id, NewUp, requester string) error {
	existing := ctx.GetData()
	if existing == nil {
		return Errorf("Key with %v doesn't exists", id)
	}
	var lead Lead
	json.Unmarshal(existing, &lead)
	if lead.General.Owner != requester { // new logic will be implemented when ecert is added
		return Errorf("Owner missmatch")
	}
	lead.Update[time.Now().Unix()] = requester

	lead.Company.Website = NewUp

	existing, _ = json.Marshal(lead)

	return ctx.GetStub().PutState(id, lead)
}

func (c *Chaincode) updateNoofEmployees(ctx CustomTransactionContextInterface, id, NewUp uint, requester string) error {
	existing := ctx.GetData()
	if existing == nil {
		return Errorf("Key with %v doesn't exists", id)
	}
	var lead Lead
	json.Unmarshal(existing, &lead)
	if lead.General.Owner != requester { // new logic will be implemented when ecert is added
		return Errorf("Owner missmatch")
	}
	lead.Update[time.Now().Unix()] = requester

	lead.Company.NoofEmployees = NewUp

	existing, _ = json.Marshal(lead)

	return ctx.GetStub().PutState(id, lead)
}

func (c *Chaincode) updateSource(ctx CustomTransactionContextInterface, id, NewUp, requester string) error {
	existing := ctx.GetData()
	if existing == nil {
		return Errorf("Key with %v doesn't exists", id)
	}
	var lead Lead
	json.Unmarshal(existing, &lead)
	if lead.General.Owner != requester { // new logic will be implemented when ecert is added
		return Errorf("Owner missmatch")
	}
	lead.Update[time.Now().Unix()] = requester

	lead.General.Source = NewUp

	existing, _ = json.Marshal(lead)

	return ctx.GetStub().PutState(id, lead)
}

func (c *Chaincode) updateStatus(ctx CustomTransactionContextInterface, id, NewUp, requester string) error {
	existing := ctx.GetData()
	if existing == nil {
		return Errorf("Key with %v doesn't exists", id)
	}
	var lead Lead
	json.Unmarshal(existing, &lead)
	if lead.General.Owner != requester { // new logic will be implemented when ecert is added
		return Errorf("Owner missmatch")
	}
	lead.Update[time.Now().Unix()] = requester

	lead.General.Status = NewUp

	existing, _ = json.Marshal(lead)

	return ctx.GetStub().PutState(id, lead)
}

func (c *Chaincode) updateRating(ctx CustomTransactionContextInterface, id, NewUp, requester string) error {
	existing := ctx.GetData()
	if existing == nil {
		return Errorf("Key with %v doesn't exists", id)
	}
	var lead Lead
	json.Unmarshal(existing, &lead)
	if lead.General.Owner != requester { // new logic will be implemented when ecert is added
		return Errorf("Owner missmatch")
	}
	lead.Update[time.Now().Unix()] = requester

	lead.General.Rating = NewUp

	existing, _ = json.Marshal(lead)

	return ctx.GetStub().PutState(id, lead)
}

func (c *Chaincode) updateContactID(ctx CustomTransactionContextInterface, id, NewUp, requester string) error {
	existing := ctx.GetData()
	if existing == nil {
		return Errorf("Key with %v doesn't exists", id)
	}
	var lead Lead
	json.Unmarshal(existing, &lead)
	if lead.General.Owner != requester { // new logic will be implemented when ecert is added
		return Errorf("Owner missmatch")
	}
	lead.Update[time.Now().Unix()] = requester

	lead.General.ContactID = NewUp

	existing, _ = json.Marshal(lead)

	return ctx.GetStub().PutState(id, lead)
}

func (c *Chaincode) updateOwner(ctx CustomTransactionContextInterface, id, NewUp, requester string) error {
	existing := ctx.GetData()
	if existing == nil {
		return Errorf("Key with %v doesn't exists", id)
	}
	var lead Lead
	json.Unmarshal(existing, &lead)
	if lead.General.Owner != requester { // new logic will be implemented when ecert is added
		return Errorf("Owner missmatch")
	}
	lead.Update[time.Now().Unix()] = requester

	lead.General.Owner = NewUp

	existing, _ = json.Marshal(lead)

	return ctx.GetStub().PutState(id, lead)
}
