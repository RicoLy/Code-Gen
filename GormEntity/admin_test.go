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

func TestAdmins_Count(t *testing.T) {
	r := &Admins{}
	r.Db = MasterDB
	count, err := r.Count("id > ? and user_name like ?", 2121231321, "王%")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(count)
}