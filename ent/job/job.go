// Code generated by entc, DO NOT EDIT.

package job

const (
	// Label holds the string label denoting the job type in the database.
	Label = "job"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"

	// EdgeSkill holds the string denoting the skill edge name in mutations.
	EdgeSkill = "skill"

	// Table holds the table name of the job in the database.
	Table = "jobs"
	// SkillTable is the table the holds the skill relation/edge.
	SkillTable = "jobs"
	// SkillInverseTable is the table name for the Skill entity.
	// It exists in this package in order to avoid circular dependency with the "skill" package.
	SkillInverseTable = "skills"
	// SkillColumn is the table column denoting the skill relation/edge.
	SkillColumn = "job_skill"
)

// Columns holds all SQL columns for job fields.
var Columns = []string{
	FieldID,
	FieldTitle,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Job type.
var ForeignKeys = []string{
	"job_skill",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}
