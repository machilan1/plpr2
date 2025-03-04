package seeder

import (
	"context"
	"embed"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/machilan1/plpr2/internal/business/sdk/sqldb"
	"github.com/machilan1/plpr2/internal/framework/logger"
)

//go:embed sql/*.sql
var sqlFiles embed.FS

func SeedSQLFiles(ctx context.Context, log *logger.Logger, db *sqldb.DB, targetVersion int) error {
	files, err := sqlFiles.ReadDir("sql")
	if err != nil {
		return fmt.Errorf("failed to read sql directory: %w", err)
	}

	log.Info(ctx, "seedSQLFiles", "status", "found files", "count", len(files))

	// sorting files by name
	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})

	// running the sql files
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if targetVersion != 0 {
			parts := strings.Split(file.Name(), "_")
			sequence := parts[0]
			ver, err := strconv.Atoi(sequence)
			if err != nil {
				return fmt.Errorf("failed to convert sequence to integer: %w", err)
			}
			if ver < targetVersion {
				log.Info(ctx, "seedSQLFiles", "status", "skipped file", "file", file.Name())
				continue
			}
		}

		data, err := sqlFiles.ReadFile(fmt.Sprintf("sql/%s", file.Name()))
		if err != nil {
			return fmt.Errorf("failed to read sql file: %w", err)
		}

		if err := sqldb.ExecContext(ctx, db, string(data)); err != nil {
			return fmt.Errorf("failed to execute sql file[%s]: %w", file.Name(), err)
		}

		log.Info(ctx, "seedSQLFiles", "status", "executed file", "file", file.Name())
	}

	return nil
}
