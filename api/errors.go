package api

import (
	"context"

	"github.com/janotchain/janeta/account"
	"github.com/janotchain/janeta/asset"
	"github.com/janotchain/janeta/blockchain/pseudohsm"
	"github.com/janotchain/janeta/blockchain/rpc"
	"github.com/janotchain/janeta/blockchain/signers"
	"github.com/janotchain/janeta/blockchain/txbuilder"
	"github.com/janotchain/janeta/contract"
	"github.com/janotchain/janeta/errors"
	"github.com/janotchain/janeta/net/http/httperror"
	"github.com/janotchain/janeta/net/http/httpjson"
	"github.com/janotchain/janeta/protocol/validation"
	"github.com/janotchain/janeta/protocol/vm"
)

var (
	// ErrDefault is default janeta API Error
	ErrDefault = errors.New("janeta API Error")
)

func isTemporary(info httperror.Info, err error) bool {
	switch info.ChainCode {
	case "JTA000": // internal server error
		return true
	case "JTA001": // request timed out
		return true
	case "JTA761": // outputs currently reserved
		return true
	case "JTA706": // 1 or more action errors
		errs := errors.Data(err)["actions"].([]httperror.Response)
		temp := true
		for _, actionErr := range errs {
			temp = temp && isTemporary(actionErr.Info, nil)
		}
		return temp
	default:
		return false
	}
}

var respErrFormatter = map[error]httperror.Info{
	ErrDefault: {500, "JTA000", "janeta API Error"},

	// Signers error namespace (2xx)
	signers.ErrBadQuorum: {400, "JTA200", "Quorum must be greater than or equal to 1, and must be less than or equal to the length of xpubs"},
	signers.ErrBadXPub:   {400, "JTA201", "Invalid xpub format"},
	signers.ErrNoXPubs:   {400, "JTA202", "At least one xpub is required"},
	signers.ErrDupeXPub:  {400, "JTA203", "Root XPubs cannot contain the same key more than once"},

	// Contract error namespace (3xx)
	contract.ErrContractDuplicated: {400, "JTA302", "Contract is duplicated"},
	contract.ErrContractNotFound:   {400, "JTA303", "Contract not found"},

	// Transaction error namespace (7xx)
	// Build transaction error namespace (70x ~ 72x)
	account.ErrInsufficient:         {400, "JTA700", "Funds of account are insufficient"},
	account.ErrImmature:             {400, "JTA701", "Available funds of account are immature"},
	account.ErrReserved:             {400, "JTA702", "Available UTXOs of account have been reserved"},
	account.ErrMatchUTXO:            {400, "JTA703", "UTXO with given hash not found"},
	ErrBadActionType:                {400, "JTA704", "Invalid action type"},
	ErrBadAction:                    {400, "JTA705", "Invalid action object"},
	ErrBadActionConstruction:        {400, "JTA706", "Invalid action construction"},
	txbuilder.ErrMissingFields:      {400, "JTA707", "One or more fields are missing"},
	txbuilder.ErrBadAmount:          {400, "JTA708", "Invalid asset amount"},
	account.ErrFindAccount:          {400, "JTA709", "Account not found"},
	asset.ErrFindAsset:              {400, "JTA710", "Asset not found"},
	txbuilder.ErrBadContractArgType: {400, "JTA711", "Invalid contract argument type"},
	txbuilder.ErrOrphanTx:           {400, "JTA712", "Transaction input UTXO not found"},
	txbuilder.ErrExtTxFee:           {400, "JTA713", "Transaction fee exceeded max limit"},
	txbuilder.ErrNoGasInput:         {400, "JTA714", "Transaction has no gas input"},

	// Submit transaction error namespace (73x ~ 79x)
	// Validation error (73x ~ 75x)
	validation.ErrTxVersion:                 {400, "JTA730", "Invalid transaction version"},
	validation.ErrWrongTransactionSize:      {400, "JTA731", "Invalid transaction size"},
	validation.ErrBadTimeRange:              {400, "JTA732", "Invalid transaction time range"},
	validation.ErrNotStandardTx:             {400, "JTA733", "Not standard transaction"},
	validation.ErrWrongCoinbaseTransaction:  {400, "JTA734", "Invalid coinbase transaction"},
	validation.ErrWrongCoinbaseAsset:        {400, "JTA735", "Invalid coinbase assetID"},
	validation.ErrCoinbaseArbitraryOversize: {400, "JTA736", "Invalid coinbase arbitrary size"},
	validation.ErrEmptyResults:              {400, "JTA737", "No results in the transaction"},
	validation.ErrMismatchedAssetID:         {400, "JTA738", "Mismatched assetID"},
	validation.ErrMismatchedPosition:        {400, "JTA739", "Mismatched value source/dest position"},
	validation.ErrMismatchedReference:       {400, "JTA740", "Mismatched reference"},
	validation.ErrMismatchedValue:           {400, "JTA741", "Mismatched value"},
	validation.ErrMissingField:              {400, "JTA742", "Missing required field"},
	validation.ErrNoSource:                  {400, "JTA743", "No source for value"},
	validation.ErrOverflow:                  {400, "JTA744", "Arithmetic overflow/underflow"},
	validation.ErrPosition:                  {400, "JTA745", "Invalid source or destination position"},
	validation.ErrUnbalanced:                {400, "JTA746", "Unbalanced asset amount between input and output"},
	validation.ErrOverGasCredit:             {400, "JTA747", "Gas credit has been spent"},
	validation.ErrGasCalculate:              {400, "JTA748", "Gas usage calculate got a math error"},

	// VM error (76x ~ 78x)
	vm.ErrAltStackUnderflow:  {400, "JTA760", "Alt stack underflow"},
	vm.ErrBadValue:           {400, "JTA761", "Bad value"},
	vm.ErrContext:            {400, "JTA762", "Wrong context"},
	vm.ErrDataStackUnderflow: {400, "JTA763", "Data stack underflow"},
	vm.ErrDisallowedOpcode:   {400, "JTA764", "Disallowed opcode"},
	vm.ErrDivZero:            {400, "JTA765", "Division by zero"},
	vm.ErrFalseVMResult:      {400, "JTA766", "False result for executing VM"},
	vm.ErrLongProgram:        {400, "JTA767", "Program size exceeds max int32"},
	vm.ErrRange:              {400, "JTA768", "Arithmetic range error"},
	vm.ErrReturn:             {400, "JTA769", "RETURN executed"},
	vm.ErrRunLimitExceeded:   {400, "JTA770", "Run limit exceeded because the JTA Fee is insufficient"},
	vm.ErrShortProgram:       {400, "JTA771", "Unexpected end of program"},
	vm.ErrToken:              {400, "JTA772", "Unrecognized token"},
	vm.ErrUnexpected:         {400, "JTA773", "Unexpected error"},
	vm.ErrUnsupportedVM:      {400, "JTA774", "Unsupported VM because the version of VM is mismatched"},
	vm.ErrVerifyFailed:       {400, "JTA775", "VERIFY failed"},

	// Mock HSM error namespace (8xx)
	pseudohsm.ErrDuplicateKeyAlias: {400, "JTA800", "Key Alias already exists"},
	pseudohsm.ErrLoadKey:           {400, "JTA801", "Key not found or wrong password"},
	pseudohsm.ErrDecrypt:           {400, "JTA802", "Could not decrypt key with given passphrase"},
}

// Map error values to standard janeta error codes. Missing entries
// will map to internalErrInfo.
//
// TODO(jackson): Share one error table across Chain
// products/services so that errors are consistent.
var errorFormatter = httperror.Formatter{
	Default:     httperror.Info{500, "JTA000", "janeta API Error"},
	IsTemporary: isTemporary,
	Errors: map[error]httperror.Info{
		// General error namespace (0xx)
		context.DeadlineExceeded: {408, "JTA001", "Request timed out"},
		httpjson.ErrBadRequest:   {400, "JTA002", "Invalid request body"},
		rpc.ErrWrongNetwork:      {502, "JTA103", "A peer core is operating on a different blockchain network"},

		//accesstoken authz err namespace (86x)
		errNotAuthenticated: {401, "JTA860", "Request could not be authenticated"},
	},
}
