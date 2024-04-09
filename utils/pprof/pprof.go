package pprof

import (
	"cloud.google.com/go/profiler"
)

func Launch() {
	if err := profiler.Start(profiler.Config{Service: "my-service", ServiceVersion: "v1"}); err != nil {
	}
	// ssh -l yangbo -p 2222 yw-jump.ttyuyin.com -i ~/.ssh/yangbo-jumpserver-1.pe
}
