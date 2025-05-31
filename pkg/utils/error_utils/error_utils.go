package errorutils

import (
	"errors"
	"strings"

	RepoErrors "be.blog/pkg/custom_errors"
	"github.com/jackc/pgx/v5"
)

func MapRepoError(err error) error {
	if err == nil {
		return nil
	}
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return errors.New(RepoErrors.RepoItemNotFound)
	case strings.Contains(err.Error(), "unique constraint"):
		return errors.New(RepoErrors.RepoItemAlreadyExists)
	case strings.Contains(err.Error(), "foreign key violation"):
		return errors.New(RepoErrors.RepoForeignKeyViolation)
	case strings.Contains(err.Error(), "id not found"):
		return errors.New(RepoErrors.RepoIDNotFound)
	case strings.Contains(err.Error(), "not found"):
		return errors.New(RepoErrors.RepoItemNotFound)
	case strings.Contains(err.Error(), "duplicate key value violates unique constraint"):
		return errors.New(RepoErrors.RepoItemAlreadyExists)
	default:
		return err
	}
}
