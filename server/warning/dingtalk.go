package warning

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

/*
钉钉报警配置  如果机器人设置了信息关键字 hook_keyword填写该关键字 否则无法推送成功
比如 我的机器人设置了"任务"2个字 只有消息里面含有任务字样才可以发送
"roboter": {
	"token": "token",
	"atMobiles": [
		"$mobile"
	],
	"isAtAll": false,
	"hook_keyword": "任务报警"
}
*/

func init() {
	Register("DingTalk", &DingTalkRoboter{})
}

//钉钉机器人post请求接口地址
var DingTalkBaseApi = "https://oapi.dingtalk.com/robot/send?access_token=%s"

type DingTalkParam struct {
	Token      string
	AtMobiles  []interface{}
	Hookeyword string
	IsAtall    bool
}

type DingTalkRoboter struct {
	p DingTalkParam
}

func (dt *DingTalkRoboter) paramTansfer(p map[string]interface{}) error {
	s, err := json.Marshal(p)
	if err != nil {
		return err
	}
	err2 := json.Unmarshal(s, &dt.p)
	if err2 != nil {
		return err2
	}
	return nil
}

func (dt *DingTalkRoboter) SendWarning(p map[string]interface{}, title string, Body string) error {
	err1 := dt.paramTansfer(p)
	if err1 != nil {
		return err1
	}
	payload, err := dt.GetPayload(Body, "NULL")
	if err != nil {
		log.Println("get dingtalk payload message error,", err)
		return err
	}
	contentType := "application/json;charset=utf-8"
	api := fmt.Sprintf(DingTalkBaseApi, dt.p.Token)
	client := &http.Client{
		Timeout: 15 * time.Second,
	}
	resp, err := client.Post(api, contentType, bytes.NewBuffer(payload))
	if err != nil {
		panic("send dingtalk msg err")
	}
	defer resp.Body.Close()
	_, err = ioutil.ReadAll(resp.Body)
	return err

}

func (dt *DingTalkRoboter) GetPayload(content, mobile string) ([]byte, error) {
	data := make(map[string]interface{})
	data["msgtype"] = "text"
	at := make(map[string]interface{})
	if mobile == "NULL" {
		at["atMobiles"] = dt.p.AtMobiles
	} else {
		atMobiles := make([]string, 1)
		atMobiles[0] = mobile
		at["atMobiles"] = atMobiles
	}
	at["isAtAll"] = dt.p.IsAtall
	data["at"] = at
	text := make(map[string]string)
	text["content"] = fmt.Sprintf("%s:%s", dt.p.Hookeyword, content)
	data["text"] = text
	return json.Marshal(data)
}
