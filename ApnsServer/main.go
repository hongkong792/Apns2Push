package main

import (
	"log"
	"fmt"

	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/certificate"
)

func main() {

	cert, err := certificate.FromP12File("/Users/yaowan/Downloads/cert.p12", "123456")
	if err != nil {
		log.Fatal("Cert Error:", err)
	}

	notification := &apns2.Notification{}
	notification.DeviceToken = "ea7f974617a338aaa84fee8b7a7d6516794f1103a2771d20133cc01ac0d3cffb"
	notification.Topic = "com.xiyouzhi.danaotianzhu"
	//notification.Payload = []byte(`{"aps":{"alert":"Hello!"}}`) // See Payload section below
	notification.Payload = []byte(`{
    "aps": {
        "category": "category",
        "content-available": 1,
        "alert": {
            "launch-image": "icon.icon",
            "action-loc-key": "LocalizedActionButtonKey",
            "loc-key": "LocalizedAlertMessage",
            "subtitle-loc-key": "LocalizedAlertSubtitle",
            "title-loc-key": "LocalizedAlertTitle"
        },
        "url": "https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1543493128&di=ec8f8beae886e173fdc763a688ce4363&imgtype=jpg&er=1&src=http%3A%2F%2Fwww.znsfagri.com%2Fuploadfile%2Feditor%2Fimage%2F20170626%2F20170626151136_11631.jpg",
        "sound": "sound.wav",
        "badge": 2,
        "mutable-content": 1
    },
    "name": "value",
    "taskId": "9988",
    "url": "https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1543493128&di=ec8f8beae886e173fdc763a688ce4363&imgtype=jpg&er=1&src=http%3A%2F%2Fwww.znsfagri.com%2Fuploadfile%2Feditor%2Fimage%2F20170626%2F20170626151136_11631.jpg"
       

}`) // See Payload section below


	client := apns2.NewClient(cert).Development()
	res, err := client.Push(notification)

	if err != nil {
		log.Fatal("push Error:", err)
	}

	fmt.Printf("%v %v %v\n", res.StatusCode, res.ApnsID, res.Reason)
}