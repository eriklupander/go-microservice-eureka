package eureka

import (
        "os"
        "io/ioutil"
        "strings"
        "fmt"
        "time"
        "github.com/eriklupander/goeureka/util"
)

var instanceId string

func Register() {
        instanceId = util.GetUUID();

        dir, _ := os.Getwd()
        data, _ := ioutil.ReadFile(dir + "/templates/regtpl.json")

        tpl := string(data)
        tpl = strings.Replace(tpl, "${ipAddress}", util.GetLocalIP(), -1)
        tpl = strings.Replace(tpl, "${port}", "8080", -1)
        tpl = strings.Replace(tpl, "${instanceId}", instanceId, -1)

        // Register.
        registerAction := HttpAction {
                Url : "http://192.168.99.100:8761/eureka/apps/vendor",
                Method: "POST",
                ContentType: "application/json",
                Body: tpl,
        }
        var result bool
        for {
                result = DoHttpRequest(registerAction)
                if result {
                        break
                } else {
                        time.Sleep(time.Second * 5)
                }
        }
}

func StartHeartbeat() {
        for {
                time.Sleep(time.Second * 30)
                heartbeat()
        }
}

func heartbeat() {
        heartbeatAction := HttpAction{
                Url : "http://192.168.99.100:8761/eureka/apps/vendor/" + util.GetLocalIP() + ":vendor:" + instanceId,
                Method: "PUT",
        }
        DoHttpRequest(heartbeatAction)
}

func Deregister() {
        // Deregister
        deregisterAction := HttpAction {
                Url : "http://192.168.99.100:8761/eureka/apps/vendor/" + util.GetLocalIP() + ":vendor:" + instanceId,
                Method: "DELETE",
        }
        DoHttpRequest(deregisterAction)
        fmt.Println("Deregistered application, exiting. Check Eureka...")
}
