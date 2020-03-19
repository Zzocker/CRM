configtxgen -profile SampleMultiNodeEtcdRaft -outputBlock ./channel-artifacts/genesis.block -channelID genesis

configtxgen -outputCreateChannelTx ./channel-artifacts/channel.tx -channelID crmchannel -profile CRM
