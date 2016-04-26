package main
import "github.com/gin-gonic/gin"
func main(){
    router := gin.Default()
    router.GET("/ping", func(c *gin.Context) {
        c.String(200, "pong")
    })
    
    router.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.tmpl", gin.H{
            "title":"polar"
        })  
    })
    
    router.GET("/scans", func(c *gin.Context){
        c.HTML(http.StatusOK, "index.tmpl", gin.H{
            "title":"scans"
        })
    })
    
    router.Run()
}
