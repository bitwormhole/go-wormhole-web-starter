package gin_starter

import (
	"github.com/bitwormhole/go-wormhole-core/application"
	"github.com/gin-gonic/gin"
)

type GinWebController interface {
	Config(engine *gin.Engine) error
}

type GinServerContainer struct {
	port        int
	engine      *gin.Engine
	controllers []GinWebController
}

func (inst *GinServerContainer) Inject(context application.RuntimeContext) error {

	components := context.GetComponents()
	keys := components.GetComponentNameList(false)
	ctrlist := []GinWebController{}

	for index := range keys {
		key := keys[index]
		com, err := components.GetComponent(key)
		if err != nil {
			return err
		}
		controller, ok := com.(GinWebController)
		if ok {
			ctrlist = append(ctrlist, controller)
		}
	}

	inst.controllers = ctrlist
	return nil
}

func (inst *GinServerContainer) initEngine() error {

	server := gin.Default()

	server.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"hello":  "world",
			"server": "gin",
			"port":   666,
		})
	})

	inst.engine = server
	return nil
}

func (inst *GinServerContainer) initControllers() error {

	controllers := inst.controllers
	engine := inst.engine

	for index := range controllers {
		ctrl := controllers[index]
		err := ctrl.Config(engine)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *GinServerContainer) Run() error {
	inst.engine.Run(":10217")
	return nil
}

func (inst *GinServerContainer) Init() error {

	err := inst.initEngine()
	if err != nil {
		return err
	}

	err = inst.initControllers()
	if err != nil {
		return err
	}

	return nil
}

func (inst *GinServerContainer) Destroy() error {

	return nil
}
