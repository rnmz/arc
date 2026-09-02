package data

import (
	"arc/app/entity"
	"context"

	"github.com/jmoiron/sqlx"
)

type PlanData struct{}

func (_ PlanData) CreateNewPlan(db *sqlx.DB, ctx context.Context, title string, size int) error {
	tx, txErr := db.BeginTxx(ctx, nil)
	if txErr != nil {
		return txErr
	}
	_, _ = tx.Exec(`INSERT INTO plans (title, size) VALUES ($1, $2)`, title, size)
	return tx.Commit()
}

func (_ PlanData) GetPlan(db *sqlx.DB, ctx context.Context, id int) (entity.PlanEntity, error) {
	var plan entity.PlanEntity
	err := db.GetContext(ctx, &plan, "SELECT 1 FROM plans WHERE id = $1", id)
	return plan, err
}

func (_ PlanData) GetAllPlans(db *sqlx.DB, ctx context.Context) ([]entity.PlanEntity, error) {
	var plans []entity.PlanEntity
	err := db.GetContext(ctx, &plans, "SELECT * FROM plans")
	return plans, err
}

func (_ PlanData) EditPlan(db *sqlx.DB, ctx context.Context, id int, title string, size int) error {
	tx, txErr := db.BeginTxx(ctx, nil)
	if txErr != nil {
		return txErr
	}
	_, _ = tx.Exec(`UPDATE plans SET title = $1 and size = $2 WHERE id = $3`, title, size, id)
	return tx.Commit()
}

func (_ PlanData) DeletePlan(db *sqlx.DB, ctx context.Context, id int) error {
	tx, txErr := db.BeginTxx(ctx, nil)
	if txErr != nil {
		return txErr
	}
	_, _ = tx.Exec(`DELETE FROM plans WHERE id = $1`, id)
	return tx.Commit()
}
