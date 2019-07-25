package fazzrepository

import (
	"context"
	"reflect"

	"github.com/payfazz/go-apt/pkg/fazzdb"
)

// RepositoryInterface is an interface for repository
type RepositoryInterface interface {
	GetQuery(ctx context.Context) (*fazzdb.Query, error)

	FindAll(ctx context.Context, conditions []fazzdb.SliceCondition, orders []fazzdb.Order, limit int, offset int) (interface{}, error)
	FindOne(ctx context.Context, conditions []fazzdb.SliceCondition, orders []fazzdb.Order) (interface{}, error)
	Find(ctx context.Context, ID interface{}) (interface{}, error)

	Create(ctx context.Context, m fazzdb.ModelInterface) (interface{}, error)
	Update(ctx context.Context, m fazzdb.ModelInterface) (bool, error)
	Delete(ctx context.Context, m fazzdb.ModelInterface) (bool, error)

	Count(ctx context.Context, conditions []fazzdb.SliceCondition) (*float64, error)
}

type Repository struct {
	model fazzdb.ModelInterface
}

// GetQuery is a function to get query from ctx
func (r *Repository) GetQuery(ctx context.Context) (*fazzdb.Query, error) {
	return fazzdb.GetTransactionOrQueryContext(ctx)
}

// FindAll is a function that used to get all data
func (r *Repository) FindAll(
	ctx context.Context,
	conditions []fazzdb.SliceCondition,
	orders []fazzdb.Order,
	limit int,
	offset int,
) (interface{}, error) {
	q, err := r.GetQuery(ctx)
	if nil != err {
		return nil, err
	}

	current := q.Use(r.model).
		WhereMany(conditions...).
		OrderByMany(orders...)

	if limit > 0 {
		current.WithLimit(limit)
	} else if limit == NO_LIMIT {
		current.WithLimit(0)
	}

	if offset > 0 {
		current.WithOffset(offset)
	}

	return current.AllCtx(ctx)
}

// FindOneBy is a function that used to find one data with some conditions
func (r *Repository) FindOne(
	ctx context.Context,
	conditions []fazzdb.SliceCondition,
	orders []fazzdb.Order,
) (interface{}, error) {
	q, err := r.GetQuery(ctx)
	if nil != err {
		return nil, err
	}

	current := q.Use(r.model).
		WhereMany(conditions...).
		OrderByMany(orders...).
		WithLimit(1)

	rows, err := current.AllCtx(ctx)
	if nil != err {
		return nil, err
	}

	val := reflect.ValueOf(rows)
	if val.Len() == 0 {
		return nil, NewEmptyResultError()
	}

	return val.Index(0).Interface(), nil
}

// Find is a function that used to find the data
func (r *Repository) Find(ctx context.Context, ID interface{}) (interface{}, error) {
	q, err := r.GetQuery(ctx)
	if nil != err {
		return nil, err
	}

	rows, err := q.Use(r.model).
		Where("id", ID).
		WithLimit(1).
		AllCtx(ctx)

	if nil != err {
		return nil, err
	}

	val := reflect.ValueOf(rows)
	if val.Len() == 0 {
		return nil, NewEmptyResultError()
	}

	return val.Index(0).Interface(), nil
}

// Create is a function that used to create the data
func (r *Repository) Create(ctx context.Context, m fazzdb.ModelInterface) (interface{}, error) {
	q, err := r.GetQuery(ctx)
	if nil != err {
		return nil, err
	}

	result, err := q.Use(m).
		InsertCtx(ctx, false)

	if nil != err {
		return nil, err
	}

	return result, nil
}

// Update is a function that used to update the data
func (r *Repository) Update(ctx context.Context, m fazzdb.ModelInterface) (bool, error) {
	q, err := r.GetQuery(ctx)
	if nil != err {
		return false, err
	}

	_, err = q.Use(m).
		UpdateCtx(ctx)

	if nil != err {
		return false, err
	}

	return true, nil
}

// Delete is a function to delete the data
func (r *Repository) Delete(ctx context.Context, m fazzdb.ModelInterface) (bool, error) {
	q, err := r.GetQuery(ctx)
	if nil != err {
		return false, err
	}

	_, err = q.Use(m).
		DeleteCtx(ctx)

	if nil != err {
		return false, err
	}

	return true, err
}

// Count is a function that used to get count of all data with given conditions
func (r *Repository) Count(ctx context.Context, conditions []fazzdb.SliceCondition) (*float64, error) {
	q, err := r.GetQuery(ctx)
	if nil != err {
		return nil, err
	}

	return q.Use(r.model).
		WhereMany(conditions...).
		WithLimit(0).
		CountCtx(ctx)
}

// NewRepository is a constructor for cash back repo
func NewRepository() RepositoryInterface {
	return &Repository{}
}
