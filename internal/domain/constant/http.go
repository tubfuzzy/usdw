package constant

var (
	HTTPStatus200 int = 200
	HTTPStatus400 int = 400
	HTTPStatus401 int = 401
	HTTPStatus403 int = 403
	HTTPStatus404 int = 404
	HTTPStatus500 int = 500
)

const (
	OK                string = "20000"
	UNKNOWN           string = "50000"
	INVALID           string = "40000"
	DEADLINE_EXCEEDED string = "50400"
	NOT_FOUND         string = "40400"
	UNAUTHENTICATED   string = "40100"
	PERMISSION_DENIED string = "40300"
	INTERNAL_ERROR    string = "50000"
	UNAVAILABLE       string = "50300"
	FRAMEWORK_ERROR   string = "50200"
)
