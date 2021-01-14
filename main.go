package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/peersafe/deployFabricTool/cmd"
)

var (
	file               = flag.String("f", "", "configtx, crypto-config, node, client, jmeter, zabbix ' create yaml file '")
	start              = flag.String("s", "", "start peer, order, explorer, zookeeper, kafka, all ,api, jmeter,nmon, zabbix 'start node or api'")
	stop               = flag.String("d", "", "stop peer , order, explorer ,kafka , zookeeper , all , api")
	create             = flag.String("c", "", "crypto, genesisblock, channel, 'create source'")
	getlog             = flag.String("g", "", "get jmeter or event or nmon logs")
	logdir             = flag.String("gn", "", "log dir name eg: 50_50  loop 50*50")
	channelname        = flag.String("n", "", "channelname")
	ccname             = flag.String("ccname", "", "chaincode name")
	ccversion          = flag.String("version", "", "chaincode version")
	ccpath             = flag.String("ccpath", "", "chaincode go path")
	testArgs           = flag.String("args", "", "test chaincode args")
	function           = flag.String("func", "invoke", "invoke or query")
	run                = flag.String("r", "", "joinchannel,  updateanchor, installchaincode, runchaincode, " +
		"createpeeryaml, addorgnodecert,putnodecrypto,runaddnode,installcctonewnode,runcctonewnode,checknode, upgradecc,testcc")
	put                = flag.String("p", "", "put all (include crypto-config and channel-artifacts to remote)")
	removeData         = flag.String("rm", "", "remove mount data")
	analyse            = flag.String("a", "", "event analyse")
	orgid              = flag.String("orgid", "", "orgid")
	nodename           = flag.String("nodename", "", "nodename")
)

func main() {
	flag.Parse()
	var err error
	cmd.GlobalConfig, err = cmd.ParseJson("node.json")
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%#v\n",cmd.GlobalConfig.Apiservers)
	if *file != "" {
		if *file == "jmeter" {
			err = cmd.CreateJmeterConfig()
		} else if *file == "haproxy" {
			err = cmd.CreateHaproxyConfig()
		} else {
			err = cmd.CreateYamlByJson(*file)
		}
	} else if *start != "" {
		if *start == "jmeter" {
			err = cmd.CreateJmeterConfig()
			if err == nil {
				err = cmd.StartJmeter()
			}
		} else if *start == "haproxy" {
			err = cmd.StartHaproxy()
		} else {
			err = cmd.HandleNode(*start, true)
		}
	} else if *create == "genesisblock" {
		err = cmd.CreateGenesisBlock()
	} else if *create == "crypto-config" {
		err = cmd.CreateCert()
	} else if *create == "channel" {
		err = cmd.CreateChannel(*channelname)
	} else if *run == "updateanchor" {
		err = cmd.UpdateAnchor(*channelname)
	} else if *run == "joinchannel" {
		err = cmd.JoinChannel(*channelname, *nodename)
	} else if *run == "installchaincode" {
		err = cmd.InstallChaincode(*ccname, *ccversion, *channelname, *ccpath)
	} else if *run == "runchaincode" {
		err = cmd.RunChaincode(*ccname, *ccversion, *channelname, "instantiate")
	} else if *run == "upgradecc" {
		err = cmd.RunChaincode(*ccname, *ccversion, *channelname, "upgrade")
	} else if *run == "testcc" {
		err = cmd.TestChaincode(*ccname, *channelname, *function, *testArgs)
	} else if *run == "checknode" {
		err = cmd.CheckNode("all")
	} else if *getlog == "jmeter" {
		err = cmd.GetJmeterLog(*logdir)
	} else if *getlog == "event" {
		err = cmd.GetEventServerLog(*logdir)
	} else if *put != "" {
		err = cmd.PutCryptoConfig(*put)
	} else if *stop != "" {
		err = cmd.HandleNode(*stop, false)
	} else if *removeData != "" {
		err = cmd.RemoveData(*removeData)
	} else if *analyse != "" {
		err = cmd.EventAnalyse(*logdir)
	} else if *run == "addorgnodecert" {
		err = cmd.AddOrgNodeCertById(*orgid)
	} else if *run == "putnodecrypto" {
		err = cmd.PutNodeCrypto(*nodename)
	} else if *run == "createpeeryaml" {
		err = cmd.CreatePeerYaml(*nodename)
	} else if *run == "runaddnode" {
		err = cmd.RunAddNode(*nodename)
	} else if *run == "installcctonewnode" {
		err = cmd.InstallCCToNewNode(*ccname, *ccversion, *ccpath, *nodename)
	} else if *run == "runcctonewnode" {
		err = cmd.RunCCToNewNode(*ccname, *ccversion, *channelname)
	} else {
		fmt.Println("Both data and file are nil.")
		flag.Usage()
		os.Exit(1)
	}
	if err != nil {
		panic(err)
	}
}
