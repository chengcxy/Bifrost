package warning_test

import (
	"github.com/brokercap/Bifrost/server/warning"
	"log"
	"testing"
)

func TestDingTalkSendWarning(t *testing.T) {
	obj := &warning.DingTalkRoboter{}
	p := make(map[string]interface{}, 0)
	p["token"] = "769cb96830c374d5a691a10b4f0efb1a004ba9d3c958046dd4d3d38657365a"
	p["atMobiles"] = []string{"18811788263"}
	p["isAtAll"] = false
	p["hook_keyword"] = "任务报警"
	err := obj.SendWarning(p, "test warning title", "it is test")
	if err != nil {
		t.Errorf(err.Error())
	} else {
		log.Println("success")
	}
}
