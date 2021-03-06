// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

const (

	// ErrCodeCertificateConflictException for service response error code
	// "CertificateConflictException".
	//
	// Unable to verify the CA certificate used to sign the device certificate you
	// are attempting to register. This is happens when you have registered more
	// than one CA certificate that has the same subject field and public key.
	ErrCodeCertificateConflictException = "CertificateConflictException"

	// ErrCodeCertificateStateException for service response error code
	// "CertificateStateException".
	//
	// The certificate operation is not allowed.
	ErrCodeCertificateStateException = "CertificateStateException"

	// ErrCodeCertificateValidationException for service response error code
	// "CertificateValidationException".
	//
	// The certificate is invalid.
	ErrCodeCertificateValidationException = "CertificateValidationException"

	// ErrCodeConflictingResourceUpdateException for service response error code
	// "ConflictingResourceUpdateException".
	//
	// A conflicting resource update exception. This exception is thrown when two
	// pending updates cause a conflict.
	ErrCodeConflictingResourceUpdateException = "ConflictingResourceUpdateException"

	// ErrCodeDeleteConflictException for service response error code
	// "DeleteConflictException".
	//
	// You can't delete the resource because it is attached to one or more resources.
	ErrCodeDeleteConflictException = "DeleteConflictException"

	// ErrCodeIndexNotReadyException for service response error code
	// "IndexNotReadyException".
	//
	// The index is not ready.
	ErrCodeIndexNotReadyException = "IndexNotReadyException"

	// ErrCodeInternalException for service response error code
	// "InternalException".
	//
	// An unexpected error has occurred.
	ErrCodeInternalException = "InternalException"

	// ErrCodeInternalFailureException for service response error code
	// "InternalFailureException".
	//
	// An unexpected error has occurred.
	ErrCodeInternalFailureException = "InternalFailureException"

	// ErrCodeInvalidQueryException for service response error code
	// "InvalidQueryException".
	//
	// The query is invalid.
	ErrCodeInvalidQueryException = "InvalidQueryException"

	// ErrCodeInvalidRequestException for service response error code
	// "InvalidRequestException".
	//
	// The request is not valid.
	ErrCodeInvalidRequestException = "InvalidRequestException"

	// ErrCodeInvalidResponseException for service response error code
	// "InvalidResponseException".
	//
	// The response is invalid.
	ErrCodeInvalidResponseException = "InvalidResponseException"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// The number of attached entities exceeds the limit.
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodeMalformedPolicyException for service response error code
	// "MalformedPolicyException".
	//
	// The policy documentation is not valid.
	ErrCodeMalformedPolicyException = "MalformedPolicyException"

	// ErrCodeNotConfiguredException for service response error code
	// "NotConfiguredException".
	//
	// The resource is not configured.
	ErrCodeNotConfiguredException = "NotConfiguredException"

	// ErrCodeRegistrationCodeValidationException for service response error code
	// "RegistrationCodeValidationException".
	//
	// The registration code is invalid.
	ErrCodeRegistrationCodeValidationException = "RegistrationCodeValidationException"

	// ErrCodeResourceAlreadyExistsException for service response error code
	// "ResourceAlreadyExistsException".
	//
	// The resource already exists.
	ErrCodeResourceAlreadyExistsException = "ResourceAlreadyExistsException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// The specified resource does not exist.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeResourceRegistrationFailureException for service response error code
	// "ResourceRegistrationFailureException".
	//
	// The resource registration failed.
	ErrCodeResourceRegistrationFailureException = "ResourceRegistrationFailureException"

	// ErrCodeServiceUnavailableException for service response error code
	// "ServiceUnavailableException".
	//
	// The service is temporarily unavailable.
	ErrCodeServiceUnavailableException = "ServiceUnavailableException"

	// ErrCodeSqlParseException for service response error code
	// "SqlParseException".
	//
	// The Rule-SQL expression can't be parsed correctly.
	ErrCodeSqlParseException = "SqlParseException"

	// ErrCodeThrottlingException for service response error code
	// "ThrottlingException".
	//
	// The rate exceeds the limit.
	ErrCodeThrottlingException = "ThrottlingException"

	// ErrCodeTransferAlreadyCompletedException for service response error code
	// "TransferAlreadyCompletedException".
	//
	// You can't revert the certificate transfer because the transfer is already
	// complete.
	ErrCodeTransferAlreadyCompletedException = "TransferAlreadyCompletedException"

	// ErrCodeTransferConflictException for service response error code
	// "TransferConflictException".
	//
	// You can't transfer the certificate because authorization policies are still
	// attached.
	ErrCodeTransferConflictException = "TransferConflictException"

	// ErrCodeUnauthorizedException for service response error code
	// "UnauthorizedException".
	//
	// You are not authorized to perform this operation.
	ErrCodeUnauthorizedException = "UnauthorizedException"

	// ErrCodeVersionConflictException for service response error code
	// "VersionConflictException".
	//
	// An exception thrown when the version of a thing passed to a command is different
	// than the version specified with the --version parameter.
	ErrCodeVersionConflictException = "VersionConflictException"

	// ErrCodeVersionsLimitExceededException for service response error code
	// "VersionsLimitExceededException".
	//
	// The number of policy versions exceeds the limit.
	ErrCodeVersionsLimitExceededException = "VersionsLimitExceededException"
)
