package params

import (
	serverconfig "github.com/cosmos/cosmos-sdk/server/config"
)

// CustomAppConfig defines custom application configuration.
type CustomAppConfig struct {
	serverconfig.Config
}
