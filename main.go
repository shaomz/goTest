package main
import "fmt"
import "strconv"
import "strings"
// import "errors"
import "net/http"
import "github.com/gin-gonic/gin"

// 定义User属性
type User struct{
	ID int
	Name string
}

func (u User) getName() string{
	return fmt.Sprintf("the name is %s and ID is %d",u.Name,u.ID)
}

// 定义通用功能
type Print interface {
	getName() string
}
// 实现通用功能
func printName(n Print) string{
	return n.getName()
}


// 数据初始化
var users = []User{
	{ID: 1,Name: "张三"},
	{ID: 2,Name: "李四"},
	{ID: 3,Name: "王五"},
}


// route配置和实现
func main() {
	r:= gin.Default()
	r.GET("/users", listUser)
	r.GET("/users/:id", getUser)
	r.POST("/users",createUser)
	r.GET("/getName/:id", getName)
	r.Run(":8080")
}



func listUser(c *gin.Context) {
	c.JSON(200, users)
}

func getUser(c *gin.Context) {
	id:=c.Param("id")
	user := User{}
	fmt.Println(user)
	for _,u:= range users {
		if strings.EqualFold(id , strconv.Itoa(u.ID)){
			user = u
			c.JSON(200,user)
			return 
		}
	}
	c.JSON(404, gin.H{"message":"用户不存在"})
}

func getName(c *gin.Context){
	id:=c.Param("id")
	user := User{}
	fmt.Println(user)
	for _,u:= range users {
		if strings.EqualFold(id , strconv.Itoa(u.ID)){
			user = u
			c.JSON(200,gin.H{"message":printName(user)})
			return 
		}
	}
	c.JSON(404, gin.H{"message":"用户不存在"})
}

func createUser(c *gin.Context) {
	name := c.DefaultPostForm("name","")
	if name != "" {
		u := User{ID:len(users)+1, Name:name}
		users = append(users, u)
		c.JSON(http.StatusCreated, u)
	} else {
		c.JSON(http.StatusOK, gin.H{"message":"用户不存在"})
	}
}