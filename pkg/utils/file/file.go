package file

import (
	log "github.com/sirupsen/logrus"
	"io/ioutil"
)

func ReadFile(path string) ([]byte, error) {
	r, err := ioutil.ReadFile(path)
	if err != nil {
		log.Errorf("Read config %s is failed, err: %s", path, err)
		return nil, err
	}
	return r, nil
}