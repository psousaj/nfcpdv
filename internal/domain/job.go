package domain

import "github.com/uptrace/bun"

type JobType string

const (
	JobTypeScanPDV       JobType = "scan_pdv"
	JobTypeMigrateRename JobType = "migrate_rename"
	JobTypeMigrateUndo   JobType = "migrate_undo_rename"
	JobTypeMySqlExec     JobType = "mysql_exec"
)

type JobStatus string

const (
	JobStatusPending   JobStatus = "pending"
	JobStatusRunning   JobStatus = "running"
	JobStatusCompleted JobStatus = "completed"
	JobStatusFailed    JobStatus = "failed"
)

type Job struct {
	bun.BaseModel `bun:"table:jobs"`

	ID      int64     `bun:"id,pk,autoincrement" json:"id"`
	JobType JobType   `bun:"notnull" json:"job_type"`
	Status  JobStatus `bun:"notnull, default:pending" json:"status"`
}
