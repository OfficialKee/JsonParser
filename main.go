package main 

import(
	"encoding/json"
	"fmt"
)

func main(){

	 fake := `{"key": "value"}`

	jsonMap := make(map[string]interface{})

	err := json.Unmarshal([]byte(fake),&jsonMap)

	if err != nil {
		fmt.Println("Err: ",err)
		return
	}
	
	fmt.Println(jsonMap["key"])
	

}