package db

import (
	"context"
	"errors"
	"strconv"
	"strings"

	"gorm.io/gorm"

	errs "github.com/cyansilver/go-libs/err"
	"github.com/cyansilver/go-libs/log"
)

var (
	// ErrRecordNotFound
	ErrRecordNotFound = errors.New("record not found")
	// ErrNotAnyRecordAffect
	ErrNotAnyRecordAffect = errors.New("not any record(s) affect")
	// ErRowsAffectNotExpected
	ErrRowsAffectNotExpected = errors.New("row(s) affect not expected")
)

var (
	ErrCheckViolation = errors.New("error: check violation")
)

const (
	ErrorCode23514 string = "23514"
)

// Repo interface for all implementations of AccountRepo
type Repo[T any] interface {
	Upsert(*T) (*T, error)
	FindOne(criteria map[string]interface{}) (T, error)
	Find(criteria map[string]interface{}) ([]T, error)
}

// LoggedRepo presents baseRepo Repo with logging feature
type LoggedRepo[T any] struct {
	baseRepo Repo[T]
}

// NewLoggedRepo returns new LoggedRepo instance
func NewLoggedRepo[T any](baseRepo Repo[T]) *LoggedRepo[T] {
	return &LoggedRepo[T]{
		baseRepo: baseRepo,
	}
}

func (r *LoggedRepo[T]) Upsert(d *T) (*T, error) {
	ret, err := r.baseRepo.Upsert(d)
	if err != nil {
		log.Logger.WithError(err).Error("Failed to upsert")
		log.Logger.WithError(err).WithField("obj", d).Trace("Failed to upsert")
	}
	return ret, err
}

func (r *LoggedRepo[T]) FindOne(criteria map[string]interface{}) (T, error) {
	ret, err := r.baseRepo.FindOne(criteria)
	if err != nil {
		log.Logger.WithError(err).Error("Failed to find one")
		log.Logger.WithError(err).WithField("criteria", criteria).Trace("Failed to find one")
	}
	return ret, err
}

func (r *LoggedRepo[T]) Find(criteria map[string]interface{}) ([]T, error) {
	ret, err := r.baseRepo.Find(criteria)
	if err != nil {
		log.Logger.WithError(err).Error("Failed to find")
		log.Logger.WithError(err).WithField("criteria", criteria).Trace("Failed to find")
	}
	return ret, err
}

// ID wrapper the relational id
type ID = uint32

type Repository[T any] struct {
	Db      *gorm.DB
	Ctx     context.Context
	tblname string
}

func NewRepository[T any](ctx context.Context, db *gorm.DB, tblname string) *Repository[T] {
	return &Repository[T]{
		Db:      db,
		Ctx:     ctx,
		tblname: tblname,
	}
}

func (r *Repository[T]) CreateBulk(m []T) ([]T, error) {
	tx := r.Db.WithContext(r.Ctx).Create(&m)
	if err := tx.Error; err != nil {
		return m, err
	}
	return m, nil
}

func (r *Repository[T]) Create(m *T) (*T, error) {
	tx := r.Db.WithContext(r.Ctx).Create(m)
	if err := tx.Error; err != nil {
		return m, err
	}
	return m, nil
}

func (r *Repository[T]) FindOne(criteria map[string]interface{}) (T, error) {
	var m T
	tx := r.Db.WithContext(r.Ctx).
		Where(criteria).
		First(&m)

	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) == true {
			return m, errs.ErrNotFound
		}
		return m, err
	}
	return m, nil
}

func (r *Repository[T]) Find(criteria map[string]interface{}) ([]T, error) {
	var m []T
	perPage, perPageOk := criteria["per_page"]
	if perPageOk == true {
		delete(criteria, "per_page")
	}

	page, pageOk := criteria["page"]
	if pageOk == true {
		delete(criteria, "page")
	}

	sort, sortOk := criteria["sort"]
	if sortOk == true {
		delete(criteria, "sort")
	} else {
		sort = "id desc"
	}

	lastId, ok := criteria["last_id"]
	if ok == true {
		if sort == "id desc" {
			criteria["id.<"] = lastId
		} else {
			criteria["id.>"] = lastId
		}
		delete(criteria, "last_id")
	}

	whereClause, newCriteria := r.GetCondition(criteria, "AND")
	// fmt.Println(whereClause)
	q := r.Db.WithContext(r.Ctx)
	if whereClause != "" {
		q = q.Where(whereClause, newCriteria)
	}
	if perPageOk == true {
		intPerPage, _ := strconv.Atoi(perPage.(string))
		q = q.Limit(intPerPage)
	}
	if pageOk == true {
		pageSize, _ := strconv.Atoi(perPage.(string))
		page, _ := strconv.Atoi(page.(string))
		offset := (page - 1) * pageSize
		q = q.Offset(offset)
	}
	q = q.Order(sort)

	tx := q.Find(&m)

	if err := tx.Error; err != nil && errors.Is(err, gorm.ErrRecordNotFound) == false {
		return nil, err
	}
	return m, nil
}

func (r *Repository[T]) Count(criteria map[string]interface{}) (int64, error) {
	var count int64
	var m T
	if _, ok := criteria["per_page"]; ok {
		delete(criteria, "per_page")
	}
	if _, ok := criteria["page"]; ok {
		delete(criteria, "page")
	}
	if _, ok := criteria["last_id"]; ok {
		delete(criteria, "last_id")
	}
	whereClause, newCriteria := r.GetCondition(criteria, "AND")
	q := r.Db.WithContext(r.Ctx)
	if whereClause != "" {
		q = q.Where(whereClause, newCriteria)
	}
	tx := q.Model(&m).Count(&count)

	if err := tx.Error; err != nil && errors.Is(err, gorm.ErrRecordNotFound) == false {
		return 0, err
	}
	return count, nil
}

func (r *Repository[T]) Update(id ID, m *T) error {
	tx := r.Db.WithContext(r.Ctx).Debug().Updates(m)
	if err := tx.Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository[T]) Delete(criteria map[string]interface{}, m *T) error {
	whereClause, newCriteria := r.GetCondition(criteria, "AND")
	tx := r.Db.WithContext(r.Ctx).
		Where(whereClause, newCriteria).Delete(m)
	if err := tx.Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository[T]) UpdateBulk(
	criteria map[string]interface{},
	data map[string]interface{},
) error {
	var m T
	whereClause, newCriteria := r.GetCondition(criteria, "AND")
	result := r.Db.WithContext(r.Ctx).Debug().
		Model(&m).Where(whereClause, newCriteria).
		Updates(data)
	return result.Error
}

// Example: ?status.in=1,2,3&name=john&name.like=abc&created_At.between=1/1/2018,1/2/2019
func (r *Repository[T]) GetCondition(criteria map[string]interface{}, oper string) (string, map[string]interface{}) {
	where := ""

	for key, v := range criteria {
		field, operator := splitString(key, ".")
		if len(field) <= 0 {
			continue
		}
		if where != "" {
			where = where + " " + oper + " "
		}
		switch operator {
		case ">=":
			where = where + "`" + field + "`" + " >= @" + key
		case ">":
			where = where + "`" + field + "`" + " > @" + key
		case "<":
			where = where + "`" + field + "`" + " < @" + key
		case "<=":
			where = where + "`" + field + "`" + " <= @" + key
		case "in":
			where = where + "`" + field + "`" + " in @" + key
		case "search":
			where = where + "MATCH(`" + field + "`)" + " AGAINST (@" + key + ")"
		case "like":
			where = where + "`" + field + "`" + " like @" + key
			criteria[key] = "%" + v.(string) + "%"
		default:
			where = where + "`" + field + "`" + " = @" + key
		}
	}

	return where, criteria
}

func splitString(field, sep string) (string, string) {
	result := strings.Split(field, sep)

	if len(result) >= 2 {
		return result[0], result[1]
	}

	return result[0], ""
}
