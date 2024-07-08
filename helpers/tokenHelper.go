package helper

import (
	jwt "github.com/dgrijalva/jwt-go"
)

type SignedDetails struct {
	Email      string
	First_name string
	Last_name  string
	Uid        string
	User_type  string
	jwt.StandardClaims
}

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

var SECRET_KEY string = os.Getenv("SECRET_KEY")

func GenerateAllTokens(email string, first_name string, last_name string, uid string, user_type string) (string, string, error) {
	claims := &SignedDetails{
		Email:      email,
		First_name: first_name,
		Last_name:  last_name,
		Uid:        uid,
		User_type:  user_type,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(1)).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", "", err
	}
	refreshToken := uuid.New().String()
	saveErr := CreateAuth(uid, refreshToken)
	if saveErr != nil {
		return "", "", saveErr
	}
	return accessToken, refreshToken, nil
}
