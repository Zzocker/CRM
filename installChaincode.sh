if ! [ -x "$(command -v jq)" ]; then
  echo "Wait Installing Jq"
  bash snap install jq
fi
ORDERER_PATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/orderer.com/orderers/orderer/msp/tlscacerts/tlsca.orderer.com-cert.pem
sequence=$(docker exec cli peer lifecycle chaincode querycommitted -C crmchannel -O json | jq .chaincode_definitions[0].sequence)
next=$(($sequence+1))
echo "Removing crm.tat.gz file"
docker exec cli  rm crm.tar.gz
echo "Packaging chaincode"
docker exec cli peer lifecycle chaincode package crm.tar.gz --label crm -p chaincode
echo "installing chaincode"
docker exec cli peer lifecycle chaincode install crm.tar.gz
PACKAGE_ID=$(docker exec cli  peer lifecycle chaincode queryinstalled -O json | jq .installed_chaincodes[0].package_id -r)
echo $PACKAGE_ID
echo "approving for my org"
docker exec cli peer lifecycle chaincode approveformyorg -C crmchannel -n crm --package-id $PACKAGE_ID -v 1.0 -o orderer:7050 --tls --cafile $ORDERER_PATH --sequence $next
echo "commiting chaincode"
docker exec cli peer lifecycle chaincode commit -C crmchannel -n crm -v 1.0 -o orderer:7050 --tls --cafile $ORDERER_PATH --sequence $next
echo "Done"