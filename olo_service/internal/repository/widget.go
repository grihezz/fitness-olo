package repository

import (
	"OLO-backend/olo_service/internal/entity"
	"OLO-backend/olo_service/internal/repository/provider"
	"fmt"
)

type WidgetRepo struct {
	mysqlProvider *provider.MySQLProvider
}

func NewWidgetRepo(mysqlProvider *provider.MySQLProvider) *WidgetRepo {
	return &WidgetRepo{mysqlProvider: mysqlProvider}
}

func (r *WidgetRepo) GetAllWidgets() ([]entity.Widget, error) {
	return r.getWidgets(`SELECT * FROM widgets`)
}

func (r *WidgetRepo) GetUserWidgets(userId int64) ([]entity.Widget, error) {
	return r.getWidgets(fmt.Sprintf("SELECT w.id as id, w.description as `description` FROM widgets as w, user_has_widget as u WHERE w.id = u.id_widget AND u.id_user = '%d'", userId))
}

func (r *WidgetRepo) getWidgets(widgetQuery string) ([]entity.Widget, error) {
	driver, err := r.mysqlProvider.Driver()
	if err != nil {
		return nil, err
	}

	rows, err := driver.Queryx(widgetQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var widgets []entity.Widget
	for rows.Next() {
		var widget entity.Widget
		if err := rows.StructScan(&widget); err != nil {
			return nil, err
		}
		widgets = append(widgets, widget)
	}
	return widgets, nil
}

func (r *WidgetRepo) AddWidgetForUser(widgetId, userId int64) error {
	driver, err := r.mysqlProvider.Driver()
	if err != nil {
		return err
	}
	_, err = driver.NamedExec("INSERT INTO `user_has_widget` (`id_widget`, `id_user`) VALUES (:id_widget, :id_user)", map[string]interface{}{
		"id_widget": widgetId,
		"id_user":   userId,
	})
	return err
}
func (r *WidgetRepo) DeleteWidgetForUser(widgetId int64, userId int64) error {
	driver, err := r.mysqlProvider.Driver()
	if err != nil {
		return err
	}
	_, err = driver.NamedExec("DELETE FROM `user_has_widget` WHERE `id_widget` = :id_widget AND `id_user` = :id_user", map[string]interface{}{
		"id_widget": widgetId,
		"id_user":   userId,
	})
	return err
}
