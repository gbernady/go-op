package op

import (
	"fmt"
	"time"
)

type Item struct {
	ID                    string    `json:"id"`
	Title                 string    `json:"title"`
	Version               int       `json:"version"`
	Vault                 Vault     `json:"vault"`
	Category              Category  `json:"category"`
	LastEditedBy          string    `json:"last_edited_by"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
	AdditionalInformation string    `json:"additional_information"`

	Sections []Section `json:"sections"`
	Tags     []string  `json:"tags"`
	Fields   []Field   `json:"fields"`
	Files    []File    `json:"files"`
	URLs     []URL     `json:"urls"`
}

func (i Item) Field(name string) *Field {
	var f *Field
	for _, field := range i.Fields {
		if field.ID == name || field.Label == name {
			f = &field
			break
		}
	}
	return f
}

func (i Item) FindFields(match func(f Field) bool) []Field {
	var f []Field
	for _, field := range i.Fields {
		if match(field) {
			f = append(f, field)
		}
	}
	return f
}

type Category string

const (
	CategoryAPICredential        = "API Credential"
	CategoryBankAccount          = "Bank Account"
	CategoryCreditCard           = "Credit Card"
	CategoryDatabase             = "Database"
	CategoryDocument             = "Document"
	CategoryDriverLicense        = "Driver License"
	CategoryEmailAccount         = "Email Account"
	CategoryIdentity             = "Identity"
	CategoryLogin                = "Login"
	CategoryMedicalRecord        = "Medical Record"
	CategoryMembership           = "Membership"
	CategoryOutdoorLicense       = "Outdoor License"
	CategoryPassport             = "Passport"
	CategoryPassword             = "Password"
	CategoryRewardProgram        = "Reward Program"
	CategorySecureNote           = "Secure Note"
	CategoryServer               = "Server"
	CategorySocialSecurityNumber = "Social Security Number"
	CategorySoftwareLicense      = "Software License"
	CategorySSHKey               = "SSH Key"
	CategoryWirelessRouter       = "Wireless Router"
)

func (c *Category) UnmarshalText(text []byte) error {
	switch string(text) {
	case "API_CREDENTIAL":
		*c = CategoryAPICredential
	case "BANK_ACCOUNT":
		*c = CategoryBankAccount
	case "CREDIT_CARD":
		*c = CategoryCreditCard
	case "DATABASE":
		*c = CategoryDatabase
	case "DOCUMENT":
		*c = CategoryDocument
	case "DRIVER_LICENSE":
		*c = CategoryDriverLicense
	case "EMAIL_ACCOUNT":
		*c = CategoryEmailAccount
	case "IDENTITY":
		*c = CategoryIdentity
	case "LOGIN":
		*c = CategoryLogin
	case "MEDICAL_RECORD":
		*c = CategoryMedicalRecord
	case "MEMBERSHIP":
		*c = CategoryMembership
	case "OUTDOOR_LICENSE":
		*c = CategoryOutdoorLicense
	case "PASSPORT":
		*c = CategoryPassport
	case "PASSWORD":
		*c = CategoryPassword
	case "REWARD_PROGRAM":
		*c = CategoryRewardProgram
	case "SECURE_NOTE":
		*c = CategorySecureNote
	case "SERVER":
		*c = CategoryServer
	case "SOCIAL_SECURITY_NUMBER":
		*c = CategorySocialSecurityNumber
	case "SOFTWARE_LICENSE":
		*c = CategorySoftwareLicense
	case "SSH_KEY":
		*c = CategorySSHKey
	case "WIRELESS_ROUTER":
		*c = CategoryWirelessRouter
	default:
		return fmt.Errorf("unrecognized category %q", string(text))
	}
	return nil
}

type Field struct {
	ID              string          `json:"id"`
	Section         Section         `json:"section"`
	Type            FieldType       `json:"type"`
	Purpose         FieldPurpose    `json:"purpose"`
	Label           string          `json:"label"`
	Value           string          `json:"value"`
	TOTP            string          `json:"totp"`
	Entropy         int64           `json:"entropy"`
	PasswordDetails PasswordDetails `json:"password_details"`
	Reference       string          `json:"reference"`
}

type FieldAssignment struct {
	Label   string              `json:"label"`
	Type    FieldAssignmentType `json:"type"`
	Value   string              `json:"value"`
	Purpose FieldPurpose        `json:"purpose"`
}

type FieldAssignmentType string

const (
	FieldAssignmentTypeConcealed = "concealed"
	FieldAssignmentTypeText      = "text"
	FieldAssignmentTypeEmail     = "email"
	FieldAssignmentTypeURL       = "url"
	FieldAssignmentTypeDate      = "date"
	FieldAssignmentTypeMonthYear = "monthYear"
	FieldAssignmentTypePhone     = "phone"
)

type FieldType string

const (
	FieldTypeAddress          = "address"
	FieldTypeConcealed        = "concealed"
	FieldTypeCreditCardNumber = "ccnum"
	FieldTypeCreditCardType   = "cctype"
	FieldTypeDate             = "date"
	FieldTypeEmail            = "email"
	FieldTypeFile             = "file"
	FieldTypeGender           = "gender"
	FieldTypeMenu             = "menu"
	FieldTypeMonthYear        = "monthYear"
	FieldTypeOTP              = "OTP"
	FieldTypePhone            = "phone"
	FieldTypeReference        = "reference"
	FieldTypeSSHKey           = "sshkey"
	FieldTypeString           = "string"
	FieldTypeUnknown          = ""
	FieldTypeURL              = "URL"
)

func (f *FieldType) UnmarshalText(text []byte) error {
	switch string(text) {
	case "ADDRESS":
		*f = FieldTypeAddress
	case "CONCEALED":
		*f = FieldTypeConcealed
	case "CREDIT_CARD_NUMBER":
		*f = FieldTypeCreditCardNumber
	case "CREDIT_CARD_TYPE":
		*f = FieldTypeCreditCardType
	case "DATE":
		*f = FieldTypeDate
	case "EMAIL":
		*f = FieldTypeEmail
	case "FILE":
		*f = FieldTypeFile
	case "GENDER":
		*f = FieldTypeGender
	case "MENU":
		*f = FieldTypeMenu
	case "MONTH_YEAR":
		*f = FieldTypeMonthYear
	case "OTP":
		*f = FieldTypeOTP
	case "PHONE":
		*f = FieldTypePhone
	case "REFERENCE":
		*f = FieldTypeReference
	case "SSHKEY":
		*f = FieldTypeSSHKey
	case "STRING":
		*f = FieldTypeString
	case "URL":
		*f = FieldTypeURL
	case "UNKNOWN":
		fallthrough
	default:
		*f = FieldTypeUnknown
	}
	return nil
}

type FieldPurpose string

const (
	FieldPurposeNotes    = "NOTES"
	FieldPurposePassword = "PASSWORD"
	FieldPurposeUsername = "USERNAME"
)

type File struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Size        int64   `json:"size"`
	ContentPath string  `json:"content_path"`
	Section     Section `json:"section"`
}

type PasswordDetails struct {
	Entropy   int64            `json:"entropy"`
	Generated bool             `json:"generated"`
	Strength  PasswordStrength `json:"strength"`
}

type PasswordStrength string

const (
	PasswordStrengthTerrible  = "TERRIBLE"
	PasswordStrengthWeak      = "WEAK"
	PasswordStrengthFair      = "FAIR"
	PasswordStrengthGood      = "GOOD"
	PasswordStrengthVeryGood  = "VERY_GOOD"
	PasswordStrengthExcellent = "EXCELLENT"
	PasswordStrengthFantastic = "FANTASTIC"
)

type Section struct {
	ID    string `json:"id"`
	Label string `json:"label"`
}

type URL struct {
	Label   string `json:"label"`
	Primary bool   `json:"primary"`
	HRef    string `json:"href"`
}

// ListItems returns a list of all items the account has read access to.
// Excludes items in the Archive by default.
//
// Supported filters:
//
//	--categories categories   Only list items in these categories (comma-separated).
//	--favorite                Only list favorite items
//	--include-archive         Include items in the Archive. Can also be set using OP_INCLUDE_ARCHIVE environment variable.
//	--tags tags               Only list items with these tags (comma-separated).
//	--vault vault             Only list items in this vault.
func (c *CLI) ListItems(filters ...Filter) ([]Item, error) {
	var val []Item
	err := c.execJSON(applyFilters([]string{"item", "list"}, filters), nil, &val)
	return val, err
}

// CreateItem creates a new item and returns it with all the fileds like ID filled.
//
//	--dry-run                      Perform a dry run of the command and output a preview of the resulting item.
//	--generate-password[=recipe]   Give the item a randomly generated password.
func (c *CLI) CreateItem(item *Item) (*Item, error) {
	return nil, nil
}

// GetItem returns the details of an item specified by its name, ID, or sharing link.
//
// Supported filters:
//
//	--include-archive         Include items in the Archive. Can also be set using OP_INCLUDE_ARCHIVE environment variable.
//	--vault vault             Only list items in this vault.
func (c *CLI) GetItem(name string, filters ...Filter) (*Item, error) {
	var val *Item
	err := c.execJSON([]string{"item", "get", sanitize(name)}, nil, &val)
	return val, err
}

// GetItemTemplate returns an item template for a given item category.
func (c *CLI) GetItemTemplate(cat Category) (*Item, error) {
	var val *Item
	err := c.execJSON([]string{"item", "template", "get", sanitize(string(cat))}, nil, &val)
	return val, err
}

// DeleteItem permanently deletes an item specified by its name, ID, or sharing link.
func (c *CLI) DeleteItem(name string) error {
	_, err := c.execRaw([]string{"item", "delete", sanitize(name)}, nil)
	return err
}

// ArchiveItem archives the item specified by its name, ID, or sharing link.
func (c *CLI) ArchiveItem(name string) error {
	_, err := c.execRaw([]string{"item", "delete", "--archive", sanitize(name)}, nil)
	return err
}
