package entity

type PlanEntity struct {
	Id    int    `db:"id"`
	Title string `db:"title"`
	Size  int    `db:"size"`
}
