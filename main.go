package main 

import(
	"encoding/json"
	"fmt"
	"os"
)

func main(){

	args := os.Args
	// fake := `{"key": "value"}`
	jsonString := args[1]
	if len(args) > 1 {
		parseJson(jsonString)
		// fmt.Println(jsonString)
	}

	
	

}

func parseJson(jsonString string) int {
	
	jsonMap := make(map[string]interface{})

	err :=  json.Unmarshal([]byte(jsonString),&jsonMap)

	if err != nil {
		fmt.Println("Err: ",err)
		return 1
	}
	
	fmt.Println(jsonMap)
	return 0
}
