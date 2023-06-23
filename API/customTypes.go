package API

const (
	DATE   = "date"
	AMOUNT = "amount"

	HIDDEN = "hidden"
	PUBLIC = "public"

	NOT_SENDED = "not_sended"
	SENDED     = "sended"

	CARD     = "bank-card"
	CARD_UAH = "bank-card-uah"
	CARD_SNG = "bank-card-sng"
	QIWI     = "qiwi"
	WEBMONEY = "webmoney"
	YOOMONEY = "yoomoney"
	MOBILE   = "mobile"

	DRAFT   = "draft"
	ACTIVE  = "active"
	ARCHIVE = "arctive"
)

type Count int
type Offset int

type StartDate int
type EndDate int

type Sort string
type Reverse bool

type OffsetID string

type ID string
type Status string
