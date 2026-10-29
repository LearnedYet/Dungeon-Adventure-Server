package controller

import (
	"github.com/LearnedYet/Dungeon-Adventure-Server/controller/chat"
	"github.com/LearnedYet/Dungeon-Adventure-Server/controller/config"
	"github.com/LearnedYet/Dungeon-Adventure-Server/controller/employment"
	"github.com/LearnedYet/Dungeon-Adventure-Server/controller/gameengines"
	"github.com/LearnedYet/Dungeon-Adventure-Server/controller/mercenaries"
	"github.com/LearnedYet/Dungeon-Adventure-Server/controller/quests"
	"github.com/LearnedYet/Dungeon-Adventure-Server/controller/rooms"
	"github.com/LearnedYet/Dungeon-Adventure-Server/controller/shop"
	"github.com/LearnedYet/Dungeon-Adventure-Server/controller/users"
	"github.com/LearnedYet/Dungeon-Adventure-Server/engine"
	"github.com/LearnedYet/Dungeon-Adventure-Server/persistance"

	log "github.com/sirupsen/logrus"
)

type GameControllerBroadcaster interface {
	BroadcastGameMessageSync(playerIds []engine.PlayerId, message []byte)
	BroadcastGameMessageAsync(playerIds []engine.PlayerId, message []byte)
}

type GameController struct {
	users           *users.Users
	rooms           *rooms.GameRooms
	engines         *gameengines.GameEngines
	shop            *shop.Shop
	quests          *quests.Quests
	mercenaries     *mercenaries.Mercenaries
	employment      *employment.Employment
	configurator    *engine.UnitConfigurator
	itemsConfig     *config.GameItemsConfig
	unitsConfig     *config.GameUnitsConfig
	scenariosConfig *config.GameScenariosConfig
	broadcaster     GameControllerBroadcaster
	persistance     *persistance.Persistance
	lobbyChat       *chat.Chat
}

func NewGameController(persistance *persistance.Persistance) *GameController {
	c := &GameController{}
	c.users = users.NewUsers()
	c.rooms = rooms.NewGameRooms()
	c.engines = gameengines.NewGameEngines()
	c.shop = shop.NewShop()
	c.quests = quests.NewQuests()
	c.mercenaries = mercenaries.NewMercenaries()
	c.employment = employment.NewEmployment()
	c.itemsConfig = config.NewGameItemsConfig()
	c.unitsConfig = config.NewGameUnitsConfig()
	c.scenariosConfig = config.NewGameScenariosConfig()
	c.configurator = engine.NewUnitConfigurator()
	c.broadcaster = c
	c.persistance = persistance
	c.lobbyChat = chat.NewChat(chat.ChatConfig{
		MaxMessages:         CHAT_MAX_MESSAGES,
		MaxMessageLength:    CHAT_MAX_MESSAGE_LENGTH,
		MessageRate:         CHAT_MESSAGE_RATE,
		MessageRateDuration: CHAT_MESSAGE_RATE_DURATION,
	}, c.broadcastLobbyChatMessage, c.broadcastLobbyChatParticipant)
	c.init()
	return c
}

func (c *GameController) init() {
	if err := c.itemsConfig.Load(ITEMS_CONFIG_PATH); err != nil {
		log.Fatal("无法加载物品配置：", err)
	}
	if err := c.shop.LoadItems(SHOP_CONFIG_PATH, c.itemsConfig); err != nil {
		log.Fatal("无法加载商店配置：", err)
	}
	if err := c.quests.LoadItems(QUESTS_CONFIG_PATH, c.itemsConfig); err != nil {
		log.Fatal("无法加载任务配置：", err)
	}
	if err := c.mercenaries.LoadItems(MERCENARIES_CONFIG_PATH, c.itemsConfig); err != nil {
		log.Fatal("无法加载佣兵配置：", err)
	}
	if err := c.unitsConfig.LoadUnits(UNITS_CONFIG_PATH, c.itemsConfig); err != nil {
		log.Fatal("无法加载单位配置：", err)
	}
	if err := c.scenariosConfig.LoadScenarios(SCENARIO_CONFIG_PATH, c.unitsConfig); err != nil {
		log.Fatal("无法加载场景配置：", err)
	}
	if err := c.employment.Load(JOBS_CONFIG_PATH); err != nil {
		log.Fatal("无法加载工作配置：", err)
	}
}
