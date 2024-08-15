package postgres

import (
	ap "auth/genproto/auth"
	"auth/verification"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	t "auth/api/token"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UsersRepo struct {
	db  *sql.DB
	rdb *redis.Client
}

func NewUsersRepo(db *sql.DB, rdb *redis.Client) *UsersRepo {
	return &UsersRepo{db: db, rdb: rdb}
}

func (u *UsersRepo) Register(req *ap.UserCreateReq) (*ap.Void, error) {
	id := uuid.New().String()

	query := `
	INSERT INTO users(
		id, 
		email,
		password,
		first_name,
		last_name,
		phone,
		role
	) VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := u.db.Exec(query, id, req.Email, req.Password, req.FirstName, req.LastName, req.Phone, req.Role)

	if err != nil {
		log.Println("Error while registering user: ", err)
		return nil, err
	}

	log.Println("Successfully registered user")

	return nil, nil
}

func (u *UsersRepo) Login(req *ap.UserLoginReq) (*ap.TokenRes, error) {
	var id string
	var email string
	var firstName string
	var lastName string
	var password string
	var role string

	query := `
	SELECT 
		id,
		email,
		password,
		first_name,
		last_name,
		role
	FROM 
		users
	WHERE
		email = $1
	AND 
		deleted_at = 0
	`

	row := u.db.QueryRow(query, req.Email)

	err := row.Scan(
		&id,
		&email,
		&password,
		&firstName,
		&lastName,
		&role,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}

	if err != nil {
		log.Println("Error while login user: ", err)
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid email or password")
	}

	token := t.GenerateJWTToken(id, email, firstName+" "+lastName, role)
	tokens := ap.TokenRes{
		Token: token.AccessToken,
		ExpAt: "1 hours",
	}

	return &tokens, nil
}

func (u *UsersRepo) Profile(req *ap.ById) (*ap.UserRes, error) {
	userData, err := u.rdb.Get(context.Background(), req.Id).Result()
	if err == redis.Nil {
		user := ap.UserRes{}

		query := `
	SELECT 
		id,
		first_name,
		last_name,
		role,
		phone
	FROM 	
		users
	WHERE
		id = $1
	AND 
		deleted_at = 0
	`
		row := u.db.QueryRow(query, req.Id)

		err := row.Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.Role,
			&user.Phone,
		)

		if err != nil {
			log.Println("Error while getting user profile: ", err)
			return nil, err
		}

		fmt.Println("Successfully got profile")

		return &user, nil
		
	} else if err != nil {
		log.Printf("Redis get error: %v", err)
		return nil, err
	}

	user := &ap.UserRes{}

	err = json.Unmarshal([]byte(userData), user)
	if err != nil {
		log.Printf("JSON unmarshalling error: %v", err)
		return nil, err
	}

	return user, nil
}

func (u *UsersRepo) UpdateProfile(req *ap.UserUpdateReq) (*ap.Void, error) {

	query := `
	UPDATE
		users
	SET 
	`

	var conditions []string
	var args []interface{}

	if req.FirstName != "" && req.FirstName != "string" {
		conditions = append(conditions, " first_name = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.FirstName)
	}
	if req.LastName != "" && req.LastName != "string" {
		conditions = append(conditions, " last_name = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.LastName)
	}
	if req.Email != "" && req.Email != "string" {
		conditions = append(conditions, " email = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Email)
	}

	if len(conditions) == 0 {
		return nil, errors.New("nothing to update")
	}

	conditions = append(conditions, " updated_at = now()")
	query += strings.Join(conditions, ", ")
	query += " WHERE id = $" + strconv.Itoa(len(args)+1) + " AND deleted_at = 0 "

	args = append(args, req.Id)

	_, err := u.db.Exec(query, args...)

	if err != nil {
		log.Println("Error while updating user profile: ", err)
		return nil, err
	}

	log.Println("Successfully updated user profile")

	return nil, nil
}

func (u *UsersRepo) DeleteProfile(req *ap.ById) (*ap.Void, error) {
	query := `
	UPDATE 
		users
	SET 
		deleted_at = EXTRACT(EPOCH FROM NOW())
	WHERE 
		id = $1
	AND 
		deleted_at = 0
	`

	res, err := u.db.Exec(query, req.Id)
	if err != nil {
		log.Println("Error while deleting user: ", err)
		return nil, err
	}

	if r, err := res.RowsAffected(); r == 0 {
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("user with id %s not found", req.Id)
	}

	log.Println("Successfully deleted user")

	return nil, nil
}

func (u *UsersRepo) RefreshToken(req *ap.ById) (*ap.TokenRes, error) {
	var id string
	var email string
	var username string
	var role string

	query := `
	SELECT 
		id,
		username,
		email,
		role
	FROM 
		users
	WHERE
		id = $1
	AND 
		deleted_at = 0
	`

	row := u.db.QueryRow(query, req.Id)

	err := row.Scan(
		&id,
		&username,
		&email,
		&role,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}

	if err != nil {
		log.Println("Error while getting user: ", err)
		return nil, err
	}

	token := t.GenerateJWTToken(id, email, username, role)
	tokens := ap.TokenRes{
		Token: token.RefreshToken,
		ExpAt: "24 hours",
	}

	return &tokens, nil
}

func (u *UsersRepo) ForgotPassword(req *ap.UsersForgotPassword) (*ap.Void, error) {
	code, err := verification.GenerateRandomCode(6)
	if err != nil {
		return nil, errors.New("failed to generate code for verification: " + err.Error())
	}

	u.rdb.Set(context.Background(), req.Email, code, time.Minute*5)

	from := "axadjonovsardorbeck@gmail.com"
	password := "ypuw yybh sqjr boww"
	err = verification.SendVerificationCode(verification.Params{
		From:     from,
		Password: password,
		To:       req.Email,
		Message:  fmt.Sprintf("Hi %s, your verification code is: %s", req.Email, code),
		Code:     code,
	})

	if err != nil {
		return nil, errors.New("failed to send verification email: " + err.Error())
	}
	return nil, nil
}

func (u *UsersRepo) ResetPassword(req *ap.UsersResetPassword) (*ap.Void, error) {
	em, err := u.rdb.Get(context.Background(), req.Email).Result()

	log.Println(em)

	if err != nil {
		return nil, errors.New("invalid code or code expired")
	}

	if em != req.ResetToken {
		return nil, errors.New("invalid code or code expired")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to reset password")
	}
	req.NewPassword = string(hashedPassword)

	query := `update users set password = $1 where email = $2 and deleted_at = 0`
	_, err = u.db.Exec(query, req.NewPassword, req.Email)
	log.Println(req.NewPassword)
	if err != nil {
		return nil, fmt.Errorf("failed to reset password: %v", err)
	}
	return nil, nil
}
func (u *UsersRepo) ChangePassword(req *ap.UsersChangePassword) (*ap.Void, error) {
	var curPass string

	queryCurrent := `SELECT password FROM users WHERE id = $1 AND deleted_at = 0`

	row := u.db.QueryRow(queryCurrent, req.Id)

	err := row.Scan(&curPass)

	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(curPass), []byte(req.CurrentPassword)); err != nil {
		return nil, errors.New("invalid current password")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to change password")
	}

	queryUpdate := `UPDATE users SET password = $1 WHERE id = $2 AND deleted_at = 0`

	_, err = u.db.Exec(queryUpdate, string(hashedPassword), req.Id)

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (u *UsersRepo) CheckEmail(req *ap.CheckEmailReq) (*ap.ById, error) {
	query := `
	SELECT 
		id
	FROM
		users
	WHERE 
		email = $1
	AND 
		deleted_at = 0
	`

	row := u.db.QueryRow(query, req.Email)

	var user_id string

	err := row.Scan(
		&user_id,
	)

	if err == sql.ErrNoRows {
		log.Println("Cart is empty")
		return nil, errors.New("cart is empty")
	}

	if err != nil {
		log.Println("Error while getting cart id: ", err)
		return nil, err
	}

	log.Println("Successfully got cart id")

	return &ap.ById{
		Id: user_id,
	}, nil
}

func (u *UsersRepo) GetAllUsers(req *ap.GetAllUsersReq) (*ap.GetAllUsersRes, error) {
	users := ap.GetAllUsersRes{}

	query := `
	SELECT 
		id,
		first_name,
		last_name,
		role,
		phone
	FROM 	
		users
	WHERE
		deleted_at = 0
	`
	var args []interface{}
	var conditions []string

	if req.Role != "" && req.Role != "string" {
		conditions = append(conditions, " role = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Role)
	}

	if len(conditions) > 0 {
		query += " AND " + strings.Join(conditions, " AND ")
	}

	var limit int32
	var offset int32

	limit = 10
	offset = (req.Page - 1) * limit

	args = append(args, limit, offset)
	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", len(args)-1, len(args))

	rows, err := u.db.Query(query, args...)

	if err == sql.ErrNoRows {
		log.Println("Users not found")
		return nil, errors.New("users not found")
	}

	if err != nil {
		log.Println("Error while retriving users: ", err)
		return nil, err
	}

	for rows.Next() {
		user := ap.UserRes{}

		err := rows.Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.Role,
			&user.Phone,
		)

		if err != nil {
			log.Println("Error while scanning all users: ", err)
			return nil, err
		}

		users.Users = append(users.Users, &user)
	}

	log.Println("Successfully fetched all orders")

	return &users, nil
}