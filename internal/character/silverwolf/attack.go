package silverwolf

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

var attackHits = []float64{0.25, 0.25, 0.5}

func (c *char) Attack(target key.TargetID, state info.ActionState) {
	for _, hitRatio := range attackHits {
		c.engine.Attack(info.Attack{
			Source:     c.id,
			Targets:    []key.TargetID{target},
			DamageType: model.DamageType_QUANTUM,
			AttackType: model.AttackType_NORMAL,
			BaseDamage: info.DamageMap{
				model.DamageFormula_BY_ATK: atk[c.info.AbilityLevel.Attack-1],
			},
			StanceDamage: 30.0,
			EnergyGain:   20.0,
			HitRatio:     hitRatio,
		})
	}

	state.EndAttack()
}
