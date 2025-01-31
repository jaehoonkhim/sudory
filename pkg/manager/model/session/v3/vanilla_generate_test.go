package sessions_test

import (
	"os"
	"testing"

	"github.com/NexClipper/sudory/pkg/manager/database/vanilla/ice_cream_maker"
	sessions "github.com/NexClipper/sudory/pkg/manager/model/session/v3"
)

var objs = []interface{}{
	// v3.Session_essential{},
	sessions.Session{},
}

func TestNoXormColumns(t *testing.T) {
	s, err := ice_cream_maker.GenerateParts(objs, ice_cream_maker.Ingredients)
	if err != nil {
		t.Fatal(err)
	}

	println(s)

	if true {
		filename := "vanilla_generated.go"
		fd, err := os.Create(filename)
		if err != nil {
			t.Fatal(err)
		}

		if _, err = fd.WriteString(s); err != nil {
			t.Fatal(err)
		}
	}
}
