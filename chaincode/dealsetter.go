package main

import (
	"encoding/json"
	"time"
	."fmt"
)

func (c *Chaincode) UpdateDealOwner(ctx CustomTransactionContextInterface, id, newUp, requester string) error {
	existing := ctx.GetData()
	if existing == nil {
		return Errorf("Key with %v doesn't exists", id)
	}
	var deal Deal
	json.Unmarshal(existing, &deal)
	if deal.DealOwner != requester { // new logic will be implemented when ecert is added
		return Errorf("Owner missmatch")
	}
	deal.UpdatedBy = requester
	deal.UpdatedDate = time.Now().Unix()

	deal.DealOwner = newUp

	existing, _ = json.Marshal(deal)

	return ctx.GetStub().PutState(id, existing)
}

func (c *Chaincode) UpdateDealName(ctx CustomTransactionContextInterface, id, newUp, requester string) error {
	existing := ctx.GetData()
	if existing == nil {
		return Errorf("Key with %v doesn't exists", id)
	}
	var deal Deal
	json.Unmarshal(existing, &deal)
	if deal.DealOwner != requester { // new logic will be implemented when ecert is added
		return Errorf("Owner missmatch")
	}
	deal.UpdatedBy = requester
	deal.UpdatedDate = time.Now().Unix()

	deal.DealName = newUp

	existing, _ = json.Marshal(deal)

	return ctx.GetStub().PutState(id, existing)
}

func (c *Chaincode) UpdateDealDate(ctx CustomTransactionContextInterface, id, newUp, requester string) error {
	existing := ctx.GetData()
	if existing == nil {
		return Errorf("Key with %v doesn't exists", id)
	}
	var deal Deal
	json.Unmarshal(existing, &deal)
	if deal.DealOwner != requester { // new logic will be implemented when ecert is added
		return Errorf("Owner missmatch")
	}
	deal.UpdatedBy = requester
	deal.UpdatedDate = time.Now().Unix()

	deal.DealDate = newUp

	existing, _ = json.Marshal(deal)

	return ctx.GetStub().PutState(id, existing)
}

func (c *Chaincode) UpdateDealPotentialName(ctx CustomTransactionContextInterface, id, newUp, requester string) error {
	existing := ctx.GetData()
	if existing == nil {
		return Errorf("Key with %v doesn't exists", id)
	}
	var deal Deal
	json.Unmarshal(existing, &deal)
	if deal.DealOwner != requester { // new logic will be implemented when ecert is added
		return Errorf("Owner missmatch")
	}
	deal.UpdatedBy = requester
	deal.UpdatedDate = time.Now().Unix()

	deal.DealPotentialName = newUp

	existing, _ = json.Marshal(deal)

	return ctx.GetStub().PutState(id, existing)
}

func (c *Chaincode) UpdateDealAccountName(ctx CustomTransactionContextInterface, id, newUp, requester string) error {
	existing := ctx.GetData()
	if existing == nil {
		return Errorf("Key with %v doesn't exists", id)
	}
	var deal Deal
	json.Unmarshal(existing, &deal)
	if deal.DealOwner != requester { // new logic will be implemented when ecert is added
		return Errorf("Owner missmatch")
	}
	deal.UpdatedBy = requester
	deal.UpdatedDate = time.Now().Unix()

	deal.DealAccountName = newUp

	existing, _ = json.Marshal(deal)

	return ctx.GetStub().PutState(id, existing)
}

func (c *Chaincode) UpdateDealType(ctx CustomTransactionContextInterface, id, newUp, requester string) error {
	existing := ctx.GetData()
	if existing == nil {
		return Errorf("Key with %v doesn't exists", id)
	}
	var deal Deal
	json.Unmarshal(existing, &deal)
	if deal.DealOwner != requester { // new logic will be implemented when ecert is added
		return Errorf("Owner missmatch")
	}
	deal.UpdatedBy = requester
	deal.UpdatedDate = time.Now().Unix()

	deal.DealType = newUp

	existing, _ = json.Marshal(deal)

	return ctx.GetStub().PutState(id, existing)
}

func (c *Chaincode) UpdateDealNextStep(ctx CustomTransactionContextInterface, id, newUp, requester string) error {
	existing := ctx.GetData()
	if existing == nil {
		return Errorf("Key with %v doesn't exists", id)
	}
	var deal Deal
	json.Unmarshal(existing, &deal)
	if deal.DealOwner != requester { // new logic will be implemented when ecert is added
		return Errorf("Owner missmatch")
	}
	deal.UpdatedBy = requester
	deal.UpdatedDate = time.Now().Unix()

	deal.DealNextStep = newUp

	existing, _ = json.Marshal(deal)

	return ctx.GetStub().PutState(id, existing)
}

func (c *Chaincode) UpdateDealNextStepDate(ctx CustomTransactionContextInterface, id, newUp, requester string) error {
	existing := ctx.GetData()
	if existing == nil {
		return Errorf("Key with %v doesn't exists", id)
	}
	var deal Deal
	json.Unmarshal(existing, &deal)
	if deal.DealOwner != requester { // new logic will be implemented when ecert is added
		return Errorf("Owner missmatch")
	}
	deal.UpdatedBy = requester
	deal.UpdatedDate = time.Now().Unix()

	deal.DealNextStepDate = newUp

	existing, _ = json.Marshal(deal)

	return ctx.GetStub().PutState(id, existing)
}

func (c *Chaincode) UpdateDealLeadSource(ctx CustomTransactionContextInterface, id, newUp, requester string) error {
	existing := ctx.GetData()
	if existing == nil {
		return Errorf("Key with %v doesn't exists", id)
	}
	var deal Deal
	json.Unmarshal(existing, &deal)
	if deal.DealOwner != requester { // new logic will be implemented when ecert is added
		return Errorf("Owner missmatch")
	}
	deal.UpdatedBy = requester
	deal.UpdatedDate = time.Now().Unix()

	deal.DealLeadSource = newUp

	existing, _ = json.Marshal(deal)

	return ctx.GetStub().PutState(id, existing)
}
func (c *Chaincode) UpdateCampaignSource(ctx CustomTransactionContextInterface, id, newUp, requester string) error {
	existing := ctx.GetData()
	if existing == nil {
		return Errorf("Key with %v doesn't exists", id)
	}
	var deal Deal
	json.Unmarshal(existing, &deal)
	if deal.DealOwner != requester { // new logic will be implemented when ecert is added
		return Errorf("Owner missmatch")
	}
	deal.UpdatedBy = requester
	deal.UpdatedDate = time.Now().Unix()

	deal.DealCampaignSource = newUp

	existing, _ = json.Marshal(deal)

	return ctx.GetStub().PutState(id, existing)
}

func (c *Chaincode) UpdateDealContactName(ctx CustomTransactionContextInterface, id, newUp, requester string) error {
	existing := ctx.GetData()
	if existing == nil {
		return Errorf("Key with %v doesn't exists", id)
	}
	var deal Deal
	json.Unmarshal(existing, &deal)
	if deal.DealOwner != requester { // new logic will be implemented when ecert is added
		return Errorf("Owner missmatch")
	}
	deal.UpdatedBy = requester
	deal.UpdatedDate = time.Now().Unix()

	deal.DealContactName = newUp

	existing, _ = json.Marshal(deal)

	return ctx.GetStub().PutState(id, existing)
}

func (c *Chaincode) UpdateDealCurrencyCode(ctx CustomTransactionContextInterface, id, newUp, requester string) error {
	existing := ctx.GetData()
	if existing == nil {
		return Errorf("Key with %v doesn't exists", id)
	}
	var deal Deal
	json.Unmarshal(existing, &deal)
	if deal.DealOwner != requester { // new logic will be implemented when ecert is added
		return Errorf("Owner missmatch")
	}
	deal.UpdatedBy = requester
	deal.UpdatedDate = time.Now().Unix()

	deal.DealCurrencyCode = newUp

	existing, _ = json.Marshal(deal)

	return ctx.GetStub().PutState(id, existing)
}

func (c *Chaincode) UpdateDealAmount(ctx CustomTransactionContextInterface, id string, newUp float64, requester string) error {
	existing := ctx.GetData()
	if existing == nil {
		return Errorf("Key with %v doesn't exists", id)
	}
	var deal Deal
	json.Unmarshal(existing, &deal)
	if deal.DealOwner != requester { // new logic will be implemented when ecert is added
		return Errorf("Owner missmatch")
	}
	deal.UpdatedBy = requester
	deal.UpdatedDate = time.Now().Unix()

	deal.DealAmount = newUp

	existing, _ = json.Marshal(deal)

	return ctx.GetStub().PutState(id, existing)
}

func (c *Chaincode) UpdateClosingDate(ctx CustomTransactionContextInterface, id, newUp, requester string) error {
	existing := ctx.GetData()
	if existing == nil {
		return Errorf("Key with %v doesn't exists", id)
	}
	var deal Deal
	json.Unmarshal(existing, &deal)
	if deal.DealOwner != requester { // new logic will be implemented when ecert is added
		return Errorf("Owner missmatch")
	}
	deal.UpdatedBy = requester
	deal.UpdatedDate = time.Now().Unix()

	deal.DealClosingDate = newUp

	existing, _ = json.Marshal(deal)

	return ctx.GetStub().PutState(id, existing)
}

func (c *Chaincode) UpdateDealStage(ctx CustomTransactionContextInterface, id, newUp, requester string) error {
	existing := ctx.GetData()
	if existing == nil {
		return Errorf("Key with %v doesn't exists", id)
	}
	var deal Deal
	json.Unmarshal(existing, &deal)
	if deal.DealOwner != requester { // new logic will be implemented when ecert is added
		return Errorf("Owner missmatch")
	}
	deal.UpdatedBy = requester
	deal.UpdatedDate = time.Now().Unix()

	deal.DealStage = newUp

	existing, _ = json.Marshal(deal)

	return ctx.GetStub().PutState(id, existing)
}

func (c *Chaincode) UpdateDealProbility(ctx CustomTransactionContextInterface, id, newUp, requester string) error {
	existing := ctx.GetData()
	if existing == nil {
		return Errorf("Key with %v doesn't exists", id)
	}
	var deal Deal
	json.Unmarshal(existing, &deal)
	if deal.DealOwner != requester { // new logic will be implemented when ecert is added
		return Errorf("Owner missmatch")
	}
	deal.UpdatedBy = requester
	deal.UpdatedDate = time.Now().Unix()

	deal.DealProbility = newUp

	existing, _ = json.Marshal(deal)

	return ctx.GetStub().PutState(id, existing)
}

func (c *Chaincode) UpdateExpectedRevenue(ctx CustomTransactionContextInterface, id, newUp, requester string) error {
	existing := ctx.GetData()
	if existing == nil {
		return Errorf("Key with %v doesn't exists", id)
	}
	var deal Deal
	json.Unmarshal(existing, &deal)
	if deal.DealOwner != requester { // new logic will be implemented when ecert is added
		return Errorf("Owner missmatch")
	}
	deal.UpdatedBy = requester
	deal.UpdatedDate = time.Now().Unix()

	deal.DealExpectedRevenue = newUp

	existing, _ = json.Marshal(deal)

	return ctx.GetStub().PutState(id, existing)
}
