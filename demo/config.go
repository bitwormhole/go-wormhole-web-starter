package demo

import (
	"github.com/bitwormhole/go-wormhole-core/application"
	"github.com/bitwormhole/go-wormhole-core/application/config"
	"github.com/bitwormhole/go-wormhole-core/lang"
	"github.com/bitwormhole/go-wormhole-web-starter/data/gorm_starter"
	"github.com/bitwormhole/go-wormhole-web-starter/web/gin_starter"
)

func Config(cfg application.ConfigBuilder) {
	gin_starter.Config(cfg)
	gorm_starter.Config(cfg)

	cfg.AddComponent(&config.ComInfo{
		OnNew: func() lang.Object {
			return &exampleController1{}
		},
	})

	cfg.AddComponent(&config.ComInfo{

		ID: "exampleController2",

		OnNew: func() lang.Object {
			return &exampleController2{}
		},

		OnInject: func(t lang.Object, context application.RuntimeContext) error {
			return t.(*exampleController2).inject(context)
		},

		OnInit: func(t lang.Object) error {
			return t.(*exampleController2).init()
		},

		OnDestroy: func(t lang.Object) error {
			return t.(*exampleController2).destroy()
		},
	})

}
