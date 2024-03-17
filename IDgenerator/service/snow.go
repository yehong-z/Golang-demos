package service

//
//type SnowFlake struct {
//	IdGenerator
//	randNum int64
//}
//
//func (s *SnowFlake) GetId() (int64, error) {
//
//	randNum := s.rand()
//
//	return s.snowId(time.Now(), 1, randNum)
//}
//
//func (s *SnowFlake) rand() int64 {
//	if s.randNum <= 0 {
//		return 0
//	}
//
//	return rand.Int63n(s.randNum)
//}
//
//func (s *SnowFlake) snowId(now time.Time, counter, randNum int64) (int64, error) {
//
//}
