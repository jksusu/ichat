package talk

type sync struct{}

var Sync = new(sync)

func (*sync) Msg() {

}
