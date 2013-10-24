package goworker

import "log"

const data = string(`
  { "class":"MyClass",
    "args":[{
      "locations":[
        { "id":"1234", "contingents":[2, 3]},
        { "id":"2afe", "contingents":[1, 0, 4]},
        { "id":"9283", "contingents":[2, 1, 4]}
      ],
      "contingents":[100,200,300,400,500,199.4],
      "clusters":[{"cca":23.4,"dqb":29.7}]
  }] }
`)


type testjob struct {
	Args  []arg
}

type arg struct {
	Locations []location
	Contingents []float64
	Clusters []cluster
}

type location struct {
	Id          string
	Contingents []int
}

type cluster struct {
	Cca float64
	Dqb  float64
}

func EnqueueJSON(json string) {
    conn, err := redisConnFromUri("redis://localhost:6379")

    if err != nil {
        log.Println(err)
    }

    for i := 0; i < 1; i++ {
        conn.Do("RPUSH", "resque:queue:gofoo", json)
    }

    if _, err := conn.Do(""); err != nil {
        log.Println(err)
        return
    }

}
