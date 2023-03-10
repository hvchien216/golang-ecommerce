package appctx

import (
	"github.com/hvchien216/golang-ecommerce/internal/pkg/db/pg"
)

type AppContext interface {
	GetMainDBConnection() pg.PGExecutor
	//UploadProvider() uploadprovider.UploadProvider
	//GetSecretKey() string
	//NewTokenConfig() TokenConfig
	//GetPubsub() pubsub.Pubsub
}

type appCtx struct {
	db pg.PGExecutor
	//uploadProvider uploadprovider.UploadProvider
	//secretKey      string
	//pb             pubsub.Pubsub
}

func NewAppContext(
	db pg.PGExecutor,
	// uploadProvider uploadprovider.UploadProvider,
	// secretKey string,
	// pb pubsub.Pubsub,
) *appCtx {
	return &appCtx{
		db: db,
		//uploadProvider: uploadProvider,
		//secretKey:      secretKey,
		//pb:             pb,
	}
}

func (ctx *appCtx) GetMainDBConnection() pg.PGExecutor {
	return ctx.db
}

//func (ctx *appCtx) UploadProvider() uploadprovider.UploadProvider {
//	return ctx.uploadProvider
//}
//
//func (ctx *appCtx) GetSecretKey() string {
//	return ctx.secretKey
//}
//
//func (ctx *appCtx) GetPubsub() pubsub.Pubsub {
//	return ctx.pb
//}
