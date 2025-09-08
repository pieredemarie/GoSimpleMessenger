package storage

import (
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage(connStr string) *PostgresStorage{
	db, err := sql.Open("postgres",connStr)
	if err != nil {
		return nil
	}	

	return &PostgresStorage{
		db: db,
	}
}

func (p *PostgresStorage) Register(newUser *User) error {
	var exists bool 
	_ = p.db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)",newUser.Email).Scan(&exists)
	if !exists {
		fmt.Errorf("email already exists")
	}
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password),bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = p.db.Exec("INSERT into user VALUES ($1,$2,$3,$4)",newUser.FirstName,newUser.LastName,newUser.Email,hashPassword)
	return err
}

func (p *PostgresStorage) Login(email string,password string) (string, error) {
	var (
		userID int 
		correctPass string
	)

	err := p.db.QueryRow("SELECT id, password FROM user WHERE email = $1",email).Scan(&userID,&correctPass)
	if err == sql.ErrNoRows {
		return "", fmt.Errorf("user not found")
	}

	isPasswordCorrect := bcrypt.CompareHashAndPassword([]byte(correctPass),[]byte(password))
	if isPasswordCorrect != nil {
		return "", fmt.Errorf("invalid password")
	}

	//token := jwt.GenerateToken()
	return "", nil
}

func (p *PostgresStorage) GetUserInfo(email string) (*User,error) {
	newUser := &User{}
	err := p.db.QueryRow("SELECT * FROM users WHERE email = $1", email).Scan(&newUser.FirstName,&newUser.LastName,&newUser.Email,&newUser.Password)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("user not found")
	}
	return newUser,nil
}
