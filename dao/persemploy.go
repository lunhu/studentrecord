package dao

import (
	"database/sql"
	"github.com/jmoiron/sqlx"

	"studentrecord/model"
)

func SelectMaxSeq(sqlxDB *sqlx.DB, pers_id string) (int32, error) {
	var maxSeq sql.NullInt32
	err := sqlxDB.Get(&maxSeq, "SELECT MAX(eff_seq) FROM e_employs_tbl WHERE pers_id = ?", pers_id)
	if err != nil {
		return 0, err
	}

	if maxSeq.Valid == false {
		return 0, nil
	}

	return maxSeq.Int32, nil
}

// InsertPersEmail 新增一条人员邮箱
func InsertPersEmploy(tx *sqlx.Tx, params *model.PersEmploy) (*model.PersEmploy, error) {

	_, err := tx.Exec(
		`insert into e_employs_tbl 
		(pers_id, eff_dt, eff_seq, employer, empl_status, hr_status, start_dt, end_dt, depart_id, position, reg_temp, full_part_time, job_indicator, job_number, work_location, depart_out, staff_indicator, u_id) 
		values 
		(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`,
		params.PersID,
		params.EffDt,
		params.EffSeq,
		params.Employer,
		params.EmplStatus,
		params.HrStatus,
		params.StartDt,
		params.EndDt,
		params.DepartID,
		params.Position,
		params.RegTemp,
		params.FullPartTime,
		params.JobIndicator,
		params.JobNumber,
		params.WorkLocation,
		params.DepartOut,
		params.StaffIndicator,
		params.UID,
	)
	if err != nil {
		return nil, err
	}

	return params, nil
}