package repository

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/tiptok/gocomm/pkg/cache"
	"github.com/tiptok/gz-blog-microsevices/app/skateboard/interanl/pkg/db/models"
	"github.com/tiptok/gz-blog-microsevices/app/skateboard/interanl/pkg/db/transaction"
	"github.com/tiptok/gz-blog-microsevices/app/skateboard/interanl/pkg/domain"
	"gorm.io/gorm"
)

type ShopRepository struct {
	*cache.CachedRepository
}

func (repository *ShopRepository) Insert(ctx context.Context, conn transaction.Conn, dm *domain.Shop) (*domain.Shop, error) {
	var (
		err error
		m   = &models.Shop{}
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

func (repository *ShopRepository) Update(ctx context.Context, conn transaction.Conn, dm *domain.Shop) (*domain.Shop, error) {
	var (
		err error
		m   *models.Shop
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

func (repository *ShopRepository) UpdateWithVersion(ctx context.Context, transaction transaction.Conn, dm *domain.Shop) (*domain.Shop, error) {
	var (
		err error
		m   *models.Shop
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

func (repository *ShopRepository) Delete(ctx context.Context, conn transaction.Conn, dm *domain.Shop) (*domain.Shop, error) {
	var (
		tx = conn.DB()
		m  = &models.Shop{Id: dm.Identify().(int64)}
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

func (repository *ShopRepository) FindOne(ctx context.Context, conn transaction.Conn, id int64) (*domain.Shop, error) {
	var (
		err error
		tx  = conn.DB()
		m   = new(models.Shop)
	)
	queryFunc := func() (interface{}, error) {
		tx = tx.Model(m).Where("id = ?", id).First(m)
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, domain.ErrNotFound
		}
		return m, tx.Error
	}
	cacheModel := new(models.Shop)
	cacheModel.Id = id
	if err = repository.QueryCache(cacheModel.CacheKeyFunc, m, queryFunc); err != nil {
		return nil, err
	}
	return repository.ModelToDomainModel(m)
}

func (repository *ShopRepository) Find(ctx context.Context, conn transaction.Conn, queryOptions map[string]interface{}) (int64, []*domain.Shop, error) {
	var (
		tx    = conn.DB()
		ms    []*models.Shop
		dms   = make([]*domain.Shop, 0)
		total int64
	)
	queryFunc := func() (interface{}, error) {
		tx = tx.Model(&ms)
		if v, ok := queryOptions["sortByRank"]; ok && len(v.(string)) > 0 {
			tx = tx.Order("rank " + v.(string))
		} else {
			tx = tx.Order("id desc")
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

func (repository *ShopRepository) ModelToDomainModel(from *models.Shop) (*domain.Shop, error) {
	to := &domain.Shop{}
	err := copier.Copy(to, from)
	return to, err
}

func (repository *ShopRepository) DomainModelToModel(from *domain.Shop) (*models.Shop, error) {
	to := &models.Shop{}
	err := copier.Copy(to, from)
	return to, err
}

func NewShopRepository(cache *cache.CachedRepository) domain.ShopRepository {
	return &ShopRepository{CachedRepository: cache}
}
