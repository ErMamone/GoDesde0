package usuarios

import (
	"fmt"
	"github.com/ErMamone/GoDesde0/modelos"
	"time"
)

func AltaUsuario() {
	u := new(modelos.User)

	u.AddUser(1, "Juan Carlos", time.Now(), true)

	fmt.Println(u)
	fmt.Printf("User:\n{\n\tId: %d \n\tName: %s \n\tCreatedAt: %s \n\tStatus: %t\n}", u.Id, u.Name, u.CreatedAt, u.Status)
}
