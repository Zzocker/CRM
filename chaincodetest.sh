
export ORDERER=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/orderer.com/orderers/orderer/msp/tlscacerts/tlsca.orderer.com-cert.pem
peer lifecycle chaincode package crm.tar.gz --label crm -p chaincode/
peer lifecycle chaincode install crm.tar.gz

packageID=$(peer lifecycle chaincode queryinstalled -O json | jq .installed_chaincodes | jq .[0] | jq .package_id)

export CC_PACKAGE=$packageID

echo $CC_PACKAGE
