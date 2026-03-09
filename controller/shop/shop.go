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
	s.mu.Lock()
	defer s.mu.Unlock()
	s.shop = engine.NewGameShop(items, itemsConfig.PopulateFromDescriptor)
	return nil
}

func (s *Shop) GetStatus(unit *domain.Unit) *engine.GameShopStatus {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.shop.GetStatus(unit)
}

func (s *Shop) ExecuteAction(action domain.Action, user *users.User) *domain.ActionResult {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.shop.ExecuteAction(action, &user.Unit.Unit, user.RndGen)
}
