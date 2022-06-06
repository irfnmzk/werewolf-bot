package werewolf

import (
	"context"
	"fmt"
	"time"

	"github.com/irfnmzk/werewolf-arena/storage"
)

type gameLoop struct {
	redis *storage.RedisInterface
}

func NewGameLoop(redis *storage.RedisInterface) *gameLoop {
	return &gameLoop{
		redis: redis,
	}
}

func (gl *gameLoop) Execute() {
	fmt.Print("Executing Gameloop")
	ctx := context.Background()
	gameIter := gl.redis.Client.Scan(ctx, 0, "game:*", 0).Iterator()

	for gameIter.Next(ctx) {
		fmt.Println("executing keys ", gameIter.Val())
	}

	if err := gameIter.Err(); err != nil {
		panic(err)
	}

	time.Sleep(time.Second)
	// TODO: handle this infinity recursive
	gl.Execute()
}
