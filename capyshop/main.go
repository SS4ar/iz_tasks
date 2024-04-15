package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB
var jwtSecret = []byte("!&DA3A&!EDb4313aT@T!Dvbf@!$4uaf")

type User struct {
	ID       int
	Username string `form:"username"`
	Password string `form:"password"`
	Balance  float64
}

type Product struct {
	ID          int
	Name        string
	Price       float64
	ImageURL    string
	Description string
}

func main() {
	setupDatabase()
	defer db.Close()

	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	r.Static("/resources", "./resources")
	r.Static("/.git", "./git")

	r.POST("/logout", logout)

	r.POST("/login", loginHandler)
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	r.GET("/", CatalogHandler)
	r.POST("/register", registerHandler)
	r.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", nil)
	})
	r.GET("/purchase", purchasePageHandler)
	r.GET("/profile", ProfileHandler)
	r.POST("/buy", purchaseProductHandler)
	r.GET("/refund", cancelProductHandler)
	r.Run(":5000")
}

func setupDatabase() {
	connStr := fmt.Sprintf("user=postgres password=3x4mP1eP4sS dbname=postgres sslmode=disable host=%s", os.Getenv("DB_HOST"))
	dbConn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	db = dbConn
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
}

func registerHandler(c *gin.Context) {
	var user User

	// Привязываем данные из формы x-www-form-urlencoded к структуре user
	if err := c.ShouldBind(&user); err != nil {
		c.HTML(http.StatusBadRequest, "register.html", gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		c.HTML(http.StatusBadRequest, "register.html", gin.H{"error": "Error registering user"})
		return
	}

	_, err = db.Exec("INSERT INTO users (username, password, balance) VALUES ($1, $2, 50)", user.Username, string(hashedPassword))
	if err != nil {
		log.Println(err)
		c.HTML(http.StatusBadRequest, "register.html", gin.H{"error": "User already exists"})
		return
	}

	c.Redirect(http.StatusSeeOther, "/login")
}

func loginHandler(c *gin.Context) {
	var user User
	if err := c.ShouldBind(&user); err != nil {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{"error": err.Error()})
		return
	}

	var dbUser User
	err := db.QueryRow("SELECT id, username, password FROM users WHERE username = $1", user.Username).Scan(&dbUser.ID, &dbUser.Username, &dbUser.Password)
	if err != nil {
		log.Println(err)
		c.HTML(http.StatusBadRequest, "login.html", gin.H{"error": "Invalid username or password"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	if err != nil {
		log.Println(err)
		c.HTML(http.StatusBadRequest, "login.html", gin.H{"error": "Invalid username or password"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": dbUser.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		log.Fatal(err)
		c.HTML(http.StatusBadRequest, "login.html", gin.H{"error": "Error generating token"})
		return
	}

	c.SetCookie("session", tokenString, 86400, "/", "", false, true)
	c.Redirect(http.StatusSeeOther, "/")
}

func logout(c *gin.Context) {
	// Получаем текущий JWT-токен из запроса (если он есть)
	tokenString, err := c.Cookie("session") // Предполагаем, что токен хранится в куках
	if err != nil {
		// Обработка ошибки (например, отсутствие токена)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Парсим токен
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		// Обработка ошибки парсинга токена
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Проверяем, что токен действителен
	if !token.Valid {
		// Токен недействителен, возможно, он уже был инвалидирован или истек
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Далее можно добавить логику инвалидации токена (например, через хранение списка инвалидированных токенов)

	// Удаляем токен из куков (если он был сохранен там)
	c.SetCookie("session", "", -1, "/", "", false, true)

	// Перенаправляем пользователя на страницу входа или другую страницу
	c.Redirect(http.StatusSeeOther, "/login")
}

func CatalogHandler(c *gin.Context) {
	user := getCurrentUserFromContext(c)

	rows, err := db.Query("SELECT * FROM products")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var products []Product

	// Проход по результатам запроса и добавление данных в массив
	for rows.Next() {
		var product Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.ImageURL, &product.Description); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		products = append(products, product)
	}

	c.HTML(http.StatusOK, "catalog.html", gin.H{"products": products, "balance": user.Balance})
}

func ProfileHandler(c *gin.Context) {
	user := getCurrentUserFromContext(c)
	products := getUserProducts(c)

	c.HTML(http.StatusOK, "profile.html", gin.H{"balance": user.Balance, "products": products, "username": user.Username})

}

func purchasePageHandler(c *gin.Context) {
	user := getCurrentUserFromContext(c)

	productID := c.Query("productID")

	product, err := getProductByID(productID)
	if err != nil {
		// Обработайте ошибку, если товар не найден
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.HTML(http.StatusOK, "purchase.html", gin.H{"ID": product.ID, "ImageURL": product.ImageURL, "Name": product.Name, "Price": product.Price, "balance": user.Balance})
}

func purchaseProductHandler(c *gin.Context) {
	// Получите текущего пользователя из контекста или сессии
	user := getCurrentUserFromContext(c)
	// Получите ID товара, который пользователь хочет купить, из параметров запроса
	productID := c.PostForm("productID")
	promo := c.PostForm("promo")
	if promo == "CAPYDISCOUNT" {
		promo = "0.5"

	} else {
		promo = "1"
	}
	discount, err := strconv.ParseFloat(promo, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parse error"})
		return
	}
	// Проверьте, есть ли такой товар в базе данных
	product, err := getProductByID(productID)
	if err != nil {
		// Обработайте ошибку, если товар не найден
		c.HTML(http.StatusNotFound, "purchase.html", gin.H{"error": "Product not found", "ID": product.ID, "ImageURL": product.ImageURL, "Name": product.Name, "Price": product.Price, "balance": user.Balance})
		return
	}

	exists, err := hasProduct(user.ID, productID)
	if err != nil {
		c.HTML(http.StatusForbidden, "purchase.html", gin.H{"error": "Database error", "ID": product.ID, "ImageURL": product.ImageURL, "Name": product.Name, "Price": product.Price, "balance": user.Balance})
		return
	}

	if exists {
		c.HTML(http.StatusForbidden, "purchase.html", gin.H{"error": "You already have", "ID": product.ID, "ImageURL": product.ImageURL, "Name": product.Name, "Price": product.Price, "balance": user.Balance})
		return
	}

	// Проверьте, достаточно ли у пользователя средств для покупки товара
	if user.Balance < product.Price {
		// Обработайте ошибку, если у пользователя недостаточно средств
		c.HTML(http.StatusForbidden, "purchase.html", gin.H{"error": "Not enough balance", "ID": product.ID, "ImageURL": product.ImageURL, "Name": product.Name, "Price": product.Price, "balance": user.Balance})
		return
	}

	// Выполните операцию покупки, уменьшите баланс пользователя и добавьте запись о покупке в таблицу user_products
	err = buyProduct(user.ID, productID, product.Price, discount)
	if err != nil {
		// Обработайте ошибку, если операция покупки не удалась
		c.HTML(http.StatusInternalServerError, "purchase.html", gin.H{"error": "Failed to purchase product", "ID": product.ID, "ImageURL": product.ImageURL, "Name": product.Name, "Price": product.Price, "balance": user.Balance})
		return
	}

	// Верните успешный ответ, чтобы показать, что операция покупки прошла успешно
	c.HTML(http.StatusOK, "purchase.html", gin.H{"error": "Product purchased successfully", "ID": product.ID, "ImageURL": product.ImageURL, "Name": product.Name, "Price": product.Price, "balance": user.Balance})
}

func getCurrentUserFromContext(c *gin.Context) User {
	// Получаем текущий JWT-токен из запроса (предполагается, что токен хранится в куках)
	tokenString, err := c.Cookie("session")
	if err != nil {
		// Обработка ошибки (например, отсутствие токена)
		c.Redirect(http.StatusSeeOther, "/login")
		return User{}
	}

	// Парсим токен и извлекаем идентификатор пользователя (userID)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil || !token.Valid {
		// Обработка ошибки парсинга токена или недействительного токена
		c.Redirect(http.StatusSeeOther, "/login")
		return User{}
	}

	// Получаем идентификатор пользователя из токена
	claims := token.Claims.(jwt.MapClaims)
	username := claims["username"]

	var user User
	err = db.QueryRow("SELECT balance,id,username FROM users WHERE username = $1", username).Scan(&user.Balance, &user.ID, &user.Username)
	if err != nil {
		// Обработка ошибки запроса к базе данных
		c.SetCookie("session", "", -1, "/", "", false, true)
		c.Redirect(http.StatusSeeOther, "/login")
		return User{}
	}
	return user
}

func getProductByID(productID string) (*Product, error) {
	var product Product
	err := db.QueryRow("SELECT id, name, price, image_url FROM products WHERE id = $1", productID).
		Scan(&product.ID, &product.Name, &product.Price, &product.ImageURL)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func buyProduct(userID int, productID string, purchasePrice float64, discount float64) error {
	// Начните транзакцию, чтобы гарантировать целостность данных
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback() // Откат транзакции в случае ошибки

	// Уменьшите баланс пользователя
	_, err = tx.Exec("UPDATE users SET balance = balance - $1::float * $2::float WHERE id = $3", purchasePrice, discount, userID)
	if err != nil {
		return err
	}

	// Добавьте запись о покупке в таблицу user_products
	_, err = tx.Exec("INSERT INTO user_products (user_id, product_id, purchase_price) VALUES ($1, $2, $3)", userID, productID, purchasePrice)
	if err != nil {
		println("insert")
		return err
	}

	// Завершите транзакцию
	err = tx.Commit()
	if err != nil {
		println("final")
		return err
	}

	return nil
}

func getUserProducts(c *gin.Context) []Product {
	user := getCurrentUserFromContext(c)
	var products []Product
	rows, err := db.Query("SELECT * FROM products WHERE id IN (SELECT product_id FROM user_products WHERE user_id = $1)", user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var product Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.ImageURL, &product.Description); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return nil
		}

		products = append(products, product)
	}

	return products
}

func hasProduct(userID int, productID string) (bool, error) {
	var exists bool

	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM user_products WHERE user_id = $1 AND product_id = $2)", userID, productID).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func cancelProductHandler(c *gin.Context) {
	user := getCurrentUserFromContext(c)
	productID := c.Query("productID")
	product, err := getProductByID(productID)
	if err != nil {
		c.HTML(http.StatusNotFound, "purchase.html", gin.H{"error": "Product not found"})
		return
	}

	exists, err := hasProduct(user.ID, productID)
	if err != nil {
		c.HTML(http.StatusForbidden, "purchase.html", gin.H{"error": "Database error"})
		return
	}
	if !exists {
		c.HTML(http.StatusForbidden, "purchase.html", gin.H{"error": "You hasn't this product"})
		return
	}

	err = cancelProduct(user.ID, productID, product.Price)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "purchase.html", gin.H{"error": "Failed to refund product"})
		return
	}

	c.HTML(http.StatusOK, "purchase.html", gin.H{"error": "Product refunded successfully", "ID": product.ID, "ImageURL": product.ImageURL, "Name": product.Name, "Price": product.Price, "balance": user.Balance})

}

func cancelProduct(userID int, productID string, purchasePrice float64) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Уменьшите баланс пользователя
	_, err = tx.Exec("UPDATE users SET balance = balance + $1 WHERE id = $2", purchasePrice, userID)
	if err != nil {
		return err
	}

	// Добавьте запись о покупке в таблицу user_products
	_, err = tx.Exec("DELETE FROM user_products where user_id=$1 and product_id=$2", userID, productID)
	if err != nil {
		println("insert")
		return err
	}

	// Завершите транзакцию
	err = tx.Commit()
	if err != nil {
		println("final")
		return err
	}

	return nil
}
