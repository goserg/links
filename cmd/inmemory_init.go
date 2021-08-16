//+build inmemory

package main

import (
	"github.com/goserg/links/domain"
	inmemoryLinkRepo "github.com/goserg/links/link/repository/inmemory"
	"github.com/sirupsen/logrus"
)

func databaseInit() (domain.LinkRepository, func() error, error) {
	logrus.Println("starting app with inmemory storage")
	return inmemoryLinkRepo.New(), func() error { return nil }, nil
}
