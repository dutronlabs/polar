package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "gopkg.in/redis.v3"
    "fmt"
    "encoding/json"
    "os"
)

type Configuration struct {
    address string
    pw string
    db int
}

func loadConfiguration() {
    file,_ := os.Open("conf.json")
    decoder := json.NewDecoder(file)
    configuration := Configuration{}
    err := decoder.Decode(&configuration)
    if err != nil {
         fmt.Println("Error loading configuration: ", err)
    }
    fmt.Println("Configuration loaded")
}


type Container struct {
    Id int
    Name string  `form:"name" json:"name" binding:"required"`
    Version string `form:"version" json:"version" binding:"required"`
    Src string `form: "src" json:"src" binding:"required"`
}

type Scan struct {
    id int
}

type Issue struct {
    id int
}


func redisClient() (*redis.Client) {
    client := redis.NewClient(&redis.Options{
        Addr:     "192.168.99.100:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })

    pong, err := client.Ping().Result()
    fmt.Println(pong, err)
    return client
}

func main(){
    router := gin.Default()
    client := redisClient()
    defer client.Close()
    // initialize the model here if this is a new instance... we should be able to check to make sure we can connect to redis
    if client.LRange("containers", 0, 1).Err() != nil {
	    client.LPush("containers", "default")
	    client.LPush("scans", "default")
	    client.LPush("issues", "default")
    }


    router.GET("/ping", func(c *gin.Context) {
        //client := redisClient()
        c.String(200, "pong")
    })

    api := router.Group("/api/v1")
    {
        api.GET("/containers", func(c *gin.Context) {
	    client := redisClient()
	    defer client.Close()
	    containers := client.LRange("containers", 0, -1).Val()
            c.JSON(http.StatusOK, containers)
        })


        // POST /api/v1/containers -d { "name" : <container name>, "version" : "<container version>", "src" : "<container URL>" }
        api.POST("/containers", func(c *gin.Context){
	    var ct Container
            if c.Bind(&ct) == nil {
	        if ct.Name != "" && ct.Version != "" {
			client := redisClient()
			defer client.Close()
			// push onto container hash
			client.RPush("containers", ct.Name).Val()
			client.HMSet("container_" + ct.Name, "name", ct.Name, "version", ct.Version, "src", ct.Src)
			c.JSON(http.StatusCreated, gin.H{"status": "OK" })
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"status":"Bad Request"})
		}
	    } else {
		c.JSON(http.StatusInternalServerError, gin.H{"status":"Not sure how we got here."})
	    }
	})

        api.GET("/containers/:id", func(c *gin.Context){
            containerID := c.Param("id")
            c.JSON(http.StatusOK, fmt.Sprintf("{ x : containers with id of %[1]d !}", containerID))
        })
        api.GET("/containers/:id/scans", func(c *gin.Context){
            c.JSON(http.StatusOK, "{ x : container with scans!}")
        })
        api.GET("/containers/:id/scans/:sid", func(c *gin.Context){
            c.JSON(http.StatusOK, "{ x : container with scan ID!}")
        })
        api.GET("/containers/:id/scans/:sid/issues", func(c *gin.Context){
            c.JSON(http.StatusOK, "{ x : container with scan ID, list of issues")
        })
        api.GET("/containers/:id/scans/:sid/issues/:iid", func(c *gin.Context){
            c.JSON(http.StatusOK, "{ x : container with scan ID, issue filtered by id")
        })
        api.GET("/vulnerabilities", func(c *gin.Context){
            c.JSON(http.StatusOK, "{ x : list of vulnerabilities}")
        })
    }

    router.Run(":9001")
}
