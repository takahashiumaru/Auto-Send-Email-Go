package helper

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func RegisterValidation(validate *validator.Validate) {
	validate.RegisterValidation("account_type", validateAccountType)
	validate.RegisterValidation("discount_proposal_confirmation_status_type", validateDiscountProposalConfirmationStatusType)
	validate.RegisterValidation("discount_proposal_confirmation_status_unit", validateDiscountProposalConfirmationStatusUnit)
	validate.RegisterValidation("discount_proposal_transfer_type", validateDiscountProposalTransferType)
	validate.RegisterValidation("gender", validateGender)
	validate.RegisterValidation("ktp", validateKtp)
	validate.RegisterValidation("npwp", validateNpwp)
	validate.RegisterValidation("period_day", validatePeriodDay)
	validate.RegisterValidation("period_month", validatePeriodMonth)
	validate.RegisterValidation("religion", validateReligion)
	validate.RegisterValidation("event_organizer_type", validateEventOrganizerType)
	validate.RegisterValidation("active_status", validateActive)
	validate.RegisterValidation("class", validateClass)
	validate.RegisterValidation("discount_proposal_category_user", validateDiscountProposalCategoryUser)
	validate.RegisterValidation("discount_proposal_type", validateDiscountProposalType)
	validate.RegisterValidation("quantity_type", validateQuantityType)
	validate.RegisterValidation("note_status", validateNoteStatus)
	validate.RegisterValidation("department", validateDepartment)
}

func validateNpwp(fieldLevel validator.FieldLevel) bool {
	value := fieldLevel.Field().String()

	if len(value) != 20 {
		return false
	}

	const pattern = "[0-9]{2}[.][0-9]{3}[.][0-9]{3}[.][0-9]{1}[-][0-9]{3}[.][0-9]{3}"
	regex := regexp.MustCompile(pattern)
	result := regex.MatchString(value)
	return result
}

func validateKtp(fieldLevel validator.FieldLevel) bool {
	value := fieldLevel.Field().String()

	if value == "" {
		return true
	}

	if len(value) != 16 {
		return false
	}

	const pattern = "(1[1-9]|21|[37][1-6]|5[1-3]|6[1-5]|[89][12])\\d{2}\\d{2}([04][1-9]|[1256][0-9]|[37][01])(0[1-9]|1[0-2])\\d{2}\\d{4}"
	regex := regexp.MustCompile(pattern)
	result := regex.MatchString(value)
	return result
}

func validateAccountType(fieldLevel validator.FieldLevel) bool {
	value := fieldLevel.Field().String()

	if value == "" {
		return true
	}

	const pattern = "(RKC|RBO)"
	regex := regexp.MustCompile(pattern)
	result := regex.MatchString(value)
	return result
}

func validateGender(fieldLevel validator.FieldLevel) bool {
	value := fieldLevel.Field().String()

	const pattern = "(Laki-laki|Perempuan)"
	regex := regexp.MustCompile(pattern)
	result := regex.MatchString(value)
	return result
}

func validateReligion(fieldLevel validator.FieldLevel) bool {
	value := fieldLevel.Field().String()

	const pattern = "(Islam|Katolik|Protestan|Hindu|Buddha|Konghucu)"
	regex := regexp.MustCompile(pattern)
	result := regex.MatchString(value)
	return result
}

func validateDiscountProposalCategoryUser(fieldLevel validator.FieldLevel) bool {
	value := fieldLevel.Field().String()

	const pattern = "(Tenaga Medis|Dokter|Non Dokter)"
	regex := regexp.MustCompile(pattern)
	result := regex.MatchString(value)
	return result
}

func validatePeriodDay(fieldLevel validator.FieldLevel) bool {
	value := fieldLevel.Field().String()

	if len(value) != 8 {
		return false
	}

	const pattern = ("\\d{4}(0[1-9]|1[012])(0[1-9]|[12][0-9]|3[01])")
	regex := regexp.MustCompile(pattern)
	result := regex.MatchString(value)
	return result
}

func validateDiscountProposalConfirmationStatusUnit(fieldLevel validator.FieldLevel) bool {
	value := fieldLevel.Field().String()

	const pattern = ("(Days|Minutes|Hours)")
	regex := regexp.MustCompile(pattern)
	result := regex.MatchString(value)
	return result
}

func validateDiscountProposalConfirmationStatusType(fieldLevel validator.FieldLevel) bool {
	value := fieldLevel.Field().String()

	const pattern = "(SKI|DPF)"
	regex := regexp.MustCompile(pattern)
	result := regex.MatchString(value)
	return result
}

func validateDiscountProposalTransferType(fieldLevel validator.FieldLevel) bool {
	value := fieldLevel.Field().String()

	const pattern = "(BO|504)"
	regex := regexp.MustCompile(pattern)
	result := regex.MatchString(value)
	return result
}

func validatePeriodMonth(fieldLevel validator.FieldLevel) bool {
	value := fieldLevel.Field().String()

	if len(value) != 6 {
		return false
	}

	const pattern = "\\d{4}(0[1-9]|1[012])"
	regex := regexp.MustCompile(pattern)
	result := regex.MatchString(value)
	return result
}

func validateClass(fieldLevel validator.FieldLevel) bool {
	value := fieldLevel.Field().String()

	const pattern = "General|Student|Specialist"
	regex := regexp.MustCompile(pattern)
	result := regex.MatchString(value)
	return result
}

func validateDiscountProposalType(fieldLevel validator.FieldLevel) bool {
	value := fieldLevel.Field().String()

	const pattern = "(SKI1|SKI2|DPL|DPF)"
	regex := regexp.MustCompile(pattern)
	result := regex.MatchString(value)
	return result
}

func validateEventOrganizerType(fieldLevel validator.FieldLevel) bool {
	value := fieldLevel.Field().String()

	const pattern = "College|Profession|Institution"
	regex := regexp.MustCompile(pattern)
	result := regex.MatchString(value)
	return result
}

func validateActive(fieldLevel validator.FieldLevel) bool {
	value := fieldLevel.Field().String()

	const pattern = "Active|Not Active"
	regex := regexp.MustCompile(pattern)
	result := regex.MatchString(value)
	return result
}

func validateQuantityType(fieldLevel validator.FieldLevel) bool {
	value := fieldLevel.Field().String()

	const pattern = "(FULL|HALF)"
	regex := regexp.MustCompile(pattern)
	result := regex.MatchString(value)
	return result
}

func validateNoteStatus(fieldLevel validator.FieldLevel) bool {
	value := fieldLevel.Field().String()

	const pattern = "(ACCEPT|REJECT)"
	regex := regexp.MustCompile(pattern)
	result := regex.MatchString(value)
	return result
}

func validateDepartment(fieldLevel validator.FieldLevel) bool {
	value := fieldLevel.Field().String()

	const pattern = "(MKT|NON MKT)"
	regex := regexp.MustCompile(pattern)
	result := regex.MatchString(value)
	return result
}
