package originstamp

import (
	"fmt"
	"strings"
)

type NotificationType int8

const (
	EMAIL NotificationType = iota
	WEBHOOK
)

func NotificationTypeValues() []NotificationType {
	return []NotificationType{
		EMAIL,
		WEBHOOK,
	}
}

func NotificationTypeStrValues() []string {
	vals := NotificationTypeValues()
	strs := make([]string, len(vals))
	for i, notificationType := range vals {
		strs[i] = notificationType.String()
	}
	return strs
}

func (c NotificationType) String() string {
	switch c {
	case EMAIL:
		return "EMAIL"
	case WEBHOOK:
		return "WEBHOOK"
	default:
		return fmt.Sprintf("NotificationType(%d)", c)
	}
}

func NotificationTypeFromString(str string) (NotificationType, error) {
	switch strings.ToUpper(str) {
	case "EMAIL":
		return EMAIL, nil
	case "WEBHOOK":
		return WEBHOOK, nil
	default:
		return 0, fmt.Errorf("%s is not a valid notification type", str)
	}
}
