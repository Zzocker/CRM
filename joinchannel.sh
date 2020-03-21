export ORDERER=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/orderer.com/orderers/orderer/msp/tlscacerts/tlsca.orderer.com-cert.pem
peer channel create -f channel.tx -o orderer:7050 -c crmchannel --tls --cafile $ORDERER
peer channel join -b crmchannel.block

echo $(peer channel list)