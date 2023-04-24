1. ##### 如果不再需要这个框架，则执行go mod tidy，如果程序中没有涉及到该框架的代码，该指令会自动将go.mod文件中的该框架删除

```
go mod tidy
```

2. `ShouldBindJSON`

   将`post`的data参数进行Josn化。

   ```tsx
   data:data:"{\"username\":\"root\",\"password\":\"root\"}"
   ```

   ```go
   type User struct {
   	ID uint `gorm:"primary_key"`
   	Username string `gorm:"username" form:"username" json:"username" `
   	Password string `gorm:"password" form:"password" json:"password" `
   }	
   //转换为josn
   var data User.User
   	err := context.ShouldBindJSON(&data)
   	if err != nil {
   		context.JSON(http.StatusBadRequest, gin.H{
   			"message": err.Error(),
   		})
   		return
   	}
   ```
3. `Update` and `Updates`

   
