package linode

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/pharmer/cloud/pkg/util"
)

func TestRegion(t *testing.T) {
	client, err := NewClient(tgetToken())
	if err != nil {
		t.Error(err)
	}
	rList, err := client.GetRegions()
	if err != nil {
		t.Error(err)
	}
	for _, r := range rList {
		fmt.Println(r.Spec.Location)
	}
}

func TestInstance(t *testing.T) {
	client, err := NewClient(tgetToken())
	if err != nil {
		t.Error(err)
	}
	iList, err := client.GetMachineTypes()
	if err != nil {
		t.Error(err)
	}
	for _, i := range iList {
		fmt.Println(i.Spec.Description)
	}
}

func tgetToken() string {
	b, _ := util.ReadFile("/home/ac/Downloads/cred/linode.json")
	v := struct {
		Token string `json:"token"`
	}{}
	fmt.Println(json.Unmarshal(b, &v))
	//fmt.Println(v)
	return v.Token
}