package mysql


type Client interface {

}

func NewClient() Client {
 	return realClient{nil}
}