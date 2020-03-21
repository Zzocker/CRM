 2067  configtxgen -profile Genesis -outputBlock ./deployment/channel-artifacts/genesis.block -channelID genesis
 2068  configtxgen -outputCreateChannelTx ./deployment/channel-artifacts/channel.tx -channelID crmchannel -profile CRM


peer chaincode invoke -C crmchannel -n crm -o orderer:7050 -c '{"args":["CreateNewLead"," ","Mr","Pritam","Singh","9119216041","21"]}' --tls --cafile $ORDERER 

 export ORDERER=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/orderer.com/orderers/orderer/msp/tlscacerts/tlsca.orderer.com-cert.pem



1. export ORDERER=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/orderer.com/orderers/orderer/msp/tlscacerts/tlsca.orderer.com-cert.pem
2. peer channel create -f channel.tx -o orderer:7050 -c crmchannel --tls --cafile $ORDERER
3. peer channel join -b crmchannel.block
3. peer lifecycle chaincode package crm.tar.gz --label crm -p chaincode/
4. peer lifecycle chaincode install crm.tar.gz
5. export CC_PACKAGE=crm:233685aeb9d4362ee1d5181713f106ac2c669297198fc9f8bb6bdbe46d677d65
6. peer lifecycle chaincode approveformyorg -C crmchannel -n crm --package-id $CC_PACKAGE -v 1.0 -o orderer:7050 --tls --cafile $ORDERER
7. peer lifecycle chaincode commit -C crmchannel -n crm -v 1.0 -o orderer:7050 --tls --cafile $ORDERER