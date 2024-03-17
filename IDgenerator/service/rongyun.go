package service

import "time"

type MessageIdGen struct {
	IdGenerator
}

func (s *MessageIdGen) GetId() (int64, error) {
	now := time.Now()

}

func (s *MessageIdGen) getMessageSeq() (int64, error) {
	return 0, nil
}
