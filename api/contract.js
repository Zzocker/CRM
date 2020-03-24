const {Wallets,Gateway} = require('fabric-network')
const yaml = require('js-yaml')
const fs = require('fs')

const walletPath = './client/wallet'
const ccp = yaml.safeLoad(fs.readFileSync('./client/connection.yaml'))

const CHAINCODE="crm"
const CHANNEL = "crmchannel"
const CLIENT = "211"

const contract = async (type,inputs,callback)=>{

    try {
        const wallet = await Wallets.newFileSystemWallet(walletPath)
        const client = await wallet.get(CLIENT)
        if (!client){
            console.log(`${CLIENT} doesn't exists`)
            return callback(`${CLIENT} doesn't exists`,null)
        }

       const gateway = new Gateway() 
       await gateway.connect(ccp,{wallet,identity:CLIENT,discovery: { enabled: false, asLocalhost: true }})
       const network = await gateway.getNetwork(CHANNEL)
       const contract = network.getContract(CHAINCODE)
       var payload
       if (type=="invoke"){
            payload = await contract.submitTransaction(...inputs)
       }
       else if (type == "query"){
        // const payload = await contract.evaluateTransaction("GetAllMyLeads","created","}}","21")
            payload = await contract.evaluateTransaction(...inputs)
       }
        return callback(null,payload)
    } catch (error) {
       return callback(error,null)
    }
}
module.exports = {contract}