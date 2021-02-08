package usecase

import (
	"errors"
	"github.com/arlenmendes/codepix-go/domain/model"
	"log"
)

type PixUseCase struct {
	PixKeyRepository model.PixKeyRepositoryInterface
}

func (p *PixUseCase) RegisterKey(key string, kind string, accountId string) (*model.PixKey, error) {
	account, err := p.PixKeyRepository.FindAccount(accountId)

	if err != nil {
		return nil, err
	}

	pixKey, err := model.NewPixKey(kind, account, key)

	if err != nil {
		return nil, err
	}

	_, errorRegister := p.PixKeyRepository.RegisterKey(pixKey)

	if errorRegister != nil {
		log.Println(errorRegister)
		return nil, errors.New("unable to process this pix")
	}

	return pixKey, nil
}

func (p *PixUseCase) FindKey(key string, kind string) (*model.PixKey, error) {
	pixKey, err := p.PixKeyRepository.FindKeyByKind(key, kind)
	if err != nil {
		return nil, err
	}
	return pixKey, nil
}