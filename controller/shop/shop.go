package shop

import (
	"github.com/LearnedYet/Dungeon-Adventure-Server/controller/users"
	"github.com/LearnedYet/Dungeon-Adventure-Server/domain"
	"github.com/LearnedYet/Dungeon-Adventure-Server/engine"
	"sync"
)

type Shop struct {
	mu   sync.RWMutex
	shop *engine.GameShop
}

func NewShop() *Shop {
	s := &Shop{}
	s.shop = engine.NewGameShop(NewTestShopItems())
	return s
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
