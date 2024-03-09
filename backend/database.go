package main

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/forms"
	"github.com/pocketbase/pocketbase/models"
)

func GetCount(app *pocketbase.PocketBase) (int, error) {
	record, err := app.Dao().FindRecordById("Counter", "loh30i7ry1384ep")
	if err != nil {
		return -1, echo.NewHTTPError(http.StatusInternalServerError, "Could not access collection 'Counter'")
	}

	value := record.GetInt("Value")
	return value, nil
}

func UpdateCount(app *pocketbase.PocketBase, newVal int) error {
	record, err := app.Dao().FindRecordById("Counter", "loh30i7ry1384ep")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Could not access collection 'Counter'")
	}

	record.Set("Value", newVal)
	if err := app.Dao().SaveRecord(record); err != nil {
		return err
	}
	return nil
}

func AddUserRecord(app *pocketbase.PocketBase, newUser User) (*models.Record, error) {
	users, err := app.Dao().FindCollectionByNameOrId("users")
	if err != nil {
		return nil, err
	}
	newUserRecord := models.NewRecord(users)
	form := forms.NewRecordUpsert(app, newUserRecord)
	form.LoadData(
		map[string]any{
			"username":      newUser.Username,
			"email":         newUser.Email,
			"password_hash": newUser.PasswordHash,
			"password_salt": newUser.PasswordSalt,
		},
	)
	if err := form.Submit(); err != nil {
		return nil, err
	}
	return newUserRecord, nil
}

func GetUserRecord(app *pocketbase.PocketBase, userId string) (*models.Record, error) {
	record, err := app.Dao().FindRecordById("users", userId)
	if err != nil {
		return nil, err
	}
	return record, nil
}

func UpdateUserRecord(
	app *pocketbase.PocketBase,
	record *models.Record,
	updates map[string]any) (*models.Record, error) {

	form := forms.NewRecordUpsert(app, record)
	form.LoadData(updates)
	if err := form.Submit(); err != nil {
		return nil, err
	}
	return record, nil
}

func DeleteUserRecord(app *pocketbase.PocketBase, record *models.Record) error {
	if err := app.Dao().DeleteRecord(record); err != nil {
		return err
	}
	return nil
}
