package utils

import (
	"fmt"
)

func ContainsStr(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func RemoveStr(slice []string, s string) (result []string) {
	for _, item := range slice {
		if item == s {
			continue
		}
		result = append(result, item)
	}
	return
}

func GetStatefulSetReplicaName(name string, replicaIdx int) string {
	return fmt.Sprintf("%s-%d", name, replicaIdx)
}
