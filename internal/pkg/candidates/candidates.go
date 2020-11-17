package candidates

import (
	"context"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/zeako/candidate-matcher/ent"
	"github.com/zeako/candidate-matcher/ent/skill"
	"github.com/zeako/candidate-matcher/internal/pkg/database"
)

// CreateRequest ..
type CreateRequest struct {
	Title  string   `json:"title"`
	Skills []string `json:"skills"`
}

func create(ctx context.Context, req CreateRequest) (int, error) {
	db := database.Get()

	if len(req.Skills) == 0 {
		_, err := db.Candidate.
			Create().
			SetTitle(req.Title).
			Save(ctx)
		return 0, err
	}

	tx, err := db.Tx(ctx)

	skillIDs := make([]int, len(req.Skills))
	for i, sk := range req.Skills {
		skillID, err := tx.Skill.Query().Where(skill.Name(sk)).FirstID(ctx)
		if err != nil && !ent.IsNotFound(err) {
			return 0, rollback(tx, err)
		}
		if ent.IsNotFound(err) {
			skill, err := tx.Skill.Create().SetName(sk).Save(ctx)
			if err != nil {
				return 0, rollback(tx, err)
			}
			skillID = skill.ID
		}
		skillIDs[i] = skillID
	}

	candidate, err := tx.Candidate.
		Create().
		SetTitle(req.Title).
		AddSkillIDs(skillIDs...).
		Save(ctx)
	if err != nil {
		return 0, rollback(tx, err)
	}

	return candidate.ID, tx.Commit()
}

// rollback calls to tx.Rollback and wraps the given error
// with the rollback error if occurred.
func rollback(tx *ent.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%v: %v", err, rerr)
	}
	return err
}

// CreateHandler ..
func CreateHandler(c *fiber.Ctx) error {
	var req CreateRequest
	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	req.Title = strings.ToLower(req.Title)
	for i := 0; i < len(req.Skills); i++ {
		req.Skills[i] = strings.ToLower(req.Skills[i])
	}

	id, err := create(c.Context(), req)
	if err != nil {
		status := 500
		if ent.IsValidationError(err) {
			status = 400
		}
		return c.Status(status).JSON(&fiber.Map{
			"error": err.Error(),
		})

	}
	return c.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"candidateId": id,
	})
}

// Route contains candidates routes
func Route(r fiber.Router) {
	r.Post("/", CreateHandler)
}
