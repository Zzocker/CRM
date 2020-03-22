const fabricCAService = require('fabric-ca-client')
const { Wallets } = require('fabric-network')
const yaml = require('js-yaml')
const fs = require('fs')

const walletpath = './wallet'
const ccpPath = './connection.yaml'

const ccp = yaml.safeLoad(fs.readFileSync(ccpPath))

const main = async () =>{
    try {
        caConfig = ccp.certificateAuthorities[ccp.organizations.Peepaltree.certificateAuthorities[0]]
        // console.log(caConfig)
        const ca = new fabricCAService(caConfig.url,{trustedRoots:caConfig.pem,verify:false},caConfig.caName)
        
        const wallet = await Wallets.newFileSystemWallet(walletpath)
        const exists = await wallet.get('admin')
        if (exists){
            console.log("admin already exists in the wallet")
        }

        const enrollment = await ca.enroll({ enrollmentID: 'admin', enrollmentSecret: 'adminpw' })
        const X509Identity = {
            credentials: {
                certificate: enrollment.certificate,
                privateKey: enrollment.key.toBytes(),
            },
            mspId: 'PeepaltreeMSP',
            type: 'X.509',
        }
        await wallet.put('admin',X509Identity)
        console.log('Successfully enrolled admin user "admin" and imported it into the wallet')
    } catch (error) {
       console.log(error) 
       
    }
    finally{
        process.exit(1)
    }
}
main()