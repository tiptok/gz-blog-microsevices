package repository

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/tiptok/gocomm/pkg/cache"
	"github.com/tiptok/gz-blog-microsevices/app/comment/interanl/pkg/db/models"
	"github.com/tiptok/gz-blog-microsevices/app/comment/interanl/pkg/db/transaction"
	"github.com/tiptok/gz-blog-microsevices/app/comment/interanl/pkg/domain"
	"gorm.io/gorm"
)

type CommentsRepository struct {
	*cache.CachedRepository
}

func (repository *CommentsRepository) Insert(ctx context.Context, conn transaction.Conn, dm *domain.Comments) (*domain.Comments, error) {
	var (
		err error
		m   = &models.Comments{}
		tx  = conn.DB()
	)
	if m, err = repository.DomainModelToModel(dm); err != nil {
		return nil, err
	}
	if tx = tx.Model(m).Save(m); tx.Error != nil {
		return nil, tx.Error
	}
	dm.Id = m.Id
	return repository.ModelToDomainModel(m)

}

func (repository *CommentsRepository) Update(ctx context.Context, conn transaction.Conn, dm *domain.Comments) (*domain.Comments, error) {
	var (
		err error
		m   *models.Comments
		tx  = conn.DB()
	)
	if m, err = repository.DomainModelToModel(dm); err != nil {
		return nil, err
	}
	queryFunc := func() (interface{}, error) {
		tx = tx.Model(m).Updates(m)
		return nil, tx.Error
	}
	if _, err = repository.Query(queryFunc, m.CacheKeyFunc()); err != nil {
		return nil, err
	}
	return dm, nil
}

func (repository *CommentsRepository) UpdateUnscoped(ctx context.Context, conn transaction.Conn, dm *domain.Comments) (*domain.Comments, error) {
	var (
		err error
		m   *models.Comments
		tx  = conn.DB()
	)
	if m, err = repository.DomainModelToModel(dm); err != nil {
		return nil, err
	}
	queryFunc := func() (interface{}, error) {
		tx = tx.Unscoped().Model(m).Updates(m)
		return nil, tx.Error
	}
	if _, err = repository.Query(queryFunc, m.CacheKeyFunc()); err != nil {
		return nil, err
	}
	return dm, nil
}

func (repository *CommentsRepository) UpdateWithVersion(ctx context.Context, transaction transaction.Conn, dm *domain.Comments) (*domain.Comments, error) {
	var (
		err error
		m   *models.Comments
		tx  = transaction.DB()
	)
	if m, err = repository.DomainModelToModel(dm); err != nil {
		return nil, err
	}
	oldVersion := dm.Version
	m.Version += 1
	queryFunc := func() (interface{}, error) {
		tx = tx.Model(m).Where("id = ?", m.Id).Where("version = ?", oldVersion).Updates(m)
		return nil, tx.Error
	}
	if _, err = repository.Query(queryFunc, m.CacheKeyFunc()); err != nil {
		return nil, err
	}
	return dm, nil
}

func (repository *CommentsRepository) Delete(ctx context.Context, conn transaction.Conn, dm *domain.Comments) (*domain.Comments, error) {
	var (
		tx = conn.DB()
		m  = &models.Comments{Id: dm.Identify().(int64)}
	)
	queryFunc := func() (interface{}, error) {
		tx = tx.Where("id = ?", m.Id).Delete(m)
		return m, tx.Error
	}
	if _, err := repository.Query(queryFunc, m.CacheKeyFunc()); err != nil {
		return dm, err
	}
	return dm, nil
}

func (repository *CommentsRepository) FindOne(ctx context.Context, conn transaction.Conn, id int64) (*domain.Comments, error) {
	var (
		err error
		tx  = conn.DB()
		m   = new(models.Comments)
	)
	queryFunc := func() (interface{}, error) {
		tx = tx.Model(m).Where("id = ?", id).First(m)
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, domain.ErrNotFound
		}
		return m, tx.Error
	}
	cacheModel := new(models.Comments)
	cacheModel.Id = id
	if err = repository.QueryCache(cacheModel.CacheKeyFunc, m, queryFunc); err != nil {
		return nil, err
	}
	return repository.ModelToDomainModel(m)
}

func (repository *CommentsRepository) FindOneUnscoped(ctx context.Context, conn transaction.Conn, id int64) (*domain.Comments, error) {
	var (
		err error
		tx  = conn.DB()
		m   = new(models.Comments)
	)
	err = tx.Unscoped().First(m, id).Error
	if err != nil {
		return nil, err
	}
	return repository.ModelToDomainModel(m)
}

func (repository *CommentsRepository) FindOneByUuid(ctx context.Context, conn transaction.Conn, uuid string) (*domain.Comments, error) {
	var (
		err error
		tx  = conn.DB()
		m   = new(models.Comments)
	)
	queryFunc := func() (interface{}, error) {
		tx = tx.Model(m).Where("uuid = ?", uuid).First(m)
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, domain.ErrNotFound
		}
		return m, tx.Error
	}
	if _, err = repository.Query(queryFunc); err != nil {
		return nil, err
	}
	return repository.ModelToDomainModel(m)
}

func (repository *CommentsRepository) Find(ctx context.Context, conn transaction.Conn, queryOptions map[string]interface{}) (int64, []*domain.Comments, error) {
	var (
		tx    = conn.DB()
		ms    []*models.Comments
		dms   = make([]*domain.Comments, 0)
		total int64
	)
	queryFunc := func() (interface{}, error) {
		tx = tx.Model(&ms).Order("id desc")
		if v, ok := queryOptions["postId"]; ok {
			tx.Where("post_id = ?", v)
		}
		if v, ok := queryOptions["unscoped"]; ok && v.(bool) {
			tx.Unscoped()
		}
		if total, tx = transaction.PaginationAndCount(ctx, tx, queryOptions, &ms); tx.Error != nil {
			return dms, tx.Error
		}
		return dms, nil
	}

	if _, err := repository.Query(queryFunc); err != nil {
		return 0, nil, err
	}

	for _, item := range ms {
		if dm, err := repository.ModelToDomainModel(item); err != nil {
			return 0, dms, err
		} else {
			dms = append(dms, dm)
		}
	}
	return total, dms, nil
}

func (repository *CommentsRepository) ModelToDomainModel(from *models.Comments) (*domain.Comments, error) {
	to := &domain.Comments{}
	err := copier.Copy(to, from)
	return to, err
}

func (repository *CommentsRepository) DomainModelToModel(from *domain.Comments) (*models.Comments, error) {
	to := &models.Comments{}
	err := copier.Copy(to, from)
	return to, err
}

func NewCommentsRepository(cache *cache.CachedRepository) domain.CommentsRepository {
	return &CommentsRepository{CachedRepository: cache}
}
