package types

import "github.com/google/uuid"

type RuleTypeDTO struct {
	Id              uuid.UUID
	Name            string
	RuleAttribute   string
	DefaultPriority int
}

type CreateRuleTypeDTO struct {
	Id              uuid.UUID
	Name            string
	RuleAttribute   string
	DefaultPriority int
}

type UpdateRuleTypeDTO struct {
	Id              uuid.UUID
	Name            string
	RuleAttribute   string
	DefaultPriority int
}

type FilterableRuleTypeProps struct {
	Id            uuid.UUIDs
	Name          []string
	RuleAttribute []string
}
