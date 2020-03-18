package main

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "gip"
	app.Usage = "This is cli tool to get your global ip address in use."
	app.Version = "0.0.1"
	app.Action = action
	app.Run(os.Args)
}

func action(c *cli.Context) error {
	ip, err := getFromMyglobalip()
	if err == nil {
		fmt.Println(ip)
		return nil
	}
	ip, err = getFromCman()
	if err == nil {
		fmt.Println(ip)
		return nil
	}
	em := "You cannot get global ip address."
	fmt.Println(em)
	return errors.New(em)
}

func getFromMyglobalip() (string, error) {
	doc, err := goquery.NewDocument("https://www.myglobalip.com/")
	if err != nil {
		return "", err
	}
	ip := doc.Find("h3#default > span.ip").Text()
	return ip, nil
}

func getFromCman() (string, error) {
	doc, err := goquery.NewDocument("https://www.cman.jp/network/support/go_access.cgi")
	if err != nil {
		return "", err
	}
	ip := doc.Find("div.inArea > div.outIp").Text()
	return ip, nil
}
