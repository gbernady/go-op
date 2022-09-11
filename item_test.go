package op

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItemField(t *testing.T) {
	tests := []struct {
		name   string
		item   Item
		key    string
		result *Field
	}{
		{
			name: "MatchID",
			item: Item{
				Fields: []Field{
					{
						ID:        "username",
						Type:      FieldTypeString,
						Label:     "user",
						Value:     "foo",
						Reference: "op://Personal/Foo/username",
					},
					{
						ID:        "password",
						Type:      FieldTypeConcealed,
						Label:     "pass",
						Value:     "bar",
						Reference: "op://Personal/Foo/password",
					},
				},
			},
			key: "username",
			result: &Field{
				ID:        "username",
				Type:      FieldTypeString,
				Label:     "user",
				Value:     "foo",
				Reference: "op://Personal/Foo/username",
			},
		},
		{
			name: "MatchLabel",
			item: Item{
				Fields: []Field{
					{
						ID:        "username",
						Type:      FieldTypeString,
						Label:     "uname",
						Value:     "foo",
						Reference: "op://Personal/Foo/username",
					},
					{
						ID:        "password",
						Type:      FieldTypeConcealed,
						Label:     "passwd",
						Value:     "bar",
						Reference: "op://Personal/Foo/password",
					},
				},
			},
			key: "uname",
			result: &Field{
				ID:        "username",
				Type:      FieldTypeString,
				Label:     "uname",
				Value:     "foo",
				Reference: "op://Personal/Foo/username",
			},
		},
		{
			name: "MatchFirst",
			item: Item{
				Fields: []Field{
					{
						ID:        "username",
						Type:      FieldTypeString,
						Label:     "uname",
						Value:     "foo",
						Reference: "op://Personal/Foo/username",
					},
					{
						ID: "06CDE696F7B54212BE47E7F99CF674F0",
						Section: Section{
							ID: "Section_FFD16B98A713452695E49DA0EB32BFD0",
						},
						Type:      FieldTypeString,
						Label:     "username",
						Value:     "wat",
						Reference: "op://Personal/Foo/Section_FFD16B98A713452695E49DA0EB32BFD0/username",
					},
				},
			},
			key: "username",
			result: &Field{
				ID:        "username",
				Type:      FieldTypeString,
				Label:     "uname",
				Value:     "foo",
				Reference: "op://Personal/Foo/username",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.result, test.item.Field(test.key))
		})
	}
}

func TestItemFindFields(t *testing.T) {
	tests := []struct {
		name   string
		item   Item
		matchF func(f Field) bool
		result []Field
	}{
		{
			name: "MatchSingle",
			item: Item{
				Fields: []Field{
					{
						ID:        "username",
						Type:      FieldTypeString,
						Label:     "username",
						Value:     "foo",
						Reference: "op://Personal/Foo/username",
					},
					{
						ID:        "password",
						Type:      FieldTypeConcealed,
						Label:     "password",
						Value:     "bar",
						Reference: "op://Personal/Foo/password",
					},
					{
						ID:        "email",
						Type:      FieldTypeString,
						Label:     "email",
						Value:     "foo@example.com",
						Reference: "op://Personal/Foo/email",
					},
				},
			},
			matchF: func(f Field) bool { return f.Label == "password" },
			result: []Field{
				{
					ID:        "password",
					Type:      FieldTypeConcealed,
					Label:     "password",
					Value:     "bar",
					Reference: "op://Personal/Foo/password",
				},
			},
		},
		{
			name: "MatchMultiple",
			item: Item{
				Fields: []Field{
					{
						ID:        "username",
						Type:      FieldTypeString,
						Label:     "username",
						Value:     "foo",
						Reference: "op://Personal/Foo/username",
					},
					{
						ID:        "password",
						Type:      FieldTypeConcealed,
						Label:     "password",
						Value:     "bar",
						Reference: "op://Personal/Foo/password",
					},
					{
						ID:        "email",
						Type:      FieldTypeString,
						Label:     "email",
						Value:     "foo@example.com",
						Reference: "op://Personal/Foo/email",
					},
				},
			},
			matchF: func(f Field) bool { return f.Type == FieldTypeString },
			result: []Field{
				{
					ID:        "username",
					Type:      FieldTypeString,
					Label:     "username",
					Value:     "foo",
					Reference: "op://Personal/Foo/username",
				},
				{
					ID:        "email",
					Type:      FieldTypeString,
					Label:     "email",
					Value:     "foo@example.com",
					Reference: "op://Personal/Foo/email",
				},
			},
		},
		{
			name: "MatchNone",
			item: Item{
				Fields: []Field{
					{
						ID:        "username",
						Type:      FieldTypeString,
						Label:     "username",
						Value:     "foo",
						Reference: "op://Personal/Foo/username",
					},
					{
						ID:        "password",
						Type:      FieldTypeConcealed,
						Label:     "password",
						Value:     "bar",
						Reference: "op://Personal/Foo/password",
					},
					{
						ID:        "email",
						Type:      FieldTypeString,
						Label:     "email",
						Value:     "foo@example.com",
						Reference: "op://Personal/Foo/email",
					},
				},
			},
			matchF: func(f Field) bool { return f.Type == FieldTypeDate },
			result: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.result, test.item.FindFields(test.matchF))
		})
	}
}

func TestCategoryUnmarshalText(t *testing.T) {
	tests := []struct {
		text   string
		result Category
	}{
		{
			text:   "API_CREDENTIAL",
			result: CategoryAPICredential,
		},
		{
			text:   "BANK_ACCOUNT",
			result: CategoryBankAccount,
		},
		{
			text:   "CREDIT_CARD",
			result: CategoryCreditCard},
		{
			text:   "DATABASE",
			result: CategoryDatabase},
		{
			text:   "DOCUMENT",
			result: CategoryDocument,
		},
		{
			text:   "DRIVER_LICENSE",
			result: CategoryDriverLicense,
		},
		{
			text:   "EMAIL_ACCOUNT",
			result: CategoryEmailAccount,
		},
		{
			text:   "IDENTITY",
			result: CategoryIdentity,
		},
		{
			text:   "LOGIN",
			result: CategoryLogin,
		},
		{
			text:   "MEDICAL_RECORD",
			result: CategoryMedicalRecord,
		},
		{
			text:   "MEMBERSHIP",
			result: CategoryMembership,
		},
		{
			text:   "OUTDOOR_LICENSE",
			result: CategoryOutdoorLicense,
		},
		{
			text:   "PASSPORT",
			result: CategoryPassport,
		},
		{
			text:   "PASSWORD",
			result: CategoryPassword,
		},
		{
			text:   "REWARD_PROGRAM",
			result: CategoryRewardProgram,
		},
		{
			text:   "SECURE_NOTE",
			result: CategorySecureNote,
		},
		{
			text:   "SERVER",
			result: CategoryServer,
		},
		{
			text:   "SOCIAL_SECURITY_NUMBER",
			result: CategorySocialSecurityNumber,
		},
		{
			text:   "SOFTWARE_LICENSE",
			result: CategorySoftwareLicense,
		},
		{
			text:   "SSH_KEY",
			result: CategorySSHKey,
		},
		{
			text:   "WIRELESS_ROUTER",
			result: CategoryWirelessRouter,
		},
		{
			text:   "invalid",
			result: Category(""),
		},
	}
	for _, test := range tests {
		var c Category
		err := c.UnmarshalText([]byte(test.text))
		assert.NoError(t, err)
		assert.Equal(t, test.result, c)
	}
}

func TestFieldTypeUnmarshalText(t *testing.T) {
	tests := []struct {
		text   string
		result FieldType
	}{
		{
			text:   "ADDRESS",
			result: FieldTypeAddress,
		},
		{
			text:   "CONCEALED",
			result: FieldTypeConcealed,
		},
		{
			text:   "CREDIT_CARD_NUMBER",
			result: FieldTypeCreditCardNumber,
		},
		{
			text:   "CREDIT_CARD_TYPE",
			result: FieldTypeCreditCardType,
		},
		{
			text:   "DATE",
			result: FieldTypeDate,
		},
		{
			text:   "EMAIL",
			result: FieldTypeEmail,
		},
		{
			text:   "FILE",
			result: FieldTypeFile,
		},
		{
			text:   "GENDER",
			result: FieldTypeGender,
		},
		{
			text:   "MENU",
			result: FieldTypeMenu,
		},
		{
			text:   "MONTH_YEAR",
			result: FieldTypeMonthYear,
		},
		{
			text:   "OTP",
			result: FieldTypeOTP,
		},
		{
			text:   "PHONE",
			result: FieldTypePhone,
		},
		{
			text:   "REFERENCE",
			result: FieldTypeReference,
		},
		{
			text:   "SSHKEY",
			result: FieldTypeSSHKey,
		},
		{
			text:   "STRING",
			result: FieldTypeString,
		},
		{
			text:   "UNKNOWN",
			result: FieldTypeUnknown,
		},
		{
			text:   "URL",
			result: FieldTypeURL,
		},
		{
			text:   "invalid",
			result: FieldTypeUnknown,
		},
	}
	for _, test := range tests {
		var ft FieldType
		err := ft.UnmarshalText([]byte(test.text))
		assert.NoError(t, err)
		assert.Equal(t, test.result, ft)
	}
}

func TestListItems(t *testing.T) {
	// FIXME: implement
}

func TestCreateItem(t *testing.T) {
	// FIXME: implement
}

func TestGetItem(t *testing.T) {
	// FIXME: implement
}

func TestGetItemTemplate(t *testing.T) {
	// FIXME: implement
}

func TestDeleteItem(t *testing.T) {
	// FIXME: implement
}

func TestArchiveItem(t *testing.T) {
	// FIXME: implement
}
