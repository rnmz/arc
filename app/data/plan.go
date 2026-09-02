package data

import (
	"arc/app/entity"
	"context"

	"github.com/jmoiron/sqlx"
)

type PlanData struct{}

func (_ PlanData) CreateNewPlan(db *sqlx.DB, ctx context.Context, name string) {

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

func (_ PlanData) EditPlan(db *sqlx.DB, ctx context.Context, id int, name string, size int) {

}

func (_ PlanData) RemovePlan(db *sqlx.DB, ctx context.Context, id int) {

}
