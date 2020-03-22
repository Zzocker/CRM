const {Wallets,Gateway} = require('fabric-network')
const yaml = require('js-yaml')
const fs = require('fs')

const walletPath = './client/wallet'
const ccp = yaml.safeLoad(fs.readFileSync('./client/connection.yaml'))

const CHAINCODE="crm"
const CHANNEL = "crmchannel"
const CLIENT = "211"

const contract = async ()=>{

    try {
        const wallet = await Wallets.newFileSystemWallet(walletPath)
        const client = await wallet.get(CLIENT)
        if (!client){
            console.log(`${CLIENT} doesn't exists`)
            return
        }

       const gateway = new Gateway() 
       await gateway.connect(ccp,{wallet,identity:CLIENT,discovery: { enabled: false, asLocalhost: true }})
       const network = await gateway.getNetwork(CHANNEL)
       const contract = network.getContract(CHAINCODE)
    //    const res= await contract.evaluateTransaction("GetAllMyLeads","created","}}","21")
    //    console.log(JSON.parse(res))
        return contract
    } catch (error) {
        console.log(error)
    }
}
module.exports = {contract}