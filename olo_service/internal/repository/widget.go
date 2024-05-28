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

func (r *WidgetRepo) GetWidgets(userId int64) ([]entity.Widget, error) {
	return r.getWidgets(fmt.Sprintf("SELECT id, data FROM widgetsUser WHERE id_user = %d", userId))
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

func (r *WidgetRepo) UpdateWidget(data string, widgetId, userId int64) error {
	driver, err := r.mysqlProvider.Driver()
	if err != nil {
		return err
	}

	_, err = driver.NamedExec("UPDATE `widgetsUser` SET `data`=:data WHERE `id`=:id_widget AND `id_user`=:id_user", map[string]interface{}{
		"data":      data,
		"id_widget": widgetId,
		"id_user":   userId,
	})
	return err
}

func (r *WidgetRepo) AddWidget(data string, userId int64) (int64, error) {
	driver, err := r.mysqlProvider.Driver()
	if err != nil {
		return 0, err
	}
	res, err := driver.NamedExec("INSERT INTO `widgetsUser` (`data`, `id_user`) VALUES (:data, :id_user)", map[string]interface{}{
		"data":    data,
		"id_user": userId,
	})
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	return id, err
}

func (r *WidgetRepo) DeleteWidget(widgetId int64, userId int64) error {
	driver, err := r.mysqlProvider.Driver()
	if err != nil {
		return err
	}
	_, err = driver.NamedExec("DELETE FROM `widgetsUser` WHERE `id` = :id_widget AND `id_user` = :id_user", map[string]interface{}{
		"id_widget": widgetId,
		"id_user":   userId,
	})
	return err
}
