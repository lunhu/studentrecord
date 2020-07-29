package util

import (
	"time"
	"net/http"
	"strings"
	"strconv"
	"math"
	"fmt"

	"studentrecord/common"

	"github.com/gbrlsnchs/jwt/v3"
)


type Auth struct {
}

func NewAuth() *Auth {
	return &Auth{
	}
}

func (auth *Auth)  Middleware( next http.Handler)  http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		azt, ok := r.Header["Authorization"];
		
		if !ok {
	
			auth.handError(w, "400", "")

			return

		}


		authValue := strings.Split(azt[0], " ")

		if len(authValue) != 2 {

			auth.handError(w, "400", "")

			return

		} 

		if authValue[0] != "Bearer"  {

			auth.handError(w, "400", "")
	
			return
		} 

		// Generate a secret uesed in JWT

		p5 := int64(time.Now().Year())
		p3 := int64(time.Now().Month())
		p7 := int64(time.Now().Day())
		p1 := int64(time.Now().Weekday())

		if p1 == 0 {
			p1 = 7
		}
		
		ymdInt, err := strconv.Atoi(time.Now().Format("20060102"))

		if err != nil {
		 
			auth.handError(w, "500", "")

			return
		}

		p2 := int64(ymdInt) * p1
		
		p4 := int64(math.Pow(float64((p1+1) * 2), float64(p1 + 1)))

		p6 := (p5 + p1) * (p3 + p7)

		secret := fmt.Sprintf("%02v%v%02v%v%04v%v%02v",p1, p2, p3, p4, p5, p6, p7)
		//panic(secret)
		//token := []byte("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.qdqBx1DjYrwEt_sU_XwBaWqKEPqUwXaRKmZH-_nykz0")

		var hs = jwt.NewHS256([]byte(secret))
		// ...

		var pl jwt.Payload

		_, err = jwt.Verify([]byte(authValue[1]), hs, &pl)

		if err != nil {

			auth.handError(w, "500", "")

			return

		}

    	// Call the next handler, which can be another middleware in the chain, or the final handler.
   		 next.ServeHTTP(w, r)	

	})
		
}


func (auth *Auth) handError(w http.ResponseWriter, entry, catalog string) {

	formatter := common.GetFormatter()

	cjError := &common.CjError{}

	var stCode int

	switch entry {
	case "400":
		cjError.Title = "The request cannot be completed"
		cjError.Code = "Bad Request Code x400"
		cjError.Message = "The request was rejected by the server. Check the parameters, please!"
		stCode = http.StatusBadRequest
	case "500":
		cjError.Title = "Unable to search item"
		cjError.Code = "Internal Error Code x500"
		cjError.Message = "The server encountered an internal error or misconfiguration and was unable to complete your request"
		stCode = http.StatusInternalServerError
	default:
		cjError.Title = "An unknown error occurred"
		cjError.Code = "Internal Error Code x000"
		cjError.Message = "The server encountered an unknown error or misconfiguration and was unable to complete your request"
		stCode = http.StatusInternalServerError
	}

	cjCollection :=&common.CjCollection{}
	
	cjCollection.Version = "1.0"

	cjCollection.Error = cjError

	cjResponse := &common.CjResponse{}

	cjResponse.Collection = cjCollection

	formatter.JSON(w, stCode, cjResponse)

	return

}