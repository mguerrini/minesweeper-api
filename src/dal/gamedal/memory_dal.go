package gamedal

import (
	"github.com/beevik/guid"
	"github.com/minesweeper/src/domain"
	"sync"
)

type GameInMemoryDal struct {
	userGames []*UserGame
	mutex     *sync.Mutex

}

type UserGame struct {
	UserId string
	Game domain.Game
}

func NewInMemoryGameDal(factoryConfigurationName string) GameDal {
	output := &GameInMemoryDal{
		userGames: make([]*UserGame, 0),
		mutex:     &sync.Mutex{},
	}
	return output
}


//Factory Method
func CreateInMemoryGameDal(configurationName string) (interface{}, error){
	return NewInMemoryGameDal(""), nil
}



func (this *GameInMemoryDal) GetGameById(userId, gameId string) (*domain.Game, error) {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	for _, user := range this.userGames {
		if user.Game.GetId() == gameId && user.UserId == userId{
			output := user.Game
			return &output, nil
		}
	}

	return nil, nil
}

func (this * GameInMemoryDal) GetGameListByUserId(userId string) ([]domain.Game, error) {
	panic("implement me")
}

func (this * GameInMemoryDal) InsertGame(userId string, game *domain.Game) (domain.Game, error) {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	id := guid.New().String()
	game.SetId(id)

	reg := UserGame{
		UserId: userId,
		Game:   *game,
	}

	this.userGames = append(this.userGames, &reg)

	return *game,nil
}

func (this *GameInMemoryDal) UpdateGame(game *domain.Game) (domain.Game, error) {

	this.mutex.Lock()
	defer this.mutex.Unlock()

	for _, user := range this.userGames {
		if user.Game.GetId() == game.GetId() {
			user.Game = *game
		}
	}

	return *game, nil
}

