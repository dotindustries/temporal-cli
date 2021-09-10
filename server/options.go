// Unless explicitly stated otherwise all files in this repository are licensed under the MIT License.
//
// This product includes software developed at Datadog (https://www.datadoghq.com/). Copyright 2021 Datadog, Inc.

package server

import (
	"github.com/DataDog/temporalite/internal/liteconfig"

	"go.temporal.io/server/common/log"
)

// WithLogger overrides the default logger.
func WithLogger(logger log.Logger) Option {
	return newApplyFuncContainer(func(cfg *liteconfig.Config) {
		cfg.Logger = logger
	})
}

// WithDatabaseFilePath persists state to the file at the specified path.
func WithDatabaseFilePath(filepath string) Option {
	return newApplyFuncContainer(func(cfg *liteconfig.Config) {
		cfg.Ephemeral = false
		cfg.DatabaseFilePath = filepath
	})
}

// WithPersistenceDisabled disables file persistence and uses the in-memory storage driver. State will be reset on each process restart.
func WithPersistenceDisabled() Option {
	return newApplyFuncContainer(func(cfg *liteconfig.Config) {
		cfg.Ephemeral = true
	})
}

// WithFrontendPort sets the listening port for the temporal-frontend GRPC service.
func WithFrontendPort(port int) Option {
	return newApplyFuncContainer(func(cfg *liteconfig.Config) {
		cfg.FrontendPort = port
	})
}

// WithDynamicPorts starts Temporal on any available ports.
func WithDynamicPorts() Option {
	return newApplyFuncContainer(func(cfg *liteconfig.Config) {
		cfg.DynamicPorts = true
	})
}

// WithNamespaces registers each namespace on Temporal start.
func WithNamespaces(namespaces ...string) Option {
	return newApplyFuncContainer(func(cfg *liteconfig.Config) {
		cfg.Namespaces = append(cfg.Namespaces, namespaces...)
	})
}

type applyFuncContainer struct {
	applyInternal func(*liteconfig.Config)
}

func (fso *applyFuncContainer) apply(cfg *liteconfig.Config) {
	fso.applyInternal(cfg)
}

func newApplyFuncContainer(apply func(*liteconfig.Config)) *applyFuncContainer {
	return &applyFuncContainer{
		applyInternal: apply,
	}
}
