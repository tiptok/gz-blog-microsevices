package repository

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/tiptok/gocomm/pkg/cache"
	"github.com/tiptok/gz-blog-microsevices/app/user/interanl/pkg/db/models"
	"github.com/tiptok/gz-blog-microsevices/app/user/interanl/pkg/db/transaction"
	"github.com/tiptok/gz-blog-microsevices/app/user/interanl/pkg/domain"
	"gorm.io/gorm"
)

type UsersRepository struct {
	*cache.CachedRepository
}

func (repository *UsersRepository) FindByIDs(ctx context.Context, conn transaction.Conn, ids []int64) (int64, []*domain.Users, error) {
	return repository.Find(ctx, conn, map[string]interface{}{"inUserIds": ids})
}

func (repository *UsersRepository) FindOneByEmail(ctx context.Context, conn transaction.Conn, email string) (*domain.Users, error) {
	var (
		err error
		tx  = conn.DB()
		m   = new(models.Users)
	)
	queryFunc := func() (interface{}, error) {
		tx = tx.Model(m).Where("email = ?", email).First(m)
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

func (repository *UsersRepository) FindOneByUsername(ctx context.Context, conn transaction.Conn, username string) (*domain.Users, error) {
	var (
		err error
		tx  = conn.DB()
		m   = new(models.Users)
	)
	queryFunc := func() (interface{}, error) {
		tx = tx.Model(m).Where("username = ?", username).First(m)
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

func (repository *UsersRepository) Insert(ctx context.Context, conn transaction.Conn, dm *domain.Users) (*domain.Users, error) {
	var (
		err error
		m   = &models.Users{}
		tx  = conn.DB()
	)
	if m, err = repository.DomainModelToModel(dm); err != nil {
		return nil, err
	}
	if tx = tx.Model(m).Save(m); tx.Error != nil { //.Omit("deleted_at")
		return nil, tx.Error
	}
	dm.Id = m.Id
	return dm, nil

}

func (repository *UsersRepository) Update(ctx context.Context, conn transaction.Conn, dm *domain.Users) (*domain.Users, error) {
	var (
		err error
		m   *models.Users
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

func (repository *UsersRepository) UpdateWithVersion(ctx context.Context, transaction transaction.Conn, dm *domain.Users) (*domain.Users, error) {
	var (
		err error
		m   *models.Users
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

func (repository *UsersRepository) Delete(ctx context.Context, conn transaction.Conn, dm *domain.Users) (*domain.Users, error) {
	var (
		tx = conn.DB()
		m  = &models.Users{Id: dm.Identify().(int64)}
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

func (repository *UsersRepository) FindOne(ctx context.Context, conn transaction.Conn, id int64) (*domain.Users, error) {
	var (
		err error
		tx  = conn.DB()
		m   = new(models.Users)
	)
	queryFunc := func() (interface{}, error) {
		tx = tx.Model(m).Where("id = ?", id).First(m)
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, domain.ErrNotFound
		}
		return m, tx.Error
	}
	cacheModel := new(models.Users)
	cacheModel.Id = id
	if err = repository.QueryCache(cacheModel.CacheKeyFunc, m, queryFunc); err != nil {
		return nil, err
	}
	return repository.ModelToDomainModel(m)
}

func (repository *UsersRepository) Find(ctx context.Context, conn transaction.Conn, queryOptions map[string]interface{}) (int64, []*domain.Users, error) {
	var (
		tx    = conn.DB()
		ms    []*models.Users
		dms   = make([]*domain.Users, 0)
		total int64
	)
	queryFunc := func() (interface{}, error) {
		tx = tx.Model(&ms).Order("id desc")
		if v, ok := queryOptions["inUserIds"]; ok && len(v.([]int64)) > 0 {
			tx.Where("id in ?", v.([]int64))
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

func (repository *UsersRepository) ModelToDomainModel(from *models.Users) (*domain.Users, error) {
	to := &domain.Users{}
	err := copier.Copy(to, from)
	return to, err
}

func (repository *UsersRepository) DomainModelToModel(from *domain.Users) (*models.Users, error) {
	to := &models.Users{}
	err := copier.Copy(to, from)
	return to, err
}

func NewUsersRepository(cache *cache.CachedRepository) domain.UsersRepository {
	return &UsersRepository{CachedRepository: cache}
}
