package filesystem

type Storage struct {
	Driver
	config  *Config
	drivers map[string]Driver
}

func NewStorage(conf *Config) (*Storage, error) {
	driver, err := conf.NewDriver(conf.Default)
	if err != nil {
		return nil, err
	}
	return &Storage{Driver: driver}, nil
}

func (s *Storage) Disk(disk string) (Driver, error) {
	if d, ok := s.drivers[disk]; ok {
		return d, nil
	}
	d, err := s.config.NewDriver(disk)
	if err != nil {
		return nil, err
	}
	s.drivers[disk] = d
	return d, nil
}
