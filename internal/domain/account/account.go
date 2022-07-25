package account

type (
	Account struct { // domain entity
		ID     uint64
		UserID uint64
	}
)

type Repo interface {
	ByUserID(userID uint64) (Account, error)
	WithdrawByID(id uint64, amount float64) error
}
