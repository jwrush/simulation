package fight

import (
	"testing"
)

func TestNewMonster(t *testing.T) {
	attacks := []Attack{
		NewAttack("Bite", "A vicious bite!", 10, 2),
	}
	monster := NewMonster("Goblin", 30, 3, attacks)

	if monster.Name != "Goblin" {
		t.Errorf("expected Name to be 'Goblin', got '%s'", monster.Name)
	}
	if monster.HitPoints != 30 {
		t.Errorf("expected HitPoints to be 30, got %d", monster.HitPoints)
	}
	if monster.BrainSpeed != 3 {
		t.Errorf("expected BrainSpeed to be 3, got %d", monster.BrainSpeed)
	}
	if len(monster.Attacks) != 1 {
		t.Errorf("expected 1 attack, got %d", len(monster.Attacks))
	}
}

func TestMonsterAttack(t *testing.T) {
	attacker := NewMonster("Orc", 40, 2, []Attack{
		NewAttack("Smash", "A powerful smash!", 12, 1),
	})
	target := NewMonster("Elf", 25, 4, nil)

	recovery, err := attacker.Attack(target, 0)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if recovery != 1 {
		t.Errorf("expected recovery to be 1, got %d", recovery)
	}
	if target.HitPoints != 13 {
		t.Errorf("expected target HitPoints to be 13, got %d", target.HitPoints)
	}
}

func TestMonsterAttack_InvalidIndex(t *testing.T) {
	attacker := NewMonster("Orc", 40, 2, []Attack{})
	target := NewMonster("Elf", 25, 4, nil)

	_, err := attacker.Attack(target, 0)
	if err == nil {
		t.Error("expected error for invalid attack index, got nil")
	}
}
