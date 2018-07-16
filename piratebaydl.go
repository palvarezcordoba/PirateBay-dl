package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"github.com/gnur/go-piratebay"
)

var pirateBay piratebay.Piratebay

func search(search string) []piratebay.Torrent {
	torrents, err := pirateBay.Search(search)
	if err != nil {
		panic(err)
	}
	return torrents
}

func init() {
	pirateBay = piratebay.New("https://thepiratebay.org")
}

func main() {
	fmt.Print("Escribe qué quieres buscar: ")
	reader := bufio.NewReader(os.Stdin)
	s, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	s = s[:len(s)-1] // Delete \n
	torrents := search(s)
	for i, torrent := range torrents {
		fmt.Printf("%d. %s\n", i+1, torrent.Title)
	}
	fmt.Print("Elige uno: ")
	s, err = reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	s = s[:len(s)-1] // Delete \n
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	magnetLink := torrents[i-1].MagnetLink
	cmd := exec.Command("xdg-open", magnetLink)
	cmd.Run()
}
