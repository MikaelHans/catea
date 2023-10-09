package signup

import (
	_ "fmt"
	"net/http"

	"github.com/MikaelHans/catea/user/pkg/structs"
	"github.com/MikaelHans/catea/user/pkg/util"
	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	var data structs.Member
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//ENCRYPT MEMBER PASSWORD//////////////////////////////////////////////////
	encrypted_pass, err := util.EncryptString(data.Pass)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	data.Pass = encrypted_pass
	//INSERT MEMBER TO DATABASE//////////////////////////////////////////////////
	rows, err := InsertIntoMember(data)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}
	var result bool
	rows.Next()
	err = rows.Scan(&result)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}
	if result == false {
		c.JSON(400, gin.H{
			"msg": "Email already used",
		})
		return
	}
	c.JSON(200, gin.H{
		"msg": "sucess",
	})
	return
}

// func add_member(newmember Member) (*sql.Rows, error) {
// 	rows, err := InsertIntoMember(newmember)
// 	return rows, err
// }
