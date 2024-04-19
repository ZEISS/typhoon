package static

import (
	"embed"
)

//go:embed output.css output.js
var Assets embed.FS
