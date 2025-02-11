package main

import (
	"encoding/json"
	"learning-fleamarket-api/infra"
	"learning-fleamarket-api/models"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

// 他のすべてのテスト関数が読み込まれる前に呼び出される
func TestMain(m *testing.M) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalln("Error loading .env.test file")
	}

	code := m.Run()
	os.Exit(code)
}

func setupTestData(db *gorm.DB) {
	items := []models.Item{
		{Name: "テストアイテム1", Price: 1000, Description: "", SoldOut: false, UserID: 1},
		{Name: "テストアイテム2", Price: 2000, Description: "テスト2", SoldOut: true, UserID: 2},
		{Name: "テストアイテム3", Price: 3000, Description: "テスト3", SoldOut: false, UserID: 3},
	}
	users := []models.User{
		{Email: "test1@example.com", Password: "test1pass"},
		{Email: "test2@example.com", Password: "test2pass"},
	}

	// DBに登録
	for _, user := range users {
		db.Create(&user)
	}
	for _, item := range items {
		db.Create(&item)
	}
}

// 各テストでsetup関数を呼び出すことでどのテストでも同じ初期データを使用することができる
func setup() *gin.Engine {
	db := infra.SetupDB()
	// DBにテーブル作成
	db.AutoMigrate(&models.Item{}, &models.User{})

	setupTestData(db)
	router := setupRouter(db)
	return router
}

func TestFindAll(t *testing.T) {
	// テストのセットアップ
	router := setup()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/items", nil)

	// APIリクエストの実行
	router.ServeHTTP(w, req)

	// APIの実行結果
	var res map[string][]models.Item
	json.Unmarshal([]byte(w.Body.String()), &res)

	// アサーション
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, 3, len(res["data"]))
}
