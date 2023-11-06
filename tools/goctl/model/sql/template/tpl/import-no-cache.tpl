import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	{{if .time}}"time"{{end}}

    {{if .containsPQ}}"github.com/lib/pq"{{end}}
	"github.com/userzhangjinlong/go-zero/core/stores/builder"
	"github.com/userzhangjinlong/go-zero/core/stores/sqlc"
	"github.com/userzhangjinlong/go-zero/core/stores/sqlx"
	"github.com/userzhangjinlong/go-zero/core/stringx"
)
