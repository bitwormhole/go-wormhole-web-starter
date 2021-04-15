package gin_starter

import (
	"github.com/bitwormhole/go-wormhole-core/application"
	"github.com/bitwormhole/go-wormhole-core/application/config"
	"github.com/bitwormhole/go-wormhole-core/lang"
)

func Config(cfg application.ConfigBuilder) {

	cfg.AddComponent(&config.ComInfo{

		ID: "gin-web-server",

		OnNew: func() lang.Object {
			return &GinServerContainer{}
		},

		OnInject: func(t lang.Object, context application.RuntimeContext) error {
			return t.(*GinServerContainer).Inject(context)
		},

		OnInit: func(t lang.Object) error {
			return t.(*GinServerContainer).Init()
		},

		OnDestroy: func(t lang.Object) error {
			return t.(*GinServerContainer).Destroy()
		},
	})

}
