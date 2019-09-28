// Go Api server
// @jeffotoni

package cf

import "time"
import "fmt"
import "math/rand"

const (
	MaxClients = 200000
)

const LayoutDateLog = "2006-01-02 15:04:05"
const LayoutDate = "2006-01-02"
const LayoutHour = "15:04:05"

var (
	HOST_MAXBYTE        = 1 << 26 // 64MB
	PORT_METRICS        = "6010"
	PORT_SERVER         = "8081"
	HOST_SERVER         = "0.0.0.0"
	HOST_CONFIG         = HOST_SERVER + ":" + PORT_SERVER
	AUTHORIZATION_REGEX = "Basic %s:%s"

	X_KEY_1 = "12345"
	X_KEY_2 = "54321"

	AUTHORIZATION = fmt.Sprintf(AUTHORIZATION_REGEX, X_KEY_1, X_KEY_2)

	READTIMEOUT  = time.Duration(rand.Intn(1000-700)+700) * time.Millisecond
	WRITETIMEOUT = time.Duration(rand.Intn(1400-900)+900) * time.Millisecond
)

// Config provides
// basic configuration
type Config struct {
	Host           string
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	MaxHeaderBytes int
}

/////// DATA BASE
var (
	DB_NAME     = "apis3produto"
	DB_HOST     = "postgres.local.com"
	DB_USER     = "gopher"
	DB_PASSWORD = "12345"
	DB_PORT     = "5432"
	DB_SSL      = "disable"
	DB_SORCE    = "postgres"
)
