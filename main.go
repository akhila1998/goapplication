package main
import (
    "fmt"
    "os"
    "encoding/json"
)
type Config struct {
    Database struct {
        Host string `json:"host"`
        Port string `json:"port"`
    } `json:"database"`
    Host string `json:"host"`
    Port string `json:"port"`
}
func LoadConfiguration(filename string) (Config,error) {
    var config Config
    configFile, err := os.Open(filename)
    defer configFile.Close()
    if err != nil {
	return config,err
    }
    jsonParser := json.NewDecoder(configFile)
    err=jsonParser.Decode(&config)
    return config,err
}
func main(){
	fmt.Println("Starting application...")
	config, _:=LoadConfiguration("config.json")
	fmt.Println(config.Database)
}