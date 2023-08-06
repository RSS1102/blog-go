package until

import "strings"

// ConvertToSnakeCase 将大写字段更新为下划线类型
func ConvertToSnakeCase(s string) string {
	var result strings.Builder
	for i, char := range s {
		if char >= 'A' && char <= 'Z' {
			if i > 0 {
				if s != "ID" {
					result.WriteRune('_')
				}
			}
			result.WriteRune(char + 32) // Convert to lowercase
		} else {
			result.WriteRune(char)
		}
	}
	return result.String()
}
