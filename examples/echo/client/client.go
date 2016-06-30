package main

import (
  "time"
  "fmt"
  "github.com/leesper/tao"
  "github.com/leesper/tao/examples/echo"
  "github.com/golang/glog"
)

func main() {
  tao.MessageMap.Register(echo.EchoMessage{}.MessageNumber(), echo.DeserializeEchoMessage)

  tcpConnection := tao.NewClientConnection(0, "127.0.0.1:18342", false)

  tcpConnection.SetOnConnectCallback(func(client tao.Connection) bool {
    glog.Infoln("On connect")
    return true
  })

  tcpConnection.SetOnErrorCallback(func() {
    glog.Infoln("On error")
  })

  tcpConnection.SetOnCloseCallback(func(client tao.Connection) {
    glog.Infoln("On close")
  })

  tcpConnection.SetOnMessageCallback(func(msg tao.Message, c tao.Connection) {
    echoMessage := msg.(echo.EchoMessage)
    fmt.Printf("%s\n", echoMessage.Message)
  })

  echoMessage := echo.EchoMessage{
    Message: "hello, world",
  }

  tcpConnection.Start()

  for i := 0; i < 100; i++ {
    time.Sleep(600 * time.Millisecond)
    err := tcpConnection.Write(echoMessage)
    if err != nil {
      glog.Errorln(err)
    }
  }
  time.Sleep(time.Second)
  tcpConnection.Close()
}
