package roles

import "fmt"

type ServiceRole string

const (
	RoleAdmin ServiceRole = "ADMIN"
	RoleUser  ServiceRole = "USER"
	RoleOrganization  ServiceRole = "ORGANIZATION"
	RoleEcosystem ServiceRole = "ECOSYSTEM"
	RoleEvaluator ServiceRole = "EVALUATOR"
	RoleBanned ServiceRole = "BANNED"
	RoleUnknown ServiceRole = "UNKNOWN"

)

func IsValidRole(role ServiceRole) (error) {
	switch role {
	case RoleUser, RoleAdmin, RoleOrganization, RoleEcosystem, RoleEvaluator:
		return nil
	case RoleUnknown:
		return fmt.Errorf("Unknown role, access denied")
	case RoleBanned:
		return fmt.Errorf("User is banned, access denied")
	default:
		return fmt.Errorf("Invalid role: %s", role)
	}
}

