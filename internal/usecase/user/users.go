package usersvc

import errs "movie-service/pkg/errors"

type Users struct {
}

func (s *Users) FindById() errs.Errors {
	return errs.WrapError(errs.MetaErrorInternalServer, "error")
}

func (s *Users) FindOne() {

}

func (s *Users) DeleteOne() {

}

func (s *Users) UpdateOne() {

}
