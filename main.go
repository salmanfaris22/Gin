package main

func main() {

	// r := gin.Default()
	// userManeger := manegers.NewUSerManager()
	// userHadler := handlers.NewUserHandler()
	// r.Run()
}

// func getData(ctx *gin.Context) {
// 	ctx.JSON(200, gin.H{
// 		"message": "HelloWorld",
// 	})
// }
// func getDataPOST(ctx *gin.Context) {
// 	body := ctx.Request.Body
// 	value, err := ioutil.ReadAll(body)
// 	if err != nil {
// 		fmt.Println("Errore")

// 	}
// 	ctx.JSON(200, gin.H{
// 		"message": "Post Data",
// 		"body":    string(value),
// 	})
// }

// // http://localhost:8080/getUrlData/name/MArk/age/30
// func getUrlData(ctx *gin.Context) {
// 	name := ctx.Param("name")
// 	age := ctx.Param("age")
// 	ctx.JSON(200, gin.H{
// 		"message": "Get IN query",
// 		"name":    name,
// 		"age":     age,
// 	})
// }

// // http://localhost:8080/getQueryString?name=MArk&age=30
// func getQueryString(ctx *gin.Context) {
// 	name := ctx.Query("name")
// 	age := ctx.Query("age")
// 	ctx.JSON(200, gin.H{
// 		"message": "Get IN query",
// 		"name":    name,
// 		"age":     age,
// 	})
// }
