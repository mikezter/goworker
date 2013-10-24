package goworker

import (
        "fmt"
        "testing"
)

func newJob() Worker {
        return Worker(&testjob{})
}

func (theJob testjob) Work() error {
        theParams := theJob.Args[0]
        fmt.Println(theParams)
        fmt.Println(theParams.Clusters)
        fmt.Println(theParams.Clusters[0].Cca)
        return nil
}

func TestStructAssignment(t *testing.T) {
        t.Skip()
        queues.Set("gofoo")
        exitOnComplete = true
        concurrency = 2

        EnqueueJSON(data)

        Register("MyClass", newJob)

        if err := Work(); err != nil {
                t.Error("Error:", err)
        }

}
