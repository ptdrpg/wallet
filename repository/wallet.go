package repository

import "github.com/ptdrpg/wallet/entity"

func (r *Repository) FindAllWallet() ([]entity.Wallet, error) {
	var wallets []entity.Wallet
	if err := r.DB.Model(&entity.Wallet{}).Order("created_at").Find(&wallets).Error; err != nil {
		return []entity.Wallet{}, err
	}
	return wallets, nil
}

func (r *Repository) FindWalletById(uid string) (entity.Wallet, error) {
	var wallet entity.Wallet
	if err := r.DB.Where("uid = ?", uid).Find(&wallet).Error; err != nil {
		return entity.Wallet{}, err
	}
	return wallet, nil
}

func (r *Repository) CreateWallet(wallet *entity.Wallet) error {
	if err := r.DB.Create(wallet).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateWallet(wallet *entity.Wallet) error {
	if err := r.DB.Model(wallet).Updates(&wallet).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) DeleteContact(uid string) error {
	var wallet entity.Wallet
	if err := r.DB.Where("uid = ?", uid).Delete(&wallet).Error; err != nil {
		return err
	}

	return nil
}
