import(
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type QRDatabase struct {
	Server   string
	Database string
}

var db *mgo.Database

func (m *QRDatabase) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}
