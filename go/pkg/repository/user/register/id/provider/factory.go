package provider

import (
	"fmt"

	domain "github.com/makocchi-git/ddd-guys/go/pkg/domain/user"
)

var valid = []string{
	"uuid",
	"random",
}

func CreateIdProvider(selector string) (domain.IIdProvider, error) {
	if err := validSelector(selector); err != nil {
		return nil, err
	}
	if selector == "random" {
		return NewRandomStringIDProvider(32), nil
	}

	// デフォルトだと uuid にしているので、なんとなく uuid をこちらに配置
	return NewUUIDIDProvider(), nil
}

// validSelector は他でも使いまわしているので、 util 化を検討かな
func validSelector(selector string) error {
	for _, v := range valid {
		if v == selector {
			return nil
		}
	}
	return fmt.Errorf("given selector isn't supported: %s", selector)
}
