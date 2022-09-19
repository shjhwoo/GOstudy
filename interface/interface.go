package main

type Sender interface {
	Send(parcel string)
}

func SendBook(name string, sender Sender) {
	sender.Send(name)
}

func main() {
	koreaPostSender := &koreaPost.PostSender()
	SendBook("1", koreaPostSender)

	fedexSender := &fedex.PostSender()
	SendBook("1", fedexSender)
}