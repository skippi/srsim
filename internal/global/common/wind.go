package common

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	WindShear      key.Modifier = "common_wind_shear"
	BreakWindShear key.Modifier = "break_wind_shear"
)

type WindShearState struct {
	DamagePercentage float64
}

func init() {
	modifier.Register(WindShear, modifier.Config{
		Stacking:          modifier.ReplaceBySource,
		TickMoment:        modifier.ModifierPhase1End,
		MaxCount:          5,
		CountAddWhenStack: 1,
		StatusType:        model.StatusType_STATUS_DEBUFF,
		BehaviorFlags: []model.BehaviorFlag{
			model.BehaviorFlag_STAT_DOT,
			model.BehaviorFlag_STAT_DOT_POISON,
		},
		Listeners: modifier.Listeners{
			OnPhase1: windShearPhase1,
		},
	})

	// TODO: Break wind shear
}

func windShearPhase1(mod *modifier.ModifierInstance) {
	state, ok := mod.State().(WindShearState)
	if !ok {
		panic("incorrect state used for wind shear modifier")
	}

	// perform wind shear damage
	mod.Engine().Attack(info.Attack{
		Source:     mod.Source(),
		Targets:    []key.TargetID{mod.Owner()},
		AttackType: model.AttackType_DOT,
		DamageType: model.DamageType_WIND,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: state.DamagePercentage * mod.Count(),
		},
		UseSnapshot: true,
	})
}
