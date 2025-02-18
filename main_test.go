package main

import (
	"bytes"
	"encoding/json"
	"learning-fleamarket-api/dto"
	"learning-fleamarket-api/infra"
	"learning-fleamarket-api/models"
	"learning-fleamarket-api/services"
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

func setupTestData(t *testing.T, db *gorm.DB, router *gin.Engine) {
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
		reqBody, _ := json.Marshal(user)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/auth/signup", bytes.NewBuffer(reqBody))
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusCreated, w.Code)
	}
	for _, item := range items {
		db.Create(&item)
	}
}

// 各テストでsetup関数を呼び出すことでどのテストでも同じ初期データを使用することができる
func setup(t *testing.T) *gin.Engine {
	db := infra.SetupDB()
	// DBにテーブル作成
	db.AutoMigrate(&models.Item{}, &models.User{})
	router := setupRouter(db)

	setupTestData(t, db, router)
	return router
}

func TestFindAll(t *testing.T) {
	// テストのセットアップ
	router := setup(t)

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

// 認証が必要なテスト
func TestCreate(t *testing.T) {
	// テストのセットアップ
	router := setup(t)
	token, err := services.CreateToken(1, "test1@example.com")
	assert.Equal(t, nil, err)

	createItemInput := dto.CreateItemInput{
		Name:        "テストアイテム4",
		Price:       4000,
		Description: "Createテスト",
	}
	reqBody, _ := json.Marshal(createItemInput)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/items", bytes.NewBuffer(reqBody))
	req.Header.Set("Authorization", "Bearer "+*token)

	// APIリクエストの実行
	router.ServeHTTP(w, req)

	// APIの実行結果を取得
	var res map[string]models.Item
	json.Unmarshal([]byte(w.Body.String()), &res)

	// アサーション
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, uint(4), res["data"].ID)
}

func TestCreateUnAuthorized(t *testing.T) {
	// テストのセットアップ
	router := setup(t)

	createItemInput := dto.CreateItemInput{
		Name:        "テストアイテム4",
		Price:       4000,
		Description: "Createテスト",
	}
	reqBody, _ := json.Marshal(createItemInput)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/items", bytes.NewBuffer(reqBody))

	// APIリクエストの実行
	router.ServeHTTP(w, req)

	// APIの実行結果を取得
	var res map[string]models.Item
	json.Unmarshal([]byte(w.Body.String()), &res)

	// アサーション
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestSignUp(t *testing.T) {
	router := setup(t)
	signUpInput := dto.SignUpInput{
		Email:    "test3@example.com",
		Password: "test3pass",
	}
	reqBody, _ := json.Marshal(signUpInput)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/auth/signup", bytes.NewBuffer(reqBody))

	// APIリクエストの実行
	router.ServeHTTP(w, req)

	// 実行結果の確認
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestLogin(t *testing.T) {
	router := setup(t)
	loginInput := dto.LoginInput{
		Email:    "test2@example.com",
		Password: "test2pass",
	}
	reqBody, _ := json.Marshal(loginInput)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(reqBody))
	router.ServeHTTP(w, req)

	// APIリクエストの実行
	router.ServeHTTP(w, req)

	// 実行結果の確認
	assert.Equal(t, http.StatusOK, w.Code)
}
