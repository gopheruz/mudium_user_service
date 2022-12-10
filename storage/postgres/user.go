package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/nurmuhammaddeveloper/mudium_user_service/pkg/utils"
	"github.com/nurmuhammaddeveloper/mudium_user_service/querys"
	"github.com/nurmuhammaddeveloper/mudium_user_service/storage/repo"
)

type userRepo struct {
	db *sqlx.DB
}

func NewUser(db *sqlx.DB) repo.UserStorageI {
	return &userRepo{
		db: db,
	}
}

func (ur *userRepo) Create(u *repo.User) (*repo.User, error) {
	hashpsw, err := utils.HashPassword(u.Password)
	if err != nil {
		return nil, err
	}

	row := ur.db.QueryRow(
		querys.UserCreateQuery,
		u.FirstName,
		u.LastName,
		utils.NullString(u.PhoneNumber),
		u.Email,
		utils.NullString(u.Gender),
		hashpsw,
		utils.NullString(u.UserName),
		utils.NullString(u.ProfileImageUrl),
		u.Type,
	)
	err = row.Scan(
		&u.ID,
		&u.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return u, err

}
func (ur *userRepo) GetByEmail(email string) (*repo.User, error) {
	return nil, nil
}
func (ur *userRepo) GetAll(params *repo.GetAllUsersParams) (*repo.GetAllUsersResult, error) {
	return nil, nil
}
func (ur *userRepo) UpdatePassword(req *repo.UpdatePassword) error {
	return nil
}
func (ur *userRepo) Update(u *repo.User) (*repo.User, error) {
	return nil, nil
}
func (ur *userRepo) Delete(id int64) error {
	return nil
}
func (ur *userRepo) Get(id int64) (*repo.User, error) {
	return nil, nil
}
