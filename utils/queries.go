package utils

const (
	// PRODUCT QUERIES
	GET_ALL_PRODUCT   = "SELECT p.* FROM m_product p join m_category c on p.product_category = c.category_id WHERE p.status = 'A' ORDER BY product_name"
	GET_PRODUCT_BY_ID = "SELECT * FROM m_product WHERE status = 'A' and product_id = ? ORDER BY product_name"
	ADD_PRODUCT       = "INSERT INTO m_product(product_id, product_code, product_name, product_category, product_price) VALUES(?, ?, ?, ?, ?)"
	UPDATE_PRODUCT    = "UPDATE m_product SET product_code = ?, product_name = ?, product_category = ?, product_price = ? where product_id = ?"
	DELETE_PRODUCT    = "UPDATE m_product SET status = 'I' where product_id = ?"

	// CATEGORY QUERIES
	GET_ALL_CATEGORY   = "SELECT * FROM m_category WHERE status = 'A' ORDER BY category_name"
	GET_CATEGORY_BY_ID = "SELECT * FROM m_category WHERE status = 'A' and category_id = ? ORDER BY category_name"
	ADD_CATEGORY       = "INSERT INTO m_category(category_id, category_name) VALUES(?, ?)"
	UPDATE_CATEGORY    = "UPDATE m_category SET category_name = ? where category_id = ?"
	DELETE_CATEGORY    = "UPDATE m_category SET status = 'I' where category_id = ?"

	// TRANSACTION QUERIES
	GET_ALL_TRANSACTION   = "SELECT * FROM m_transaction ORDER BY transaction_date"
	GET_TRANSACTION_BY_ID = "SELECT * FROM m_transaction WHERE id = ? ORDER BY transaction_date"
	ADD_TRANSACTION       = "INSERT INTO m_transaction(id, transaction_date, product_id, qty) VALUES(?, ?, ?, ?)"
	UPDATE_TRANSACTION    = "UPDATE m_transaction SET product_id= ?, qty = ? where id = ?"
	DELETE_TRANSACTION    = "DELETE from m_transaction where id = ?"

	// REPORT QUERIES
	GET_DAILY_REPORT   = "SELECT * FROM v_report where year(transaction_date) =? and month(transaction_date) =? and day(transaction_date) =? "
	GET_MONTHLY_REPORT = "select * from v_report where month(transaction_date) = ?"
)
