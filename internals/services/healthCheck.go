package service

import "perpus-app/internals/interfaces"

type HealthCheck struct {
	HealtyCheckRepository interfaces.IHealthCheckRepository
}

func (s *HealthCheck) HealthCheckService() (string, error) {
	return "servie is healthy", nil
}
