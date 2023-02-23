package test

import "database/sql"

func recordStats(db *sql.DB, userID, productID int64) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return
	}
	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()
	if _, err = tx.Exec("update products set views = views + 1"); err != nil {
		return
	}
	if _, err = tx.Exec("insert into product_viewers(user_id, product_id) values(?, ?)", userID, productID); err != nil {
		return
	}
	return
}
