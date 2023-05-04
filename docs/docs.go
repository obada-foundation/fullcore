package docs

import "embed"

// Docs contains the embedded Swagger UI documentation.
//
//go:embed static
var Docs embed.FS
