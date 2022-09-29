package test

import (
	"os"
	"testing"

	"github.com/i18n/pkg/core"
)

func TestExport(t *testing.T) {
	file_name := "./test.csv"
	os.Remove(file_name)
	defer os.Remove(file_name)
	err := core.Export([]string{"zh_cn", "en_us", "key"},[][]string{
		{"zh_cn0", "en_us0", "key0"},
		{"zh_cn1", "en_us1", "key1"},
		{"zh_cn2", "en_us2", "key2"},
	},file_name)
	if err != nil{
		t.Errorf(err.Error());
	}

	_, e := os.Stat(file_name)
	if os.IsNotExist(e) {
		t.Errorf(e.Error());
	}
	
}