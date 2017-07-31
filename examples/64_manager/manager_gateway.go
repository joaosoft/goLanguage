package pm

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"golang-learn/examples/64_manager/gateway"
	"io"
)

// -------------- GATEWAY --------------
// NewGateway ... creates a new web server
func (instance *manager) NewGateway(config *gateway.Config) (*gateway.Gateway, error) {
	log.Infof(fmt.Sprintf("gateway, creating gateway"))
	return gateway.NewGateway(config), nil
}

// -------------- METHODS --------------
// GetGateway ... get a gateway by key
func (instance *manager) GetGateway(key string) (*gateway.Gateway, error) {
	return instance.gatewayController[key], nil
}

// AddGateway, add a new gateway
func (instance *manager) AddGateway(key string, gateway *gateway.Gateway) error {
	log.Infof(fmt.Sprintf("gateway, add a new gateway '%s'", key))
	instance.gatewayController[key] = gateway
	return nil
}

// RemGateway, remove a gateway by key
func (instance *manager) RemGateway(key string) (*gateway.Gateway, error) {
	log.Infof(fmt.Sprintf("gateway, remove the gateway '%s'", key))

	// get gateway
	controller := instance.gatewayController[key]

	// delete gateway
	delete(instance.gatewayController, key)
	log.Infof(fmt.Sprintf("gateway, gateway '%s' removed", key))

	return controller, nil
}
func (instance *manager) RequestGateway(key string, method string, endpoint string, headers map[string]string, body io.Reader) (int, []byte, error) {
	log.Infof(fmt.Sprintf("gateway, request the gateway '%s'", key))
	return instance.gatewayController[key].Request(method, endpoint, headers, body)
}