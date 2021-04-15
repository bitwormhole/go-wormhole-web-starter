package demo

import (
	"github.com/bitwormhole/go-wormhole-core/application"
	"github.com/gin-gonic/gin"
)

type exampleController2 struct {
}

func (inst *exampleController2) Config(engine *gin.Engine) error {

	path := "/api/v1/example2"

	engine.GET(path, func(ctx *gin.Context) { inst.httpGetEx2(ctx) })
	engine.PUT(path, func(ctx *gin.Context) { inst.httpGetEx2(ctx) })
	engine.POST(path, func(ctx *gin.Context) { inst.httpGetEx2(ctx) })
	engine.DELETE(path, func(ctx *gin.Context) { inst.httpGetEx2(ctx) })

	return nil
}

func (inst *exampleController2) httpGetEx2(ctx *gin.Context) {

	ctx.JSON(200, gin.H{})
}

func (inst *exampleController2) httpPostEx2(ctx *gin.Context) {

	ctx.JSON(200, gin.H{})
}

func (inst *exampleController2) httpPutEx2(ctx *gin.Context) {

	ctx.JSON(200, gin.H{})
}

func (inst *exampleController2) httpDeleteEx2(ctx *gin.Context) {

	ctx.JSON(200, gin.H{})
}

func (inst *exampleController2) inject(ctx application.RuntimeContext) error {

	return nil
}

func (inst *exampleController2) init() error {

	return nil
}

func (inst *exampleController2) destroy() error {

	return nil
}
