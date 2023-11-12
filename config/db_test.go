package config

import (
	"fmt"
	"testing"
)

func TestCreateMariaConn(t *testing.T) {
	conn := CreateMariaGormConnection(Stringmaria)
	fmt.Printf("%+v", conn)
}
