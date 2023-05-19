package gate

type gate struct{}

var Gateway = new(gate)

func (*gate) Notice(uid string) {

}

func (*gate) Resp() {

}

func (*gate) Push() {

}

func (*gate) Send() {

}
