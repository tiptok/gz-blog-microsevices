package repository

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/tiptok/gocomm/pkg/cache"
	"github.com/tiptok/gz-blog-microsevices/app/product/interanl/pkg/db/models"
	"github.com/tiptok/gz-blog-microsevices/app/product/interanl/pkg/db/transaction"
	"github.com/tiptok/gz-blog-microsevices/app/product/interanl/pkg/domain"
	"gorm.io/gorm"
)

type ProductsRepository struct {
	*cache.CachedRepository
}

func (repository *ProductsRepository) Insert(ctx context.Context, conn transaction.Conn, dm *domain.Products) (*domain.Products, error) {
	var (
		err error
		m   = &models.Products{}
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

func (repository *ProductsRepository) Update(ctx context.Context, conn transaction.Conn, dm *domain.Products) (*domain.Products, error) {
	var (
		err error
		m   *models.Products
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

func (repository *ProductsRepository) UpdateWithVersion(ctx context.Context, transaction transaction.Conn, dm *domain.Products) (*domain.Products, error) {
	var (
		err error
		m   *models.Products
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

func (repository *ProductsRepository) Delete(ctx context.Context, conn transaction.Conn, dm *domain.Products) (*domain.Products, error) {
	var (
		tx = conn.DB()
		m  = &models.Products{Id: dm.Identify().(int64)}
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

func (repository *ProductsRepository) FindOne(ctx context.Context, conn transaction.Conn, id int64) (*domain.Products, error) {
	var (
		err error
		tx  = conn.DB()
		m   = new(models.Products)
	)
	queryFunc := func() (interface{}, error) {
		tx = tx.Model(m).Where("id = ?", id).First(m)
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, domain.ErrNotFound
		}
		return m, tx.Error
	}
	cacheModel := new(models.Products)
	cacheModel.Id = id
	if err = repository.QueryCache(cacheModel.CacheKeyFunc, m, queryFunc); err != nil {
		return nil, err
	}
	return repository.ModelToDomainModel(m)
}

func (repository *ProductsRepository) Find(ctx context.Context, conn transaction.Conn, queryOptions map[string]interface{}) (int64, []*domain.Products, error) {
	var (
		tx    = conn.DB()
		ms    []*models.Products
		dms   = make([]*domain.Products, 0)
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

func (repository *ProductsRepository) ModelToDomainModel(from *models.Products) (*domain.Products, error) {
	to := &domain.Products{}
	err := copier.Copy(to, from)
	return to, err
}

func (repository *ProductsRepository) DomainModelToModel(from *domain.Products) (*models.Products, error) {
	to := &models.Products{}
	err := copier.Copy(to, from)
	return to, err
}

func NewProductsRepository(cache *cache.CachedRepository) domain.ProductsRepository {
	return &ProductsRepository{CachedRepository: cache}
}
