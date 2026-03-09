package quests

import (
	"github.com/LearnedYet/Dungeon-Adventure-Server/controller/config"
	"github.com/LearnedYet/Dungeon-Adventure-Server/controller/users"
	"github.com/LearnedYet/Dungeon-Adventure-Server/domain"
	"github.com/LearnedYet/Dungeon-Adventure-Server/engine"
	"github.com/LearnedYet/Dungeon-Adventure-Server/util"
	"sync"
)

type Quests struct {
	mu     sync.RWMutex
	quests *engine.GameQuests
}

func NewQuests() *Quests {
	s := &Quests{}
	return s
}

func (q *Quests) LoadItems(path string, itemsConfig *config.GameItemsConfig) error {
	items, err := util.ReadJsonFile(&[]domain.Quest{}, path)
	if err != nil {
		return err
	}
	q.mu.Lock()
	defer q.mu.Unlock()
	q.quests = engine.NewGameQuests(items, itemsConfig.PopulateFromDescriptor)
	return nil
}

func (q *Quests) GetStatus(unit *domain.Unit) *engine.GameQuestsStatus {
	q.mu.RLock()
	defer q.mu.RUnlock()
	return q.quests.GetStatus(unit)
}

func (q *Quests) ExecuteAction(action domain.Action, user *users.User) *domain.ActionResult {
	q.mu.RLock()
	defer q.mu.RUnlock()
	return q.quests.ExecuteAction(action, &user.Unit.Unit, user.RndGen)
}
