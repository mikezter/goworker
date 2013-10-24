package goworker

import (
        "fmt"
        "testing"
)

type ifaceJob struct {
        Args []interface{}
}

func (theJob ifaceJob) Work() error {
        arg := theJob.Args[0]
        one := arg.(map[string]interface{})

        var conts []float64
        for _, i := range one["contingents"].([]interface{}) {
                conts = append(conts, i.(float64))
        }

        var locs []location
        for _, i := range one["locations"].([]interface{}) {
                val := i.(map[string]interface{})
                id := val["id"].(string)

                var contingents []int
                for _, i := range val["contingents"].([]interface{}) {
                        contingents = append(contingents, int(i.(float64)))
                }

                locs = append(locs, location{Id: id, Contingents: contingents})
        }

        var clusts []cluster

        for _, i := range one["clusters"].([]interface{}) {
                val := i.(map[string]interface{})
                dqw := val["dqw"].(float64)
                cpqa := val["cpqa"].(float64)
                clusts = append(clusts, cluster{Cca: dqw, Dqb: cpqa})
        }

        fmt.Println(conts)
        fmt.Println(locs)
        fmt.Println(clusts)

        return nil
}

func newIfaceJob() Worker {
        return Worker(&ifaceJob{})
}

func TestEmptyInterface(t *testing.T) {

        queues.Set("gofoo")
        exitOnComplete = true
        concurrency = 2

        EnqueueJSON(data)

        Register("MyClass", newIfaceJob)

        if err := Work(); err != nil {
                t.Error("Error:", err)
        }

}
