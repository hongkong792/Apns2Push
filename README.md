# Apns2Push
iOS Apps Push 推送 高效 简单
iOS服务端批量推送
#####1.官方文档：
https://developer.apple.com/library/archive/documentation/NetworkingInternet/Conceptual/RemoteNotificationsPG/CommunicatingwithAPNs.html#//apple_ref/doc/uid/TP40008194-CH11-SW1
### 2.踩坑经历，项目为了提高推送效率，之前推送用的socket长链接，不断的向apple的网关写数据，这种情况很糟糕，只有出错的情况，才会有反馈

##3.苹果新的解决方法

http2，对http2，关于http2请自己维基百科
## 4.理解原理
####问：批量推送你是采用的http，还是tcp的socket
###答:如果你是从0开始的话，建议你直接采用APNs最新的https/2.0协议，支持长连接，而且每次调用都有相应的应答,不要再用以前的tcp socket方式，有很多坑要躺，我是一路躺过来的 

####问：socket方式，是不是效率更好，更节省时间
###答:https/2.0协议具备了tcp和http的特性，效率也不差;而且支持长连接，每个请求APNs还会给你应答；以前的tcp socket协议，只有出错了才会给你应答，巨坑

####问：好像是，如果用户卸载了app，就是token失效了，socket就直接出错，断了
###答:遇到失效的socket，也不一定会socket直接退出，得看APNs的处理机制；APNs应该是当出错达到一定次数后，直接掐掉socket；很多坑要填，建议你直接用https/2.0，我都打算升级上去了。

####问：这个http2,每一个的返回值，能不能确定消息到达端了，方便统计
###答:不是到端，而是代表APNs正确接收到你的请求了，至于能否到app端，就得看实际情况了，这个是没有返回，也没有统计的，需要app自己做
![image.png](https://upload-images.jianshu.io/upload_images/2581808-6aa59bed198cc190.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

##5.明白了原理，推送很简单，几行代码：
```
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
```
完整可运行源码：别忘了加星哦
https://github.com/hongkong792/Apns2Push.git
