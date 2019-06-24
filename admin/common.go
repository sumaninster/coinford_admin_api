package admin

import (
	"coinford_admin_api/admin_models"
	"coinford_admin_api/admin_configs"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
	"errors"
)

func apiAuthenticateAdmin(tokenString string) (CommonResponse, *admin_models.Admin, *admin_models.AdminGroup, bool, error) {
	err := parseAdminToken(tokenString)
	if err == nil {
		authToken := admin_models.AdminToken{Token: tokenString}
		err = authToken.Read("Token")
		if err == nil {
			if time.Now().Before(authToken.ExpirationTime) {
				admin := admin_models.Admin{Id: authToken.AdminId}
				err = admin.Read("id")
				if err == nil {
					if admin.DeletedAt == *admin_configs.NullTime {
						adminGroup := admin_models.AdminGroup{Id: admin.AdminGroupId}
						err = adminGroup.Read()
						if err == nil {
							return CommonResponse{}, &admin, &adminGroup, true, nil
						} else {
							return CommonResponse{ResponseCode: 403, ResponseDescription: "Admin group not found."}, nil,  nil, false, errors.New("Admin group not found.")
						}
					} else {
						debugMessage("Authenticate Error 2: ", errors.New("Admin deleted"))
						return CommonResponse{ResponseCode: 403, ResponseDescription: "Admin deleted."}, nil, nil, false, errors.New("Admin deleted.")
					}
				} else {
					debugMessage("Authenticate Error 3: ", err)
					return CommonResponse{ResponseCode: 403, ResponseDescription: "Unable to read or write data. Please relogin."}, nil, nil, false, err
				}
			} else {
				debugMessage("Authenticate Error 4: ", errors.New("Token expired"))
				return CommonResponse{ResponseCode: 403, ResponseDescription: "Token expired. Please relogin."}, nil, nil, false, errors.New("Token expired.")
			} 
		} else {
			debugMessage("Authenticate Error 5: ", err)
			return CommonResponse{ResponseCode: 403, ResponseDescription: "Authentication failed. Please relogin."}, nil, nil, false, err
		} 
	} else {
		debugMessage("Authenticate Error 6: ", err)
		return CommonResponse{ResponseCode: 403, ResponseDescription: "Invalid token. Please relogin."}, nil, nil, false, err
	}
	return CommonResponse{ResponseCode: 403, ResponseDescription: "Unrecognized error."}, nil, nil, false, errors.New("Unrecognized error.")
}
/*
func apiAdminAuthenticate(tokenString string) (CommonResponse, *admin_models.Admin, bool, bool, error) {
	jres, admin, isLogin, _ := apiAuthenticate(rqd.Token)
	if isLogin {

	}
}*/

func saveAdminToken(admin *admin_models.Admin) (CommonResponse, string, error){
	expirationTime := time.Now().Add(time.Hour * time.Duration(admin_configs.PostLoginTokenTime))
	tokenString, _ := tokenAdmin(expirationTime.Unix())
	authToken := admin_models.AdminToken{AdminId: admin.Id, Token: tokenString, ExpirationTime: expirationTime}
	_, _, err := authToken.ReadOrCreate("admin_id")
	if err == nil {
		authToken.Token = tokenString
		authToken.ExpirationTime = expirationTime
		authToken.UpdatedAt = time.Now()
		authToken.DeletedAt = *admin_configs.NullTime
		err = authToken.Update()
		if err == nil {
			return CommonResponse{}, tokenString, nil		
		} else {
			debugMessage("saveAdminToken Error 1: ", err)
			return CommonResponse{ResponseCode: 403, ResponseDescription: "Unable to read or write data. Please relogin."}, *admin_configs.NullString, err
		}
	} else {
		debugMessage("saveAdminToken Error 2: ", err)
		return CommonResponse{ResponseCode: 403, ResponseDescription: "Unable to read or write data. Please relogin."}, *admin_configs.NullString, err
	}
	return CommonResponse{ResponseCode: 403, ResponseDescription: "Unrecognized error"}, *admin_configs.NullString, errors.New("Unrecognized error")
}

func tokenAdmin(expirationTime int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS512, 
    jwt.MapClaims{
        "exp": expirationTime,
        "iat": time.Now().Unix() })
	//fmt.Println(admin_configs.SignBytes)
    privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(admin_configs.SignBytes)
	debugMessage("ParseRSAPrivateKeyFromPEM Error: ", err)
	if err != nil {return "Invalid Token", err}

    tokenString, err := token.SignedString(privateKey)
	debugMessage("SignedString Error: ", err)
	if err != nil {return "Invalid Token", err}

	return tokenString, nil
}

func parseAdminToken(tokenString string) error {
	//fmt.Println(tokenString)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        publicKey, err := jwt.ParseRSAPublicKeyFromPEM(admin_configs.VerifyBytes)
        return publicKey, err
    })

    if err == nil && token.Valid {
        debugMessage("ParseToken Success: Your token is valid.", nil)
        return nil
    }
    debugMessage("ParseToken Error: ", err)
    return err
}

func debugMessage(tag string, err error) {
	if err != nil && admin_models.Runmode == "dev" {
		fmt.Println(tag, err)
	}
}