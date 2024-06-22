package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xams-backend/services/auth-service/src/internal/models"
	"github.com/xams-backend/services/auth-service/src/packages/utils/mock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	dsn_test := "root:123@tcp(localhost:3306)/xams?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn_test), &gorm.Config{})
	return db
}

func TestRepositoryFunctions(t *testing.T) {
    testCases := []struct {
        testName string
        setupFunc func() (*gorm.DB, *UserRepository)
        action func(repo *UserRepository, mockUser *models.User) error
        assertFunc func(t *testing.T, db *gorm.DB, mockUser *models.User)
    }{
        {
            testName: "CreateUser",
            setupFunc: func() (*gorm.DB, *UserRepository) {
                db := setupTestDB()
                return db, NewUserRepository(db)
            },
            action: func(repo *UserRepository, mockUser *models.User) error {
                return repo.CreateUser(mockUser)
            },
            assertFunc: func(t *testing.T, db *gorm.DB, mockUser *models.User) {
                var user models.User
                result := db.First(&user, "user_id = ?", mockUser.UserID)

                assert.Nil(t, result.Error)
                assert.Equal(t, mockUser.UserID, user.UserID)
                assert.Equal(t, mockUser.Email, user.Email)
                assert.Equal(t, mockUser.FirstName, user.FirstName)
                assert.Equal(t, mockUser.LastName, user.LastName)
            },
        },
        {
            testName: "GetUser",
            setupFunc: func() (*gorm.DB, *UserRepository) {
                db := setupTestDB()
                repo := NewUserRepository(db)
                return db, repo
            },
            action: func(repo *UserRepository, mockUser *models.User) error {
                _, err := repo.GetUser("39035909950590")
                return err
            },
            assertFunc: func(t *testing.T, db *gorm.DB, mockUser *models.User) {
                // Additional assertions can be added here for GetUser
                // Example:
                // assert.Equal(t, mockUser.Email, user.Email)
            },
        },
        // Add more test cases as needed
    }

    for _, tc := range testCases {
        t.Run(tc.testName, func(t *testing.T) {
            db, repo := tc.setupFunc()
            mockUser := mock.GenerateMockUser()

            err := tc.action(repo, mockUser)
            assert.Nil(t, err)

            tc.assertFunc(t, db, mockUser)
        })
    }
}
