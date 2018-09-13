package middleware

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

func TransactionMiddleware(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(c echo.Context) error {
			tx := db.Begin()
			c.Set("DB", tx)
			if err := next(c); err != nil {
				tx.Rollback()
				return err
			}
			tx.Commit()
			return nil
		})
	}
}
