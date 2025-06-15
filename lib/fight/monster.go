package fight

import "fmt"

type Narratable interface {
	Narrate() string
}

type Attack struct {
	Name string
	Narration string
	Power int
	Recovery int
}

func (a Attack) Narrate() string {
	return a.Narration
}


func NewAttack(name string, narration string, power int, recovery int) Attack {
	return Attack{
		Name: name,
		Narration: narration,
		Power: power,
		Recovery: recovery,
	}
}

type Monster struct {
	Name  string
	HitPoints int
	BrainSpeed int
	Attacks []Attack
}

func NewMonster(name string, hitPoints int, brainSpeed int, attacks []Attack) *Monster {
	return &Monster{
		Name: name,
		HitPoints: hitPoints,
		BrainSpeed: brainSpeed,
		Attacks: attacks,
	}
}


func (m *Monster) IsAlive() bool {
	return m.HitPoints > 0
}

func (m *Monster) Attack(target *Monster, attackIndex int) (int, error) {
	if attackIndex < 0 || attackIndex >= len(m.Attacks) {
		return 0, fmt.Errorf("invalid attack index %d: must be between 0 and %d", attackIndex, len(m.Attacks)-1)
	}
	
	if !m.IsAlive() {
		return 0, fmt.Errorf("monster %s is not alive and cannot attack", m.Name)
	}
	
	if target == nil {
		return 0, fmt.Errorf("target cannot be nil")
	}
	
	attack := m.Attacks[attackIndex]
	target.HitPoints -= attack.Power
	fmt.Printf("%s uses %s on %s! %s\n", m.Name, attack.Name, target.Name, attack.Narrate())
	return attack.Recovery, nil
}