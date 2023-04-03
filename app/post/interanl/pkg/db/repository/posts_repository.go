package repository

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/tiptok/gocomm/pkg/cache"
	"github.com/tiptok/gz-blog-microsevices/app/post/interanl/pkg/db/models"
	"github.com/tiptok/gz-blog-microsevices/app/post/interanl/pkg/db/transaction"
	"github.com/tiptok/gz-blog-microsevices/app/post/interanl/pkg/domain"
	"gorm.io/gorm"
)

type PostsRepository struct {
	*cache.CachedRepository
}

func (repository *PostsRepository) Insert(ctx context.Context, conn transaction.Conn, dm *domain.Posts) (*domain.Posts, error) {
	var (
		err error
		m   = &models.Posts{}
		tx  = conn.DB()
	)
	if m, err = repository.DomainModelToModel(dm); err != nil {
		return nil, err
	}
	if tx = tx.Model(m).Save(m); tx.Error != nil {
		return nil, tx.Error
	}
	dm.Id = m.Id
	return dm, nil

}

func (repository *PostsRepository) Update(ctx context.Context, conn transaction.Conn, dm *domain.Posts) (*domain.Posts, error) {
	var (
		err error
		m   *models.Posts
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

func (repository *PostsRepository) UpdateUnscoped(ctx context.Context, conn transaction.Conn, dm *domain.Posts) (*domain.Posts, error) {
	var (
		err error
		m   *models.Posts
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

func (repository *PostsRepository) UpdateWithVersion(ctx context.Context, transaction transaction.Conn, dm *domain.Posts) (*domain.Posts, error) {
	var (
		err error
		m   *models.Posts
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

func (repository *PostsRepository) Delete(ctx context.Context, conn transaction.Conn, dm *domain.Posts) (*domain.Posts, error) {
	var (
		tx = conn.DB()
		m  = &models.Posts{Id: dm.Identify().(int64)}
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

func (repository *PostsRepository) FindOne(ctx context.Context, conn transaction.Conn, id int64) (*domain.Posts, error) {
	var (
		err error
		tx  = conn.DB()
		m   = new(models.Posts)
	)
	queryFunc := func() (interface{}, error) {
		tx = tx.Model(m).Where("id = ?", id).First(m)
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, domain.ErrNotFound
		}
		return m, tx.Error
	}
	cacheModel := new(models.Posts)
	cacheModel.Id = id
	if err = repository.QueryCache(cacheModel.CacheKeyFunc, m, queryFunc); err != nil {
		return nil, err
	}
	return repository.ModelToDomainModel(m)
}

func (repository *PostsRepository) FindOneUnscoped(ctx context.Context, conn transaction.Conn, id int64) (*domain.Posts, error) {
	var (
		err error
		tx  = conn.DB()
		m   = new(models.Posts)
	)
	err = tx.Unscoped().First(m, id).Error
	if err != nil {
		return nil, err
	}
	return repository.ModelToDomainModel(m)
}

func (repository *PostsRepository) Find(ctx context.Context, conn transaction.Conn, queryOptions map[string]interface{}) (int64, []*domain.Posts, error) {
	var (
		tx    = conn.DB()
		ms    []*models.Posts
		dms   = make([]*domain.Posts, 0)
		total int64
	)
	queryFunc := func() (interface{}, error) {
		tx = tx.Model(&ms).Order("id desc")
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

func (repository *PostsRepository) ModelToDomainModel(from *models.Posts) (*domain.Posts, error) {
	to := &domain.Posts{}
	err := copier.Copy(to, from)
	return to, err
}

func (repository *PostsRepository) DomainModelToModel(from *domain.Posts) (*models.Posts, error) {
	to := &models.Posts{}
	err := copier.Copy(to, from)
	return to, err
}

func NewPostsRepository(cache *cache.CachedRepository) domain.PostsRepository {
	return &PostsRepository{CachedRepository: cache}
}
