package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var version = "0.0.1"

type Player struct {
	Name string
	Wfp  int64
}

func main() {
	file, err := os.Open("wfp.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	players := []*Player{}
	for scanner.Scan() {
		text := scanner.Text()
		i := strings.Index(text, "[")
		for i != -1 {
			k := strings.Index(text, "]")

			text = text[:i] + text[k+1:]
			i = strings.Index(text, "[")
		}
		fields := strings.Fields(text)

		wfp, _ := strconv.ParseInt(strings.Replace(fields[3], ",", "", -1), 10, 64)
		player := &Player{Name: fields[1], Wfp: wfp}
		players = append(players, player)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	tickets := []*Player{}
	for _, player := range players {
		numberOfTickets := player.Wfp / 500

		fmt.Printf(player.Name+" %v\n", numberOfTickets)
		for i := int64(0); i < numberOfTickets; i++ {
			tickets = append(tickets, player)
		}
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(tickets), func(i, j int) { tickets[i], tickets[j] = tickets[j], tickets[i] })

	fmt.Println("\nResults\n")

	numberOfWinners := 3
	previousWinners := []int{}
	prizes := []int{3000, 1000, 1000}

	for i := 0; i < numberOfWinners; i++ {
		for len(previousWinners) <= i {
			winner := random(0, len(tickets))
			if !stringInSlice(winner, previousWinners) {
				previousWinners = append(previousWinners, winner)
				fmt.Println(fmt.Sprintf("%v", prizes[i]) + "c " + tickets[winner].Name)
				continue
			}
		}
	}
	fmt.Println("\nGenerated by [url]https://github.com/benjih/weeky-wfp-calculator[/url] " + version)
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func stringInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
