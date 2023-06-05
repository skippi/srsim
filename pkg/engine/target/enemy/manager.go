package enemy

import (
	"fmt"

	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/attribute"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
)

type Manager struct {
	engine    engine.Engine
	attr      attribute.AttributeModifier
	instances map[key.TargetID]EnemyInstance
	info      map[key.TargetID]info.Enemy
}

func New(engine engine.Engine, attr attribute.AttributeModifier) *Manager {
	return &Manager{
		engine:    engine,
		attr:      attr,
		instances: make(map[key.TargetID]EnemyInstance, 5),
		info:      make(map[key.TargetID]info.Enemy, 5),
	}
}

func (mgr *Manager) Get(id key.TargetID) (EnemyInstance, error) {
	if instance, ok := mgr.instances[id]; ok {
		return instance, nil
	}
	return nil, fmt.Errorf("target is not an enemy: %v", id)
}

func (mgr *Manager) Info(id key.TargetID) (info.Enemy, error) {
	if char, ok := mgr.info[id]; ok {
		return char, nil
	}
	return info.Enemy{}, fmt.Errorf("target is not an enemy: %v", id)
}
