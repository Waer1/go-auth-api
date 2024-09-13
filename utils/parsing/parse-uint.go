package parsing

import (
	"api-auth/utils"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetParamUint extracts a uint parameter from the gin.Context
func GetParamUint(c *gin.Context, param string) (uint, error) {
	paramStr := c.Param(param)
	paramUint, err := strconv.ParseUint(paramStr, 10, 32)
	if err != nil {
		c.Error(utils.NewServiceErr(400, map[string]string{param: fmt.Sprintf("Invalid %s", param)}))
		return 0, utils.NewServiceErr(400, map[string]string{param: fmt.Sprintf("Invalid %s", param)})
	}
	return uint(paramUint), nil
}
