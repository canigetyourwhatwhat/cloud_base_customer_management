package database

import (
	"database/sql"
	"erply/entity"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"time"
)

func InsertCustomer(db *sqlx.DB, input map[string]string) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("insert into customers(customer_id, first_name, last_name, company_name, email) values (?, ?, ?, ?, ?)",
		input["customerID"],
		input["firstName"],
		input["lastName"],
		input["companyName"],
		input["email"])
	if err != nil {
		if err = tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	if err = tx.Commit(); err != nil {
		if err = tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	return nil
}

func GetCustomerByCustomerID(db *sqlx.DB, id string) (*entity.Customer, error) {
	var customer entity.Customer
	err := db.Get(&customer, `select * from customers where customer_id = ?`, &id)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, entity.ErrCustomerNotFound
	}
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func UpdateCustomerByCustomerID(db *sqlx.DB, input map[string]string) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("update customers set first_name = ?,  last_name = ?,  company_name = ?, email = ?, updated_at = ? where customer_id = ?",
		input["firstName"],
		input["lastName"],
		input["companyName"],
		input["email"],
		time.Now(),
		input["customerID"])
	if err != nil {
		if err = tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	if err = tx.Commit(); err != nil {
		if err = tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	return nil
}

func DeleteCustomerByCustomerID(db *sqlx.DB, customerID string) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("delete from customers where customer_id = ? ", customerID)
	if err != nil {
		if err = tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	if err = tx.Commit(); err != nil {
		if err = tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	return nil
}
