package shop

import (
	"github.com/LearnedYet/Dungeon-Adventure-Server/controller/config"
	"github.com/LearnedYet/Dungeon-Adventure-Server/controller/users"
	"github.com/LearnedYet/Dungeon-Adventure-Server/domain"
	"github.com/LearnedYet/Dungeon-Adventure-Server/engine"
	"github.com/LearnedYet/Dungeon-Adventure-Server/util"
	"sync"
)

type Shop struct {
	mu   sync.RWMutex
	shop *engine.GameShop
}

func NewShop() *Shop {
	s := &Shop{}
	return s
}

func (s *Shop) LoadItems(path string, itemsConfig *config.GameItemsConfig) error {
	items, err := util.ReadJsonFile(&domain.UnitInventory{}, path)
	if err != nil {
		return err
	}
	defer s.mu.Unlock()
	s.mu.Lock()
	itemsConfig.PopulateFromDescriptor(items)
	s.shop = engine.NewGameShop(items)
	return nil
}

func (s *Shop) GetStatus() engine.GameShop {
	defer s.mu.RUnlock()
	s.mu.RLock()
	return *s.shop
}

func (s *Shop) ExecuteAction(action domain.Action, user *users.User) *domain.ActionResult {
	defer s.mu.RUnlock()
	s.mu.RLock()
	return s.shop.ExecuteAction(action, &user.Unit.Unit, user.RndGen)
}
