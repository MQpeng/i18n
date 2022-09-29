package test

import (
	"log"
	"testing"

	"github.com/i18n/pkg/core"
)

func TestByGlob(t *testing.T) {
	source, _ := core.ByGlob("../pkg/**/*.go")
	log.Println(source)
	t.Error("asdf")
}