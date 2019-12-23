package verification

import "testing"

func TestConsume(t *testing.T) {
	t.Log(Consume("13216817777"))
}

func TestProduce(t *testing.T) {
	key := "13216817777"
	value := Consume(key)
	t.Log(Produce(key, value))
}
