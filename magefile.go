// +build mage

package go_minecraft

import (
	"go-minecraft/internal"
	"os"
)

// Clean up after yourself
func Generate() error {
	//f, err := os.Create("items.go")
	//if err != nil {
	//	return err
	//}

	items := internal.GetItems()
	if err := items.GenerateCode(os.Stdout); err != nil {
		return err
	}

	return nil
}
