package main

import "fmt"

type Json struct {
}

func (json *Json) ConvertToXml() string {
	return "converted to xml"
}

type JsonAdapter struct {
	json *Json
}

func (adapter *JsonAdapter) SendData() {
	fmt.Println(adapter.json.ConvertToXml())
	fmt.Println("send data json converted to xml")
}

type XML struct {
}

func (xml *XML) SendData() {
	fmt.Println("send data xml")
}

type DataService interface {
	SendData()
}

func MakeDataServiceJson(json *Json) DataService {
	return &JsonAdapter{json: json}
}

func MakeDataServiceXml() DataService {
	return &XML{}
}

func main() {
	json := Json{}
	jsonDataService := MakeDataServiceJson(&json)
	xmlDataService := MakeDataServiceXml()
	jsonDataService.SendData()
	xmlDataService.SendData()
}
