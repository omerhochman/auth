package models

import (
	"database/sql"
	"fmt"
	"github.com/netlify/gotrue/crypto"
	"github.com/netlify/gotrue/storage"
	"github.com/pkg/errors"
	"time"
)

type Challenge struct {
	ID        string     `json:"challenge_id" db:"id"`
	FactorID  string     `json:"factor_id" db:"factor_id"`
	CreatedAt *time.Time `json:"created_at" db:"created_at"`
}

const CHALLENGE_PREFIX = "challenge"

func NewChallenge(factorID string) (*Challenge, error) {
	challenge := &Challenge{
		ID:       fmt.Sprintf("%s_%s", CHALLENGE_PREFIX, crypto.SecureToken()),
		FactorID: factorID,
	}
	return challenge, nil
}

func FindChallengesByFactorID(tx *storage.Connection, factorID string) ([]*Challenge, error) {
	challenges := []*Challenge{}
	if err := tx.Q().Where("factor_id = ?", factorID, true).All(&challenges); err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return challenges, nil
		}
		return nil, errors.Wrap(err, "Error finding MFA Challenges for factor")
	}
	return challenges, nil
}
