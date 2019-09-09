package utils

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	datahub_v1alpha1 "github.com/containers-ai/api/alameda_api/v1alpha1/datahub"
	"github.com/spf13/viper"
)

func RemoveEmptyStr(strList []string) []string {
	ret := []string{}
	for _, str := range strList {
		if str != "" {
			ret = append(ret, str)
		}
	}
	return ret
}

func EventEmailSubject(evt *datahub_v1alpha1.Event) string {
	msg := evt.GetMessage()
	levelMap := viper.GetStringMap("eventLevel")
	level := evt.GetLevel()
	return fmt.Sprintf("Federator.ai Notification: %s - %s",
		strings.Title(levelMap[strconv.FormatInt(int64(level), 10)].(string)), msg)
}

func EventHTMLMsg(evt *datahub_v1alpha1.Event) string {
	levelMap := viper.GetStringMap("eventLevel")
	eventMap := viper.GetStringMap("eventType")
	return fmt.Sprintf(`
	<html>
		<body>
			Federator.ai Event Notification<br>
			###############################################################<br>
			<table cellspacing="5" cellpadding="0">
				<tr><td align="right">Time:</td><td>%s</td></tr>
				<tr><td align="right">Level:</td><td>%s</td></tr>
				<tr><td align="right">Message:</td><td>%s</td></tr>
				<tr><td align="right">Event Type:</td><td>%s</td></tr>
				<tr><td align="right">Resource Name:</td><td>%s</td></tr>
				<tr><td align="right">Resource Kind:</td><td>%s</td></tr>
				<tr><td align="right">Namespace:</td><td>%s</td></tr>
			</table>
		</body>
	</html>
	`, time.Unix(evt.Time.GetSeconds(), 0).Format(time.RFC3339), strings.Title(levelMap[strconv.FormatInt(int64(evt.Level), 10)].(string)), evt.Message,
		eventMap[strconv.FormatInt(int64(evt.Type), 10)].(string), evt.Subject.Name, evt.Subject.Kind, evt.Subject.Namespace)
}