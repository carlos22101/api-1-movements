package repositories

import (
    "database/sql"
    "recibe/src/domain/entities"
    "recibe/src/domain/repositories"
)

type MovementMySQLRepo struct {
    DB *sql.DB
}

func NewMovementMySQLRepo(db *sql.DB) repositories.MovementRepository {
    return &MovementMySQLRepo{DB: db}
}

func (repo *MovementMySQLRepo) Create(m *entities.Movement) error {
    query := "INSERT INTO movements (sensor_id, timestamp, motion) VALUES (?, ?, ?)"
    res, err := repo.DB.Exec(query, m.SensorID, m.Timestamp, m.Motion)
    if err != nil {
        return err
    }
    id, err := res.LastInsertId()
    if err != nil {
        return err
    }
    m.ID = int(id)
    return nil
}

func (repo *MovementMySQLRepo) GetMovements() ([]entities.Movement, error) {
    rows, err := repo.DB.Query("SELECT id, sensor_id, timestamp, motion FROM movements")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var movements []entities.Movement
    for rows.Next() {
        var movement entities.Movement
        if err := rows.Scan(&movement.ID, &movement.SensorID, &movement.Timestamp, &movement.Motion); err != nil {
            return nil, err
        }
        movements = append(movements, movement)
    }
    return movements, nil
}