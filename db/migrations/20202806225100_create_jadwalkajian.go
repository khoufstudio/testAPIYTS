package migrations

import (
	"github.com/go-rel/rel"
)

//MigrateCreateJadwalKajian Migrate Table Jadwal Kajian
func MigrateCreateJadwalKajian(schema *rel.Schema) {
	schema.CreateTable("jadwal_kajian", func(t *rel.Table) {
		t.ID("id")
		t.String("title")
		t.String("lecturer")
		t.DateTime("start_date")
		t.DateTime("end_date")
		t.Date("event_date")
		t.DateTime("created_at")
		t.DateTime("updated_at")
	})
}

//RollbackCreateJadwalKajian Drop Table Jadwal Kajian
func RollbackCreateJadwalKajian(schema *rel.Schema) {
	schema.DropTable("jadwal_kajian")
}
