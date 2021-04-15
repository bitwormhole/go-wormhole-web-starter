package demo

import "github.com/gin-gonic/gin"

type exampleController1 struct {
}

func (inst *exampleController1) Config(engine *gin.Engine) error {

	path := "/api/v1/example1"

	engine.GET(path, func(ctx *gin.Context) { inst.httpGetEx1(ctx) })
	engine.PUT(path, func(ctx *gin.Context) { inst.httpGetEx1(ctx) })
	engine.POST(path, func(ctx *gin.Context) { inst.httpGetEx1(ctx) })
	engine.DELETE(path, func(ctx *gin.Context) { inst.httpGetEx1(ctx) })

	return nil
}

func (inst *exampleController1) httpGetEx1(ctx *gin.Context) {

	ctx.JSON(200, gin.H{})
}

func (inst *exampleController1) httpPostEx1(ctx *gin.Context) {

	ctx.JSON(200, gin.H{})
}

func (inst *exampleController1) httpPutEx1(ctx *gin.Context) {

	ctx.JSON(200, gin.H{})
}

func (inst *exampleController1) httpDeleteEx1(ctx *gin.Context) {

	ctx.JSON(200, gin.H{})
}
