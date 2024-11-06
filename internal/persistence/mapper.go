package persistence

import (
	"database/sql"
	"github.com/relvacode/iso8601"
	"strconv"
	"strings"
	"ytlog/internal/tools"
	"ytlog/internal/ytracker"
)

func MapUsers(ytUsers *[]ytracker.User) *[]User {
	users := tools.Map(*ytUsers, func(ytUser ytracker.User) User {
		return User{
			Self:                 ytUser.Self,
			Uid:                  ytUser.Uid,
			Login:                ytUser.Login,
			TrackerUid:           ytUser.TrackerUid,
			PassportUid:          ytUser.PassportUid,
			CloudUid:             ytUser.CloudUid,
			FirstName:            ytUser.FirstName,
			LastName:             ytUser.LastName,
			Display:              ytUser.Display,
			Email:                ytUser.Email,
			External:             ytUser.External,
			HasLicense:           ytUser.HasLicense,
			Dismissed:            ytUser.Dismissed,
			UseNewFilters:        ytUser.UseNewFilters,
			DisableNotifications: ytUser.DisableNotifications,
			FirstLoginDate:       ytUser.FirstLoginDate,
			LastLoginDate:        ytUser.LastLoginDate,
		}
	})
	return &users
}

func MapIssues(ytIssues *[]ytracker.Issue) *[]Issue {
	issues := tools.Map(*ytIssues, func(ytIssue ytracker.Issue) Issue {
		createdAt, _ := iso8601.ParseString(ytIssue.CreatedAt)
		return Issue{
			Key:        ytIssue.Key,
			Type:       ytIssue.Type.Key,
			Status:     ytIssue.Status.Key,
			Summary:    ytIssue.Summary,
			Complexity: ytIssue.Complexity,
			Priority:   ytIssue.Priority.Key,
			Spent:      ytIssue.Spent,
			Queue:      ytIssue.Queue.Key,
			CreatedAt:  createdAt,
		}
	})
	return &issues
}

func mapChangedFiledValue(fieldType string, field interface{}) sql.NullString {
	if field == nil {
		return sql.NullString{}
	}
	var value interface{}
	switch fieldType {
	case "status":
		value = field.(map[string]interface{})["key"]
	case "assignee":
		value = field.(map[string]interface{})["id"]
	case "spent":
		value = field
	case "priority":
		value = field.(map[string]interface{})["key"]
	case "originalEstimation":
		value = field
	case "tags":
		tags := tools.Map(field.([]interface{}), func(t interface{}) string {
			return t.(string)
		})
		value = strings.Join(tags, ",")
	case "boards":
		boards := tools.Map(field.([]interface{}), func(t interface{}) string {
			return strconv.FormatFloat(t.(map[string]interface{})["id"].(float64), 'f', -1, 64)
		})
		value = strings.Join(boards, ",")
	case "complexity":
		value = field
	case "type":
		value = field.(map[string]interface{})["key"]
	default:
		value = ""
	}
	return sql.NullString{
		String: value.(string),
		Valid:  true,
	}
}

func MapIssueLog(log *[]ytracker.IssueLog) *[]IssueLog {
	result := make([]IssueLog, 0)
	for _, issueLog := range *log {
		if issueLog.Fields == nil {
			continue
		}
		for _, field := range issueLog.Fields {
			key := field.Field.Id
			isNeededFields := func(key string) bool {
				switch key {
				case "status", //
					"assignee",
					"spent",
					"priority",
					"originalEstimation",
					"tags",
					"boards",
					"complexity",
					"type":
					return true
				}
				return false
			}
			if !isNeededFields(key) {
				continue
			}
			updatedAt, _ := iso8601.ParseString(issueLog.UpdatedAt)

			result = append(result, IssueLog{
				IssueKey:   issueLog.Issue.Key,
				UpdatedAt:  updatedAt,
				UpdatedBy:  issueLog.UpdatedBy.Id,
				ChangeType: issueLog.Type,
				Field:      field.Field.Id,
				FromValue:  mapChangedFiledValue(field.Field.Id, field.From),
				ToValue:    mapChangedFiledValue(field.Field.Id, field.To),
			})
		}
	}
	return &result
}
