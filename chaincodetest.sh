
export ORDERER=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/orderer.com/orderers/orderer/msp/tlscacerts/tlsca.orderer.com-cert.pem
peer lifecycle chaincode package crm.tar.gz --label crm -p chaincode/
peer lifecycle chaincode install crm.tar.gz

packageID=$(peer lifecycle chaincode queryinstalled -O json | jq .installed_chaincodes | jq .[0] | jq .package_id)

export CC_PACKAGE=$packageID

echo $CC_PACKAGE

# peer lifecycle chaincode approveformyorg -C crmchannel -n crm --package-id $CC_PACKAGE -v 1.0 -o orderer:7050 --tls --cafile $ORDERER

# peer lifecycle chaincode commit -C crmchannel -n crm -v 1.0 -o orderer:7050 --tls --cafile $ORDERER

# peer chaincode invoke -C crmchannel -n crm -o orderer:7050 -c '{"args":["CreateNewLead"," ","Mr","Pritam","Singh","9119216041","21"]}' --tls --cafile $ORDERER 

# peer chaincode query -C crmchannel -n crm -c '{"args":["GetAllMyLeads","start","21"]}'
