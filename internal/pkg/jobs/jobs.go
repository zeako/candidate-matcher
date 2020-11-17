package jobs

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/zeako/candidate-matcher/ent"
	"github.com/zeako/candidate-matcher/ent/job"
	"github.com/zeako/candidate-matcher/ent/skill"
	"github.com/zeako/candidate-matcher/internal/pkg/database"
)

// CreateRequest ..
type CreateRequest struct {
	Title string `json:"title"`
	Skill string `json:"skill"`
}

func create(ctx context.Context, req CreateRequest) (int, error) {
	db := database.Get()

	tx, err := db.Tx(ctx)
	if err != nil {
		return 0, err
	}

	skillID, err := tx.Skill.Query().Where(skill.Name(req.Skill)).FirstID(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return 0, rollback(tx, err)
	}
	if ent.IsNotFound(err) {
		skill, err := tx.Skill.Create().SetName(req.Skill).Save(ctx)
		if err != nil {
			return 0, rollback(tx, err)
		}
		skillID = skill.ID
	}

	job, err := tx.Job.Create().SetTitle(req.Title).SetSkillID(skillID).Save(ctx)
	if err != nil {
		return 0, rollback(tx, err)
	}

	return job.ID, tx.Commit()
}

// rollback calls to tx.Rollback and wraps the given error
// with the rollback error if occurred.
func rollback(tx *ent.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%v: %v", err, rerr)
	}
	return err
}

// CreateHandler creates Job entry.
func CreateHandler(c *fiber.Ctx) error {
	var req CreateRequest
	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	req.Title = strings.ToLower(req.Title)
	req.Skill = strings.ToLower(req.Skill)

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
		"jobId": id,
	})
}

func find(ctx context.Context, id int) ([]int, error) {
	db := database.Get()

	job, err := db.Job.Query().Where(job.ID(id)).Only(ctx)
	if err != nil {
		return nil, err
	}

	candidates, err := job.
		QuerySkill().
		QueryCandidate().
		All(ctx)
	if err != nil {
		return nil, err
	}

	var matchingIDs []int

	for _, candidate := range candidates {
		// full match - put first
		// happens once since candidate titles are unique
		if candidate.Title == job.Title {
			matchingIDs = append([]int{candidate.ID}, matchingIDs...)
			continue
		}

		// if any word in the candidate title matches - add it sequentially
		for _, s := range strings.Split(candidate.Title, " ") {
			if strings.Contains(job.Title, s) {
				matchingIDs = append(matchingIDs, candidate.ID)
				break
			}
		}
	}
	return matchingIDs, nil

}

// FindCandidates returns ids of matching candidates ids for job title.
func FindCandidates(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	ids, err := find(c.Context(), id)
	if err != nil {
		status := 500
		if ent.IsValidationError(err) {
			status = 400
		}
		return c.Status(status).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"matchingCandidateIds": ids,
	})
}

// Route contains jobs routes
func Route(r fiber.Router) {
	r.Post("/", CreateHandler)
	r.Get("/:id/candidates", FindCandidates)
}
