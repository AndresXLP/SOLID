/*
Dependency Inversion Principle (DIP)

Este principio nos dice que deberíamos depender de abstracciones en lugar de implementaciones concretas.

Piensa en una puerta. No necesitas saber cómo están hechas las bisagras o cómo funciona la cerradura para poder
abrir la puerta. Solo necesitas la abstracción de “puerta” para hacerlo.

En programación, esto significa que deberíamos programar contra interfaces en lugar de clases concretas.
Esto hace que nuestro código sea más flexible y fácil de probar.
*/

package main

import "fmt"

type UserRepository interface {
	GetUserNameByID(id int) string
}

type DBUserRepository struct {
	//db gorm, sqlx... etc
}

func (dbRepo DBUserRepository) GetUserNameByID(id int) string {
	// Simulando una consulta a una base de datos

	// user := model.User
	// if err := repo.db.Table("users").
	//		Where("id = ?", id).
	//		Scan(&user).Error; err != nil {
	//		return "", err
	//	}

	// userName := user.Name
	userName := "Gopherito"
	return fmt.Sprintf("UserName is %s", userName)
}

type UserService struct {
	repo UserRepository
}

func (userService *UserService) GetUserNameByID(id int) string {
	return userService.repo.GetUserNameByID(id)
}

func main() {
	userRepository := DBUserRepository{}
	userService := UserService{repo: userRepository}

	result := userService.GetUserNameByID(123)
	fmt.Println(result)
}
