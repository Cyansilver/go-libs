package err

const (
	ERR_FAILED_AUTH_CODE        = 1
	ERR_WRONG_PASSWORD_CODE     = 2
	ERR_NOT_FOUND_USER_CODE     = 3
	ERR_DUPLICATE_LOGIN_CODE    = 4
	ERR_EXPIRED_TOKEN_CODE      = 6
	ERR_INVALID_CODE_CODE       = 7
	ERR_NOT_VERIFIED_CODE       = 8
	ERR_DUPLICATE_ACCOUNT_CODE  = 9
	ERR_INVALID_USERNAME_CODE   = 10
	ERR_INVALID_PASSWORD_CODE   = 11
	ERR_FAILED_PERMISSION_CODE  = 12
	ERR_NOT_FOUND_ACCOUNT_CODE  = 13
	ERR_NOT_FOUND_ACCOUNTS_CODE = 14
	ERR_NOT_FOUND_ROLE_CODE     = 15
	ERR_NOT_FOUND_ROLES_CODE    = 16

	ERR_MISSING_PARAMS_CODE = 600
	ERR_INVALID_JSON_CODE   = 601
	ERR_INVALID_DATA_CODE   = 602
	ERR_INTERNAL_ERROR_CODE = 603
	ERR_LOAD_CONFIG_CODE    = 604
	ERROR_NOT_FOUND         = 605
	ERR_INVALID_TOKEN_CODE  = 606

	ERR_FAILED_AUTH_MSG        = "Authentication failed. Please provide valid credentials"
	ERR_WRONG_PASSWORD_MSG     = "Id/Password does not match"
	ERR_NOT_FOUND_USER_MSG     = "User does not exist"
	ERR_DUPLICATE_LOGIN_MSG    = "Duplicate Login"
	ERR_INVALID_TOKEN_MSG      = "Invalid Token"
	ERR_EXPIRED_TOKEN_MSG      = "Expired Token"
	ERR_INVALID_CODE_MSG       = "Invalid Code"
	ERR_NOT_VERIFIED_MSG       = "User hasn't verify the account"
	ERR_DUPLICATE_ACCOUNT_MSG  = "User name already exists"
	ERR_INVALID_USERNAME_MSG   = "User name is invalid"
	ERR_INVALID_PASSWORD_MSG   = "Password is invalid"
	ERR_FAILED_PERMISSION_MSG  = "Access to this resource has been restricted"
	ERR_NOT_FOUND_ACCOUNT_MSG  = "Account does not exist"
	ERR_NOT_FOUND_ACCOUNTS_MSG = "Cannot find accounts"

	ERR_FREE_LIMITED_MSG         = "Cannot create a new one. You reached the free account limit"
	ERR_NOT_FOUND_ALERT_COND_MSG = "Cannot find Alert Condition"
	ERR_NOT_FOUND_ALERT_MSG      = "Cannot find Alert"

	ERR_NOT_FOUND_ROLE_MSG       = "Role does not exist"
	ERR_NOT_FOUND_ROLES_MSG      = "Cannot find roles"
	ERR_NOT_FOUND_SHEET_CELL_MSG = "Cannot find the sheet cell"
	ERR_MISSING_PARAMS_MSG       = "Missing parameters"
	ERR_INVALID_JSON_MSG         = "Invalid Json"
	ERR_INVALID_DATA_MSG         = "Invalid Data"
	ERR_INTERNAL_ERROR_MSG       = "Internal Server Error"
	ERR_LOAD_CONFIG_MSG          = "Cannot load the config"
	ERROR_NOT_FOUND_MSG          = "Cannot find the record"
)

var (
	ErrInternal         = New(ERR_INTERNAL_ERROR_CODE, ERR_INTERNAL_ERROR_MSG)
	ErrInvalidToken     = New(ERR_INVALID_TOKEN_CODE, ERR_INVALID_TOKEN_MSG)
	ErrInvalidCode      = New(ERR_INVALID_CODE_CODE, ERR_INVALID_CODE_MSG)
	ErrFailedAuth       = New(ERR_FAILED_AUTH_CODE, ERR_FAILED_AUTH_MSG)
	ErrFailedPermission = New(ERR_FAILED_PERMISSION_CODE, ERR_FAILED_PERMISSION_MSG)
	ErrInvalidData      = New(ERR_INVALID_DATA_CODE, ERR_INVALID_DATA_MSG)
	ErrInvalidJson      = New(ERR_INVALID_JSON_CODE, ERR_INVALID_JSON_MSG)
	ErrMissingParams    = New(ERR_MISSING_PARAMS_CODE, ERR_MISSING_PARAMS_MSG)
	ErrLoadConfig       = New(ERR_LOAD_CONFIG_CODE, ERR_LOAD_CONFIG_MSG)
	ErrNotFound         = New(ERROR_NOT_FOUND, ERROR_NOT_FOUND_MSG)
)
