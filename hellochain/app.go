package hellochain

import (
	"log"

	"github.com/cosmos/sdk-tutorials/hellochain/starter"
	abci "github.com/tendermint/tendermint/abci/types"
	dbm "github.com/tendermint/tm-db"
)

const appName = "hellochain"

var (
	// ModuleBasics holds the AppModuleBasic struct of all modules included in the app
	ModuleBasics = starter.ModuleBasics
)

type helloChainApp struct {
	*starter.AppStarter
}

func NewHelloChainApp(logger log.Logger, db dbm.DB) abci.Application {
	appStarter := starter.NewAppStarter(appName, logger, db)

	var app = &helloChainApp{
		appStarter,
	}

	app.InitializeStarter()

	return app
}
