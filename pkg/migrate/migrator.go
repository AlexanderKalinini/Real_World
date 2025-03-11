package migrate

import (
	"database/sql"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"slices"
	"strconv"
	"strings"
	"time"
)

var (
	MigrationsNotFound = errors.New("migrations not found")
	NotSqlFile         = errors.New("not sql file in migrations dir")
)

type Migrator struct {
	db *sql.DB
}

func NewMigrator(db *sql.DB) *Migrator {
	return &Migrator{db: db}
}

func (m *Migrator) Up(migrationDir string) error {
	files, err := os.ReadDir(migrationDir)
	if err != nil {
		return err
	}

	if len(files) == 0 {
		return MigrationsNotFound
	}

	for _, fileEntity := range files {
		if !strings.HasSuffix(fileEntity.Name(), ".sql") {
			continue
		}

		if strings.Contains(fileEntity.Name(), "down") {
			continue
		}

		file, err := os.Open(path.Join(migrationDir, fileEntity.Name()))
		if err != nil {
			return err
		}
		query, err := io.ReadAll(file)
		if err != nil {
			return err
		}
		_, err = m.db.Exec(string(query))

		if err != nil {
			return err
		}
		err = file.Close()
		if err != nil {
			return err
		}
	}
	err = m.db.Close()
	if err != nil {
		return err
	}
	fmt.Println("migrations up success")
	return nil
}

func (m *Migrator) Down(migrationDir string) error {

	files, err := os.ReadDir(migrationDir)
	if err != nil {
		return err
	}

	if len(files) == 0 {
		return MigrationsNotFound
	}

	slices.Reverse(files)

	for _, fileEntity := range files {
		if !strings.HasSuffix(fileEntity.Name(), ".sql") {
			continue
		}

		if !strings.Contains(fileEntity.Name(), "down") {
			continue
		}

		file, err := os.Open(path.Join(migrationDir, fileEntity.Name()))
		if err != nil {
			return err
		}
		query, err := io.ReadAll(file)
		if err != nil {
			return err
		}

		_, err = m.db.Exec(string(query))

		if err != nil {
			return err
		}
		err = file.Close()
		if err != nil {
			return err
		}
	}

	err = m.db.Close()
	if err != nil {
		return err
	}

	fmt.Println("migrations down success")
	return nil
}

func (m *Migrator) Create(migrationDir, fileName string) error {
	fileName = strings.TrimSuffix(fileName, ".sql")
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	_, err := os.Create(path.Join(migrationDir, timestamp+"_"+fileName+".sql"))
	if err != nil {
		return err
	}
	_, err = os.Create(path.Join(migrationDir, timestamp+"_"+fileName+".down.sql"))
	if err != nil {
		return err
	}
	return nil
}
