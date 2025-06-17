package handler

import "github.com/gin-gonic/gin"

func GetHoldings(c *gin.Context) {
	c.JSON(200, gin.H{"data": []gin.H{
		{"stock": "TCS", "qty": 10, "avg_price": 3200},
	}})
}

func GetOrderbook(c *gin.Context) {
	c.JSON(200, gin.H{"orders": []gin.H{
		{"id": 1, "stock": "INFY", "pnl": "Unrealized"},
	}})
}

func GetPositions(c *gin.Context) {
	c.JSON(200, gin.H{"positions": []gin.H{
		{"id": 1, "stock": "RELIANCE", "pnl": "Realized"},
	}})
}
