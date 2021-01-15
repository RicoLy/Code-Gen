package entity

import (
	"fmt"
	"testing"
)

func TestAdmins_Delete(t *testing.T) {
	r := &Admins{}
	r.Db = MasterDB

	err := r.Delete("id = ?", 321231321)
	fmt.Println(err)

}
