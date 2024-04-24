package main

import (
	"encoding/xml"
	"fmt"
	"github.com/H5Q4/tibrv-go"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("message.xml")
	if err != nil {
		log.Fatal(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	var msg tibrv.RvMessage

	if err := msg.Create(); err != nil {
		panic(err)
	}
	defer func(msg *tibrv.RvMessage) {
		err := msg.Destroy()
		if err != nil {
			log.Fatal(err)
		}
	}(&msg)

	if err := msg.SetXML("msg", data); err != nil {
		panic(err)
	}
	ret, err := msg.GetXML("msg")
	if err != nil {
		panic(err)
	}
	m := Message{}
	err = xml.Unmarshal(ret, &m)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", m)
}

type Message struct {
	XMLName xml.Name `xml:"Message"`
	Header  Header   `xml:"Header"`
	Body    Body     `xml:"Body"`
	Return  Return   `xml:"Return"`
	Raw     string   `xml:",innerxml"`
}

type Header struct {
	XMLName                   xml.Name `xml:"Header"`
	MessageName               string   `xml:"MESSAGENAME"`
	ShopName                  string   `xml:"SHOPNAME"`
	MachineName               string   `xml:"MACHINENAME"`
	TransactionID             string   `xml:"TRANSACTIONID"`
	OriginalSourceSubjectName string   `xml:"ORIGINALSOURCESUBJECTNAME"`
	SourceSubjectName         string   `xml:"SOURCESUBJECTNAME"`
	TargetSubjectName         string   `xml:"TARGETSUBJECTNAME"`
	EventUser                 string   `xml:"EVENTUSER"`
	EventComment              string   `xml:"EVENTCOMMENT"`
}

type Body struct {
	XMLName     xml.Name `xml:"Body"`
	MachineName string   `xml:"MACHINENAME"`
}

type Return struct {
	XMLName       xml.Name `xml:"Return"`
	ReturnCode    int      `xml:"RETURNCODE"`
	ReturnMessage string   `xml:"RETURNMESSAGE"`
}
