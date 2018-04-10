import(
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// struct to define the Database fields
type QRDatabase struct {
	Server   string
	Database string
	db *mgo.Database
}

//database object to be returned
var db *mgo.Database

func (m *QRDatabase) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	
	m.db =  session.DB(m.Database)
}
