package attendanceRepository

import (
	"database/sql"
	"fmt"
	"sirclo/project-capstone/entities/attendanceEntities"
)

type attendanceRepo struct {
	db *sql.DB
}

func NewMySQLDayRepository(db *sql.DB) AttendanceRepoInterface {
	return &attendanceRepo{
		db: db,
	}
}

func (ar *attendanceRepo) GetAttendances() ([]attendanceEntities.Attendance, error) {
	var attendances []attendanceEntities.Attendance

	result, err := ar.db.Query(`
	SELECT
		attendances.id, day.date AS date, office.name, user.name, attendances.status, (COALESCE(NULLIF(attendances.notes,''), '-')) AS notes, (COALESCE(NULLIF(admin.name,''), '-')) AS admin 
	FROM 
		attendances
	LEFT JOIN
		days AS day ON day.id = attendances.day_id
	LEFT JOIN
		offices AS office ON office.id = day.office_id
	LEFT JOIN
		users AS user ON user.id = attendances.user_id
	LEFT JOIN
		users AS admin ON admin.id = attendances.admin_id`)
	if err != nil {
		return attendances, err
	}

	for result.Next() {
		var attendance attendanceEntities.Attendance

		errScan := result.Scan(&attendance.ID, &attendance.Day.Date, &attendance.Office, &attendance.Employee.Name, &attendance.Status, &attendance.Notes, &attendance.Admin.Name)

		if errScan != nil {
			return attendances, errScan
		}
		attendances = append(attendances, attendance)
	}
	fmt.Println("att repo: ", attendances)
	return attendances, nil
}

func (ar *attendanceRepo) CreateAttendance(att attendanceEntities.Attendance) (attendanceEntities.Attendance, error) {
	query := `INSERT INTO attendances (id, day_id, user_id, created_at) VALUES (?, ?, ?, now())`

	statement, err := ar.db.Prepare(query)
	if err != nil {
		return att, err
	}

	_, errExec := statement.Exec(att.ID, att.Day.ID, att.Employee.ID)
	if errExec != nil {
		return att, errExec
	}
	return att, nil
}

func (ar *attendanceRepo) UpdateAttendance(att attendanceEntities.Attendance) (attendanceEntities.Attendance, error) {
	query := `UPDATE attendances SET status = ?, notes = ?, admin_id = ?, updated_at = now() WHERE id = ?`

	statement, err := ar.db.Prepare(query)
	if err != nil {
		return att, err
	}

	_, errExec := statement.Exec(att.Status, att.Notes, att.Admin.ID, att.ID)
	if errExec != nil {
		return att, errExec
	}
	return att, nil
}
