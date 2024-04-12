package internal

import (
	"fmt"
)

func IntIdtoUint(num int) (uint, *Error) {
	if num >= 0 {
		return uint(num), nil
	}

	return 0, DetailError(400, fmt.Errorf("id conversion to unsigned integer, negative id"))
}

/*
func GetAccessTokenFromCtx(c *gin.Context) (*jwt.Token, *Claims) {
	token, _ := c.Get(CAccessToken)
	claims, _ := c.Get(CAccessClaims)

	return token.(*jwt.Token), claims.(*Claims)
}
*/
