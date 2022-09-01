package service

import (
	"broker/pkg/model"
	"broker/pkg/server"

	"github.com/labstack/echo/v4"
)

type ServiceBroker interface {
	server.ServerInterface
}

type serviceBroker struct {
	availableServices []model.Service
	availablePlans    []model.Plan
}

func NewServiceBroker() ServiceBroker {
	services := []model.Service{}
	free, notFree := true, false
	plans := []model.Plan{}
	plans = append(plans,
		model.Plan{
			Id:          "pA",
			Description: "Plan A",
			Name:        "planA",
			Free:        &free,
		}, model.Plan{
			Id:          "pB",
			Description: "Plan B",
			Name:        "planB",
			Free:        &notFree,
		})

	services = append(services,
		model.Service{
			Bindable:    true,
			Id:          "A",
			Name:        "serviceA",
			Description: "Service A",
			Plans:       plans,
		},
		model.Service{
			Bindable:    false,
			Id:          "B",
			Name:        "serviceB",
			Description: "Service B",
			Plans:       plans,
		})
	return &serviceBroker{
		availableServices: services,
		availablePlans:    plans,
	}
}

// get the catalog of services that the service broker offers
// (GET /v2/catalog)
func (s *serviceBroker) CatalogGet(ctx echo.Context, params model.CatalogGetParams) error {

	catalog := &model.Catalog{Services: &s.availableServices}
	ctx.JSON(200, catalog)
	return nil
}

// deprovision a service instance
// (DELETE /v2/service_instances/{instance_id})
func (s *serviceBroker) ServiceInstanceDeprovision(ctx echo.Context, instanceId string, params model.ServiceInstanceDeprovisionParams) error {
	return nil
}

// get a service instance
// (GET /v2/service_instances/{instance_id})
func (s *serviceBroker) ServiceInstanceGet(ctx echo.Context, instanceId string, params model.ServiceInstanceGetParams) error {
	return nil
}

// update a service instance
// (PATCH /v2/service_instances/{instance_id})
func (s *serviceBroker) ServiceInstanceUpdate(ctx echo.Context, instanceId string, params model.ServiceInstanceUpdateParams) error {
	return nil
}

// provision a service instance
// (PUT /v2/service_instances/{instance_id})
func (s *serviceBroker) ServiceInstanceProvision(ctx echo.Context, instanceId string, params model.ServiceInstanceProvisionParams) error {
	return nil
}

// get the last requested operation state for service instance
// (GET /v2/service_instances/{instance_id}/last_operation)
func (s *serviceBroker) ServiceInstanceLastOperationGet(ctx echo.Context, instanceId string, params model.ServiceInstanceLastOperationGetParams) error {
	return nil
}

// deprovision a service binding
// (DELETE /v2/service_instances/{instance_id}/service_bindings/{binding_id})
func (s *serviceBroker) ServiceBindingUnbinding(ctx echo.Context, instanceId string, bindingId string, params model.ServiceBindingUnbindingParams) error {
	return nil
}

// get a service binding
// (GET /v2/service_instances/{instance_id}/service_bindings/{binding_id})
func (s *serviceBroker) ServiceBindingGet(ctx echo.Context, instanceId string, bindingId string, params model.ServiceBindingGetParams) error {
	return nil
}

// generate a service binding
// (PUT /v2/service_instances/{instance_id}/service_bindings/{binding_id})
func (s *serviceBroker) ServiceBindingBinding(ctx echo.Context, instanceId string, bindingId string, params model.ServiceBindingBindingParams) error {
	return nil
}

// get the last requested operation state for service binding
// (GET /v2/service_instances/{instance_id}/service_bindings/{binding_id}/last_operation)
func (s *serviceBroker) ServiceBindingLastOperationGet(ctx echo.Context, instanceId string, bindingId string, params model.ServiceBindingLastOperationGetParams) error {
	return nil
}
