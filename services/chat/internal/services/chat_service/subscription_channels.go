package chatservice

import "github.com/BeInBloom/grpc-chat/services/chat/internal/models"

type subscriptionChannels struct {
	historical chan models.Event
	out        chan models.Event
	err        chan error
}

func (c *subscriptionChannels) closeAll() {
	close(c.out)
	close(c.err)
	close(c.historical)
}
