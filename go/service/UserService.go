package service

import (
	"go-hello/go/dao"
	"go-hello/go/entity"
)

// create
func CreateUser(user *entity.User) (err error) {
	if err = dao.SqlSession.Create(user).Error; err != nil {
		return err
	}
	return
}

// read
func GetUserById(id string) (user *entity.User, err error) {
	if err = dao.SqlSession.Where("id=?", id).First(user).Error; err != nil {
		return nil, err
	}
	return
}

// read
func GetAllUser() (userList []*entity.User, err error) {
	if err := dao.SqlSession.Find(&userList).Error; err != nil {
		return nil, err
	}
	return
}

// update
func UpdateUser(user *entity.User) (err error) {
	err = dao.SqlSession.Save(user).Error
	return
}

// delete
func DeleteUserById(id string) (err error) {
	err = dao.SqlSession.Where("id=?", id).Delete(&entity.User{}).Error
	return
}

// create
// func (m *dao.DBClient) Insert(user entity.User) error {
// 	res := m.client.Create(&user)
// 	if res.Error != nil {
// 		return res.Error
// 	}
// 	return nil
// }

//  read
// func (m *dao.DBClient) Get() ([]entity.User, error) {
// 	users := []entity.User{}
// 	res := m.client.Order("id").Find(&users)
// 	if res.Error != nil {
// 		return nil, res.Error
// 	}
// 	return users, nil
// }

// update
// func (m *dao.DBClient) Update(user entity.User) error {
// 	res := m.client.Save(&user)
// 	if res.Error != nil {
// 		return res.Error
// 	}
// 	return nil
// }

//  delete
// func (m *dao.DBClient) Delete(user entity.User) error {
// 	res := m.client.Delete(&user)
// 	if res.Error != nil {
// 		return res.Error
// 	}
// 	return nil
// }
