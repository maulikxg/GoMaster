package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type Player struct {
	health int32
}

func (p *Player) getHealth() int {
	return int(atomic.LoadInt32(&p.health))
}

func (p *Player) takeDamage(value int) {
	atomic.AddInt32(&p.health, int32(-value)) // Atomic subtraction
}

func Newplayer() *Player {
	return &Player{
		health: 100,
	}
}

func startUIloop(p *Player) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop() // Ensure ticker is stopped

	for {
		fmt.Printf("Player health: %d\n", p.getHealth()) // Use atomic method
		<-ticker.C
	}
}

func gameOver(p *Player) {
	ticker := time.NewTicker(300 * time.Millisecond)
	defer ticker.Stop()

	for {
		damage := rand.Intn(20)
		p.takeDamage(damage) // Use atomic method

		if p.getHealth() <= 0 {
			fmt.Println("Game Over: Player has died.")
			break
		}
		<-ticker.C
	}
}

func main() {
	player := Newplayer()
	go startUIloop(player) // Run UI loop concurrently
	gameOver(player)       // Run game logic
}
