# CRM (Customer Relationship Management Software)

>## Table of Contents

- [Installation](#installation)
- [Status](#status)
- [Documentation](#documentation)

## Installation

### Clone
- Clone this repo to your local machine using `https://github.com/Zzocker/CRM.git`

### Install docker (Optional)
- chmod +x ./docker.sh
- ./docker.sh
- usermod -a -G docker ${USER}

### Fabric network

#### Configuration of fabric network (v2.0)
> 2 Orderer using **raft** consensus

> 1 org named **Peepaltree** with one peer name devpeer. devpeer is using **Couchdb** as worldstate.

> Chaincode written in **Golang**

#### Start Docker Containers and etup the peers

    > cd network/deployment/
    > docker-compose -f docker-compose-cli.yaml -f docker-compose-orderer.yaml -f docker-compose-peerca.yaml up -d
    > docker exec -it cli bash
    > cd channel-artifacts/
    > chmod +x chaincodetest.sh
    > ./chaincodetest.sh

#####  Install chaincode on peer

    > export ORDERER=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/orderer.com/orderers/orderer/msp/tlscacerts/tlsca.orderer.com-cert.pem
    > peer lifecycle chaincode package crm.tar.gz --label crm -p 
    > peer lifecycle chaincode install crm.tar.gz
    > export CC_PACKAGE=$(peer lifecycle chaincode queryinstalled -O json | jq .installed_chaincodes | jq .[0] | jq .package_id)
    > peer lifecycle chaincode approveformyorg -C crmchannel -n crm --package-id $CC_PACKAGE -v 1.0 -o orderer:7050 --tls --cafile $ORDERER
    > peer lifecycle chaincode commit -C crmchannel -n crm -v 1.0 -o orderer:7050 --tls --cafile $ORDERER

### Fire up the APIs
    
    > cd api
    > npm i
    > npm run dev

## Status

:white_large_square: Chaincode development
- :heavy_check_mark: Lead model
- :heavy_check_mark: Deal model
- :heavy_check_mark: Lead and Deal Function 
- :white_large_square: Access control of Admin, salesperson,manger on deals and leads 

:heavy_check_mark: Network configured

